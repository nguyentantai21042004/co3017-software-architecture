package co3017.microservices.content_service.usecase;

import co3017.microservices.content_service.models.Question;
import co3017.microservices.content_service.usecase.types.CreateQuestionCommand;
import co3017.microservices.content_service.usecase.types.QuestionQuery;
import co3017.microservices.content_service.usecase.types.UpdateQuestionCommand;

import java.util.List;
import java.util.Optional;

/**
 * Question Use Case Interface
 * Defines all business operations for Question domain
 * Application layer interface that adapters depend on
 */
public interface QuestionUseCase {

    /**
     * Create a new question
     * @param command the creation command with question data
     * @return the created question with generated ID
     */
    Question createQuestion(CreateQuestionCommand command);

    /**
     * Update an existing question
     * @param command the update command with question data
     * @return the updated question
     * @throws IllegalArgumentException if question not found
     */
    Question updateQuestion(UpdateQuestionCommand command);

    /**
     * Get a question by ID
     * @param id the question ID
     * @return Optional containing the question if found
     */
    Optional<Question> getQuestionById(Integer id);

    /**
     * Get all questions with optional filtering
     * @param query the query parameters for filtering
     * @return list of questions matching the criteria
     */
    List<Question> getAllQuestions(QuestionQuery query);

    /**
     * Delete a question by ID
     * @param id the question ID
     * @throws IllegalArgumentException if question not found
     */
    void deleteQuestion(Integer id);
}
