package co3017.microservices.content_service.repository;

import co3017.microservices.content_service.models.MetadataTag;
import co3017.microservices.content_service.repository.postgresql.entity.MetadataTagEntity;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.domain.Specification;

import java.util.List;
import java.util.Optional;

/**
 * MetadataTag Repository Interface
 * Repository pattern cho MetadataTag domain
 */
public interface MetadataTagRepository {
    MetadataTag save(MetadataTag metadataTag);
    Optional<MetadataTag> findById(Integer tagId);
    List<MetadataTag> findAllByIds(List<Integer> tagIds);
    Page<MetadataTag> findAll(Specification<MetadataTagEntity> spec, Pageable pageable);
    boolean existsByTagName(String tagName);
    int deleteByIds(List<Integer> tagIds);
    boolean existsById(Integer tagId);
}
