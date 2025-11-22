package co3017.microservices.content_service.repository.postgresql.mapper;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.util.List;

/**
 * Helper class for JSONB conversion
 * Handles serialization/deserialization of List<String> to/from JSON
 */
public class JsonbHelper {

    private static final ObjectMapper objectMapper = new ObjectMapper();

    /**
     * Convert List<String> to JSON string
     */
    public static String toJson(List<String> list) {
        if (list == null) {
            return null;
        }
        try {
            return objectMapper.writeValueAsString(list);
        } catch (JsonProcessingException e) {
            throw new RuntimeException("Failed to convert list to JSON", e);
        }
    }

    /**
     * Convert JSON string to List<String>
     */
    public static List<String> fromJson(String json) {
        if (json == null || json.trim().isEmpty()) {
            return null;
        }
        try {
            return objectMapper.readValue(json, new TypeReference<List<String>>() {});
        } catch (JsonProcessingException e) {
            throw new RuntimeException("Failed to parse JSON to list", e);
        }
    }
}
