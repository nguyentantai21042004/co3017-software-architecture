package co3017.microservices.content_service.usecase.service;

import co3017.microservices.content_service.models.Question;
import co3017.microservices.content_service.repository.QuestionRepository;
import co3017.microservices.content_service.usecase.types.CreateQuestionCommand;
import co3017.microservices.content_service.usecase.types.QuestionQuery;
import co3017.microservices.content_service.usecase.types.UpdateQuestionCommand;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class QuestionServiceTest {

    @Mock
    private QuestionRepository questionRepository;

    @InjectMocks
    private QuestionService questionService;

    private Question testQuestion;
    private LocalDateTime now;

    @BeforeEach
    void setUp() {
        now = LocalDateTime.now();
        testQuestion = new Question(
                1,
                "What is 2 + 2?",
                1, // difficultyLevel (Integer)
                "math_arithmetic",
                "4", // correctAnswer
                false, // isRemedial
                now);
    }

    // CREATE QUESTION TESTS

    @Test
    void createQuestion_Success() {
        // Arrange
        CreateQuestionCommand command = new CreateQuestionCommand(
                "What is 2 + 2?",
                1, // difficultyLevel
                "math_arithmetic",
                "4", // correctAnswer
                false // isRemedial
        );
        when(questionRepository.save(any(Question.class))).thenReturn(testQuestion);

        // Act
        Question result = questionService.createQuestion(command);

        // Assert
        assertNotNull(result);
        assertEquals("What is 2 + 2?", result.getContent());
        assertEquals(1, result.getDifficultyLevel());
        assertEquals("math_arithmetic", result.getSkillTag());
        assertEquals("4", result.getCorrectAnswer());
        assertEquals(false, result.getIsRemedial());
        verify(questionRepository, times(1)).save(any(Question.class));
    }

    @Test
    void createQuestion_InvalidData_ThrowsException() {
        // Arrange
        CreateQuestionCommand command = new CreateQuestionCommand("", null, "", "", null);

        // Act & Assert
        IllegalArgumentException exception = assertThrows(
                IllegalArgumentException.class,
                () -> questionService.createQuestion(command));
        assertEquals("Invalid question data: all fields are required", exception.getMessage());
        verify(questionRepository, never()).save(any(Question.class));
    }

    @Test
    void createQuestion_NullContent_ThrowsException() {
        // Arrange
        CreateQuestionCommand command = new CreateQuestionCommand(null, 1, "math", "4", false);

        // Act & Assert
        assertThrows(IllegalArgumentException.class, () -> questionService.createQuestion(command));
        verify(questionRepository, never()).save(any(Question.class));
    }

    // UPDATE QUESTION TESTS

    @Test
    void updateQuestion_Success_AllFields() {
        // Arrange
        UpdateQuestionCommand command = new UpdateQuestionCommand(
                1,
                "What is 3 + 3?",
                2, // difficultyLevel
                "math_advanced",
                "6", // correctAnswer
                true // isRemedial
        );
        Question updatedQuestion = new Question(
                1,
                "What is 3 + 3?",
                2,
                "math_advanced",
                "6",
                true,
                now);
        when(questionRepository.findById(1)).thenReturn(Optional.of(testQuestion));
        when(questionRepository.update(any(Question.class))).thenReturn(updatedQuestion);

        // Act
        Question result = questionService.updateQuestion(command);

        // Assert
        assertNotNull(result);
        assertEquals("What is 3 + 3?", result.getContent());
        assertEquals(2, result.getDifficultyLevel());
        assertEquals("math_advanced", result.getSkillTag());
        assertEquals("6", result.getCorrectAnswer());
        assertEquals(true, result.getIsRemedial());
        verify(questionRepository, times(1)).findById(1);
        verify(questionRepository, times(1)).update(any(Question.class));
    }

    @Test
    void updateQuestion_PartialUpdate_ContentOnly() {
        // Arrange
        UpdateQuestionCommand command = new UpdateQuestionCommand(1, "New content", null, null, null, null);
        when(questionRepository.findById(1)).thenReturn(Optional.of(testQuestion));
        when(questionRepository.update(any(Question.class))).thenReturn(testQuestion);

        // Act
        Question result = questionService.updateQuestion(command);

        // Assert
        assertNotNull(result);
        verify(questionRepository, times(1)).findById(1);
        verify(questionRepository, times(1)).update(any(Question.class));
    }

    @Test
    void updateQuestion_QuestionNotFound_ThrowsException() {
        // Arrange
        UpdateQuestionCommand command = new UpdateQuestionCommand(999, "New content", 1, "math", "4", false);
        when(questionRepository.findById(999)).thenReturn(Optional.empty());

        // Act & Assert
        IllegalArgumentException exception = assertThrows(
                IllegalArgumentException.class,
                () -> questionService.updateQuestion(command));
        assertEquals("Question not found with id: 999", exception.getMessage());
        verify(questionRepository, times(1)).findById(999);
        verify(questionRepository, never()).update(any(Question.class));
    }

    @Test
    void updateQuestion_InvalidData_ThrowsException() {
        // Arrange
        UpdateQuestionCommand command = new UpdateQuestionCommand(1, "", null, "", "", null);
        when(questionRepository.findById(1)).thenReturn(Optional.of(testQuestion));

        // Act & Assert
        assertThrows(IllegalArgumentException.class, () -> questionService.updateQuestion(command));
        verify(questionRepository, times(1)).findById(1);
    }

    // GET QUESTION TESTS

    @Test
    void getQuestionById_Success() {
        // Arrange
        when(questionRepository.findById(1)).thenReturn(Optional.of(testQuestion));

        // Act
        Optional<Question> result = questionService.getQuestionById(1);

        // Assert
        assertTrue(result.isPresent());
        assertEquals(testQuestion.getId(), result.get().getId());
        assertEquals(testQuestion.getContent(), result.get().getContent());
        verify(questionRepository, times(1)).findById(1);
    }

    @Test
    void getQuestionById_NotFound() {
        // Arrange
        when(questionRepository.findById(999)).thenReturn(Optional.empty());

        // Act
        Optional<Question> result = questionService.getQuestionById(999);

        // Assert
        assertFalse(result.isPresent());
        verify(questionRepository, times(1)).findById(999);
    }

    // GET ALL QUESTIONS TESTS

    @Test
    void getAllQuestions_NoFilters() {
        // Arrange
        List<Question> questions = Arrays.asList(testQuestion, testQuestion);
        QuestionQuery query = new QuestionQuery(null, null);
        when(questionRepository.findAll()).thenReturn(questions);

        // Act
        List<Question> result = questionService.getAllQuestions(query);

        // Assert
        assertNotNull(result);
        assertEquals(2, result.size());
        verify(questionRepository, times(1)).findAll();
        verify(questionRepository, never()).findByDifficultyLevel(any());
        verify(questionRepository, never()).findBySkillTag(any());
        verify(questionRepository, never()).findByDifficultyLevelAndSkillTag(any(), any());
    }

    @Test
    void getAllQuestions_FilterByDifficultyLevel() {
        // Arrange
        List<Question> questions = Arrays.asList(testQuestion);
        QuestionQuery query = new QuestionQuery(1, null);
        when(questionRepository.findByDifficultyLevel(1)).thenReturn(questions);

        // Act
        List<Question> result = questionService.getAllQuestions(query);

        // Assert
        assertNotNull(result);
        assertEquals(1, result.size());
        verify(questionRepository, times(1)).findByDifficultyLevel(1);
        verify(questionRepository, never()).findAll();
    }

    @Test
    void getAllQuestions_FilterBySkillTag() {
        // Arrange
        List<Question> questions = Arrays.asList(testQuestion);
        QuestionQuery query = new QuestionQuery(null, "math_arithmetic");
        when(questionRepository.findBySkillTag("math_arithmetic")).thenReturn(questions);

        // Act
        List<Question> result = questionService.getAllQuestions(query);

        // Assert
        assertNotNull(result);
        assertEquals(1, result.size());
        verify(questionRepository, times(1)).findBySkillTag("math_arithmetic");
        verify(questionRepository, never()).findAll();
    }

    @Test
    void getAllQuestions_FilterByBothDifficultyLevelAndSkillTag() {
        // Arrange
        List<Question> questions = Arrays.asList(testQuestion);
        QuestionQuery query = new QuestionQuery(1, "math_arithmetic");
        when(questionRepository.findByDifficultyLevelAndSkillTag(1, "math_arithmetic")).thenReturn(questions);

        // Act
        List<Question> result = questionService.getAllQuestions(query);

        // Assert
        assertNotNull(result);
        assertEquals(1, result.size());
        verify(questionRepository, times(1)).findByDifficultyLevelAndSkillTag(1, "math_arithmetic");
        verify(questionRepository, never()).findAll();
    }

    @Test
    void getAllQuestions_EmptyResult() {
        // Arrange
        QuestionQuery query = new QuestionQuery(null, null);
        when(questionRepository.findAll()).thenReturn(Arrays.asList());

        // Act
        List<Question> result = questionService.getAllQuestions(query);

        // Assert
        assertNotNull(result);
        assertTrue(result.isEmpty());
        verify(questionRepository, times(1)).findAll();
    }

    // DELETE QUESTION TESTS

    @Test
    void deleteQuestion_Success() {
        // Arrange
        when(questionRepository.existsById(1)).thenReturn(true);
        doNothing().when(questionRepository).deleteById(1);

        // Act
        questionService.deleteQuestion(1);

        // Assert
        verify(questionRepository, times(1)).existsById(1);
        verify(questionRepository, times(1)).deleteById(1);
    }

    @Test
    void deleteQuestion_NotFound_ThrowsException() {
        // Arrange
        when(questionRepository.existsById(999)).thenReturn(false);

        // Act & Assert
        IllegalArgumentException exception = assertThrows(
                IllegalArgumentException.class,
                () -> questionService.deleteQuestion(999));
        assertEquals("Question not found with id: 999", exception.getMessage());
        verify(questionRepository, times(1)).existsById(999);
        verify(questionRepository, never()).deleteById(any());
    }
}
