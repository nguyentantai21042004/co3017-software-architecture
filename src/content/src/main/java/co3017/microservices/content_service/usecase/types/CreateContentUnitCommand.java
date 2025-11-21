package co3017.microservices.content_service.usecase.types;

import co3017.microservices.content_service.models.ContentUnit;
import com.fasterxml.jackson.databind.JsonNode;

import java.util.UUID;

/**
 * Command object cho việc tạo content unit mới
 */
public class CreateContentUnitCommand {
    private final UUID chapterId;
    private final ContentUnit.UnitType unitType;
    private final JsonNode metadataConfig;

    public CreateContentUnitCommand(UUID chapterId, ContentUnit.UnitType unitType, JsonNode metadataConfig) {
        this.chapterId = chapterId;
        this.unitType = unitType;
        this.metadataConfig = metadataConfig;
    }

    public UUID getChapterId() {
        return chapterId;
    }

    public ContentUnit.UnitType getUnitType() {
        return unitType;
    }

    public JsonNode getMetadataConfig() {
        return metadataConfig;
    }
}
