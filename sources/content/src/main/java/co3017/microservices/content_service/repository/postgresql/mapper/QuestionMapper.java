package co3017.microservices.content_service.repository.postgresql.mapper;

import co3017.microservices.content_service.models.Question;
import co3017.microservices.content_service.repository.postgresql.entity.QuestionEntity;
import org.springframework.stereotype.Component;

/**
 * Question Mapper
 * Converts between domain models and JPA entities
 * Infrastructure layer component
 */
@Component
public class QuestionMapper {

    /**
     * Convert domain model to JPA entity
     * 
     * @param question the domain model
     * @return the JPA entity
     */
    public QuestionEntity toEntity(Question question) {
        if (question == null) {
            return null;
        }

        QuestionEntity entity = new QuestionEntity();
        entity.setId(question.getId());
        entity.setContent(question.getContent());
        entity.setOptions(JsonbHelper.toJson(question.getOptions()));
        entity.setDifficultyLevel(question.getDifficultyLevel());
        entity.setSkillTag(question.getSkillTag());
        entity.setCorrectAnswer(question.getCorrectAnswer());
        entity.setIsRemedial(question.getIsRemedial());
        entity.setCreatedAt(question.getCreatedAt());

        return entity;
    }

    /**
     * Convert JPA entity to domain model
     * 
     * @param entity the JPA entity
     * @return the domain model
     */
    public Question toDomain(QuestionEntity entity) {
        if (entity == null) {
            return null;
        }

        return new Question(
                entity.getId(),
                entity.getContent(),
                JsonbHelper.fromJson(entity.getOptions()),
                entity.getDifficultyLevel(),
                entity.getSkillTag(),
                entity.getCorrectAnswer(),
                entity.getIsRemedial(),
                entity.getCreatedAt());
    }
}
