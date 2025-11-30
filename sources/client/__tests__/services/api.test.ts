// Mock axios module with factory function
const mockGet = jest.fn()
const mockPost = jest.fn()

jest.mock('axios', () => {
  const mockGet = jest.fn()
  const mockPost = jest.fn()
  return {
    __esModule: true,
    default: {
      create: jest.fn(() => ({
        get: mockGet,
        post: mockPost,
        interceptors: {
          response: {
            use: jest.fn(),
          },
        },
      })),
    },
    // Export mocks for use in tests
    _mockGet: mockGet,
    _mockPost: mockPost,
  }
})

import axios from 'axios'
import { api } from '@/services/api'

// Get the mocked functions
const mockedAxios = axios as any
const getMockGet = () => {
  const instance = mockedAxios.create()
  return instance.get
}
const getMockPost = () => {
  const instance = mockedAxios.create()
  return instance.post
}

describe('API Service', () => {
  let mockGetFn: jest.Mock
  let mockPostFn: jest.Mock

  beforeEach(() => {
    jest.clearAllMocks()
    // Get fresh mock instances
    mockGetFn = getMockGet()
    mockPostFn = getMockPost()
  })

  describe('getAvailableSkills', () => {
    it('should fetch available skills from content service', async () => {
      const mockResponse = {
        data: {
          error_code: 0,
          message: 'Success',
          data: ['math_algebra', 'science_physics'],
        },
      }
      mockGetFn.mockResolvedValue(mockResponse)

      const result = await api.getAvailableSkills()

      expect(mockGetFn).toHaveBeenCalledWith(
        'http://localhost:8081/api/content/skills'
      )
      expect(result.data.data).toEqual(['math_algebra', 'science_physics'])
    })

    it('should handle errors gracefully', async () => {
      mockGetFn.mockRejectedValue(new Error('Network error'))

      await expect(api.getAvailableSkills()).rejects.toThrow()
    })
  })

  describe('getMastery', () => {
    it('should fetch mastery score for user and skill', async () => {
      const mockResponse = {
        data: {
          error_code: 0,
          message: 'Success',
          data: {
            user_id: 'user_01',
            skill_tag: 'math_algebra',
            mastery_score: 75,
            last_updated: '2025-01-01T00:00:00Z',
          },
        },
      }
      mockGetFn.mockResolvedValue(mockResponse)

      const result = await api.getMastery('user_01', 'math_algebra')

      expect(mockGetFn).toHaveBeenCalledWith(
        'http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra'
      )
      expect(result.data.data.mastery_score).toBe(75)
    })
  })

  describe('getNextLesson', () => {
    it('should get next lesson recommendation', async () => {
      const mockResponse = {
        data: {
          error_code: 0,
          message: 'Success',
          data: {
            next_lesson_id: 123,
            reason: 'Continue learning',
            mastery_score: 75,
            content_type: 'standard',
          },
        },
      }
      mockPostFn.mockResolvedValue(mockResponse)

      const result = await api.getNextLesson('user_01', 'math_algebra')

      expect(mockPostFn).toHaveBeenCalledWith(
        'http://localhost:8084/api/adaptive/next-lesson',
        {
          user_id: 'user_01',
          current_skill: 'math_algebra',
        }
      )
      expect(result.data.data.next_lesson_id).toBe(123)
      expect(result.data.data.content_type).toBe('standard')
    })
  })

  describe('getQuestion', () => {
    it('should fetch question by ID', async () => {
      const mockResponse = {
        data: {
          error_code: 0,
          message: 'Success',
          data: {
            id: 123,
            content: 'What is 2+2?',
            options: ['A. 3', 'B. 4', 'C. 5'],
            skill_tag: 'math',
            correct_answer: 'B',
            is_remedial: false,
            difficulty_level: 1,
          },
        },
      }
      mockGetFn.mockResolvedValue(mockResponse)

      const result = await api.getQuestion(123)

      expect(mockGetFn).toHaveBeenCalledWith(
        'http://localhost:8081/api/content/123'
      )
      expect(result.data.data.id).toBe(123)
      expect(result.data.data.content).toBe('What is 2+2?')
    })
  })

  describe('submitAnswer', () => {
    it('should submit answer and get feedback', async () => {
      const mockResponse = {
        data: {
          error_code: 0,
          message: 'Success',
          data: {
            correct: true,
            score: 100,
            feedback: 'Correct! Well done.',
          },
        },
      }
      mockPostFn.mockResolvedValue(mockResponse)

      const result = await api.submitAnswer('user_01', 123, 'B')

      expect(mockPostFn).toHaveBeenCalledWith(
        'http://localhost:8082/api/scoring/submit',
        {
          user_id: 'user_01',
          question_id: 123,
          answer: 'B',
        }
      )
      expect(result.data.data.correct).toBe(true)
      expect(result.data.data.score).toBe(100)
    })
  })

  describe('getAnsweredQuestions', () => {
    it('should fetch answered question IDs for user and skill', async () => {
      const mockResponse = {
        data: {
          error_code: 0,
          message: 'Success',
          data: [1, 2, 3, 5, 8],
        },
      }
      mockGetFn.mockResolvedValue(mockResponse)

      const result = await api.getAnsweredQuestions('user_01', 'math_algebra')

      expect(mockGetFn).toHaveBeenCalledWith(
        'http://localhost:8082/api/scoring/answered-questions?user_id=user_01&skill=math_algebra'
      )
      expect(result.data.data).toEqual([1, 2, 3, 5, 8])
    })
  })
})
