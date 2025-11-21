package co3017.microservices.content_service.adapter.http.response;

import co3017.microservices.content_service.adapter.http.dto.TestResponse;
import co3017.microservices.content_service.models.Test;

import java.util.List;
import java.util.stream.Collectors;

/**
 * Response Builder cho Test
 * Chuyển đổi từ Domain → Response DTO
 */
public class TestResponseBuilder {

    /**
     * Chuyển Domain Test → TestResponse DTO
     */
    public static TestResponse toResponse(Test test) {
        return new TestResponse(
            test.getId(),
            test.getTitle(),
            test.getDescription(),
            test.getDuration(),
            test.getMaxScore(),
            test.getCreatedAt(),
            test.getUpdatedAt()
        );
    }

    /**
     * Chuyển List Domain Test → List TestResponse DTO
     */
    public static List<TestResponse> toResponseList(List<Test> tests) {
        return tests.stream()
            .map(TestResponseBuilder::toResponse)
            .collect(Collectors.toList());
    }
}

