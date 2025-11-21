package co3017.microservices.content_service.repository;

import co3017.microservices.content_service.models.Course;
import co3017.microservices.content_service.usecase.types.CourseSearchCriteria;
import co3017.microservices.content_service.usecase.types.CoursePageResult;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Port Out - Repository interface cho Course
 * Infrastructure layer sẽ implement interface này
 */
public interface CourseRepository {
    Course save(Course course);
    Optional<Course> findById(UUID courseId);
    CoursePageResult search(CourseSearchCriteria criteria);
    void deleteByIds(List<UUID> courseIds);
    boolean existsById(UUID courseId);
    boolean existsByTitle(String title);
}
