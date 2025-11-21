package co3017.microservices.content_service.usecase.service;

import co3017.microservices.content_service.models.ContentUnit;
import co3017.microservices.content_service.repository.ContentUnitRepository;
import co3017.microservices.content_service.usecase.ContentUnitUseCase;
import co3017.microservices.content_service.usecase.types.CreateContentUnitCommand;
import co3017.microservices.content_service.usecase.types.UpdateContentUnitCommand;
import co3017.microservices.content_service.usecase.types.ContentUnitSearchCriteria;
import co3017.microservices.content_service.usecase.types.ContentUnitPageResult;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.domain.Specification;
import co3017.microservices.content_service.repository.postgresql.entity.ContentUnitEntity;
import co3017.microservices.content_service.repository.postgresql.specification.ContentUnitSpecification;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * ContentUnit Service Implementation
 * Chứa business logic cho ContentUnit
 */
@Service
public class ContentUnitService implements ContentUnitUseCase {
    private final ContentUnitRepository repository;

    public ContentUnitService(ContentUnitRepository repository) {
        this.repository = repository;
    }

    @Override
    public ContentUnit create(CreateContentUnitCommand command) {
        // Validation
        if (command.getChapterId() == null) {
            throw new IllegalArgumentException("Chapter ID is required");
        }
        if (command.getUnitType() == null) {
            throw new IllegalArgumentException("Unit type is required");
        }
        if (command.getMetadataConfig() == null || command.getMetadataConfig().isEmpty()) {
            throw new IllegalArgumentException("Metadata config is required");
        }

        // Create domain object
        ContentUnit contentUnit = new ContentUnit(
            command.getChapterId(),
            command.getUnitType(),
            command.getMetadataConfig()
        );

        // Business validation
        if (!contentUnit.isValid()) {
            throw new IllegalArgumentException("Invalid content unit data");
        }

        // Save to database
        return repository.save(contentUnit);
    }

    @Override
    public Optional<ContentUnit> detail(UUID unitId) {
        if (unitId == null) {
            throw new IllegalArgumentException("Unit ID is required");
        }
        return repository.findById(unitId);
    }

    @Override
    public ContentUnitPageResult list(ContentUnitSearchCriteria criteria) {
        if (criteria == null) {
            criteria = new ContentUnitSearchCriteria();
        }

        // Create specification
        ContentUnitSpecification.Builder specBuilder = new ContentUnitSpecification.Builder();
        if (criteria.getKeyword() != null) {
            specBuilder.keyword(criteria.getKeyword());
        }
        if (criteria.getUnitIds() != null) {
            specBuilder.unitIds(criteria.getUnitIds());
        }
        if (criteria.getChapterId() != null) {
            specBuilder.chapterId(criteria.getChapterId());
        }
        if (criteria.getUnitType() != null) {
            specBuilder.unitType(criteria.getUnitType());
        }
        Specification<ContentUnitEntity> spec = specBuilder.build();

        // Create pageable
        Pageable pageable = PageRequest.of(criteria.getPage(), criteria.getSize(), criteria.getSort());

        // Query with pagination
        Page<ContentUnit> page = repository.findAll(spec, pageable);

        return new ContentUnitPageResult(
            page.getContent(),
            page.getTotalElements(),
            page.getTotalPages(),
            page.getNumber(),
            page.getSize(),
            page.isFirst(),
            page.isLast()
        );
    }

    @Override
    public Optional<ContentUnit> update(UUID unitId, UpdateContentUnitCommand command) {
        if (unitId == null) {
            throw new IllegalArgumentException("Unit ID is required");
        }
        if (command == null) {
            throw new IllegalArgumentException("Command is required");
        }

        // Find existing content unit
        Optional<ContentUnit> existingOpt = repository.findById(unitId);
        if (existingOpt.isEmpty()) {
            return Optional.empty();
        }

        ContentUnit existing = existingOpt.get();

        // Update fields
        if (command.getUnitType() != null) {
            existing.updateUnitType(command.getUnitType());
        }
        if (command.getMetadataConfig() != null) {
            existing.updateMetadataConfig(command.getMetadataConfig());
        }

        // Save updated content unit
        ContentUnit updated = repository.save(existing);
        return Optional.of(updated);
    }

    @Override
    public int deletes(List<UUID> unitIds) {
        if (unitIds == null || unitIds.isEmpty()) {
            return 0;
        }

        // Validate all IDs exist
        List<ContentUnit> existingUnits = repository.findAllByIds(unitIds);
        if (existingUnits.size() != unitIds.size()) {
            throw new IllegalArgumentException("Some content units not found");
        }

        return repository.deleteByIds(unitIds);
    }

    @Override
    public boolean existsById(UUID unitId) {
        if (unitId == null) {
            return false;
        }
        return repository.existsById(unitId);
    }
}
