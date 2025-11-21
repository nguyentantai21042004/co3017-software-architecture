package co3017.microservices.content_service.usecase;

import co3017.microservices.content_service.models.Chapter;
import co3017.microservices.content_service.usecase.types.CreateChapterCommand;
import co3017.microservices.content_service.usecase.types.UpdateChapterCommand;
import co3017.microservices.content_service.usecase.types.ChapterSearchCriteria;
import co3017.microservices.content_service.usecase.types.ChapterPageResult;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Chapter Use Case Interface - Tất cả use cases cho Chapter domain
 */
public interface ChapterUseCase {
    // Create
    Chapter create(CreateChapterCommand command);
    
    // Read
    Optional<Chapter> detail(UUID chapterId);
    ChapterPageResult list(ChapterSearchCriteria criteria);
    
    // Update
    Optional<Chapter> update(UUID chapterId, UpdateChapterCommand command);
    
    // Delete
    int deletes(List<UUID> chapterIds);
    
    // Existence check
    boolean existsById(UUID chapterId);
    boolean existsByCourseIdAndSequence(UUID courseId, Integer sequenceNumber);
}
