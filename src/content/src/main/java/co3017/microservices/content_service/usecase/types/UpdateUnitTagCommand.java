package co3017.microservices.content_service.usecase.types;

/**
 * Command object cho việc cập nhật unit tag
 */
public class UpdateUnitTagCommand {
    private final Float relevanceScore;

    public UpdateUnitTagCommand(Float relevanceScore) {
        this.relevanceScore = relevanceScore;
    }

    public Float getRelevanceScore() {
        return relevanceScore;
    }
}
