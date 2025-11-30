import { defineConfig, devices } from '@playwright/test'

/**
 * Playwright Configuration with Environment Support
 * 
 * Supports environment-based configuration via environment variables:
 * - NEXT_PUBLIC_CLIENT_URL: Base URL for the client application (default: http://localhost:3000)
 * - NEXT_PUBLIC_ENV: Environment name (local, test, staging)
 * 
 * See https://playwright.dev/docs/test-configuration.
 */

// Get base URL from environment variable with fallback
const getBaseURL = (): string => {
  const envBaseURL = process.env.NEXT_PUBLIC_CLIENT_URL
  if (envBaseURL && envBaseURL.trim() !== '') {
    return envBaseURL.trim()
  }
  return 'http://localhost:3000'
}

// Get environment name
const getEnvironment = (): string => {
  return process.env.NEXT_PUBLIC_ENV || process.env.NODE_ENV || 'local'
}

const baseURL = getBaseURL()
const environment = getEnvironment()

// Log configuration on load
console.log(`[Playwright] Environment: ${environment}`)
console.log(`[Playwright] Base URL: ${baseURL}`)

export default defineConfig({
  testDir: './e2e',
  /* Global setup and teardown */
  globalSetup: require.resolve('./e2e/global-setup.ts'),
  globalTeardown: require.resolve('./e2e/global-teardown.ts'),
  /* Run tests in files in parallel */
  fullyParallel: true,
  /* Fail the build on CI if you accidentally left test.only in the source code. */
  forbidOnly: !!process.env.CI,
  /* Retry on CI only */
  retries: process.env.CI ? 2 : 0,
  /* Opt out of parallel tests on CI. */
  workers: process.env.CI ? 1 : undefined,
  /* Reporter to use. See https://playwright.dev/docs/test-reporters */
  reporter: 'html',
  /* Shared settings for all the projects below. See https://playwright.dev/docs/api/class-testoptions. */
  use: {
    /* Base URL to use in actions like `await page.goto('/')`. */
    baseURL: baseURL,
    /* Collect trace when retrying the failed test. See https://playwright.dev/docs/trace-viewer */
    trace: 'on-first-retry',
    /* Screenshot on failure */
    screenshot: 'only-on-failure',
    /* Video on failure */
    video: 'retain-on-failure',
  },

  /* Configure projects for major browsers */
  projects: [
    {
      name: 'chromium',
      use: { ...devices['Desktop Chrome'] },
    },

    {
      name: 'firefox',
      use: { ...devices['Desktop Firefox'] },
    },

    {
      name: 'webkit',
      use: { ...devices['Desktop Safari'] },
    },
  ],

  /* Run your local dev server before starting the tests */
  webServer: environment === 'local' || !process.env.NEXT_PUBLIC_CLIENT_URL
    ? {
        command: 'npm run dev',
        url: baseURL,
        reuseExistingServer: !process.env.CI,
        timeout: 120 * 1000,
      }
    : undefined, // Don't start web server for test/staging environments
})

