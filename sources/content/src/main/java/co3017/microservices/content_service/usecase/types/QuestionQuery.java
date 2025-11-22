package co3017.microservices.content_service.usecase.types;

/**
 * Query parameters for filtering questions
 * Input data structure for search/filter use cases
 */
public class QuestionQuery {
    private final Integer difficultyLevel;
    private final String skillTag;

    public QuestionQuery(Integer difficultyLevel, String skillTag) {
        this.difficultyLevel = difficultyLevel;
        this.skillTag = skillTag;
    }

    public Integer getDifficultyLevel() {
        return difficultyLevel;
    }

    public String getSkillTag() {
        return skillTag;
    }

    public boolean hasDifficultyLevel() {
        return difficultyLevel != null;
    }

    public boolean hasSkillTag() {
        return skillTag != null && !skillTag.trim().isEmpty();
    }
}
