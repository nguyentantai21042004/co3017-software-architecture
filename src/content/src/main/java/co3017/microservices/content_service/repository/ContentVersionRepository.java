package co3017.microservices.content_service.repository;

import co3017.microservices.content_service.models.ContentVersion;
import co3017.microservices.content_service.repository.postgresql.entity.ContentVersionEntity;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.domain.Specification;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * ContentVersion Repository Interface
 * Repository pattern cho ContentVersion domain
 */
public interface ContentVersionRepository {
    ContentVersion save(ContentVersion contentVersion);
    Optional<ContentVersion> findById(Long versionId);
    List<ContentVersion> findAllByIds(List<Long> versionIds);
    Page<ContentVersion> findAll(Specification<ContentVersionEntity> spec, Pageable pageable);
    List<ContentVersion> findByUnitId(UUID unitId);
    Optional<ContentVersion> findActiveVersionByUnitId(UUID unitId);
    void deactivateAllVersionsByUnitId(UUID unitId);
    int deleteByIds(List<Long> versionIds);
    boolean existsById(Long versionId);
}
