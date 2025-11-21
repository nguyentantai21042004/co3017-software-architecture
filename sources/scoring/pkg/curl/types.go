package curl

// ContentQuestionResponse represents the response from Content Service
type ContentQuestionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID            int64  `json:"id"`
		Content       string `json:"content"`
		CorrectAnswer string `json:"correct_answer"`
		SkillTag      string `json:"skill_tag"`
		IsRemedial    bool   `json:"is_remedial"`
	} `json:"data"`
}
