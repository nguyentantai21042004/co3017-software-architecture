package co3017.microservices.content_service.usecase;

import co3017.microservices.content_service.models.ContentVersion;
import co3017.microservices.content_service.usecase.types.CreateContentVersionCommand;
import co3017.microservices.content_service.usecase.types.UpdateContentVersionCommand;
import co3017.microservices.content_service.usecase.types.ContentVersionSearchCriteria;
import co3017.microservices.content_service.usecase.types.ContentVersionPageResult;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * ContentVersion Use Case Interface - Tất cả use cases cho ContentVersion domain
 */
public interface ContentVersionUseCase {
    // Create
    ContentVersion create(CreateContentVersionCommand command);
    
    // Read
    Optional<ContentVersion> detail(Long versionId);
    ContentVersionPageResult list(ContentVersionSearchCriteria criteria);
    List<ContentVersion> findByUnitId(UUID unitId);
    Optional<ContentVersion> findActiveVersionByUnitId(UUID unitId);
    
    // Update
    Optional<ContentVersion> update(Long versionId, UpdateContentVersionCommand command);
    Optional<ContentVersion> setActiveVersion(Long versionId, UUID unitId);
    
    // Delete
    int deletes(List<Long> versionIds);
    
    // Existence check
    boolean existsById(Long versionId);
}
