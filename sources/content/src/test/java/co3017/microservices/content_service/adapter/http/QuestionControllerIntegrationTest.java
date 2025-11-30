package co3017.microservices.content_service.adapter.http;

import co3017.microservices.content_service.ContentServiceApplication;
import co3017.microservices.content_service.adapter.http.dto.CreateQuestionRequest;
import co3017.microservices.content_service.models.Question;
import co3017.microservices.content_service.repository.QuestionRepository;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.web.client.RestTemplate;

import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import java.util.Optional;

import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.delete;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest(classes = ContentServiceApplication.class)
@AutoConfigureMockMvc
public class QuestionControllerIntegrationTest {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private QuestionRepository questionRepository;

    @MockBean
    private RestTemplate restTemplate;

    @Autowired
    private ObjectMapper objectMapper;

    private Question question1;
    private Question question2;

    @BeforeEach
    void setUp() {
        question1 = new Question(
                1,
                "What is 2 + 2?",
                Arrays.asList("3", "4", "5"),
                1,
                "math_arithmetic",
                "4",
                false,
                LocalDateTime.now());
        question2 = new Question(
                2,
                "What is the capital of France?",
                Arrays.asList("Berlin", "Madrid", "Paris"),
                1,
                "geography",
                "Paris",
                false,
                LocalDateTime.now());
    }

