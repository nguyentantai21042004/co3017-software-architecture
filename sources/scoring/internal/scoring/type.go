package scoring

// SubmitInput is the input for submitting an answer
type SubmitInput struct {
	UserID     string
	QuestionID int64
	Answer     string
}

// SubmitOutput is the output with scoring results
type SubmitOutput struct {
	Correct  bool
	Score    int
	Feedback string
}
