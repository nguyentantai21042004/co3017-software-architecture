package co3017.microservices.content_service.adapter.http.response;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * Standard API response wrapper
 * Matches the structure used across all microservices
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
@JsonInclude(JsonInclude.Include.NON_NULL)
public class ApiResponse<T> {

    @JsonProperty("error_code")
    private int errorCode;

    @JsonProperty("message")
    private String message;

    @JsonProperty("data")
    private T data;

    @JsonProperty("errors")
    private Object errors;

    // Success responses (error_code = 0)
    public static <T> ApiResponse<T> success(T data) {
        return new ApiResponse<>(0, "Success", data, null);
    }

    public static <T> ApiResponse<T> success(String message, T data) {
        return new ApiResponse<>(0, message, data, null);
    }

    // Error responses (error_code > 0)
    public static <T> ApiResponse<T> error(int errorCode, String message) {
        return new ApiResponse<>(errorCode, message, null, null);
    }

    public static <T> ApiResponse<T> error(int errorCode, String message, Object errors) {
        return new ApiResponse<>(errorCode, message, null, errors);
    }

    // Convenience methods for common HTTP status codes
    public static <T> ApiResponse<T> badRequest(String message) {
        return error(400, message);
    }

    public static <T> ApiResponse<T> notFound(String message) {
        return error(404, message);
    }

    public static <T> ApiResponse<T> internalError(String message) {
        return error(500, message);
    }
}
