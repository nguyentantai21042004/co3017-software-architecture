package usecase

import (
	"context"
	"fmt"

	"adaptive-engine/internal/adaptive"
)

const (
	MASTERY_THRESHOLD = 50
)

// RecommendNextLesson recommends the next lesson based on user's mastery
func (uc *usecase) RecommendNextLesson(ctx context.Context, input adaptive.RecommendInput) (adaptive.RecommendOutput, error) {
	uc.l.Infof(ctx, "adaptive.usecase.RecommendNextLesson: starting | %s=%s | %s=%s",
		ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.CurrentSkill)

	// Fetch mastery from learner service using curl client
	masteryResp, err := uc.learnerClient.GetMastery(ctx, input.UserID, input.CurrentSkill)
	if err != nil {
		uc.l.Errorf(ctx, "adaptive.usecase.RecommendNextLesson: %s | %s=%s | %s=%s | error=%v",
			ErrMsgFetchMasteryFailed, ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.CurrentSkill, err)
		return adaptive.RecommendOutput{}, fmt.Errorf("%s: %w", ErrMsgFetchMasteryFailed, err)
	}

	uc.l.Infof(ctx, "adaptive.usecase.RecommendNextLesson: mastery fetched | %s=%s | %s=%s | mastery_score=%d",
		ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.CurrentSkill, masteryResp.MasteryScore)

	// Determine content type based on mastery score
	var contentType, reason string
	if masteryResp.MasteryScore < MASTERY_THRESHOLD {
		contentType = "remedial"
		reason = fmt.Sprintf("Your mastery is %d%%. Let's review the basics.", masteryResp.MasteryScore)
		uc.l.Infof(ctx, "adaptive.usecase.RecommendNextLesson: recommending REMEDIAL | %s=%s | score=%d < threshold=%d",
			ErrCtxUserID, input.UserID, masteryResp.MasteryScore, MASTERY_THRESHOLD)
	} else {
		contentType = "standard"
		reason = fmt.Sprintf("Great! Your mastery is %d%%. Continue with the next challenge.", masteryResp.MasteryScore)
		uc.l.Infof(ctx, "adaptive.usecase.RecommendNextLesson: recommending STANDARD | %s=%s | score=%d >= threshold=%d",
			ErrCtxUserID, input.UserID, masteryResp.MasteryScore, MASTERY_THRESHOLD)
	}

	// Fetch content from content service using curl client
	contentResp, err := uc.contentClient.GetRecommendation(ctx, input.CurrentSkill, contentType)
	if err != nil {
		uc.l.Errorf(ctx, "adaptive.usecase.RecommendNextLesson: %s | %s=%s | %s=%s | %s=%s | error=%v",
			ErrMsgFetchContentFailed, ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.CurrentSkill,
			ErrCtxContentType, contentType, err)
		return adaptive.RecommendOutput{}, fmt.Errorf("%s: %w", ErrMsgFetchContentFailed, err)
	}

	uc.l.Infof(ctx, "adaptive.usecase.RecommendNextLesson: content fetched | %s=%s | question_id=%d | %s=%s",
		ErrCtxUserID, input.UserID, contentResp.Data.ID, ErrCtxContentType, contentType)

	return adaptive.RecommendOutput{
		NextLessonID: int(contentResp.Data.ID),
		Reason:       reason,
		MasteryScore: masteryResp.MasteryScore,
		ContentType:  contentType,
	}, nil
}
