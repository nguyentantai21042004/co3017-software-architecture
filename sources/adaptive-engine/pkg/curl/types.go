package curl

// MasteryResponse represents the response from Learner Service
type MasteryResponse struct {
	UserID       string `json:"user_id"`
	SkillTag     string `json:"skill_tag"`
	MasteryScore int    `json:"mastery_score"`
	LastUpdated  string `json:"last_updated"`
}

// ContentResponse represents the response from Content Service
type ContentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID              int64  `json:"id"`
		Content         string `json:"content"`
		CorrectAnswer   string `json:"correct_answer"`
		SkillTag        string `json:"skill_tag"`
		DifficultyLevel int    `json:"difficulty_level"`
		IsRemedial      bool   `json:"is_remedial"`
	} `json:"data"`
}
