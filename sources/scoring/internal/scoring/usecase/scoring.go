package usecase

import (
	"context"
	"fmt"
	"scoring/internal/model"
	"scoring/internal/scoring"
	"time"
)

// SubmitAnswer submits and scores a user's answer
func (uc *usecase) SubmitAnswer(ctx context.Context, input scoring.SubmitInput) (scoring.SubmitOutput, error) {
	uc.l.Infof(ctx, "scoring.usecase.SubmitAnswer: starting | %s=%s | %s=%d | %s=%s",
		ErrCtxUserID, input.UserID, ErrCtxQuestionID, input.QuestionID, ErrCtxAnswer, input.Answer)

	// Step 1: Fetch correct answer from Content Service using curl client
	questionResp, err := uc.contentClient.GetQuestion(ctx, input.QuestionID)
	if err != nil {
		uc.l.Errorf(ctx, "scoring.usecase.SubmitAnswer: %s | %s=%s | %s=%d | error=%v",
			ErrMsgFetchQuestionFailed, ErrCtxUserID, input.UserID, ErrCtxQuestionID, input.QuestionID, err)
		return scoring.SubmitOutput{}, fmt.Errorf("%s: %w", ErrMsgFetchQuestionFailed, err)
	}

	correctAnswer := questionResp.Data.CorrectAnswer
	skillTag := questionResp.Data.SkillTag

	uc.l.Infof(ctx, "scoring.usecase.SubmitAnswer: question fetched | %s=%d | %s=%s | correct_answer=%s",
		ErrCtxQuestionID, input.QuestionID, ErrCtxSkillTag, skillTag, correctAnswer)

	// Step 2: Score the answer
	isCorrect := input.Answer == correctAnswer
	score := 0
	if isCorrect {
		score = 100
	}
	isPassed := score >= 50

	feedback := "Incorrect answer. Please try again!"
	if isCorrect {
		feedback = "Correct! Well done."
	}

	uc.l.Infof(ctx, "scoring.usecase.SubmitAnswer: scoring complete | %s=%s | %s=%d | %s=%v | %s=%d",
		ErrCtxUserID, input.UserID, ErrCtxQuestionID, input.QuestionID, ErrCtxIsCorrect, isCorrect, ErrCtxScore, score)

	// Step 3: Save to database
	submission := &model.Submission{
		UserID:          input.UserID,
		QuestionID:      input.QuestionID,
		SubmittedAnswer: input.Answer,
		ScoreAwarded:    score,
		IsPassed:        isPassed,
	}

	err = uc.repo.Create(submission)
	if err != nil {
		uc.l.Errorf(ctx, "scoring.usecase.SubmitAnswer: %s | %s=%s | %s=%d | error=%v",
			ErrMsgSaveSubmissionFailed, ErrCtxUserID, input.UserID, ErrCtxQuestionID, input.QuestionID, err)
		return scoring.SubmitOutput{}, fmt.Errorf("%s: %w", ErrMsgSaveSubmissionFailed, err)
	}

	uc.l.Infof(ctx, "scoring.usecase.SubmitAnswer: submission saved | %s=%d | %s=%s",
		ErrCtxSubmissionID, submission.ID, ErrCtxUserID, input.UserID)

	// Step 4: Publish event to RabbitMQ (async, don't block response)
	go func() {
		bgCtx := context.Background()
		event := model.SubmissionEvent{
			Event:         "SubmissionCompleted",
			UserID:        input.UserID,
			SkillTag:      skillTag,
			ScoreObtained: score,
			Timestamp:     time.Now().Format(time.RFC3339),
		}

		err := uc.publisher.PublishSubmissionEvent(event)
		if err != nil {
			uc.l.Errorf(bgCtx, "scoring.usecase.SubmitAnswer: %s | %s=%s | %s=%s | %s=%d | error=%v",
				ErrMsgPublishEventFailed, ErrCtxUserID, input.UserID, ErrCtxSkillTag, skillTag, ErrCtxScore, score, err)
		} else {
			uc.l.Infof(bgCtx, "scoring.usecase.SubmitAnswer: event published | %s=%s | %s=%s | %s=%d",
				ErrCtxUserID, input.UserID, ErrCtxSkillTag, skillTag, ErrCtxScore, score)
		}
	}()

	return scoring.SubmitOutput{
		Correct:  isCorrect,
		Score:    score,
		Feedback: feedback,
	}, nil
}
