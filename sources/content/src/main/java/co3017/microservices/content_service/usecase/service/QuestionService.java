package co3017.microservices.content_service.usecase.service;

import co3017.microservices.content_service.models.Question;
import co3017.microservices.content_service.repository.QuestionRepository;
import co3017.microservices.content_service.usecase.QuestionUseCase;
import co3017.microservices.content_service.usecase.types.CreateQuestionCommand;
import co3017.microservices.content_service.usecase.types.QuestionQuery;
import co3017.microservices.content_service.usecase.types.UpdateQuestionCommand;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.client.RestTemplate;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.*;
import java.util.stream.Collectors;

/**
 * Question Use Case Implementation
 * Contains business logic for question operations
 * Orchestrates domain models and repository calls
 */
@Service
@Transactional
public class QuestionService implements QuestionUseCase {

    private static final Logger logger = LoggerFactory.getLogger(QuestionService.class);

    private final QuestionRepository questionRepository;
    private final RestTemplate restTemplate;
    private final Random random;

    @Value("${services.scoring.url:http://localhost:8082}")
    private String scoringServiceUrl;

    public QuestionService(QuestionRepository questionRepository, RestTemplate restTemplate) {
        this.questionRepository = questionRepository;
        this.restTemplate = restTemplate;
        this.random = new Random();
    }

    @Override
    public Question createQuestion(CreateQuestionCommand command) {
        // Create domain model
        Question question = new Question(
                command.getContent(),
                command.getOptions(),
                command.getDifficultyLevel(),
                command.getSkillTag(),
                command.getCorrectAnswer(),
                command.getIsRemedial());

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
        if (command.getDifficultyLevel() != null) {
            question.updateDifficultyLevel(command.getDifficultyLevel());
        }
        if (command.getSkillTag() != null) {
            question.updateSkillTag(command.getSkillTag());
        }
        if (command.getCorrectAnswer() != null) {
            question.updateCorrectAnswer(command.getCorrectAnswer());
        }
        if (command.getIsRemedial() != null) {
            question.updateIsRemedial(command.getIsRemedial());
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
        if (query.hasDifficultyLevel() && query.hasSkillTag()) {
            return questionRepository.findByDifficultyLevelAndSkillTag(
                    query.getDifficultyLevel(),
                    query.getSkillTag());
        } else if (query.hasDifficultyLevel()) {
            return questionRepository.findByDifficultyLevel(query.getDifficultyLevel());
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

    @Override
    @Transactional(readOnly = true)
    public Question recommendQuestion(String skillTag, String type, String userId) {
        boolean isRemedial = "remedial".equalsIgnoreCase(type);
        List<Question> questions = questionRepository.findBySkillTagAndIsRemedial(skillTag, isRemedial);

        if (questions.isEmpty()) {
            throw new IllegalArgumentException("No questions found for skill: " + skillTag + " and type: " + type);
        }

        // If userId is provided, exclude already answered questions
        if (userId != null && !userId.isEmpty()) {
            Set<Integer> answeredQuestionIds = getAnsweredQuestionIds(userId, skillTag);
            questions = questions.stream()
                    .filter(q -> !answeredQuestionIds.contains(q.getId()))
                    .collect(Collectors.toList());

            if (questions.isEmpty()) {
                throw new IllegalArgumentException("No unanswered questions found for skill: " + skillTag + " and type: " + type);
            }
        }

        // Return a random question from the available ones
        return questions.get(random.nextInt(questions.size()));
    }

    /**
     * Fetch answered question IDs from Scoring Service
     */
    private Set<Integer> getAnsweredQuestionIds(String userId, String skillTag) {
        try {
            String url = scoringServiceUrl + "/api/scoring/answered-questions?user_id=" + userId + "&skill=" + skillTag;
            ResponseEntity<Map<String, Object>> response = restTemplate.exchange(
                    url,
                    HttpMethod.GET,
                    null,
                    new ParameterizedTypeReference<Map<String, Object>>() {}
            );

            if (response.getBody() != null && response.getBody().get("data") != null) {
                @SuppressWarnings("unchecked")
                List<Integer> questionIds = (List<Integer>) response.getBody().get("data");
                return new HashSet<>(questionIds);
            }
        } catch (Exception e) {
            // If Scoring Service is unavailable, log and continue without filtering
            logger.error("Failed to fetch answered questions from Scoring Service: {}", e.getMessage());
        }
        return new HashSet<>();
    }

    @Override
    public List<String> getAvailableSkills() {
        return questionRepository.findDistinctSkillTags();
    }
}
