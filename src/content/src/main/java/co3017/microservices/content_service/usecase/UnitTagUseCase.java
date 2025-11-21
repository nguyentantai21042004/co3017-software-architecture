package co3017.microservices.content_service.usecase;

import co3017.microservices.content_service.models.UnitTag;
import co3017.microservices.content_service.usecase.types.CreateUnitTagCommand;
import co3017.microservices.content_service.usecase.types.UpdateUnitTagCommand;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * UnitTag Use Case Interface - Tất cả use cases cho UnitTag domain
 */
public interface UnitTagUseCase {
    // Create
    UnitTag create(CreateUnitTagCommand command);
    
    // Read
    Optional<UnitTag> findByUnitIdAndTagId(UUID unitId, Integer tagId);
    List<UnitTag> findByUnitId(UUID unitId);
    List<UnitTag> findByTagId(Integer tagId);
    
    // Update
    Optional<UnitTag> update(UUID unitId, Integer tagId, UpdateUnitTagCommand command);
    
    // Delete
    int deleteByUnitIdAndTagId(UUID unitId, Integer tagId);
    int deleteByUnitId(UUID unitId);
    int deleteByTagId(Integer tagId);
    
    // Existence check
    boolean existsByUnitIdAndTagId(UUID unitId, Integer tagId);
}
