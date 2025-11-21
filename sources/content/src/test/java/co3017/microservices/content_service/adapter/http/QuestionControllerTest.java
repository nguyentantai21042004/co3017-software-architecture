package co3017.microservices.content_service.adapter.http;

import co3017.microservices.content_service.adapter.http.dto.CreateQuestionRequest;
import co3017.microservices.content_service.adapter.http.dto.UpdateQuestionRequest;
import co3017.microservices.content_service.models.Question;
import co3017.microservices.content_service.usecase.QuestionUseCase;
import co3017.microservices.content_service.usecase.types.CreateQuestionCommand;
import co3017.microservices.content_service.usecase.types.QuestionQuery;
import co3017.microservices.content_service.usecase.types.UpdateQuestionCommand;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.*;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@WebMvcTest(QuestionController.class)
class QuestionControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private ObjectMapper objectMapper;

    @MockBean
    private QuestionUseCase questionUseCase;

    private Question testQuestion;
    private LocalDateTime now;

    @BeforeEach
    void setUp() {
        now = LocalDateTime.now();
        testQuestion = new Question(
            1,
            "What is 2 + 2?",
            "easy",
            "math_arithmetic",
            now,
            now
        );
    }

    // CREATE QUESTION TESTS

    @Test
    void createQuestion_Success() throws Exception {
        // Arrange
        CreateQuestionRequest request = new CreateQuestionRequest(
            "What is 2 + 2?",
            "easy",
            "math_arithmetic"
        );
        when(questionUseCase.createQuestion(any(CreateQuestionCommand.class)))
            .thenReturn(testQuestion);

        // Act & Assert
        mockMvc.perform(post("/api/questions")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(request)))
            .andExpect(status().isCreated())
            .andExpect(jsonPath("$.error_code").value(0))
            .andExpect(jsonPath("$.message").value("Question created successfully"))
            .andExpect(jsonPath("$.data.id").value(1))
            .andExpect(jsonPath("$.data.content").value("What is 2 + 2?"))
            .andExpect(jsonPath("$.data.difficulty").value("easy"))
            .andExpect(jsonPath("$.data.skill_tag").value("math_arithmetic"));

        verify(questionUseCase, times(1)).createQuestion(any(CreateQuestionCommand.class));
    }

    @Test
    void createQuestion_InvalidData_BadRequest() throws Exception {
        // Arrange
        CreateQuestionRequest request = new CreateQuestionRequest("", "", "");
        when(questionUseCase.createQuestion(any(CreateQuestionCommand.class)))
            .thenThrow(new IllegalArgumentException("Invalid question data: all fields are required"));

        // Act & Assert
        mockMvc.perform(post("/api/questions")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(request)))
            .andExpect(status().isBadRequest())
            .andExpect(jsonPath("$.error_code").value(400))
            .andExpect(jsonPath("$.message").value("Invalid question data: all fields are required"));
    }

    // UPDATE QUESTION TESTS

    @Test
    void updateQuestion_Success() throws Exception {
        // Arrange
        UpdateQuestionRequest request = new UpdateQuestionRequest(
            "What is 3 + 3?",
            "medium",
            "math_advanced"
        );
        Question updatedQuestion = new Question(
            1,
            "What is 3 + 3?",
            "medium",
            "math_advanced",
            now,
            LocalDateTime.now()
        );
        when(questionUseCase.updateQuestion(any(UpdateQuestionCommand.class)))
            .thenReturn(updatedQuestion);

        // Act & Assert
        mockMvc.perform(put("/api/questions/1")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(request)))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.error_code").value(0))
            .andExpect(jsonPath("$.message").value("Question updated successfully"))
            .andExpect(jsonPath("$.data.id").value(1))
            .andExpect(jsonPath("$.data.content").value("What is 3 + 3?"))
            .andExpect(jsonPath("$.data.difficulty").value("medium"));

        verify(questionUseCase, times(1)).updateQuestion(any(UpdateQuestionCommand.class));
    }

    @Test
    void updateQuestion_NotFound() throws Exception {
        // Arrange
        UpdateQuestionRequest request = new UpdateQuestionRequest(
            "New content",
            "easy",
            "math"
        );
        when(questionUseCase.updateQuestion(any(UpdateQuestionCommand.class)))
            .thenThrow(new IllegalArgumentException("Question not found with id: 999"));

        // Act & Assert
        mockMvc.perform(put("/api/questions/999")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(request)))
            .andExpect(status().isNotFound())
            .andExpect(jsonPath("$.error_code").value(404))
            .andExpect(jsonPath("$.message").value("Question not found with id: 999"));
    }

    // GET QUESTION BY ID TESTS

    @Test
    void getQuestionById_Success() throws Exception {
        // Arrange
        when(questionUseCase.getQuestionById(1)).thenReturn(Optional.of(testQuestion));

        // Act & Assert
        mockMvc.perform(get("/api/questions/1"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.error_code").value(0))
            .andExpect(jsonPath("$.data.id").value(1))
            .andExpect(jsonPath("$.data.content").value("What is 2 + 2?"))
            .andExpect(jsonPath("$.data.difficulty").value("easy"))
            .andExpect(jsonPath("$.data.skill_tag").value("math_arithmetic"));

        verify(questionUseCase, times(1)).getQuestionById(1);
    }

    @Test
    void getQuestionById_NotFound() throws Exception {
        // Arrange
        when(questionUseCase.getQuestionById(999)).thenReturn(Optional.empty());

        // Act & Assert
        mockMvc.perform(get("/api/questions/999"))
            .andExpect(status().isNotFound())
            .andExpect(jsonPath("$.error_code").value(404))
            .andExpect(jsonPath("$.message").value("Question not found with id: 999"));

        verify(questionUseCase, times(1)).getQuestionById(999);
    }

    // GET ALL QUESTIONS TESTS

    @Test
    void getAllQuestions_NoFilters() throws Exception {
        // Arrange
        List<Question> questions = Arrays.asList(testQuestion);
        when(questionUseCase.getAllQuestions(any(QuestionQuery.class)))
            .thenReturn(questions);

        // Act & Assert
        mockMvc.perform(get("/api/questions"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.error_code").value(0))
            .andExpect(jsonPath("$.data").isArray())
            .andExpect(jsonPath("$.data.length()").value(1))
            .andExpect(jsonPath("$.data[0].id").value(1))
            .andExpect(jsonPath("$.data[0].content").value("What is 2 + 2?"));

        verify(questionUseCase, times(1)).getAllQuestions(any(QuestionQuery.class));
    }

    @Test
    void getAllQuestions_FilterByDifficulty() throws Exception {
        // Arrange
        List<Question> questions = Arrays.asList(testQuestion);
        when(questionUseCase.getAllQuestions(any(QuestionQuery.class)))
            .thenReturn(questions);

        // Act & Assert
        mockMvc.perform(get("/api/questions")
                .param("difficulty", "easy"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.error_code").value(0))
            .andExpect(jsonPath("$.data").isArray())
            .andExpect(jsonPath("$.data.length()").value(1));

        verify(questionUseCase, times(1)).getAllQuestions(any(QuestionQuery.class));
    }

    @Test
    void getAllQuestions_FilterBySkillTag() throws Exception {
        // Arrange
        List<Question> questions = Arrays.asList(testQuestion);
        when(questionUseCase.getAllQuestions(any(QuestionQuery.class)))
            .thenReturn(questions);

        // Act & Assert
        mockMvc.perform(get("/api/questions")
                .param("skillTag", "math_arithmetic"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.error_code").value(0))
            .andExpect(jsonPath("$.data").isArray())
            .andExpect(jsonPath("$.data.length()").value(1));

        verify(questionUseCase, times(1)).getAllQuestions(any(QuestionQuery.class));
    }

    @Test
    void getAllQuestions_FilterByBoth() throws Exception {
        // Arrange
        List<Question> questions = Arrays.asList(testQuestion);
        when(questionUseCase.getAllQuestions(any(QuestionQuery.class)))
            .thenReturn(questions);

        // Act & Assert
        mockMvc.perform(get("/api/questions")
                .param("difficulty", "easy")
                .param("skillTag", "math_arithmetic"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.error_code").value(0))
            .andExpect(jsonPath("$.data").isArray())
            .andExpect(jsonPath("$.data.length()").value(1));

        verify(questionUseCase, times(1)).getAllQuestions(any(QuestionQuery.class));
    }

    @Test
    void getAllQuestions_EmptyResult() throws Exception {
        // Arrange
        when(questionUseCase.getAllQuestions(any(QuestionQuery.class)))
            .thenReturn(Arrays.asList());

        // Act & Assert
        mockMvc.perform(get("/api/questions"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.error_code").value(0))
            .andExpect(jsonPath("$.data").isArray())
            .andExpect(jsonPath("$.data.length()").value(0));
    }

    // DELETE QUESTION TESTS

    @Test
    void deleteQuestion_Success() throws Exception {
        // Arrange
        doNothing().when(questionUseCase).deleteQuestion(1);

        // Act & Assert
        mockMvc.perform(delete("/api/questions/1"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.error_code").value(0))
            .andExpect(jsonPath("$.message").value("Question deleted successfully"));

        verify(questionUseCase, times(1)).deleteQuestion(1);
    }

    @Test
    void deleteQuestion_NotFound() throws Exception {
        // Arrange
        doThrow(new IllegalArgumentException("Question not found with id: 999"))
            .when(questionUseCase).deleteQuestion(999);

        // Act & Assert
        mockMvc.perform(delete("/api/questions/999"))
            .andExpect(status().isNotFound())
            .andExpect(jsonPath("$.error_code").value(404))
            .andExpect(jsonPath("$.message").value("Question not found with id: 999"));

        verify(questionUseCase, times(1)).deleteQuestion(999);
    }
}
