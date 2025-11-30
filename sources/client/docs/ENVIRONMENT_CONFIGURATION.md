# Environment Configuration Guide

This document describes how to configure the application for different environments (local, test, staging, production).

## Overview

The application uses environment variables to configure API endpoints and other environment-specific settings. Configuration is managed through `.env` files that are loaded by Next.js.

## Environment Files

### Available Templates

- `.env.local.example` - Template for local development
- `.env.test.example` - Template for test environment
- `.env.staging.example` - Template for staging environment

### Creating Environment Files

1. **Copy the appropriate template:**
   ```bash
   cp .env.local.example .env.local
   ```

2. **Or use the switch script:**
   ```bash
   ./scripts/switch-env.sh local
   ./scripts/switch-env.sh test
   ./scripts/switch-env.sh staging
   ```

3. **Update the values** in the `.env.local` file with your actual URLs and configuration.

## Environment Variables

### Required Variables

All environment variables must be prefixed with `NEXT_PUBLIC_` to be accessible in the browser.

| Variable | Description | Example |
|----------|-------------|---------|
| `NEXT_PUBLIC_CONTENT_API_URL` | Content Service API URL | `http://localhost:8081` |
| `NEXT_PUBLIC_SCORING_API_URL` | Scoring Service API URL | `http://localhost:8082` |
| `NEXT_PUBLIC_LEARNER_API_URL` | Learner Model Service API URL | `http://localhost:8083` |
| `NEXT_PUBLIC_ADAPTIVE_API_URL` | Adaptive Engine Service API URL | `http://localhost:8084/api/adaptive` |
| `NEXT_PUBLIC_CLIENT_URL` | Client Application URL (for E2E tests) | `http://localhost:3000` |

### Optional Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `NEXT_PUBLIC_ENV` | Environment name (local, test, staging, production) | `local` |
| `NODE_ENV` | Node.js environment (development, production, test) | `development` |

## Environment-Specific Configuration

### Local Development

**File:** `.env.local`

```bash
NEXT_PUBLIC_CONTENT_API_URL=http://localhost:8081
NEXT_PUBLIC_SCORING_API_URL=http://localhost:8082
NEXT_PUBLIC_LEARNER_API_URL=http://localhost:8083
NEXT_PUBLIC_ADAPTIVE_API_URL=http://localhost:8084/api/adaptive
NEXT_PUBLIC_CLIENT_URL=http://localhost:3000
```

**Usage:**
- Development server automatically loads `.env.local`
- Used for local development and E2E testing
- Services run on localhost

### Test Environment

**File:** `.env.test` (or `.env.local` with test values)

```bash
NEXT_PUBLIC_ENV=test
NEXT_PUBLIC_CONTENT_API_URL=http://test-content-service:8081
NEXT_PUBLIC_SCORING_API_URL=http://test-scoring-service:8082
NEXT_PUBLIC_LEARNER_API_URL=http://test-learner-service:8083
NEXT_PUBLIC_ADAPTIVE_API_URL=http://test-adaptive-engine:8084/api/adaptive
NEXT_PUBLIC_CLIENT_URL=http://test-client:3000
```

**Usage:**
- Used for running E2E tests against test environment
- Set `NEXT_PUBLIC_ENV=test` to enable test mode
- Services deployed to test infrastructure

### Staging Environment

**File:** `.env.staging` (or `.env.local` with staging values)

```bash
NEXT_PUBLIC_ENV=staging
NEXT_PUBLIC_CONTENT_API_URL=https://staging-content.example.com
NEXT_PUBLIC_SCORING_API_URL=https://staging-scoring.example.com
NEXT_PUBLIC_LEARNER_API_URL=https://staging-learner.example.com
NEXT_PUBLIC_ADAPTIVE_API_URL=https://staging-adaptive.example.com/api/adaptive
NEXT_PUBLIC_CLIENT_URL=https://staging-client.example.com
```

**Usage:**
- Used for staging/pre-production testing
- Services deployed to staging infrastructure
- Typically uses HTTPS

## Switching Environments

### Using the Switch Script

The easiest way to switch environments:

