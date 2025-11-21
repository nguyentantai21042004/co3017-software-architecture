package co3017.microservices.content.service;

import co3017.microservices.content.dto.QuestionResponse;
import co3017.microservices.content.entity.Question;
import co3017.microservices.content.repository.QuestionRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.Optional;

/**
 * Service layer for Question operations
 * Implements business logic for content recommendation
 */
@Service
@RequiredArgsConstructor
@Slf4j
@Transactional(readOnly = true)
public class QuestionService {

    private final QuestionRepository questionRepository;

    /**
     * Get question by ID
     */
    public Optional<QuestionResponse> getQuestionById(Long id) {
        log.debug("Fetching question with id: {}", id);
        return questionRepository.findById(id)
                .map(this::mapToResponse);
    }

    /**
     * Recommend a question based on skill and type
     * This is the core method used by Adaptive Engine
     *
     * @param skill      skill tag (e.g., "math_algebra")
     * @param type       "remedial" or "standard"
     * @param useRandom  if true, return random question; if false, return first
     * @return recommended question
     */
    public Optional<QuestionResponse> recommendQuestion(String skill, String type, boolean useRandom) {
        boolean isRemedial = "remedial".equalsIgnoreCase(type);

        log.info("Recommending question for skill: {}, type: {}, random: {}",
                skill, type, useRandom);

        Optional<Question> question;
        if (useRandom) {
            question = questionRepository.findRandomBySkillTagAndIsRemedial(skill, isRemedial);
        } else {
            question = questionRepository.findFirstBySkillTagAndIsRemedial(skill, isRemedial);
        }

        if (question.isEmpty()) {
            log.warn("No question found for skill: {}, isRemedial: {}", skill, isRemedial);
        }

        return question.map(this::mapToResponse);
    }

    /**
     * Get all questions for a skill
     */
    public List<QuestionResponse> getQuestionsBySkill(String skill) {
        log.debug("Fetching all questions for skill: {}", skill);
        return questionRepository.findBySkillTag(skill).stream()
                .map(this::mapToResponse)
                .toList();
    }

    /**
     * Get questions by skill and difficulty
     */
    public List<QuestionResponse> getQuestionsBySkillAndDifficulty(String skill, Integer difficulty) {
        log.debug("Fetching questions for skill: {}, difficulty: {}", skill, difficulty);
        return questionRepository.findBySkillTagAndDifficultyLevel(skill, difficulty).stream()
                .map(this::mapToResponse)
                .toList();
    }

    /**
     * Map Question entity to DTO
     */
    private QuestionResponse mapToResponse(Question question) {
        return QuestionResponse.builder()
                .id(question.getId())
                .content(question.getContent())
                .options(question.getOptions())
                .correctAnswer(question.getCorrectAnswer())
                .skillTag(question.getSkillTag())
                .difficultyLevel(question.getDifficultyLevel())
                .isRemedial(question.getIsRemedial())
                .createdAt(question.getCreatedAt())
                .build();
    }
}
