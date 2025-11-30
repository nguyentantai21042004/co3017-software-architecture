/**
 * TypeScript types for backend API contracts
 * These types match the backend service response structures
 */

// Common API Response Wrapper
export interface ApiResponse<T> {
    error_code: number
    message?: string
    data: T
}

// Content Service Types
export interface QuestionResponse {
    id: number
    content: string
    options: string[]
    difficulty_level: number
    skill_tag: string
    correct_answer: string
    is_remedial: boolean
    created_at?: string
}

export interface ContentApiResponse extends ApiResponse<QuestionResponse> { }

export interface SkillsApiResponse extends ApiResponse<string[]> { }

// Learner Model Service Types
export interface MasteryData {
    user_id: string
    skill_tag: string
    mastery_score: number
    last_updated: string
}

export interface MasteryApiResponse extends ApiResponse<MasteryData> { }

// Adaptive Engine Service Types
export interface NextLessonData {
    next_lesson_id: number
    reason: string
    mastery_score: number
    content_type: "remedial" | "standard"
}

export interface NextLessonApiResponse extends ApiResponse<NextLessonData> { }

// Scoring Service Types
export interface SubmitAnswerData {
    correct: boolean
    score: number
    feedback: string
}

export interface SubmitAnswerApiResponse extends ApiResponse<SubmitAnswerData> { }

export interface AnsweredQuestionsApiResponse extends ApiResponse<number[]> { }

// Request Types
export interface NextLessonRequest {
    user_id: string
    current_skill: string
}

export interface SubmitAnswerRequest {
    user_id: string
    question_id: number
    answer: string
}

