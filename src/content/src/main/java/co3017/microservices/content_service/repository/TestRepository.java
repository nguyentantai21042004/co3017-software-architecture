package co3017.microservices.content_service.repository;

import co3017.microservices.content_service.models.Test;

import java.util.List;
import java.util.Optional;

/**
 * Port Out - Repository interface cho Test
 * Infrastructure layer sẽ implement interface này
 */
public interface TestRepository {
    Test save(Test test);
    Optional<Test> findById(Long id);
    List<Test> findAll();
    List<Test> findByTitleContaining(String title);
    void deleteById(Long id);
    boolean existsByTitle(String title);
}

