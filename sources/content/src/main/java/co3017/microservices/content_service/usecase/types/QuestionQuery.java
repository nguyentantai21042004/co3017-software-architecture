package co3017.microservices.content_service.usecase.types;

/**
 * Query parameters for filtering questions
 * Input data structure for search/filter use cases
 */
public class QuestionQuery {
    private final String difficulty;
    private final String skillTag;

    public QuestionQuery(String difficulty, String skillTag) {
        this.difficulty = difficulty;
        this.skillTag = skillTag;
    }

    public String getDifficulty() {
        return difficulty;
    }

    public String getSkillTag() {
        return skillTag;
    }

    public boolean hasDifficulty() {
        return difficulty != null && !difficulty.trim().isEmpty();
    }

    public boolean hasSkillTag() {
        return skillTag != null && !skillTag.trim().isEmpty();
    }
}