    @Test
    void recommendQuestion_Success_NoUserId() throws Exception {
        when(questionRepository.findBySkillTagAndIsRemedial("math_arithmetic", false))
                .thenReturn(Arrays.asList(question1));

        mockMvc.perform(get("/api/content/recommend")
                        .param("skill", "math_arithmetic")
                        .param("type", "standard")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error_code").value(0))
                .andExpect(jsonPath("$.data.id").value(question1.getId()))
                .andExpect(jsonPath("$.data.content").value(question1.getContent()));
    }

    @Test
    void recommendQuestion_Success_WithUserId_UnansweredQuestion() throws Exception {
        when(questionRepository.findBySkillTagAndIsRemedial("math_arithmetic", false))
                .thenReturn(Arrays.asList(question1, question2)); // Assume question2 is also math_arithmetic for this test

        // Mock response from scoring service: userId has answered question1
        Map<String, Object> scoringResponse = new HashMap<>();
        scoringResponse.put("data", Collections.singletonList(question1.getId()));
        when(restTemplate.exchange(
                anyString(),
                any(org.springframework.http.HttpMethod.class),
                any(),
                any(org.springframework.core.ParameterizedTypeReference.class)))
                .thenReturn(org.springframework.http.ResponseEntity.ok(scoringResponse));

        mockMvc.perform(get("/api/content/recommend")
                        .param("skill", "math_arithmetic")
                        .param("type", "standard")
                        .param("userId", "user123")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error_code").value(0))
                .andExpect(jsonPath("$.data.id").value(question2.getId())); // Should recommend question2 as question1 is answered
    }

    @Test
    void recommendQuestion_NotFound_NoQuestionsForSkill() throws Exception {
        when(questionRepository.findBySkillTagAndIsRemedial("non_existent_skill", false))
                .thenReturn(Collections.emptyList());

        mockMvc.perform(get("/api/content/recommend")
                        .param("skill", "non_existent_skill")
                        .param("type", "standard")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error_code").value(404))
                .andExpect(jsonPath("$.message").value("No questions found for skill: non_existent_skill and type: standard"));
    }

    @Test
    void recommendQuestion_NotFound_AllQuestionsAnswered() throws Exception {
        when(questionRepository.findBySkillTagAndIsRemedial("math_arithmetic", false))
                .thenReturn(Arrays.asList(question1));

        // Mock response from scoring service: userId has answered question1
        Map<String, Object> scoringResponse = new HashMap<>();
        scoringResponse.put("data", Collections.singletonList(question1.getId()));
        when(restTemplate.exchange(
                anyString(),
                any(org.springframework.http.HttpMethod.class),
                any(),
                any(org.springframework.core.ParameterizedTypeReference.class)))
                .thenReturn(org.springframework.http.ResponseEntity.ok(scoringResponse));

        mockMvc.perform(get("/api/content/recommend")
                        .param("skill", "math_arithmetic")
                        .param("type", "standard")
                        .param("userId", "user123")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error_code").value(404))
                .andExpect(jsonPath("$.message").value("No unanswered questions found for skill: math_arithmetic and type: standard"));
    }

    @Test
    void recommendQuestion_ScoringServiceUnavailable_ReturnsUnfilteredQuestion() throws Exception {
        when(questionRepository.findBySkillTagAndIsRemedial("math_arithmetic", false))
                .thenReturn(Arrays.asList(question1, question2)); // Assume two questions available

        // Simulate Scoring Service being unavailable (e.g., connection refused)
        when(restTemplate.exchange(
                anyString(),
                any(org.springframework.http.HttpMethod.class),
                any(),
                any(org.springframework.core.ParameterizedTypeReference.class)))
                .thenThrow(new RuntimeException("Connection refused")); // Simulate connection error

        mockMvc.perform(get("/api/content/recommend")
                        .param("skill", "math_arithmetic")
                        .param("type", "standard")
                        .param("userId", "user123")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error_code").value(0))
                .andExpect(jsonPath("$.data").exists()); // Should return a question without filtering
    }

    @Test
    void createQuestion_Success() throws Exception {
        // Mock the repository save operation
        when(questionRepository.save(any(Question.class))).thenReturn(question1);

        // Create a request body
        CreateQuestionRequest createRequest = new CreateQuestionRequest(
                "What is 2 + 2?",
                Arrays.asList("3", "4", "5"),
                1,
                "math_arithmetic",
                "4",
                false
        );

        mockMvc.perform(post("/api/content")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(createRequest)))
                .andExpect(status().isCreated())
                .andExpect(jsonPath("$.error_code").value(0))
                .andExpect(jsonPath("$.data.id").value(question1.getId()))
                .andExpect(jsonPath("$.data.content").value(question1.getContent()));
    }

    @Test
    void getQuestionById_Success() throws Exception {
        when(questionRepository.findById(1)).thenReturn(Optional.of(question1));

        mockMvc.perform(get("/api/content/1")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error_code").value(0))
                .andExpect(jsonPath("$.data.id").value(question1.getId()))
                .andExpect(jsonPath("$.data.content").value(question1.getContent()));
    }

    @Test
    void getQuestionById_NotFound() throws Exception {
        when(questionRepository.findById(999)).thenReturn(Optional.empty());

        mockMvc.perform(get("/api/content/999")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error_code").value(404))
                .andExpect(jsonPath("$.message").value("Question not found with id: 999"));
    }

    @Test
    void getAllQuestions_NoFilters_Success() throws Exception {
        when(questionRepository.findAll()).thenReturn(Arrays.asList(question1, question2));

        mockMvc.perform(get("/api/content")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error_code").value(0))
                .andExpect(jsonPath("$.data.length()").value(2))
                .andExpect(jsonPath("$.data[0].id").value(question1.getId()))
                .andExpect(jsonPath("$.data[1].id").value(question2.getId()));
    }

    @Test
    void deleteQuestion_Success() throws Exception {
        when(questionRepository.existsById(1)).thenReturn(true);

        mockMvc.perform(delete("/api/content/1")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error_code").value(0))
                .andExpect(jsonPath("$.message").value("Question deleted successfully"));
    }

    @Test
    void deleteQuestion_NotFound() throws Exception {
        when(questionRepository.existsById(999)).thenReturn(false);

        mockMvc.perform(delete("/api/content/999")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error_code").value(404))
                .andExpect(jsonPath("$.message").value("Question not found with id: 999"));
    }
}
