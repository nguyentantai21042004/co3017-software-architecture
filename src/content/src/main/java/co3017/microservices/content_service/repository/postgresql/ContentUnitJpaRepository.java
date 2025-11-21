package co3017.microservices.content_service.repository.postgresql;

import co3017.microservices.content_service.repository.postgresql.entity.ContentUnitEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

import java.util.UUID;

/**
 * JPA Repository Interface cho ContentUnitEntity
 */
@Repository
public interface ContentUnitJpaRepository extends JpaRepository<ContentUnitEntity, UUID>, JpaSpecificationExecutor<ContentUnitEntity> {
}
