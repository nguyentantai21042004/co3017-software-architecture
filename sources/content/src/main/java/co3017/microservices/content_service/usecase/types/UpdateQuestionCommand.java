package co3017.microservices.content_service.usecase.types;

/**
 * Command for updating an existing question
 * Input data structure for the update use case
 */
public class UpdateQuestionCommand {
    private final Integer id;
    private final String content;
    private final String difficulty;
    private final String skillTag;

    public UpdateQuestionCommand(Integer id, String content, String difficulty, String skillTag) {
        this.id = id;
        this.content = content;
        this.difficulty = difficulty;
        this.skillTag = skillTag;
    }

    public Integer getId() {
        return id;
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
