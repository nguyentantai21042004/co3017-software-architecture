/**
 * Helper functions for API error handling and response validation
 */

import type { ApiResponse } from '@/types/api'

/**
 * Validates API response and throws error if error_code is not 0
 */
export function validateApiResponse<T>(response: { data: ApiResponse<T> }): T {
  if (response.data.error_code !== 0) {
    throw new Error(response.data.message || 'API request failed')
  }
  return response.data.data
}

/**
 * Creates a standardized error message from API error
 */
export function getApiErrorMessage(error: unknown): string {
  if (error && typeof error === 'object' && 'userMessage' in error) {
    return (error as { userMessage: string }).userMessage
  }
  if (error instanceof Error) {
    return error.message
  }
  return 'An unexpected error occurred'
}

