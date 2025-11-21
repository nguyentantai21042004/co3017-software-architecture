package co3017.microservices.content_service.repository.postgresql.mapper;

import co3017.microservices.content_service.models.Chapter;
import co3017.microservices.content_service.repository.postgresql.entity.ChapterEntity;

/**
 * Mapper chuyển đổi giữa Domain Chapter và ChapterEntity (JPA)
 */
public class ChapterMapper {

    /**
     * Chuyển từ Domain Chapter sang ChapterEntity để lưu vào database
     */
    public static ChapterEntity toEntity(Chapter chapter) {
        ChapterEntity entity = new ChapterEntity();
        entity.setChapterId(chapter.getChapterId());
        entity.setCourseId(chapter.getCourseId());
        entity.setSequenceNumber(chapter.getSequenceNumber());
        entity.setCreatedAt(chapter.getCreatedAt());
        entity.setUpdatedAt(chapter.getUpdatedAt());
        return entity;
    }

    /**
     * Chuyển từ ChapterEntity sang Domain Chapter sau khi đọc từ database
     */
    public static Chapter toDomain(ChapterEntity entity) {
        return new Chapter(
            entity.getChapterId(),
            entity.getCourseId(),
            entity.getSequenceNumber(),
            entity.getCreatedAt(),
            entity.getUpdatedAt()
        );
    }
}
