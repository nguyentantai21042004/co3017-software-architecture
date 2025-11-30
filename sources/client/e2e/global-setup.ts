/**
 * Playwright Global Setup
 * 
 * This file runs before all tests to:
 * 1. Start backend services (Docker containers)
 * 2. Verify services are healthy
 * 3. Setup test data
 * 
 * Only runs for local environment. For test/staging environments,
 * services should already be deployed.
 */

import { execSync } from 'child_process'
import { existsSync } from 'fs'
import { join } from 'path'

async function globalSetup() {
  const environment = process.env.NEXT_PUBLIC_ENV || process.env.NODE_ENV || 'local'
  
  // Only start services for local environment
  if (environment !== 'local' && process.env.NEXT_PUBLIC_CLIENT_URL) {
    console.log(`[Global Setup] Skipping service startup for ${environment} environment`)
    console.log(`[Global Setup] Assuming services are already deployed at ${process.env.NEXT_PUBLIC_CLIENT_URL}`)
    return
  }

  console.log('[Global Setup] Starting backend services for E2E tests...')
  
  const scriptsDir = join(__dirname, '..', 'scripts')
  const startServicesScript = join(scriptsDir, 'start-services.sh')
  const verifyServicesScript = join(scriptsDir, 'verify-services.sh')
  const setupDataScript = join(scriptsDir, 'setup-test-data.sh')

  try {
    // Step 1: Start services
    if (existsSync(startServicesScript)) {
      console.log('[Global Setup] Starting backend services...')
      execSync(`bash ${startServicesScript}`, {
        stdio: 'inherit',
        cwd: join(__dirname, '..'),
        env: { ...process.env },
      })
      console.log('[Global Setup] ✓ Services started')
    } else {
      console.warn(`[Global Setup] ⚠ Service startup script not found: ${startServicesScript}`)
      console.warn('[Global Setup] Please start services manually before running tests')
    }

    // Step 2: Verify services are healthy
    if (existsSync(verifyServicesScript)) {
      console.log('[Global Setup] Verifying services are healthy...')
      try {
        execSync(`bash ${verifyServicesScript}`, {
          stdio: 'inherit',
          cwd: join(__dirname, '..'),
          env: { ...process.env },
        })
        console.log('[Global Setup] ✓ All services are healthy')
      } catch (error) {
        console.error('[Global Setup] ✗ Service verification failed')
        console.error('[Global Setup] Some services may not be ready. Tests may fail.')
        // Don't throw - let tests run and fail with clear error messages
      }
    }

    // Step 3: Setup test data
    if (existsSync(setupDataScript)) {
      console.log('[Global Setup] Setting up test data...')
      try {
        execSync(`bash ${setupDataScript}`, {
          stdio: 'inherit',
          cwd: join(__dirname, '..'),
          env: { ...process.env },
        })
        console.log('[Global Setup] ✓ Test data setup complete')
      } catch (error) {
        console.warn('[Global Setup] ⚠ Test data setup failed or data already exists')
        // Don't throw - tests may still work with existing data
      }
    }

    console.log('[Global Setup] ✓ Global setup complete. Ready for E2E tests.')
  } catch (error) {
    console.error('[Global Setup] ✗ Failed to setup test environment')
    console.error(error)
    throw error
  }
}

export default globalSetup

