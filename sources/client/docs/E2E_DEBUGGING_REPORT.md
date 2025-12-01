# E2E Test Debugging Report

## Issue Description
The Dashboard E2E tests (`e2e/dashboard.spec.ts`) are consistently failing due to a timeout. The test expects the "Welcome back" text to appear, indicating that the dashboard has loaded data, but it seems to remain in a loading state.

## Symptoms
- **Test Failure:** `Timeout 30000ms exceeded` waiting for `locator('text=Welcome back, text=My Learning Dashboard')`.
- **Visuals:** Screenshots (if available) likely show the loading skeleton or spinner.
- **Logs:** Browser console logs show the page loads (`http://localhost:3001/dashboard`), but there are no successful data rendering logs.

## Verification Steps Taken
1.  **Backend Health:**
    - Confirmed all services (Content, Scoring, Learner, Adaptive) are running via Docker.
    - Verified health endpoints return 200 OK.
    - Verified direct API access via `curl` works (e.g., `curl http://localhost:8081/api/content/skills` returns data).

2.  **Test Data:**
    - Verified `setup-test-data.sh` ran successfully.
    - Confirmed data exists in PostgreSQL (`questions` table has 32 rows, `skill_mastery` has 2 rows).

3.  **Frontend Configuration:**
    - **Port Conflict:** Resolved port 3000 conflict by moving Client to port 3001.
    - **Environment Variables:**
        - Checked `.env.local`.
        - Updated `localhost` to `127.0.0.1` to rule out Node.js v17+ IP resolution issues (IPv4 vs IPv6).
    - **API Client:**
        - Added logging to `services/api.ts`.
        - Confirmed `getAvailableSkills` is being called with the correct URL: `http://127.0.0.1:8081/api/content/skills`.

4.  **Test Environment:**
    - **Fixture:** Fixed import path for `antigravity-fixture`.
    - **LocalStorage:** Confirmed `user_id` is correctly set in the browser context before navigation.

## Critical Findings (Component Instrumentation)

### Observed Behavior
From `test_failure_12.log`, the browser console shows:
```
[Browser Console] log: Dashboard: Calling getAvailableSkills
[Browser Console] log: Dashboard: getAvailableSkills response {data: undefined}
[Browser Console] log: Dashboard: fetchData finally block - setLoading(false)
```

### Root Cause Analysis
1. **Mock Data Not Activated:** Despite setting `USE_MOCK_DATA = true`, the log shows `Fetching skills from: http://127.0.0.1:8081/api/content/skills`, indicating the code is still attempting real API calls.
2. **Response is Undefined:** The `getAvailableSkills` function returns `{data: undefined}`, which means the API call is failing silently (no error thrown, but no data returned).
3. **Server Caching Issue:** The Next.js dev server appears to be serving stale JavaScript code despite file changes. This is evidenced by:
   - Mock mode flag changes not taking effect
   - Multiple server restarts (`pkill`) not resolving the issue
   - `reuseExistingServer: true` in Playwright config keeping old server alive

### Why Tests Fail
- The component reaches `finally` block and sets `loading = false`
- However, `skills` array remains empty because `skillsResponse.data` is undefined
- The UI renders the "empty state" or remains in skeleton/loading view
- The test selector `text=Welcome back, text=My Learning Dashboard` never appears because this text only shows when data is successfully loaded

## Recommended Solutions

### Immediate Fix
1. **Disable Server Reuse:** Set `reuseExistingServer: false` in `playwright.config.ts` to force fresh server on each test run
2. **Clear Next.js Cache:** Delete `.next` directory before running tests
3. **Verify Mock Mode:** Add explicit check at module level to log `USE_MOCK_DATA` value on import

### Long-term Fix
1. **Fix Test Selector:** The selector `text=Welcome back, text=My Learning Dashboard` appears to be incorrect. Should be separate selectors or use `getByText` with regex.
2. **Add Error Boundaries:** Component should handle `undefined` responses gracefully
3. **Improve Test Resilience:** Wait for specific skill cards instead of generic "Welcome" text

## Next Steps
1. Stop all running tests
2. Clear `.next` cache
3. Update `playwright.config.ts` to disable server reuse
4. Fix test selectors in `dashboard.spec.ts`
5. Re-run tests with fresh server
