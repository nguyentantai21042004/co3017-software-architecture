package co3017.microservices.content_service.repository.postgresql.specification;

import co3017.microservices.content_service.repository.postgresql.entity.MetadataTagEntity;
import org.springframework.data.jpa.domain.Specification;

import java.util.List;

/**
 * MetadataTag Specification cho JPA queries
 */
public class MetadataTagSpecification {
    
    public static Specification<MetadataTagEntity> builder() {
        return Specification.where(null);
    }

    public static Specification<MetadataTagEntity> keyword(String keyword) {
        if (keyword == null || keyword.trim().isEmpty()) {
            return null;
        }
        String searchPattern = "%" + keyword.toLowerCase() + "%";
        return (root, query, cb) -> cb.like(cb.lower(root.get("tagName").as(String.class)), searchPattern);
    }

    public static Specification<MetadataTagEntity> tagIds(List<Integer> tagIds) {
        if (tagIds == null || tagIds.isEmpty()) {
            return null;
        }
        return (root, query, cb) -> root.get("tagId").in(tagIds);
    }

    public static class Builder {
        private Specification<MetadataTagEntity> spec = Specification.where(null);

        public Builder keyword(String keyword) {
            Specification<MetadataTagEntity> keywordSpec = MetadataTagSpecification.keyword(keyword);
            if (keywordSpec != null) {
                spec = spec.and(keywordSpec);
            }
            return this;
        }

        public Builder tagIds(List<Integer> tagIds) {
            Specification<MetadataTagEntity> tagIdsSpec = MetadataTagSpecification.tagIds(tagIds);
            if (tagIdsSpec != null) {
                spec = spec.and(tagIdsSpec);
            }
            return this;
        }

        public Specification<MetadataTagEntity> build() {
            return spec;
        }
    }
}
