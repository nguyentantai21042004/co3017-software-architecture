package co3017.microservices.content_service.models;

import java.time.LocalDateTime;
import java.util.Objects;

/**
 * Question Domain Model
 * Pure business entity with no framework dependencies
 * Represents a question in the learning system
 */
public class Question {
    private Integer id;
    private String content;
    private Integer difficultyLevel;
    private String skillTag;
    private String correctAnswer;
    private Boolean isRemedial;
    private LocalDateTime createdAt;

    // Constructor for creating new questions (no ID)
    public Question(String content, Integer difficultyLevel, String skillTag, String correctAnswer,
            Boolean isRemedial) {
        this.content = content;
        this.difficultyLevel = difficultyLevel;
        this.skillTag = skillTag;
        this.correctAnswer = correctAnswer;
        this.isRemedial = isRemedial != null ? isRemedial : false;
        this.createdAt = LocalDateTime.now();
    }

    // Constructor for existing questions (with ID)
    public Question(Integer id, String content, Integer difficultyLevel, String skillTag, String correctAnswer,
            Boolean isRemedial,
            LocalDateTime createdAt) {
        this.id = id;
        this.content = content;
        this.difficultyLevel = difficultyLevel;
        this.skillTag = skillTag;
        this.correctAnswer = correctAnswer;
        this.isRemedial = isRemedial != null ? isRemedial : false;
        this.createdAt = createdAt;
    }

    // Business logic methods

    /**
     * Validates if the question is complete and valid
     */
    public boolean isValid() {
        return content != null && !content.trim().isEmpty() &&
                difficultyLevel != null &&
                skillTag != null && !skillTag.trim().isEmpty() &&
                correctAnswer != null && !correctAnswer.trim().isEmpty();
    }

    /**
     * Updates the question content
     */
    public void updateContent(String newContent) {
        if (newContent == null || newContent.trim().isEmpty()) {
            throw new IllegalArgumentException("Content cannot be empty");
        }
        this.content = newContent;
    }

    /**
     * Updates the difficulty level
     */
    public void updateDifficultyLevel(Integer newDifficultyLevel) {
        if (newDifficultyLevel == null) {
            throw new IllegalArgumentException("Difficulty level cannot be empty");
        }
        this.difficultyLevel = newDifficultyLevel;
    }

    /**
     * Updates the skill tag
     */
    public void updateSkillTag(String newSkillTag) {
        if (newSkillTag == null || newSkillTag.trim().isEmpty()) {
            throw new IllegalArgumentException("Skill tag cannot be empty");
        }
        this.skillTag = newSkillTag;
    }

    /**
     * Updates the correct answer
     */
    public void updateCorrectAnswer(String newCorrectAnswer) {
        if (newCorrectAnswer == null || newCorrectAnswer.trim().isEmpty()) {
            throw new IllegalArgumentException("Correct answer cannot be empty");
        }
        this.correctAnswer = newCorrectAnswer;
    }

    /**
     * Updates the remedial status
     */
    public void updateIsRemedial(Boolean isRemedial) {
        this.isRemedial = isRemedial != null ? isRemedial : false;
    }

    /**
     * Checks if the question matches a specific difficulty level
     */
    public boolean hasDifficultyLevel(Integer difficultyLevel) {
        return this.difficultyLevel != null && this.difficultyLevel.equals(difficultyLevel);
    }

    /**
     * Checks if the question belongs to a specific skill
     */
    public boolean hasSkillTag(String skillTag) {
        return this.skillTag != null && this.skillTag.equalsIgnoreCase(skillTag);
    }

    // Getters
    public Integer getId() {
        return id;
    }

    public String getContent() {
        return content;
    }

    public Integer getDifficultyLevel() {
        return difficultyLevel;
    }

    public String getSkillTag() {
        return skillTag;
    }

    public String getCorrectAnswer() {
        return correctAnswer;
    }

    public Boolean getIsRemedial() {
        return isRemedial;
    }

    public LocalDateTime getCreatedAt() {
        return createdAt;
    }

    // No setters for immutability - use business methods instead

    @Override
    public boolean equals(Object o) {
        if (this == o)
            return true;
        if (o == null || getClass() != o.getClass())
            return false;
        Question question = (Question) o;
        return Objects.equals(id, question.id);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id);
    }

    @Override
    public String toString() {
        return "Question{" +
                "id=" + id +
                ", content='" + content + '\'' +
                ", difficultyLevel=" + difficultyLevel +
                ", skillTag='" + skillTag + '\'' +
                ", correctAnswer='" + correctAnswer + '\'' +
                ", isRemedial=" + isRemedial +
                ", createdAt=" + createdAt +
                '}';
    }
}
