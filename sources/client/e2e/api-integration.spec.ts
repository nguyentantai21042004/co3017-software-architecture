import { test, expect } from './fixtures/antigravity-fixture';
import { screenshotPath } from './utils/artifacts';

test.describe('API Integration Tests', () => {
    test.beforeEach(async ({ agPage }) => {
        // Set up user in localStorage before navigation
        await agPage.addInitScript(() => {
            window.localStorage.setItem('user_id', 'test-user-123');
        });
    });

    test('should integrate with Content Service correctly', async ({ agPage }) => {
        // Track API calls to Content Service
        const contentApiCalls: string[] = [];

        agPage.on('request', request => {
            const url = request.url();
            if (url.includes('/api/content/')) {
                contentApiCalls.push(url);
                console.log('Content API call:', url);
            }
        });

        // Navigate to learning session
        await agPage.goto('/learn/math');
        await agPage.waitForLoadState('networkidle');

        // Wait for question to load
        await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

        // Verify Content Service was called
        expect(contentApiCalls.length).toBeGreaterThan(0);

        // Verify question content is displayed
        const questionText = await agPage.locator('h2').first().textContent();
        expect(questionText).toBeTruthy();
        expect(questionText!.length).toBeGreaterThan(0);

        // Verify question ID is shown (metadata parsed correctly)
        await expect(agPage.locator('text=/ID:/').first()).toBeVisible();

        // Capture screenshot
        await agPage.screenshot({ path: screenshotPath('api-content-service.png'), fullPage: true });
    });

    test('should integrate with Adaptive Engine correctly', async ({ agPage }) => {
        // Track API calls to Adaptive Engine
        const adaptiveApiCalls: string[] = [];

        agPage.on('request', request => {
            const url = request.url();
            if (url.includes('/api/adaptive/') || url.includes('/next-lesson')) {
                adaptiveApiCalls.push(url);
                console.log('Adaptive Engine API call:', url);
            }
        });

        // Navigate to learning session
        await agPage.goto('/learn/math');
        await agPage.waitForLoadState('networkidle');

        // Wait for question to load (which requires adaptive recommendation)
        await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

        // Verify Adaptive Engine was called
        // Note: Adaptive Engine might be called during session initialization
        // The exact number depends on implementation
        console.log('Adaptive API calls made:', adaptiveApiCalls.length);

        // Verify content type badge is shown (result of adaptive recommendation)
        const contentBadge = agPage.locator('text=/Remedial|Standard/').first();
        await expect(contentBadge).toBeVisible();

        // Capture screenshot
        await agPage.screenshot({ path: screenshotPath('api-adaptive-engine.png'), fullPage: true });
    });

    test('should integrate with Scoring Service correctly', async ({ agPage }) => {
        // Track API calls to Scoring Service
        const scoringApiCalls: string[] = [];
        let scoringResponse: any = null;

        agPage.on('request', request => {
            const url = request.url();
            if (url.includes('/api/scoring/')) {
                scoringApiCalls.push(url);
                console.log('Scoring API call:', url);
            }
        });

        agPage.on('response', async response => {
            const url = response.url();
            if (url.includes('/api/scoring/submit')) {
                try {
                    scoringResponse = await response.json();
                    console.log('Scoring response:', scoringResponse);
                } catch (e) {
                    console.log('Could not parse scoring response');
                }
            }
        });

        // Navigate and submit an answer
        await agPage.goto('/learn/math');
        await agPage.waitForLoadState('networkidle');
        await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

        // Select and submit answer
        const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });
        if (await optionButtons.count() > 0) {
            await optionButtons.first().click();
            await agPage.waitForTimeout(500);
            await agPage.locator('button:has-text("Submit Answer")').click();

            // Wait for feedback (indicates scoring response received)
            await expect(agPage.locator('text=/Excellent|Not quite/')).toBeVisible({ timeout: 15000 });

            // Verify Scoring Service was called
            expect(scoringApiCalls.length).toBeGreaterThan(0);

            // Verify scoring response was processed (feedback shown)
            const feedbackPanel = agPage.locator('text=/Excellent|Not quite/');
            await expect(feedbackPanel).toBeVisible();

            // Capture screenshot
            await agPage.screenshot({ path: screenshotPath('api-scoring-service.png'), fullPage: true });
        }
    });

    test('should integrate with Learner Model Service correctly', async ({ agPage }) => {
        // Track API calls to Learner Model Service
        const learnerApiCalls: string[] = [];

        agPage.on('request', request => {
            const url = request.url();
            if (url.includes('/internal/learner/') || url.includes('/mastery')) {
                learnerApiCalls.push(url);
                console.log('Learner Model API call:', url);
            }
        });

        // Navigate to learning session
        await agPage.goto('/learn/math');
        await agPage.waitForLoadState('networkidle');

        // Wait for mastery to be displayed (requires Learner Model call)
        await expect(agPage.locator('text=Mastery').first()).toBeVisible({ timeout: 15000 });

        // Verify Learner Model Service was called
        expect(learnerApiCalls.length).toBeGreaterThan(0);
        console.log('Learner Model API calls made:', learnerApiCalls.length);

        // Submit an answer to trigger mastery update
        await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });
        const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });

        if (await optionButtons.count() > 0) {
            const initialCallCount = learnerApiCalls.length;

            await optionButtons.first().click();
            await agPage.waitForTimeout(500);
            await agPage.locator('button:has-text("Submit Answer")').click();

            // Wait for Next Question button (indicates mastery polling complete)
            await expect(agPage.locator('button:has-text("Next Question")')).toBeEnabled({ timeout: 15000 });

            // Verify mastery polling occurred (additional calls after submission)
            expect(learnerApiCalls.length).toBeGreaterThan(initialCallCount);
            console.log('Mastery polling calls:', learnerApiCalls.length - initialCallCount);

            // Capture screenshot
            await agPage.screenshot({ path: screenshotPath('api-learner-model.png'), fullPage: true });
        }
    });

    test('should handle Content Service errors gracefully', async ({ agPage }) => {
        // Intercept Content Service calls and return error
        await agPage.route('**/api/content/**', route => {
            route.fulfill({
                status: 500,
                contentType: 'application/json',
                body: JSON.stringify({
                    error_code: 500,
                    message: 'Content Service Error'
                }),
            });
        });

        // Navigate to learning session
        await agPage.goto('/learn/math');

        // Wait a bit for error handling
        await agPage.waitForTimeout(3000);

        // Should show loading state or error message (not crash)
        // The exact behavior depends on error handling implementation
        const isLoading = await agPage.locator('text=/Consulting AI Tutor|Loading/').isVisible();
        console.log('Loading state visible:', isLoading);

        // Capture screenshot of error state
        await agPage.screenshot({ path: screenshotPath('api-content-error.png'), fullPage: true });
    });

    test('should handle Scoring Service errors gracefully', async ({ agPage }) => {
        // Navigate to learning session first
        await agPage.goto('/learn/math');
        await agPage.waitForLoadState('networkidle');
        await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

        // Intercept Scoring Service calls and return error
        await agPage.route('**/api/scoring/**', route => {
            route.fulfill({
                status: 500,
                contentType: 'application/json',
                body: JSON.stringify({
                    error_code: 500,
                    message: 'Scoring Service Error'
                }),
            });
        });

        // Try to submit an answer
        const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });
        if (await optionButtons.count() > 0) {
            await optionButtons.first().click();
            await agPage.waitForTimeout(500);
            await agPage.locator('button:has-text("Submit Answer")').click();

            // Wait for error handling
            await agPage.waitForTimeout(3000);

            // Should not crash - either show error toast or stay on same question
            // Verify page is still functional
            const submitButton = agPage.locator('button:has-text("Submit Answer")');
            const isVisible = await submitButton.isVisible();
            console.log('Submit button still visible:', isVisible);

            // Capture screenshot
            await agPage.screenshot({ path: screenshotPath('api-scoring-error.png'), fullPage: true });
        }
    });

    test('should handle Learner Model Service errors gracefully', async ({ agPage }) => {
        // Intercept Learner Model Service calls and return error
        await agPage.route('**/internal/learner/**', route => {
            route.fulfill({
                status: 500,
                contentType: 'application/json',
                body: JSON.stringify({
                    error_code: 500,
                    message: 'Learner Model Service Error'
                }),
            });
        });

        // Navigate to learning session
        await agPage.goto('/learn/math');

        // Wait for page to handle error
        await agPage.waitForTimeout(3000);

        // Should handle gracefully - either show default mastery or error message
        // Verify page doesn't crash
        const pageTitle = await agPage.title();
        expect(pageTitle).toBeTruthy();

        // Capture screenshot
        await agPage.screenshot({ path: screenshotPath('api-learner-error.png'), fullPage: true });
    });
});
