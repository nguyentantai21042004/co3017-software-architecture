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
  return 'http://localhost:3001'
}

// Get environment name
const getEnvironment = (): string => {
  return process.env.NEXT_PUBLIC_ENV || process.env.NODE_ENV || 'local'
}

// Derive run identifier for per-run artifacts
const getRunId = (): string => {
  if (process.env.PW_RUN_ID && process.env.PW_RUN_ID.trim() !== '') {
    return process.env.PW_RUN_ID.trim()
  }
  // ISO timestamp without characters that are invalid in folder names
  return new Date().toISOString().replace(/[:.]/g, '-')
}

const baseURL = getBaseURL()
const environment = getEnvironment()
const runId = getRunId()

// Base directories for artifacts (per-run)
const artifactsBaseDir = process.env.PW_ARTIFACTS_DIR || `test-results/${runId}`
const reportBaseDir = process.env.PW_REPORT_DIR || `playwright-report/${runId}`

// Expose to tests so they can build screenshot paths
process.env.PW_RUN_ID = runId
process.env.PW_ARTIFACTS_DIR = artifactsBaseDir
process.env.PW_REPORT_DIR = reportBaseDir

// Log configuration on load
console.log(`[Playwright] Environment: ${environment}`)
console.log(`[Playwright] Base URL: ${baseURL}`)
console.log(`[Playwright] Run ID: ${runId}`)
console.log(`[Playwright] Artifacts dir: ${artifactsBaseDir}`)
console.log(`[Playwright] Report dir: ${reportBaseDir}`)

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
  reporter: [
    [
      'html',
      {
        outputFolder: reportBaseDir,
        open: 'never',
      },
    ],
    ['json', { outputFile: `${artifactsBaseDir}/results.json` }],
  ],
  /* Shared settings for all the projects below. See https://playwright.dev/docs/api/class-testoptions. */
  use: {
    /* Base URL to use in actions like `await page.goto('/')`. */
    baseURL,

    /* Where to put per-test artifacts (traces, screenshots, videos) */
    // Note: manual screenshots use PW_ARTIFACTS_DIR to compute their own paths
    // but Playwright's own artifacts will respect this outputDir.
    // @ts-expect-error - Playwright types allow outputDir at top-level, but keeping here for clarity in config
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    outputDir: artifactsBaseDir,

    /* Collect trace when retrying the failed test. See https://playwright.dev/docs/trace-viewer */
    trace: 'on-first-retry',

    /* Screenshot configuration */
    screenshot: {
      mode: 'only-on-failure', // Capture screenshots on test failure
      fullPage: true, // Capture full page screenshot
    },

    /* Video configuration */
    video: {
      mode: 'retain-on-failure', // Keep videos only for failed tests
      size: { width: 1280, height: 720 },
    },
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
      command: 'npm run dev -- -p 3001',
      url: baseURL,
      reuseExistingServer: false, // Always start fresh server to avoid stale code
      timeout: 120 * 1000,
    }
    : undefined, // Don't start web server for test/staging environments
})

