package co3017.microservices.content_service.adapter.http.dto;

import co3017.microservices.content_service.models.Question;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;
import java.util.List;

/**
 * Response DTO for question data
 * HTTP layer - formats outgoing JSON
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class QuestionResponse {

    private Integer id;
    private String content;
    private List<String> options;
    private Integer difficultyLevel;
    private String skillTag;
    private String correctAnswer;
    private Boolean isRemedial;
    private LocalDateTime createdAt;

    /**
     * Create response from domain model
     */
    public static QuestionResponse fromDomain(Question question) {
        return new QuestionResponse(
                question.getId(),
                question.getContent(),
                question.getOptions(),
                question.getDifficultyLevel(),
                question.getSkillTag(),
                question.getCorrectAnswer(),
                question.getIsRemedial(),
                question.getCreatedAt());
    }
}
