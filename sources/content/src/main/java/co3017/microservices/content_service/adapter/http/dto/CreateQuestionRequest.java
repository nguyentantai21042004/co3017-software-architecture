package co3017.microservices.content_service.adapter.http.dto;

import jakarta.validation.constraints.NotBlank;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * Request DTO for creating a question
 * HTTP layer - handles incoming JSON
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class CreateQuestionRequest {

    @NotBlank(message = "Content is required")
    private String content;

    private Integer difficultyLevel;

    @NotBlank(message = "Skill tag is required")
    private String skillTag;

    @NotBlank(message = "Correct answer is required")
    private String correctAnswer;

    private Boolean isRemedial;
}
