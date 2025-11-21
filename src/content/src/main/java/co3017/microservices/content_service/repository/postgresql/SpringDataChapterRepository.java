package co3017.microservices.content_service.repository.postgresql;

import co3017.microservices.content_service.repository.postgresql.entity.ChapterEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;

import java.util.UUID;

/**
 * Spring Data JPA Repository cho Chapter với dynamic query support
 */
public interface SpringDataChapterRepository extends JpaRepository<ChapterEntity, UUID>, JpaSpecificationExecutor<ChapterEntity> {
    
    boolean existsByCourseIdAndSequenceNumber(UUID courseId, Integer sequenceNumber);
}
