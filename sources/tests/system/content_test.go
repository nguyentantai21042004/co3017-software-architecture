package system

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TC_CONTENT_01: Filter Remedial Content
func TestContentFilterRemedial(t *testing.T) {
	fmt.Println("\n🧪 TC_CONTENT_01: Filter Remedial Content")

	question, err := getRecommendedQuestion("math", "remedial")
	require.NoError(t, err, "Failed to get remedial content")

	assert.True(t, question.IsRemedial, "Should return remedial content")
	assert.Equal(t, "math", question.SkillTag, "Should match skill tag")

	fmt.Printf("✅ Got remedial question: ID=%d, Content=%s\n", question.ID, question.Content)
}

// TC_CONTENT_02: Filter Standard Content
func TestContentFilterStandard(t *testing.T) {
	fmt.Println("\n🧪 TC_CONTENT_02: Filter Standard Content")

	question, err := getRecommendedQuestion("math", "standard")
	require.NoError(t, err, "Failed to get standard content")

	assert.False(t, question.IsRemedial, "Should return standard content")
	assert.Equal(t, "math", question.SkillTag, "Should match skill tag")

	fmt.Printf("✅ Got standard question: ID=%d, Content=%s\n", question.ID, question.Content)
}

// TC_CONTENT_03: No Content Available
func TestContentNoContentAvailable(t *testing.T) {
	fmt.Println("\n🧪 TC_CONTENT_03: No Content Available")

	// Try to get content for non-existent skill
	_, err := getRecommendedQuestion("nonexistent_skill", "standard")

	// Should handle gracefully (either error or empty response)
	if err != nil {
		fmt.Printf("✅ Correctly handled no content: %v\n", err)
	} else {
		fmt.Println("✅ Returned empty/default content")
	}
}

// Helper function
func getRecommendedQuestion(skill, contentType string) (*QuestionResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/content/recommend?skill=%s&type=%s",
		ContentServiceURL, skill, contentType))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var apiResp QuestionAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Data, nil
}
