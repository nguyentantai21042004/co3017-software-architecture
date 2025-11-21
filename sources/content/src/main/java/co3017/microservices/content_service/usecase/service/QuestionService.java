package co3017.microservices.content_service.usecase.service;

import co3017.microservices.content_service.models.Question;
import co3017.microservices.content_service.repository.QuestionRepository;
import co3017.microservices.content_service.usecase.QuestionUseCase;
import co3017.microservices.content_service.usecase.types.CreateQuestionCommand;
import co3017.microservices.content_service.usecase.types.QuestionQuery;
import co3017.microservices.content_service.usecase.types.UpdateQuestionCommand;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.Optional;

/**
 * Question Use Case Implementation
 * Contains business logic for question operations
 * Orchestrates domain models and repository calls
 */
@Service
@Transactional
public class QuestionService implements QuestionUseCase {

    private final QuestionRepository questionRepository;

    public QuestionService(QuestionRepository questionRepository) {
        this.questionRepository = questionRepository;
    }

    @Override
    public Question createQuestion(CreateQuestionCommand command) {
        // Create domain model
        Question question = new Question(
            command.getContent(),
            command.getDifficulty(),
            command.getSkillTag()
        );

        // Validate using domain logic
        if (!question.isValid()) {
            throw new IllegalArgumentException("Invalid question data: all fields are required");
        }

        // Persist through repository
        return questionRepository.save(question);
    }

    @Override
    public Question updateQuestion(UpdateQuestionCommand command) {
        // Fetch existing question
        Question question = questionRepository.findById(command.getId())
            .orElseThrow(() -> new IllegalArgumentException("Question not found with id: " + command.getId()));

        // Update using domain methods
        if (command.getContent() != null) {
            question.updateContent(command.getContent());
        }
        if (command.getDifficulty() != null) {
            question.updateDifficulty(command.getDifficulty());
        }
        if (command.getSkillTag() != null) {
            question.updateSkillTag(command.getSkillTag());
        }

        // Validate and persist
        if (!question.isValid()) {
            throw new IllegalArgumentException("Invalid question data after update");
        }

        return questionRepository.update(question);
    }

    @Override
    @Transactional(readOnly = true)
    public Optional<Question> getQuestionById(Integer id) {
        return questionRepository.findById(id);
    }

    @Override
    @Transactional(readOnly = true)
    public List<Question> getAllQuestions(QuestionQuery query) {
        // Handle different query scenarios
        if (query.hasDifficulty() && query.hasSkillTag()) {
            return questionRepository.findByDifficultyAndSkillTag(
                query.getDifficulty(),
                query.getSkillTag()
            );
        } else if (query.hasDifficulty()) {
            return questionRepository.findByDifficulty(query.getDifficulty());
        } else if (query.hasSkillTag()) {
            return questionRepository.findBySkillTag(query.getSkillTag());
        } else {
            return questionRepository.findAll();
        }
    }

    @Override
    public void deleteQuestion(Integer id) {
        if (!questionRepository.existsById(id)) {
            throw new IllegalArgumentException("Question not found with id: " + id);
        }
        questionRepository.deleteById(id);
    }
}
