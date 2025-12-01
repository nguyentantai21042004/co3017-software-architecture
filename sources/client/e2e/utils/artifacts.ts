import { mkdirSync } from 'fs'
import { dirname, join } from 'path'

/**
 * Build a screenshot path that is scoped to the current Playwright run.
 *
 * The base directory is derived from:
 * - PW_ARTIFACTS_DIR (set in playwright.config.ts), or
 * - falls back to "test-results" if not set.
 *
 * The helper ensures that the parent directory exists before returning the path.
 */
export const screenshotPath = (fileName: string): string => {
  const baseDir = process.env.PW_ARTIFACTS_DIR || 'test-results'
  const fullPath = join(baseDir, 'screenshots', fileName)

  // Ensure directory exists to avoid ENOENT on first run
  mkdirSync(dirname(fullPath), { recursive: true })

  return fullPath
}


