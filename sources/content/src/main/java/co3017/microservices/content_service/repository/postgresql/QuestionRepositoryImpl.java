package co3017.microservices.content_service.repository.postgresql;

import co3017.microservices.content_service.models.Question;
import co3017.microservices.content_service.repository.QuestionRepository;
import co3017.microservices.content_service.repository.postgresql.entity.QuestionEntity;
import co3017.microservices.content_service.repository.postgresql.mapper.QuestionMapper;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

/**
 * Question Repository Implementation
 * Infrastructure layer - implements domain repository interface
 * Uses JPA repository and mapper to handle persistence
 */
@Repository
public class QuestionRepositoryImpl implements QuestionRepository {

    private final QuestionJpaRepository jpaRepository;
    private final QuestionMapper mapper;

    public QuestionRepositoryImpl(QuestionJpaRepository jpaRepository, QuestionMapper mapper) {
        this.jpaRepository = jpaRepository;
        this.mapper = mapper;
    }

    @Override
    public Question save(Question question) {
        QuestionEntity entity = mapper.toEntity(question);
        QuestionEntity saved = jpaRepository.save(entity);
        return mapper.toDomain(saved);
    }

    @Override
    public Question update(Question question) {
        QuestionEntity entity = mapper.toEntity(question);
        QuestionEntity updated = jpaRepository.save(entity);
        return mapper.toDomain(updated);
    }

    @Override
    public Optional<Question> findById(Integer id) {
        return jpaRepository.findById(id)
                .map(mapper::toDomain);
    }

    @Override
    public List<Question> findAll() {
        return jpaRepository.findAll().stream()
                .map(mapper::toDomain)
                .collect(Collectors.toList());
    }

    @Override
    public List<Question> findByDifficultyLevel(Integer difficultyLevel) {
        return jpaRepository.findByDifficultyLevel(difficultyLevel).stream()
                .map(mapper::toDomain)
                .collect(Collectors.toList());
    }

    @Override
    public List<Question> findBySkillTag(String skillTag) {
        return jpaRepository.findBySkillTag(skillTag).stream()
                .map(mapper::toDomain)
                .collect(Collectors.toList());
    }

    @Override
    public List<Question> findByDifficultyLevelAndSkillTag(Integer difficultyLevel, String skillTag) {
        return jpaRepository.findByDifficultyLevelAndSkillTag(difficultyLevel, skillTag).stream()
                .map(mapper::toDomain)
                .collect(Collectors.toList());
    }

    @Override
    public List<Question> findBySkillTagAndIsRemedial(String skillTag, Boolean isRemedial) {
        return jpaRepository.findBySkillTagAndIsRemedial(skillTag, isRemedial).stream()
                .map(mapper::toDomain)
                .collect(Collectors.toList());
    }

    @Override
    public void deleteById(Integer id) {
        jpaRepository.deleteById(id);
    }

    @Override
    public boolean existsById(Integer id) {
        return jpaRepository.existsById(id);
    }

    @Override
    public List<String> findDistinctSkillTags() {
        return jpaRepository.findDistinctSkillTags();
    }
}
