package co3017.microservices.content_service.usecase.types;

import co3017.microservices.content_service.models.ContentUnit;
import com.fasterxml.jackson.databind.JsonNode;

/**
 * Command object cho việc cập nhật content unit
 */
public class UpdateContentUnitCommand {
    private final ContentUnit.UnitType unitType;
    private final JsonNode metadataConfig;

    public UpdateContentUnitCommand(ContentUnit.UnitType unitType, JsonNode metadataConfig) {
        this.unitType = unitType;
        this.metadataConfig = metadataConfig;
    }

    public ContentUnit.UnitType getUnitType() {
        return unitType;
    }

    public JsonNode getMetadataConfig() {
        return metadataConfig;
    }
}
