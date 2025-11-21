package co3017.microservices.content_service.adapter.http.response;

import co3017.microservices.content_service.adapter.http.dto.ContentUnitResponse;
import co3017.microservices.content_service.models.ContentUnit;
import org.springframework.stereotype.Component;

/**
 * ContentUnit Response Builder
 */
@Component
public class ContentUnitResponseBuilder {
    
    public ContentUnitResponse toResponse(ContentUnit domain) {
        if (domain == null) {
            return null;
        }
        
        return new ContentUnitResponse(
            domain.getUnitId(),
            domain.getChapterId(),
            domain.getUnitType(),
            domain.getMetadataConfig(),
            domain.getCreatedAt(),
            domain.getUpdatedAt()
        );
    }
}
