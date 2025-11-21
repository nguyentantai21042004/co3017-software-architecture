package co3017.microservices.content_service.adapter.http.response;

import co3017.microservices.content_service.adapter.http.dto.CreateTestRequest;
import co3017.microservices.content_service.adapter.http.dto.CreateCourseRequest;
import co3017.microservices.content_service.adapter.http.dto.UpdateCourseRequest;
import co3017.microservices.content_service.adapter.http.dto.CreateChapterRequest;
import co3017.microservices.content_service.adapter.http.dto.UpdateChapterRequest;
import co3017.microservices.content_service.adapter.http.dto.CreateContentUnitRequest;
import co3017.microservices.content_service.adapter.http.dto.UpdateContentUnitRequest;
import co3017.microservices.content_service.adapter.http.dto.CreateContentVersionRequest;
import co3017.microservices.content_service.adapter.http.dto.UpdateContentVersionRequest;
import co3017.microservices.content_service.usecase.types.CreateTestCommand;
import co3017.microservices.content_service.usecase.types.CreateCourseCommand;
import co3017.microservices.content_service.usecase.types.UpdateCourseCommand;
import co3017.microservices.content_service.usecase.types.CreateChapterCommand;
import co3017.microservices.content_service.usecase.types.UpdateChapterCommand;
import co3017.microservices.content_service.usecase.types.CreateContentUnitCommand;
import co3017.microservices.content_service.usecase.types.UpdateContentUnitCommand;
import co3017.microservices.content_service.usecase.types.CreateContentVersionCommand;
import co3017.microservices.content_service.usecase.types.UpdateContentVersionCommand;

/**
 * Command Builder
 * Chuyển đổi từ Request DTO → Command
 */
public class CommandBuilder {

    /**
     * CreateTestRequest → CreateTestCommand
     */
    public static CreateTestCommand toCreateTestCommand(CreateTestRequest request) {
        return new CreateTestCommand(
                request.getTitle(),
                request.getDescription(),
                request.getDuration(),
                request.getMaxScore());
    }

    /**
     * CreateCourseRequest → CreateCourseCommand
     */
    public static CreateCourseCommand toCreateCourseCommand(CreateCourseRequest request) {
        return new CreateCourseCommand(
                request.getTitle(),
                request.getDescription(),
                request.getInstructorId(),
                request.getStructureType());
    }

    /**
     * UpdateCourseRequest → UpdateCourseCommand
     */
    public static UpdateCourseCommand toUpdateCourseCommand(UpdateCourseRequest request) {
        return new UpdateCourseCommand(
                request.getTitle(),
                request.getDescription(),
                request.getStructureType());
    }

    /**
     * CreateChapterRequest → CreateChapterCommand
     */
    public static CreateChapterCommand toCreateChapterCommand(CreateChapterRequest request) {
        return new CreateChapterCommand(
                request.getCourseId(),
                request.getSequenceNumber());
    }

    /**
     * UpdateChapterRequest → UpdateChapterCommand
     */
    public static UpdateChapterCommand toUpdateChapterCommand(UpdateChapterRequest request) {
        return new UpdateChapterCommand(
                request.getSequenceNumber());
    }

    /**
     * CreateContentUnitRequest → CreateContentUnitCommand
     */
    public static CreateContentUnitCommand toCreateContentUnitCommand(CreateContentUnitRequest request) {
        return new CreateContentUnitCommand(
                request.getChapterId(),
                request.getUnitType(),
                request.getMetadataConfig());
    }

    /**
     * UpdateContentUnitRequest → UpdateContentUnitCommand
     */
    public static UpdateContentUnitCommand toUpdateContentUnitCommand(UpdateContentUnitRequest request) {
        return new UpdateContentUnitCommand(
                request.getUnitType(),
                request.getMetadataConfig());
    }

    /**
     * CreateContentVersionRequest → CreateContentVersionCommand
     */
    public static CreateContentVersionCommand toCreateContentVersionCommand(CreateContentVersionRequest request) {
        return new CreateContentVersionCommand(
                request.getUnitId(),
                request.getVersionNumber(),
                request.getContentData(),
                request.isActive());
    }

    /**
     * UpdateContentVersionRequest → UpdateContentVersionCommand
     */
    public static UpdateContentVersionCommand toUpdateContentVersionCommand(UpdateContentVersionRequest request) {
        return new UpdateContentVersionCommand(
                request.getVersionNumber(),
                request.getContentData(),
                request.getIsActive());
    }
}
