package co3017.microservices.content.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;
import java.util.List;

/**
 * DTO for Question responses
 * Used in REST API responses (snake_case as per Jackson config)
 */
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class QuestionResponse {

    private Long id;

    private String content;

    private List<String> options;

    @JsonProperty("correct_answer")
    private String correctAnswer;

    @JsonProperty("skill_tag")
    private String skillTag;

    @JsonProperty("difficulty_level")
    private Integer difficultyLevel;

    @JsonProperty("is_remedial")
    private Boolean isRemedial;

    @JsonProperty("created_at")
    private LocalDateTime createdAt;
}
