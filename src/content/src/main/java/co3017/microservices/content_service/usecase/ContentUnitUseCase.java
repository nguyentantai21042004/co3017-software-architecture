package co3017.microservices.content_service.usecase;

import co3017.microservices.content_service.models.ContentUnit;
import co3017.microservices.content_service.usecase.types.CreateContentUnitCommand;
import co3017.microservices.content_service.usecase.types.UpdateContentUnitCommand;
import co3017.microservices.content_service.usecase.types.ContentUnitSearchCriteria;
import co3017.microservices.content_service.usecase.types.ContentUnitPageResult;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * ContentUnit Use Case Interface - Tất cả use cases cho ContentUnit domain
 */
public interface ContentUnitUseCase {
    // Create
    ContentUnit create(CreateContentUnitCommand command);
    
    // Read
    Optional<ContentUnit> detail(UUID unitId);
    ContentUnitPageResult list(ContentUnitSearchCriteria criteria);
    
    // Update
    Optional<ContentUnit> update(UUID unitId, UpdateContentUnitCommand command);
    
    // Delete
    int deletes(List<UUID> unitIds);
    
    // Existence check
    boolean existsById(UUID unitId);
}
