package co3017.microservices.content_service.usecase.types;

/**
 * Command for updating an existing question
 * Input data structure for the update use case
 */
public class UpdateQuestionCommand {
    private final Integer id;
    private final String content;
    private final Integer difficultyLevel;
    private final String skillTag;
    private final String correctAnswer;
    private final Boolean isRemedial;

    public UpdateQuestionCommand(Integer id, String content, Integer difficultyLevel, String skillTag,
            String correctAnswer, Boolean isRemedial) {
        this.id = id;
        this.content = content;
        this.difficultyLevel = difficultyLevel;
        this.skillTag = skillTag;
        this.correctAnswer = correctAnswer;
        this.isRemedial = isRemedial;
    }

    public Integer getId() {
        return id;
    }

    public String getContent() {
        return content;
    }

    public Integer getDifficultyLevel() {
        return difficultyLevel;
    }

    public String getSkillTag() {
        return skillTag;
    }

    public String getCorrectAnswer() {
        return correctAnswer;
    }

    public Boolean getIsRemedial() {
        return isRemedial;
    }
}
