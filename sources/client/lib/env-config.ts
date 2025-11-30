/**
 * Environment Configuration
 * 
 * This module provides environment-based configuration for API URLs.
 * It reads from Next.js environment variables (NEXT_PUBLIC_*) and provides
 * fallback to default localhost URLs if not set.
 */

// Default localhost URLs (fallback)
const DEFAULT_API_URLS = {
  content: "http://localhost:8081",
  scoring: "http://localhost:8082",
  learner: "http://localhost:8083",
  adaptive: "http://localhost:8084/api/adaptive",
  client: "http://localhost:3000",
}

/**
 * Get API URL from environment variable with fallback
 */
function getEnvUrl(envVar: string | undefined, fallback: string): string {
  if (!envVar || envVar.trim() === "") {
    return fallback
  }
  return envVar.trim()
}

/**
 * Environment-based API URLs configuration
 * 
 * Reads from Next.js environment variables (NEXT_PUBLIC_*):
 * - NEXT_PUBLIC_CONTENT_API_URL
 * - NEXT_PUBLIC_SCORING_API_URL
 * - NEXT_PUBLIC_LEARNER_API_URL
 * - NEXT_PUBLIC_ADAPTIVE_API_URL
 * - NEXT_PUBLIC_CLIENT_URL
 * 
 * Falls back to localhost URLs if environment variables are not set.
 */
export const API_URLS = {
  content: getEnvUrl(
    process.env.NEXT_PUBLIC_CONTENT_API_URL,
    DEFAULT_API_URLS.content
  ),
  scoring: getEnvUrl(
    process.env.NEXT_PUBLIC_SCORING_API_URL,
    DEFAULT_API_URLS.scoring
  ),
  learner: getEnvUrl(
    process.env.NEXT_PUBLIC_LEARNER_API_URL,
    DEFAULT_API_URLS.learner
  ),
  adaptive: getEnvUrl(
    process.env.NEXT_PUBLIC_ADAPTIVE_API_URL,
    DEFAULT_API_URLS.adaptive
  ),
  client: getEnvUrl(
    process.env.NEXT_PUBLIC_CLIENT_URL,
    DEFAULT_API_URLS.client
  ),
}

/**
 * Validate environment configuration
 * 
 * Checks if all required environment variables are set (or have valid defaults).
 * Logs warnings if using fallback values in non-local environments.
 * 
 * @returns Object with validation results
 */
export function validateEnvConfig(): {
  isValid: boolean
  warnings: string[]
  errors: string[]
} {
  const warnings: string[] = []
  const errors: string[] = []

  // Check if we're in a non-local environment but using localhost URLs
  const isUsingLocalhost =
    API_URLS.content.includes("localhost") ||
    API_URLS.scoring.includes("localhost") ||
    API_URLS.learner.includes("localhost") ||
    API_URLS.adaptive.includes("localhost")

  // Detect environment (basic check)
  const isProduction = process.env.NODE_ENV === "production"
  const isTest = process.env.NODE_ENV === "test" || process.env.NEXT_PUBLIC_ENV === "test"
  const isStaging = process.env.NEXT_PUBLIC_ENV === "staging"

  if ((isTest || isStaging || isProduction) && isUsingLocalhost) {
    warnings.push(
      "Using localhost URLs in non-local environment. " +
      "Please set appropriate NEXT_PUBLIC_*_API_URL environment variables."
    )
  }

  // Validate URL format (basic check)
  const urlPattern = /^https?:\/\/.+/i
  const urlsToCheck = [
    { name: "Content API", url: API_URLS.content },
    { name: "Scoring API", url: API_URLS.scoring },
    { name: "Learner API", url: API_URLS.learner },
    { name: "Adaptive API", url: API_URLS.adaptive },
    { name: "Client URL", url: API_URLS.client },
  ]

  urlsToCheck.forEach(({ name, url }) => {
    if (!urlPattern.test(url)) {
      errors.push(`${name} URL is invalid: ${url}`)
    }
  })

  return {
    isValid: errors.length === 0,
    warnings,
    errors,
  }
}

/**
 * Get current environment name
 */
export function getEnvironment(): "local" | "test" | "staging" | "production" | "unknown" {
  if (process.env.NEXT_PUBLIC_ENV === "test") return "test"
  if (process.env.NEXT_PUBLIC_ENV === "staging") return "staging"
  if (process.env.NODE_ENV === "production") return "production"
  if (process.env.NODE_ENV === "development" || !process.env.NODE_ENV) return "local"
  return "unknown"
}

// Validate on module load (only in development/test)
if (typeof window === "undefined" && (process.env.NODE_ENV === "development" || process.env.NODE_ENV === "test")) {
  const validation = validateEnvConfig()
  if (validation.errors.length > 0) {
    console.error("Environment Configuration Errors:", validation.errors)
  }
  if (validation.warnings.length > 0) {
    console.warn("Environment Configuration Warnings:", validation.warnings)
  }
}

