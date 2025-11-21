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
    private String difficulty;
    private String skillTag;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;

    // Constructor for creating new questions (no ID)
    public Question(String content, String difficulty, String skillTag) {
        this.content = content;
        this.difficulty = difficulty;
        this.skillTag = skillTag;
        this.createdAt = LocalDateTime.now();
        this.updatedAt = LocalDateTime.now();
    }

    // Constructor for existing questions (with ID)
    public Question(Integer id, String content, String difficulty, String skillTag,
                   LocalDateTime createdAt, LocalDateTime updatedAt) {
        this.id = id;
        this.content = content;
        this.difficulty = difficulty;
        this.skillTag = skillTag;
        this.createdAt = createdAt;
        this.updatedAt = updatedAt;
    }

    // Business logic methods

    /**
     * Validates if the question is complete and valid
     */
    public boolean isValid() {
        return content != null && !content.trim().isEmpty() &&
               difficulty != null && !difficulty.trim().isEmpty() &&
               skillTag != null && !skillTag.trim().isEmpty();
    }

    /**
     * Updates the question content and marks as updated
     */
    public void updateContent(String newContent) {
        if (newContent == null || newContent.trim().isEmpty()) {
            throw new IllegalArgumentException("Content cannot be empty");
        }
        this.content = newContent;
        this.updatedAt = LocalDateTime.now();
    }

    /**
     * Updates the difficulty level
     */
    public void updateDifficulty(String newDifficulty) {
        if (newDifficulty == null || newDifficulty.trim().isEmpty()) {
            throw new IllegalArgumentException("Difficulty cannot be empty");
        }
        this.difficulty = newDifficulty;
        this.updatedAt = LocalDateTime.now();
    }

    /**
     * Updates the skill tag
     */
    public void updateSkillTag(String newSkillTag) {
        if (newSkillTag == null || newSkillTag.trim().isEmpty()) {
            throw new IllegalArgumentException("Skill tag cannot be empty");
        }
        this.skillTag = newSkillTag;
        this.updatedAt = LocalDateTime.now();
    }

    /**
     * Checks if the question matches a specific difficulty level
     */
    public boolean hasDifficulty(String difficulty) {
        return this.difficulty != null && this.difficulty.equalsIgnoreCase(difficulty);
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

    public String getDifficulty() {
        return difficulty;
    }

    public String getSkillTag() {
        return skillTag;
    }

    public LocalDateTime getCreatedAt() {
        return createdAt;
    }

    public LocalDateTime getUpdatedAt() {
        return updatedAt;
    }

    // No setters for immutability - use business methods instead

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
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
                ", difficulty='" + difficulty + '\'' +
                ", skillTag='" + skillTag + '\'' +
                ", createdAt=" + createdAt +
                ", updatedAt=" + updatedAt +
                '}';
    }
}
