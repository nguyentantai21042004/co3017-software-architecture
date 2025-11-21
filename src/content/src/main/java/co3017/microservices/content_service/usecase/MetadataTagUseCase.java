package co3017.microservices.content_service.usecase;

import co3017.microservices.content_service.models.MetadataTag;
import co3017.microservices.content_service.usecase.types.CreateMetadataTagCommand;
import co3017.microservices.content_service.usecase.types.UpdateMetadataTagCommand;
import co3017.microservices.content_service.usecase.types.MetadataTagSearchCriteria;
import co3017.microservices.content_service.usecase.types.MetadataTagPageResult;

import java.util.List;
import java.util.Optional;

/**
 * MetadataTag Use Case Interface - Tất cả use cases cho MetadataTag domain
 */
public interface MetadataTagUseCase {
    // Create
    MetadataTag create(CreateMetadataTagCommand command);
    
    // Read
    Optional<MetadataTag> detail(Integer tagId);
    MetadataTagPageResult list(MetadataTagSearchCriteria criteria);
    
    // Update
    Optional<MetadataTag> update(Integer tagId, UpdateMetadataTagCommand command);
    
    // Delete
    int deletes(List<Integer> tagIds);
    
    // Existence check
    boolean existsById(Integer tagId);
    boolean existsByTagName(String tagName);
}
