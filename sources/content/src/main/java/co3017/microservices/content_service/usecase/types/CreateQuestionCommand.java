package co3017.microservices.content_service.usecase.types;

/**
 * Command for creating a new question
 * Input data structure for the create use case
 */
public class CreateQuestionCommand {
    private final String content;
    private final String difficulty;
    private final String skillTag;

    public CreateQuestionCommand(String content, String difficulty, String skillTag) {
        this.content = content;
        this.difficulty = difficulty;
        this.skillTag = skillTag;
    }

    public String getContent() {
        return content;
    }

    public String getDifficulty() {
        return difficulty;
    }

    public String getSkillTag() {
        return skillTag;
    }
}
