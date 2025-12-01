import { test, expect } from './fixtures/antigravity-fixture';

test.describe('Error Handling and Edge Cases', () => {
    test.beforeEach(async ({ agPage }) => {
        // Set up user in localStorage before navigation
        await agPage.addInitScript(() => {
            window.localStorage.setItem('user_id', 'test-user-123');
        });
    });

    test.describe('Network Failures', () => {
        test('should handle API timeouts gracefully', async ({ agPage }) => {
            // Simulate timeout for Content Service
            await agPage.route('**/api/content/**', async route => {
                // Delay request indefinitely or for a long time to simulate timeout
                await new Promise(resolve => setTimeout(resolve, 10000));
                await route.abort('timeout');
            });

            await agPage.goto('/learn/math');

            // Should show loading state initially
            await expect(agPage.locator('text=/Consulting AI Tutor|Loading/')).toBeVisible();

            // Should eventually show error or retry option (depending on implementation)
            // For now, verify it doesn't crash to white screen
            const pageContent = await agPage.content();
            expect(pageContent.length).toBeGreaterThan(100);

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/error-network-timeout.png', fullPage: true });
        });

        test('should handle 503 Service Unavailable', async ({ agPage }) => {
            await agPage.route('**/api/content/**', route => {
                route.fulfill({
                    status: 503,
                    contentType: 'application/json',
                    body: JSON.stringify({ message: 'Service Unavailable' })
                });
            });

            await agPage.goto('/learn/math');
            await agPage.waitForTimeout(2000);

            // Verify page handles it (e.g., shows error message)
            // We expect some visible text indicating an issue or at least the header
            await expect(agPage.locator('header')).toBeVisible();

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/error-network-503.png', fullPage: true });
        });
    });

    test.describe('Invalid Data Handling', () => {
        test('should handle malformed API responses', async ({ agPage }) => {
            // Return malformed JSON (or valid JSON with missing fields)
            await agPage.route('**/api/content/**', route => {
                route.fulfill({
                    status: 200,
                    contentType: 'application/json',
                    body: JSON.stringify({
                        // Missing 'content', 'id', etc.
                        some_random_field: 'invalid'
                    })
                });
            });

            await agPage.goto('/learn/math');
            await agPage.waitForTimeout(2000);

            // Should not crash. Check for header or error boundary
            await expect(agPage.locator('header')).toBeVisible();

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/error-data-malformed.png', fullPage: true });
        });

        test('should handle empty data gracefully', async ({ agPage }) => {
            await agPage.route('**/api/content/**', route => {
                route.fulfill({
                    status: 200,
                    contentType: 'application/json',
                    body: JSON.stringify({})
                });
            });

            await agPage.goto('/learn/math');
            await agPage.waitForTimeout(2000);

            // Verify UI resilience
            await expect(agPage.locator('header')).toBeVisible();

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/error-data-empty.png', fullPage: true });
        });
    });

    test.describe('Session State Management', () => {
        test('should persist session state on page refresh', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');

            // Wait for question
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            // Reload page
            await agPage.reload();
            await agPage.waitForLoadState('networkidle');

            // Should still be in learning session
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });
            await expect(agPage).toHaveURL(/\/learn\/math/);

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/error-session-refresh.png', fullPage: true });
        });

        test('should clear session state on logout', async ({ agPage }) => {
            await agPage.goto('/dashboard');
            await agPage.waitForLoadState('networkidle');

            // Click logout
            await agPage.click('button:has-text("Logout")');

            // Verify localStorage is cleared (or at least user_id)
            const userId = await agPage.evaluate(() => window.localStorage.getItem('user_id'));
            expect(userId).toBeNull();

            // Verify redirection to home
            await expect(agPage).toHaveURL('/');

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/error-session-logout.png', fullPage: true });
        });
    });

    test.describe('Offline Mode (Enhancement)', () => {
        test('should handle offline state', async ({ agPage }) => {
            await agPage.goto('/dashboard');
            await agPage.waitForLoadState('networkidle');

            // Go offline
            await agPage.context().setOffline(true);

            // Try to navigate
            try {
                await agPage.click('text=Math');
                // It might fail to navigate or show offline message
                // We just want to ensure it doesn't crash the browser context
            } catch (e) {
                console.log('Navigation failed as expected in offline mode');
            }

            // Verify we are still on a valid page (either dashboard or error page)
            const title = await agPage.title();
            expect(title).toBeTruthy();

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/error-offline-mode.png', fullPage: true });

            // Restore online state
            await agPage.context().setOffline(false);
        });
    });
});
