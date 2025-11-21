package co3017.microservices.content_service.repository.postgresql.entity;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;

/**
 * Question JPA Entity
 * Infrastructure layer - maps to database table
 * Contains JPA/Hibernate annotations for ORM
 */
@Entity
@Table(name = "questions")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class QuestionEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    @Column(name = "content", nullable = false, columnDefinition = "TEXT")
    private String content;

    @Column(name = "difficulty", nullable = false, length = 50)
    private String difficulty;

    @Column(name = "skill_tag", nullable = false, length = 100)
    private String skillTag;

    @Column(name = "created_at", nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @Column(name = "updated_at", nullable = false)
    private LocalDateTime updatedAt;

    @PrePersist
    protected void onCreate() {
        createdAt = LocalDateTime.now();
        updatedAt = LocalDateTime.now();
    }

    @PreUpdate
    protected void onUpdate() {
        updatedAt = LocalDateTime.now();
    }
}
