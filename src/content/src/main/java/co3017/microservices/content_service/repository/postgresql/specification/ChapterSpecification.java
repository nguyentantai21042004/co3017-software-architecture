package co3017.microservices.content_service.repository.postgresql.specification;

import co3017.microservices.content_service.repository.postgresql.entity.ChapterEntity;
import co3017.microservices.content_service.usecase.types.ChapterSearchCriteria;
import jakarta.persistence.criteria.Predicate;
import org.springframework.data.jpa.domain.Specification;

import java.util.ArrayList;
import java.util.List;

/**
 * Specification cho dynamic query Chapter
 */
public class ChapterSpecification {

    public static Specification<ChapterEntity> createSpecification(ChapterSearchCriteria criteria) {
        return (root, query, criteriaBuilder) -> {
            List<Predicate> predicates = new ArrayList<>();

            // Search by course ID
            if (criteria.getCourseId() != null) {
                predicates.add(criteriaBuilder.equal(root.get("courseId"), criteria.getCourseId()));
            }

            // Search by sequence number
            if (criteria.getSequenceNumber() != null) {
                predicates.add(criteriaBuilder.equal(root.get("sequenceNumber"), criteria.getSequenceNumber()));
            }

            return criteriaBuilder.and(predicates.toArray(new Predicate[0]));
        };
    }
}
