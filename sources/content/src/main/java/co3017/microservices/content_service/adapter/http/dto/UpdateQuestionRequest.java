package co3017.microservices.content_service.adapter.http.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * Request DTO for updating a question
 * HTTP layer - handles incoming JSON
 * All fields are optional for partial updates
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class UpdateQuestionRequest {

    private String content;
    private String difficulty;
    private String skillTag;
}
