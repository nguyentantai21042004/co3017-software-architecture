package learner

// GetMasteryInput is the input for getting mastery
type GetMasteryInput struct {
	UserID   string
	SkillTag string
}

// MasteryOutput is the output with mastery details
type MasteryOutput struct {
	UserID       string
	SkillTag     string
	MasteryScore int
	LastUpdated  string
}

// UpdateMasteryInput is the input for updating mastery from submission event
type UpdateMasteryInput struct {
	UserID        string
	SkillTag      string
	ScoreObtained int
}
