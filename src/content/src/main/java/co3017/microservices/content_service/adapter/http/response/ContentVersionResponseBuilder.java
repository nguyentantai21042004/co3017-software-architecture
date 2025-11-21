package co3017.microservices.content_service.adapter.http.response;

import co3017.microservices.content_service.adapter.http.dto.ContentVersionResponse;
import co3017.microservices.content_service.models.ContentVersion;
import org.springframework.stereotype.Component;

/**
 * ContentVersion Response Builder
 */
@Component
public class ContentVersionResponseBuilder {
    
    public ContentVersionResponse toResponse(ContentVersion domain) {
        if (domain == null) {
            return null;
        }
        
        return new ContentVersionResponse(
            domain.getVersionId(),
            domain.getUnitId(),
            domain.getVersionNumber(),
            domain.getContentData(),
            domain.isActive(),
            domain.getCreatedAt(),
            domain.getUpdatedAt()
        );
    }
}
