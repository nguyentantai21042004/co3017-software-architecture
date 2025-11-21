package co3017.microservices.content_service.repository;

import co3017.microservices.content_service.models.Chapter;
import co3017.microservices.content_service.usecase.types.ChapterSearchCriteria;
import co3017.microservices.content_service.usecase.types.ChapterPageResult;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Port Out - Repository interface cho Chapter
 * Infrastructure layer sẽ implement interface này
 */
public interface ChapterRepository {
    Chapter save(Chapter chapter);
    Optional<Chapter> findById(UUID chapterId);
    ChapterPageResult search(ChapterSearchCriteria criteria);
    void deleteByIds(List<UUID> chapterIds);
    boolean existsById(UUID chapterId);
    boolean existsByCourseIdAndSequence(UUID courseId, Integer sequenceNumber);
}
