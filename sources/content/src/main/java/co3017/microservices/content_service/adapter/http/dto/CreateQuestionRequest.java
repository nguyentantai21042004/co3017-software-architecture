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

    @NotBlank(message = "Difficulty is required")
    private String difficulty;

    @NotBlank(message = "Skill tag is required")
    private String skillTag;
}
