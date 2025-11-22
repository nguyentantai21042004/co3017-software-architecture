import axios from "axios"
import { delay } from "@/lib/utils"

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

const mockNextLesson = (skill: string, currentScore: number) => ({
  error_code: 0,
  data: {
    next_lesson_id: Math.floor(Math.random() * 1000),
    reason: `Great! Your mastery is ${currentScore}%. Continue with the next challenge.`,
    mastery_score: currentScore,
    content_type: currentScore < 50 ? "remedial" : "standard",
  },
})

const mockQuestion = (id: number, type: string) => ({
  error_code: 0,
  data: {
    id: id,
    content:
      type === "remedial"
        ? `(Remedial) Basic concept check: What is the foundation of ${id}?`
        : `(Standard) Advanced problem solving: Analyze the case of ${id}.`,
    correct_answer: "A",
    skill_tag: "general",
    is_remedial: type === "remedial",
    options: ["A", "B", "C", "D"], // Added for UI rendering
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

// API Functions
export const api = {
  getMastery: async (userId: string, skill: string) => {
    if (USE_MOCK_DATA) {
      await delay(500)
      // Return a stable random score based on skill length for consistency in demo
      const mockScore = skill === "math" ? 45 : skill === "science" ? 75 : 85
      return { data: { error_code: 0, data: { ...mockMastery(skill).data, mastery_score: mockScore } } }
    }
    return apiClient.get(`${API_URLS.learner}/internal/learner/${userId}/mastery?skill=${skill}`)
  },

  getNextLesson: async (userId: string, skill: string) => {
    if (USE_MOCK_DATA) {
      await delay(800)
      const randomScore = Math.floor(Math.random() * 100)
      return { data: mockNextLesson(skill, randomScore) }
    }
    return apiClient.post(`${API_URLS.adaptive}/next-lesson`, {
      user_id: userId,
      current_skill: skill,
    })
  },

  getQuestion: async (questionId: number) => {
    if (USE_MOCK_DATA) {
      await delay(500)
      // Randomly decide if it's remedial based on id for demo variety
      const type = questionId % 2 === 0 ? "standard" : "remedial"
      return { data: mockQuestion(questionId, type) }
    }
    return apiClient.get(`${API_URLS.content}/api/content/${questionId}`)
  },

  submitAnswer: async (userId: string, questionId: number, answer: string) => {
    if (USE_MOCK_DATA) {
      await delay(1000)
      // Mock: Answer "A" is always correct
      const isCorrect = answer === "A"
      return { data: mockSubmit(isCorrect) }
    }
    return apiClient.post(`${API_URLS.scoring}/api/scoring/submit`, {
      user_id: userId,
      question_id: questionId,
      answer: answer,
    })
  },
}
