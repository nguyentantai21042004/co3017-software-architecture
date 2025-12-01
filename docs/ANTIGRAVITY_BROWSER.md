# Antigravity Browser Integration

This project uses a custom "Antigravity Browser" fixture for Playwright E2E tests to provide enhanced capabilities and consistent test environments.

## Setup

The Antigravity Browser is configured via a Playwright fixture located in `sources/client/e2e/fixtures/antigravity-fixture.ts`.

### Usage in Tests

To use the Antigravity Browser in your tests, import `test` and `expect` from the fixture instead of `@playwright/test`:

```typescript
import { test, expect } from '../fixtures/antigravity-fixture';

test('My Antigravity Test', async ({ agPage }) => {
  await agPage.goto('/');
  // Your test code here
});
```

## Features

The Antigravity Browser fixture provides:

1.  **Custom User Agent**: Identifies automation traffic (`Antigravity-Browser/1.0`).
2.  **Console Log Capturing**: Automatically captures and logs browser console messages to the test output.
3.  **Initialization Scripts**: Injects `window.__ANTIGRAVITY_ENABLED__` for client-side detection.
4.  **Video Recording**: Automatically records videos of test executions (saved in `test-results/videos`).
5.  **Standardized Viewport**: Sets a consistent viewport of 1280x720.

## Configuration

Configuration can be adjusted in `sources/client/e2e/fixtures/antigravity-fixture.ts`.
Global Playwright settings are managed in `playwright.config.ts`.
