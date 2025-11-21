package co3017.microservices.content.repository;

import co3017.microservices.content.entity.Question;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

/**
 * Repository for Question entity
 * Provides methods to query questions based on skill and type
 */
@Repository
public interface QuestionRepository extends JpaRepository<Question, Long> {

    /**
     * Find a question by skill tag and remedial status
     * Used by Adaptive Engine to recommend appropriate content
     *
     * @param skillTag   the skill tag (e.g., "math_algebra")
     * @param isRemedial true for remedial questions, false for standard
     * @return first matching question
     */
    Optional<Question> findFirstBySkillTagAndIsRemedial(String skillTag, Boolean isRemedial);

    /**
     * Find all questions for a specific skill
     *
     * @param skillTag the skill tag
     * @return list of questions
     */
    List<Question> findBySkillTag(String skillTag);

    /**
     * Find questions by skill and difficulty level
     *
     * @param skillTag        the skill tag
     * @param difficultyLevel difficulty (1=Easy, 2=Medium, 3=Hard)
     * @return list of matching questions
     */
    List<Question> findBySkillTagAndDifficultyLevel(String skillTag, Integer difficultyLevel);

    /**
     * Get random question for a skill and type
     * Useful for generating diverse test sets
     */
    @Query(value = "SELECT * FROM questions WHERE skill_tag = :skillTag AND is_remedial = :isRemedial ORDER BY RANDOM() LIMIT 1", nativeQuery = true)
    Optional<Question> findRandomBySkillTagAndIsRemedial(
            @Param("skillTag") String skillTag,
            @Param("isRemedial") Boolean isRemedial
    );
}
