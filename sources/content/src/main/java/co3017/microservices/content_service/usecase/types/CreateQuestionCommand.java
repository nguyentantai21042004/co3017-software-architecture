package co3017.microservices.content_service.usecase.types;

/**
 * Command for creating a new question
 * Input data structure for the create use case
 */
public class CreateQuestionCommand {
    private final String content;
    private final Integer difficultyLevel;
    private final String skillTag;
    private final String correctAnswer;
    private final Boolean isRemedial;

    public CreateQuestionCommand(String content, Integer difficultyLevel, String skillTag, String correctAnswer,
            Boolean isRemedial) {
        this.content = content;
        this.difficultyLevel = difficultyLevel;
        this.skillTag = skillTag;
        this.correctAnswer = correctAnswer;
        this.isRemedial = isRemedial;
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
