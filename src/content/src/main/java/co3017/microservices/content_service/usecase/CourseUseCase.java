package co3017.microservices.content_service.usecase;

import co3017.microservices.content_service.models.Course;
import co3017.microservices.content_service.usecase.types.CreateCourseCommand;
import co3017.microservices.content_service.usecase.types.UpdateCourseCommand;
import co3017.microservices.content_service.usecase.types.CourseSearchCriteria;
import co3017.microservices.content_service.usecase.types.CoursePageResult;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Course Use Case Interface - Tất cả use cases cho Course domain
 */
public interface CourseUseCase {
    // Create
    Course create(CreateCourseCommand command);
    
    // Read
    Optional<Course> detail(UUID courseId);
    CoursePageResult list(CourseSearchCriteria criteria);
    
    // Update
    Optional<Course> update(UUID courseId, UpdateCourseCommand command);
    
    // Delete
    int deletes(List<UUID> courseIds);
    
    // Existence check
    boolean existsById(UUID courseId);
    boolean existsByTitle(String title);
}
