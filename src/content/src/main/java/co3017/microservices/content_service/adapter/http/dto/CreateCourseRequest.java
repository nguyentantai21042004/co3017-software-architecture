package co3017.microservices.content_service.adapter.http.dto;

import co3017.microservices.content_service.models.Course;

import java.util.UUID;

/**
 * DTO cho request tạo course mới
 */
public class CreateCourseRequest {
    private String title;
    private String description;
    private UUID instructorId;
    private Course.StructureType structureType;

    public CreateCourseRequest() {
    }

    public CreateCourseRequest(String title, String description, UUID instructorId,
            Course.StructureType structureType) {
        this.title = title;
        this.description = description;
        this.instructorId = instructorId;
        this.structureType = structureType;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public UUID getInstructorId() {
        return instructorId;
    }

    public void setInstructorId(UUID instructorId) {
        this.instructorId = instructorId;
    }

    public Course.StructureType getStructureType() {
        return structureType;
    }

    public void setStructureType(Course.StructureType structureType) {
        this.structureType = structureType;
    }
}
