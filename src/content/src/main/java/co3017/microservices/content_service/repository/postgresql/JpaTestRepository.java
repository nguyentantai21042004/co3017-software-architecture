package co3017.microservices.content_service.repository.postgresql;

import co3017.microservices.content_service.repository.TestRepository;
import co3017.microservices.content_service.models.Test;
import co3017.microservices.content_service.repository.postgresql.entity.TestEntity;
import co3017.microservices.content_service.repository.postgresql.mapper.TestMapper;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

/**
 * Infrastructure Repository Implementation cho Test
 * Adapter cho persistence layer, implement Port Out interface
 */
@Repository
public class JpaTestRepository implements TestRepository {

    private final SpringDataTestRepository springDataTestRepository;

    public JpaTestRepository(SpringDataTestRepository springDataTestRepository) {
        this.springDataTestRepository = springDataTestRepository;
    }

    @Override
    public Test save(Test test) {
        TestEntity entity = TestMapper.toEntity(test);
        TestEntity savedEntity = springDataTestRepository.save(entity);
        return TestMapper.toDomain(savedEntity);
    }

    @Override
    public Optional<Test> findById(Long id) {
        return springDataTestRepository.findById(id)
            .map(TestMapper::toDomain);
    }

    @Override
    public List<Test> findAll() {
        return springDataTestRepository.findAll().stream()
            .map(TestMapper::toDomain)
            .collect(Collectors.toList());
    }

    @Override
    public List<Test> findByTitleContaining(String title) {
        return springDataTestRepository.findByTitleContaining(title).stream()
            .map(TestMapper::toDomain)
            .collect(Collectors.toList());
    }

    @Override
    public void deleteById(Long id) {
        springDataTestRepository.deleteById(id);
    }

    @Override
    public boolean existsByTitle(String title) {
        return springDataTestRepository.existsByTitle(title);
    }
}

