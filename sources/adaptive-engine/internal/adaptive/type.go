package adaptive

// RecommendInput is the input for recommending next lesson
type RecommendInput struct {
	UserID       string
	CurrentSkill string
}

// RecommendOutput is the output with recommendation details
type RecommendOutput struct {
	NextLessonID int
	Reason       string
	MasteryScore int
	ContentType  string // "remedial" or "standard"
}
