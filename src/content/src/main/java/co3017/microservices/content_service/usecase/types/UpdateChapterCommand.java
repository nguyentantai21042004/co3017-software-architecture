package co3017.microservices.content_service.usecase.types;

/**
 * Command object cho việc cập nhật chapter
 */
public class UpdateChapterCommand {
    private final Integer sequenceNumber;

    public UpdateChapterCommand(Integer sequenceNumber) {
        this.sequenceNumber = sequenceNumber;
    }

    public Integer getSequenceNumber() {
        return sequenceNumber;
    }
}
