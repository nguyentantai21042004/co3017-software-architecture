package co3017.microservices.content_service.usecase.types;

/**
 * Command object cho việc cập nhật path condition
 */
public class UpdatePathConditionCommand {
    private final Integer requiredScorePct;

    public UpdatePathConditionCommand(Integer requiredScorePct) {
        this.requiredScorePct = requiredScorePct;
    }

    public Integer getRequiredScorePct() {
        return requiredScorePct;
    }
}
