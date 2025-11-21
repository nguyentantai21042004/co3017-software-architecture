package co3017.microservices.content_service.usecase;

import co3017.microservices.content_service.models.PathCondition;
import co3017.microservices.content_service.usecase.types.CreatePathConditionCommand;
import co3017.microservices.content_service.usecase.types.UpdatePathConditionCommand;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * PathCondition Use Case Interface - Tất cả use cases cho PathCondition domain
 */
public interface PathConditionUseCase {
    // Create
    PathCondition create(CreatePathConditionCommand command);
    
    // Read
    Optional<PathCondition> detail(UUID conditionId);
    List<PathCondition> findBySourceUnitId(UUID sourceUnitId);
    List<PathCondition> findByTargetUnitId(UUID targetUnitId);
    
    // Update
    Optional<PathCondition> update(UUID conditionId, UpdatePathConditionCommand command);
    
    // Delete
    int deletes(List<UUID> conditionIds);
    
    // Existence check
    boolean existsById(UUID conditionId);
}
