package co3017.microservices.content_service.repository.postgresql;

import co3017.microservices.content_service.repository.postgresql.entity.QuestionEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * Question JPA Repository
 * Spring Data JPA repository interface
 * Infrastructure layer - provides database operations
 */
@Repository
public interface QuestionJpaRepository extends JpaRepository<QuestionEntity, Integer> {

    /**
     * Find questions by difficulty level
     */
    List<QuestionEntity> findByDifficulty(String difficulty);

    /**
     * Find questions by skill tag
     */
    List<QuestionEntity> findBySkillTag(String skillTag);

    /**
     * Find questions by both difficulty and skill tag
     */
    List<QuestionEntity> findByDifficultyAndSkillTag(String difficulty, String skillTag);
}
