package co3017.microservices.content_service.usecase.types;

/**
 * Command object cho việc cập nhật metadata tag
 */
public class UpdateMetadataTagCommand {
    private final String tagName;

    public UpdateMetadataTagCommand(String tagName) {
        this.tagName = tagName;
    }

    public String getTagName() {
        return tagName;
    }
}
