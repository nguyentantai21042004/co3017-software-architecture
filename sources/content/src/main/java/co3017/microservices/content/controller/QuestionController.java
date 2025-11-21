package co3017.microservices.content.controller;

import co3017.microservices.content.dto.ApiResponse;
import co3017.microservices.content.dto.QuestionResponse;
import co3017.microservices.content.service.QuestionService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

/**
 * REST Controller for Question/Content endpoints
 * Provides APIs for:
 * 1. Getting questions by ID
 * 2. Recommending questions based on skill and type (for Adaptive Engine)
 */
@RestController
@RequestMapping("/api/content")
@RequiredArgsConstructor
@Slf4j
public class QuestionController {

    private final QuestionService questionService;

    /**
     * Get question by ID
     *
     * GET /api/content/{id}
     *
     * @param id question ID
     * @return question details
     */
    @GetMapping("/{id}")
    public ResponseEntity<ApiResponse<QuestionResponse>> getQuestionById(@PathVariable Long id) {
        log.info("GET /api/content/{} - Fetching question", id);

        return questionService.getQuestionById(id)
                .map(question -> ResponseEntity.ok(ApiResponse.success(question)))
                .orElseGet(() -> ResponseEntity.notFound().build());
    }

    /**
     * Recommend a question based on skill and type
     *
     * GET /api/content/recommend?skill=math_algebra&type=remedial
     *
     * This is the core API used by Adaptive Engine to get appropriate content
     * based on learner's mastery level
     *
     * @param skill      skill tag (e.g., "math_algebra", "math_geometry")
     * @param type       "remedial" for review/practice, "standard" for normal progression
     * @param random     optional: if true, return random question (default: false)
     * @return recommended question
     */
    @GetMapping("/recommend")
    public ResponseEntity<ApiResponse<QuestionResponse>> recommendQuestion(
            @RequestParam String skill,
            @RequestParam(defaultValue = "standard") String type,
            @RequestParam(defaultValue = "false") boolean random
    ) {
        log.info("GET /api/content/recommend - skill: {}, type: {}, random: {}",
                skill, type, random);

        return questionService.recommendQuestion(skill, type, random)
                .map(question -> ResponseEntity.ok(ApiResponse.success(
                        String.format("Recommended %s question for skill: %s", type, skill),
                        question
                )))
                .orElseGet(() -> ResponseEntity.ok(ApiResponse.error(
                        String.format("No %s question found for skill: %s", type, skill)
                )));
    }

    /**
     * Get all questions for a specific skill
     *
     * GET /api/content/skill/{skill}
     *
     * @param skill skill tag
     * @return list of questions
     */
    @GetMapping("/skill/{skill}")
    public ResponseEntity<ApiResponse<List<QuestionResponse>>> getQuestionsBySkill(
            @PathVariable String skill
    ) {
        log.info("GET /api/content/skill/{} - Fetching all questions", skill);

        List<QuestionResponse> questions = questionService.getQuestionsBySkill(skill);
        return ResponseEntity.ok(ApiResponse.success(
                String.format("Found %d questions for skill: %s", questions.size(), skill),
                questions
        ));
    }

    /**
     * Get questions by skill and difficulty
     *
     * GET /api/content/skill/{skill}/difficulty/{difficulty}
     *
     * @param skill      skill tag
     * @param difficulty 1=Easy, 2=Medium, 3=Hard
     * @return list of questions
     */
    @GetMapping("/skill/{skill}/difficulty/{difficulty}")
    public ResponseEntity<ApiResponse<List<QuestionResponse>>> getQuestionsBySkillAndDifficulty(
            @PathVariable String skill,
            @PathVariable Integer difficulty
    ) {
        log.info("GET /api/content/skill/{}/difficulty/{} - Fetching questions",
                skill, difficulty);

        List<QuestionResponse> questions = questionService
                .getQuestionsBySkillAndDifficulty(skill, difficulty);

        return ResponseEntity.ok(ApiResponse.success(
                String.format("Found %d questions", questions.size()),
                questions
        ));
    }
}
