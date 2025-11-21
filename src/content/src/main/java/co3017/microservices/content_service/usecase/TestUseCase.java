package co3017.microservices.content_service.usecase;

import co3017.microservices.content_service.models.Test;
import co3017.microservices.content_service.usecase.types.CreateTestCommand;

import java.util.List;
import java.util.Optional;

/**
 * Test Use Case Interface - Tất cả use cases cho Test domain
 */
public interface TestUseCase {
    // Create
    Test createTest(CreateTestCommand command);
    
    // Query
    Optional<Test> getTestById(Long id);
    List<Test> getAllTests();
    List<Test> getTestsByTitle(String title);
}

