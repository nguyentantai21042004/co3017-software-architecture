package co3017.microservices.content_service.repository.postgresql;

import co3017.microservices.content_service.repository.postgresql.entity.MetadataTagEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

/**
 * JPA Repository Interface cho MetadataTagEntity
 */
@Repository
public interface MetadataTagJpaRepository
        extends JpaRepository<MetadataTagEntity, Integer>, JpaSpecificationExecutor<MetadataTagEntity> {

    boolean existsByTagName(String tagName);

    Optional<MetadataTagEntity> findByTagName(String tagName);

    List<MetadataTagEntity> findByTagIdIn(List<Integer> tagIds);
}
