package co3017.microservices.content_service.adapter.http;

import co3017.microservices.content_service.adapter.http.dto.CreateQuestionRequest;
import co3017.microservices.content_service.adapter.http.dto.QuestionResponse;
import co3017.microservices.content_service.adapter.http.dto.UpdateQuestionRequest;
import co3017.microservices.content_service.adapter.http.response.ApiResponse;
import co3017.microservices.content_service.models.Question;
import co3017.microservices.content_service.usecase.QuestionUseCase;
import co3017.microservices.content_service.usecase.types.CreateQuestionCommand;
import co3017.microservices.content_service.usecase.types.QuestionQuery;
import co3017.microservices.content_service.usecase.types.UpdateQuestionCommand;
import jakarta.validation.Valid;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.stream.Collectors;

/**
 * Question REST Controller
 * HTTP Adapter - handles HTTP requests and responses
 * Depends on use case interface, not implementation
 */
@RestController
@RequestMapping("/api/content")
public class QuestionController {

    private final QuestionUseCase questionUseCase;

    public QuestionController(QuestionUseCase questionUseCase) {
        this.questionUseCase = questionUseCase;
    }

    /**
     * Create a new question
     */
    @PostMapping
    public ResponseEntity<ApiResponse<QuestionResponse>> createQuestion(
            @Valid @RequestBody CreateQuestionRequest request) {
        try {
            CreateQuestionCommand command = new CreateQuestionCommand(
                    request.getContent(),
                    request.getOptions(),
                    request.getDifficultyLevel(),
                    request.getSkillTag(),
                    request.getCorrectAnswer(),
                    request.getIsRemedial());
            Question question = questionUseCase.createQuestion(command);
            QuestionResponse response = QuestionResponse.fromDomain(question);

            return ResponseEntity
                    .status(HttpStatus.CREATED)
                    .body(ApiResponse.success("Question created successfully", response));
        } catch (IllegalArgumentException e) {
            return ResponseEntity
                    .status(HttpStatus.BAD_REQUEST)
                    .body(ApiResponse.badRequest(e.getMessage()));
        }
    }

    /**
     * Update an existing question
     */
    @PutMapping("/{id}")
    public ResponseEntity<ApiResponse<QuestionResponse>> updateQuestion(
            @PathVariable Integer id,
            @Valid @RequestBody UpdateQuestionRequest request) {
        try {
            UpdateQuestionCommand command = new UpdateQuestionCommand(
                    id,
                    request.getContent(),
                    request.getOptions(),
                    request.getDifficultyLevel(),
                    request.getSkillTag(),
                    request.getCorrectAnswer(),
                    request.getIsRemedial());
            Question question = questionUseCase.updateQuestion(command);
            QuestionResponse response = QuestionResponse.fromDomain(question);

            return ResponseEntity.ok(ApiResponse.success("Question updated successfully", response));
        } catch (IllegalArgumentException e) {
            return ResponseEntity
                    .status(HttpStatus.NOT_FOUND)
                    .body(ApiResponse.notFound(e.getMessage()));
        }
    }

    /**
     * Get a question by ID
     */
    @GetMapping("/{id}")
    public ResponseEntity<ApiResponse<QuestionResponse>> getQuestionById(@PathVariable Integer id) {
        return questionUseCase.getQuestionById(id)
                .map(question -> ResponseEntity.ok(
                        ApiResponse.success(QuestionResponse.fromDomain(question))))
                .orElse(ResponseEntity
                        .status(HttpStatus.NOT_FOUND)
                        .body(ApiResponse.notFound("Question not found with id: " + id)));
    }

    /**
     * Get all questions with optional filtering
     */
    @GetMapping
    public ResponseEntity<ApiResponse<List<QuestionResponse>>> getAllQuestions(
            @RequestParam(required = false) Integer difficultyLevel,
            @RequestParam(required = false) String skillTag) {

        QuestionQuery query = new QuestionQuery(difficultyLevel, skillTag);
        List<Question> questions = questionUseCase.getAllQuestions(query);

        List<QuestionResponse> responses = questions.stream()
                .map(QuestionResponse::fromDomain)
                .collect(Collectors.toList());

        return ResponseEntity.ok(ApiResponse.success(responses));
    }

    /**
     * Recommend a question
     */
    @GetMapping("/recommend")
    public ResponseEntity<ApiResponse<QuestionResponse>> recommendQuestion(
            @RequestParam String skill,
            @RequestParam String type) {
        try {
            Question question = questionUseCase.recommendQuestion(skill, type);
            return ResponseEntity.ok(ApiResponse.success(QuestionResponse.fromDomain(question)));
        } catch (IllegalArgumentException e) {
            return ResponseEntity
                    .status(HttpStatus.NOT_FOUND)
                    .body(ApiResponse.notFound(e.getMessage()));
        }
    }

    /**
     * Delete a question
     */
    @DeleteMapping("/{id}")
    public ResponseEntity<ApiResponse<Void>> deleteQuestion(@PathVariable Integer id) {
        try {
            questionUseCase.deleteQuestion(id);
            return ResponseEntity.ok(ApiResponse.success("Question deleted successfully", null));
        } catch (IllegalArgumentException e) {
            return ResponseEntity
                    .status(HttpStatus.NOT_FOUND)
                    .body(ApiResponse.notFound(e.getMessage()));
        }
    }
}
