package co3017.microservices.content_service.repository;

import co3017.microservices.content_service.models.Question;

import java.util.List;
import java.util.Optional;

/**
 * Question Repository Interface
 * Defines data persistence operations for Question domain
 * This is a domain interface that infrastructure layer implements
 */
public interface QuestionRepository {

    /**
     * Save a new question
     * 
     * @param question the question to save
     * @return the saved question with generated ID
     */
    Question save(Question question);

    /**
     * Update an existing question
     * 
     * @param question the question to update
     * @return the updated question
     */
    Question update(Question question);

    /**
     * Find a question by ID
     * 
     * @param id the question ID
     * @return Optional containing the question if found
     */
    Optional<Question> findById(Integer id);

    /**
     * Find all questions
     * 
     * @return list of all questions
     */
    List<Question> findAll();

    /**
     * Find questions by difficulty level
     * 
     * @param difficultyLevel the difficulty level
     * @return list of questions with the specified difficulty
     */
    List<Question> findByDifficultyLevel(Integer difficultyLevel);

    /**
     * Find questions by skill tag
     * 
     * @param skillTag the skill tag
     * @return list of questions with the specified skill tag
     */
    List<Question> findBySkillTag(String skillTag);

    /**
     * Find questions by both difficulty level and skill tag
     * 
     * @param difficultyLevel the difficulty level
     * @param skillTag        the skill tag
     * @return list of questions matching both criteria
     */
    List<Question> findByDifficultyLevelAndSkillTag(Integer difficultyLevel, String skillTag);

    /**
     * Find questions by skill tag and remedial status
     * 
     * @param skillTag   the skill tag
     * @param isRemedial the remedial status
     * @return list of questions matching both criteria
     */
    List<Question> findBySkillTagAndIsRemedial(String skillTag, Boolean isRemedial);

    /**
     * Delete a question by ID
     * 
     * @param id the question ID
     */
    void deleteById(Integer id);

    /**
     * Check if a question exists by ID
     *
     * @param id the question ID
     * @return true if exists, false otherwise
     */
    boolean existsById(Integer id);

    /**
     * Get all distinct skill tags
     *
     * @return list of distinct skill tags
     */
    List<String> findDistinctSkillTags();
}
