package co3017.microservices.content_service.mappers;

import co3017.microservices.content_service.models.ContentVersion;
import co3017.microservices.content_service.repository.postgresql.entity.ContentVersionEntity;
import org.springframework.stereotype.Component;

/**
 * ContentVersion Mapper - Chuyển đổi giữa Domain và Entity
 */
@Component
public class ContentVersionMapper {
    
    /**
     * Chuyển từ Domain sang Entity
     */
    public ContentVersionEntity toEntity(ContentVersion domain) {
        if (domain == null) {
            return null;
        }
        
        return new ContentVersionEntity(
            domain.getVersionId(),
            domain.getUnitId(),
            domain.getVersionNumber(),
            domain.getContentData(),
            domain.isActive(),
            domain.getCreatedAt(),
            domain.getUpdatedAt()
        );
    }
    
    /**
     * Chuyển từ Entity sang Domain
     */
    public ContentVersion toDomain(ContentVersionEntity entity) {
        if (entity == null) {
            return null;
        }
        
        return new ContentVersion(
            entity.getVersionId(),
            entity.getUnitId(),
            entity.getVersionNumber(),
            entity.getContentData(),
            entity.isActive(),
            entity.getCreatedAt(),
            entity.getUpdatedAt()
        );
    }
}
