import axios from "axios"
import { delay } from "@/lib/utils"
import type {
  QuestionResponse,
  ContentApiResponse,
  SkillsApiResponse,
  MasteryApiResponse,
  NextLessonApiResponse,
  SubmitAnswerApiResponse,
  AnsweredQuestionsApiResponse,
} from "@/types/api"

// Mock mode flag - set to true to simulate backend
const USE_MOCK_DATA = false

const API_URLS = {
  content: "http://localhost:8081",
  scoring: "http://localhost:8082",
  learner: "http://localhost:8083",
  adaptive: "http://localhost:8084/api/adaptive",
}

const apiClient = axios.create({
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
})

// Add response interceptor for better error handling
apiClient.interceptors.response.use(
  (response) => {
    // Check for error_code in response data (backend services use this pattern)
    if (response.data?.error_code && response.data.error_code !== 0) {
      const errorMessage = response.data.message || "An error occurred"
      console.error("API Error Response:", errorMessage)
      return Promise.reject({
        response: {
          ...response,
          data: { ...response.data, message: errorMessage },
        },
        userMessage: errorMessage,
      })
    }
    return response
  },
  (error) => {
    let message = "An unexpected error occurred"
    if (error.response) {
      // Server responded with a status code out of 2xx range
      const errorData = error.response.data
      message =
        errorData?.message ||
        `Server Error: ${error.response.status} - ${error.message}` ||
        "An error occurred on the server"
    } else if (error.request) {
      // Request was made but no response received
      message = "Network Error: Unable to reach the server. Please check your connection or if the services are running."
    } else {
      // Something happened in setting up the request
      message = `Request Error: ${error.message}`
    }
    console.error("API Error:", message, error)
    // Reject with enhanced error object
    return Promise.reject({ ...error, userMessage: message })
  }
)

// Mock Data Generators
const mockMastery = (skill: string) => ({
  error_code: 0,
  message: "Success",
  data: {
    user_id: "student-123",
    skill_tag: skill,
    mastery_score: Math.floor(Math.random() * 100),
    last_updated: new Date().toISOString(),
  },
})

const mockNextLesson = (skill: string, currentScore: number): NextLessonApiResponse => ({
  error_code: 0,
  message: "Success",
  data: {
    next_lesson_id: Math.floor(Math.random() * 1000),
    reason: `Great! Your mastery is ${currentScore}%. Continue with the next challenge.`,
    mastery_score: currentScore,
    content_type: (currentScore < 50 ? "remedial" : "standard") as "remedial" | "standard",
  },
})

const mockQuestion = (id: number, type: string): ContentApiResponse => ({
  error_code: 0,
  message: "Success",
  data: {
    id: id,
    content:
      type === "remedial"
        ? `(Remedial) Basic concept check: What is the foundation of ${id}?`
        : `(Standard) Advanced problem solving: Analyze the case of ${id}.`,
    options: ["A", "B", "C", "D"],
    difficulty_level: type === "remedial" ? 1 : 3,
    skill_tag: "general",
    correct_answer: "A",
    is_remedial: type === "remedial",
  },
})

const mockSubmit = (isCorrect: boolean) => ({
  error_code: 0,
  data: {
    correct: isCorrect,
    score: isCorrect ? 100 : 0,
    feedback: isCorrect ? "Correct! Well done." : "Incorrect. Let's review this concept.",
  },
})

// API Functions with proper TypeScript types
export const api = {
  /**
   * Get all available skills from Content Service
   * @returns List of skill tags
   */
  getAvailableSkills: async (): Promise<{ data: SkillsApiResponse }> => {
    if (USE_MOCK_DATA) {
      await delay(300)
      return { data: { error_code: 0, message: "Success", data: ["math_algebra", "science_physics", "history_world"] } }
    }
    return apiClient.get<SkillsApiResponse>(`${API_URLS.content}/api/content/skills`)
  },

  /**
   * Get mastery score for a user and skill from Learner Model Service
   * @param userId - User identifier
   * @param skill - Skill tag (e.g., "math_algebra")
   * @returns Mastery data including score and last updated timestamp
   */
  getMastery: async (userId: string, skill: string): Promise<{ data: MasteryApiResponse }> => {
    if (USE_MOCK_DATA) {
      await delay(500)
      // Return a stable random score based on skill length for consistency in demo
      const mockScore = skill === "math" ? 45 : skill === "science" ? 75 : 85
      return {
        data: {
          error_code: 0,
          message: "Success",
          data: { ...mockMastery(skill).data, mastery_score: mockScore },
        },
      }
    }
    return apiClient.get<MasteryApiResponse>(
      `${API_URLS.learner}/internal/learner/${userId}/mastery?skill=${skill}`
    )
  },

  /**
   * Get next lesson recommendation from Adaptive Engine Service
   * @param userId - User identifier
   * @param skill - Current skill tag
   * @returns Next lesson recommendation with reason and content type
   */
  getNextLesson: async (userId: string, skill: string): Promise<{ data: NextLessonApiResponse }> => {
    if (USE_MOCK_DATA) {
      await delay(800)
      const randomScore = Math.floor(Math.random() * 100)
      return { data: mockNextLesson(skill, randomScore) }
    }
    return apiClient.post<NextLessonApiResponse>(`${API_URLS.adaptive}/next-lesson`, {
      user_id: userId,
      current_skill: skill,
    })
  },

  /**
   * Get question details from Content Service
   * @param questionId - Question identifier
   * @returns Question data including content, options, and metadata
   */
  getQuestion: async (questionId: number): Promise<{ data: ContentApiResponse }> => {
    if (USE_MOCK_DATA) {
      await delay(500)
      // Randomly decide if it's remedial based on id for demo variety
      const type = questionId % 2 === 0 ? "standard" : "remedial"
      return { data: mockQuestion(questionId, type) }
    }
    return apiClient.get<ContentApiResponse>(`${API_URLS.content}/api/content/${questionId}`)
  },

  /**
   * Submit an answer for scoring via Scoring Service
   * @param userId - User identifier
   * @param questionId - Question identifier
   * @param answer - User's answer
   * @returns Scoring result with correctness, score, and feedback
   */
  submitAnswer: async (
    userId: string,
    questionId: number,
    answer: string
  ): Promise<{ data: SubmitAnswerApiResponse }> => {
    if (USE_MOCK_DATA) {
      await delay(1000)
      // Mock: Answer "A" is always correct
      const isCorrect = answer === "A"
      return { data: mockSubmit(isCorrect) }
    }
    return apiClient.post<SubmitAnswerApiResponse>(`${API_URLS.scoring}/api/scoring/submit`, {
      user_id: userId,
      question_id: questionId,
      answer: answer,
    })
  },

  /**
   * Get list of answered question IDs for a user and skill
   * NEW endpoint from Scoring Service stabilization
   * @param userId - User identifier
   * @param skill - Skill tag
   * @returns Array of question IDs that the user has already answered
   */
  getAnsweredQuestions: async (
    userId: string,
    skill: string
  ): Promise<{ data: AnsweredQuestionsApiResponse }> => {
    if (USE_MOCK_DATA) {
      await delay(300)
      return {
        data: {
          error_code: 0,
          message: "Success",
          data: [1, 2, 3], // Mock answered question IDs
        },
      }
    }
    return apiClient.get<AnsweredQuestionsApiResponse>(
      `${API_URLS.scoring}/api/scoring/answered-questions?user_id=${userId}&skill=${skill}`
    )
  },
}