```bash
# Switch to local
./scripts/switch-env.sh local

# Switch to test
./scripts/switch-env.sh test

# Switch to staging
./scripts/switch-env.sh staging
```

This script copies the appropriate `.env.*.example` file to `.env.local`.

### Manual Switching

1. **Backup current configuration** (if needed):
   ```bash
   cp .env.local .env.local.backup
   ```

2. **Copy the desired template:**
   ```bash
   cp .env.test.example .env.local
   ```

3. **Update values** in `.env.local` as needed

4. **Validate configuration:**
   ```bash
   npm run e2e:validate-env
   ```

## Validation

### Validate Environment Configuration

```bash
npm run e2e:validate-env
# Or directly:
./scripts/validate-env.sh
```

The validation script:
- Checks if all required variables are set
- Validates URL format
- Warns if using localhost in non-local environments
- Shows current configuration

### Runtime Validation

The `lib/env-config.ts` module automatically validates configuration on load (in development/test mode):
- Logs warnings if using localhost in non-local environments
- Logs errors for invalid URL formats
- Provides fallback to localhost URLs if variables are not set

## Environment Detection

The application automatically detects the environment:

1. **From `NEXT_PUBLIC_ENV` variable:**
   - `test` → Test environment
   - `staging` → Staging environment

2. **From `NODE_ENV` variable:**
   - `production` → Production environment
   - `development` or unset → Local environment

3. **Default:** Local environment

## Playwright Configuration

Playwright automatically uses environment variables:

- **Base URL:** Reads from `NEXT_PUBLIC_CLIENT_URL` (default: `http://localhost:3000`)
- **Web Server:** Only starts for local environment (when `NEXT_PUBLIC_CLIENT_URL` is not set or points to localhost)

### Running E2E Tests in Different Environments

```bash
# Local environment (default)
npm run test:e2e

# Test environment
NEXT_PUBLIC_ENV=test NEXT_PUBLIC_CLIENT_URL=http://test-client:3000 npm run test:e2e

# Staging environment
NEXT_PUBLIC_ENV=staging NEXT_PUBLIC_CLIENT_URL=https://staging-client.example.com npm run test:e2e
```

## Best Practices

1. **Never commit `.env.local`, `.env.test`, or `.env.staging` files**
   - These files are excluded in `.gitignore`
   - Only commit `.env.*.example` template files

2. **Use environment-specific URLs**
   - Local: `http://localhost:*`
   - Test: `http://test-*:808*` or internal URLs
   - Staging: `https://staging-*.example.com`
   - Production: `https://*.example.com`

3. **Validate before running tests**
   ```bash
   npm run e2e:validate-env
   ```

4. **Document environment-specific requirements**
   - Update this document when adding new environments
   - Keep `.env.*.example` files up to date

5. **Use the switch script for convenience**
   ```bash
   ./scripts/switch-env.sh <environment>
   ```

## Troubleshooting

### Environment Variables Not Loading

**Problem:** Environment variables are not being read.

**Solutions:**
1. Ensure variables are prefixed with `NEXT_PUBLIC_`
2. Restart the Next.js dev server after changing `.env.local`
3. Check that `.env.local` file exists and is in the correct location
4. Verify file syntax (no spaces around `=`)

### Wrong Environment Detected

**Problem:** Application is using the wrong environment configuration.

**Solutions:**
1. Check `NEXT_PUBLIC_ENV` variable is set correctly
2. Verify `.env.local` file has correct values
3. Run validation: `npm run e2e:validate-env`
4. Check for conflicting environment variables

### Playwright Using Wrong Base URL

**Problem:** E2E tests are running against the wrong URL.

**Solutions:**
1. Set `NEXT_PUBLIC_CLIENT_URL` environment variable
2. Check `playwright.config.ts` is reading the variable correctly
3. Verify the URL is accessible
4. For test/staging, ensure webServer is disabled (set `NEXT_PUBLIC_CLIENT_URL`)

## Related Files

- `lib/env-config.ts` - Environment configuration module
- `services/api.ts` - API service using environment URLs
- `playwright.config.ts` - Playwright configuration with environment support
- `scripts/switch-env.sh` - Environment switching script
- `scripts/validate-env.sh` - Environment validation script

