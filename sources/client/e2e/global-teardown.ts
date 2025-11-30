/**
 * Playwright Global Teardown
 * 
 * This file runs after all tests to:
 * 1. Optionally stop backend services (if SKIP_SERVICE_TEARDOWN is not set)
 * 
 * Note: By default, services are NOT stopped to allow for:
 * - Faster re-runs of tests
 * - Manual inspection of services after tests
 * 
 * Set SKIP_SERVICE_TEARDOWN=true to keep services running.
 */

import { execSync } from 'child_process'
import { existsSync } from 'fs'
import { join } from 'path'

async function globalTeardown() {
  const skipTeardown = process.env.SKIP_SERVICE_TEARDOWN === 'true'
  const environment = process.env.NEXT_PUBLIC_ENV || process.env.NODE_ENV || 'local'

  // Only teardown for local environment
  if (environment !== 'local' || skipTeardown) {
    if (skipTeardown) {
      console.log('[Global Teardown] Skipping service teardown (SKIP_SERVICE_TEARDOWN=true)')
    } else {
      console.log(`[Global Teardown] Skipping service teardown for ${environment} environment`)
    }
    return
  }

  console.log('[Global Teardown] Cleaning up test environment...')

  const scriptsDir = join(__dirname, '..', 'scripts')
  const stopServicesScript = join(scriptsDir, 'stop-services.sh')

  try {
    if (existsSync(stopServicesScript)) {
      console.log('[Global Teardown] Stopping backend services...')
      execSync(`bash ${stopServicesScript}`, {
        stdio: 'inherit',
        cwd: join(__dirname, '..'),
        env: { ...process.env },
      })
      console.log('[Global Teardown] ✓ Services stopped')
    } else {
      console.warn(`[Global Teardown] ⚠ Service stop script not found: ${stopServicesScript}`)
    }

    console.log('[Global Teardown] ✓ Global teardown complete')
  } catch (error) {
    console.error('[Global Teardown] ✗ Failed to teardown test environment')
    console.error(error)
    // Don't throw - teardown failures shouldn't fail the test run
  }
}

export default globalTeardown

