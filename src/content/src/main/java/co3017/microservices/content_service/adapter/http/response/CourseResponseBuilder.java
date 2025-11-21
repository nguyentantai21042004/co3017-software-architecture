package co3017.microservices.content_service.adapter.http.response;

import co3017.microservices.content_service.adapter.http.dto.CourseResponse;
import co3017.microservices.content_service.models.Course;

import java.util.List;
import java.util.stream.Collectors;

/**
 * Response Builder cho Course
 * Chuyển đổi từ Domain → Response DTO
 */
public class CourseResponseBuilder {

    /**
     * Chuyển Domain Course → CourseResponse DTO
     */
    public static CourseResponse toResponse(Course course) {
        return new CourseResponse(
                course.getCourseId(),
                course.getTitle(),
                course.getDescription(),
                course.getInstructorId(),
                course.getStructureType(),
                course.getCreatedAt(),
                course.getUpdatedAt());
    }

    /**
     * Chuyển List Domain Course → List CourseResponse DTO
     */
    public static List<CourseResponse> toResponseList(List<Course> courses) {
        return courses.stream()
                .map(CourseResponseBuilder::toResponse)
                .collect(Collectors.toList());
    }
}
