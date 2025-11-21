package co3017.microservices.content_service.usecase.types;

import java.util.UUID;

/**
 * Command object cho việc tạo chapter mới
 */
public class CreateChapterCommand {
    private final UUID courseId;
    private final Integer sequenceNumber;

    public CreateChapterCommand(UUID courseId, Integer sequenceNumber) {
        this.courseId = courseId;
        this.sequenceNumber = sequenceNumber;
    }

    public UUID getCourseId() {
        return courseId;
    }

    public Integer getSequenceNumber() {
        return sequenceNumber;
    }
}
