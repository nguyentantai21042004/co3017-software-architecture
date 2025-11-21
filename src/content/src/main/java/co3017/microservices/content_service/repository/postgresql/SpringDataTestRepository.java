package co3017.microservices.content_service.repository.postgresql;

import co3017.microservices.content_service.repository.postgresql.entity.TestEntity;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

/**
 * Spring Data JPA Repository cho Test
 */
public interface SpringDataTestRepository extends JpaRepository<TestEntity, Long> {
    
    List<TestEntity> findByTitleContaining(String title);
    
    boolean existsByTitle(String title);
}

