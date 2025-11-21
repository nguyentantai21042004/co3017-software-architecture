package co3017.microservices.content_service.adapter.http.dto;

import co3017.microservices.content_service.models.Question;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;

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
    private String difficulty;
    private String skillTag;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;

    /**
     * Create response from domain model
     */
    public static QuestionResponse fromDomain(Question question) {
        return new QuestionResponse(
            question.getId(),
            question.getContent(),
            question.getDifficulty(),
            question.getSkillTag(),
            question.getCreatedAt(),
            question.getUpdatedAt()
        );
    }
}
