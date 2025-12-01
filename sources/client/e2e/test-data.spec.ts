import { test, expect } from './fixtures/antigravity-fixture';
import { MOCK_QUESTION } from './fixtures/mock-data';

test.describe('Test Data and Mocking', () => {

    test.describe('Real Data Verification', () => {
        test.beforeEach(async ({ agPage }) => {
            await agPage.addInitScript(() => {
                window.localStorage.setItem('user_id', 'test-user-123');
            });
        });

        test('should verify test user exists and has valid profile', async ({ agPage }) => {
            await agPage.goto('/dashboard');
            await agPage.waitForLoadState('networkidle');

            // Verify user is logged in (dashboard visible)
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible();

            // Verify skills are loaded (implies user exists in DB and has skills)
            const skillCards = agPage.locator('button:has-text("Continue Learning")');
            await expect(skillCards.first()).toBeVisible();

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/data-real-user.png', fullPage: true });
        });

        test('should verify test questions exist in database', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');

            // Verify a question loads
            const questionText = agPage.locator('h2').first();
            await expect(questionText).toBeVisible();

            // Verify it has content (not empty)
            const text = await questionText.textContent();
            expect(text?.length).toBeGreaterThan(0);

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/data-real-question.png', fullPage: true });
        });
    });

    test.describe('Mock Mode Verification (Enhancement)', () => {
        test.beforeEach(async ({ agPage }) => {
            await agPage.addInitScript(() => {
                window.localStorage.setItem('user_id', 'mock-user-999');
            });
        });

        test('should switch to mock mode and return mocked question', async ({ agPage }) => {
            // Mock Content Service response for getQuestion
            await agPage.route('**/api/content/**', async route => {
                const url = route.request().url();
                console.log('Intercepted Content Service URL:', url);

                if (url.includes('/api/content/skills')) {
                    // Pass through skills or mock if needed
                    await route.continue();
                    return;
                }

                // Mock getQuestion response
                // The app expects { error_code: 0, message: "Success", data: { ... } }
                await route.fulfill({
                    status: 200,
                    contentType: 'application/json',
                    body: JSON.stringify({
                        error_code: 0,
                        message: "Success",
                        data: MOCK_QUESTION
                    })
                });
            });

            // Also need to mock getNextLesson to return the mock question ID
            await agPage.route('**/next-lesson', async route => {
                await route.fulfill({
                    status: 200,
                    contentType: 'application/json',
                    body: JSON.stringify({
                        error_code: 0,
                        message: "Success",
                        data: {
                            next_lesson_id: MOCK_QUESTION.id,
                            reason: "Mocked recommendation",
                            mastery_score: 50,
                            content_type: "standard"
                        }
                    })
                });
            });

            // Also mock getMastery to avoid errors
            await agPage.route('**/mastery**', async route => {
                await route.fulfill({
                    status: 200,
                    contentType: 'application/json',
                    body: JSON.stringify({
                        error_code: 0,
                        message: "Success",
                        data: {
                            user_id: 'mock-user-999',
                            skill_tag: 'math',
                            mastery_score: 50,
                            last_updated: new Date().toISOString()
                        }
                    })
                });
            });

            await agPage.goto('/learn/math');
            // Wait for question to appear
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 10000 });

            // Verify MOCKED question is displayed
            const questionText = agPage.locator('h2').first();
            await expect(questionText).toHaveText(MOCK_QUESTION.content);

            // Verify MOCKED options
            // MOCK_QUESTION.options is ['A. 3', 'B. 4', ...]
            // The UI parses this. 'B. 4' -> key 'B', text '4'
            const optionB = agPage.locator('button').filter({ hasText: '4' });
            await expect(optionB).toBeVisible();

            // Capture screenshot showing mocked data
            await agPage.screenshot({ path: 'test-results/screenshots/data-mock-mode.png', fullPage: true });
        });

        test('should mock scoring service for consistent testing', async ({ agPage }) => {
            // Setup all mocks needed for the flow

            // 1. Mastery
            await agPage.route('**/mastery**', async route => {
                await route.fulfill({
                    status: 200,
                    contentType: 'application/json',
                    body: JSON.stringify({
                        error_code: 0,
                        message: "Success",
                        data: {
                            user_id: 'mock-user-999',
                            skill_tag: 'math',
                            mastery_score: 50,
                            last_updated: new Date().toISOString()
                        }
                    })
                });
            });

            // 2. Next Lesson
            await agPage.route('**/next-lesson', async route => {
                await route.fulfill({
                    status: 200,
                    contentType: 'application/json',
                    body: JSON.stringify({
                        error_code: 0,
                        message: "Success",
                        data: {
                            next_lesson_id: MOCK_QUESTION.id,
                            reason: "Mocked recommendation",
                            mastery_score: 50,
                            content_type: "standard"
                        }
                    })
                });
            });

            // 3. Question Content
            await agPage.route('**/api/content/**', async route => {
                const url = route.request().url();
                if (url.includes('/api/content/skills')) { await route.continue(); return; }

                await route.fulfill({
                    status: 200,
                    contentType: 'application/json',
                    body: JSON.stringify({
                        error_code: 0,
                        message: "Success",
                        data: MOCK_QUESTION
                    })
                });
            });

            // 4. Scoring Service (Submit)
            await agPage.route('**/api/scoring/submit', async route => {
                console.log('Mocking Scoring Service response');
                await route.fulfill({
                    status: 200,
                    contentType: 'application/json',
                    body: JSON.stringify({
                        error_code: 0,
                        data: {
                            correct: true,
                            score: 100,
                            feedback: 'Correct! (MOCKED FEEDBACK)'
                        }
                    })
                });
            });

            await agPage.goto('/learn/math');

            // Select correct answer (B which corresponds to '4')
            await agPage.click('button:has-text("4")');
            await agPage.waitForTimeout(500);
            await agPage.click('button:has-text("Submit Answer")');

            // Verify MOCKED feedback
            await expect(agPage.getByText('Correct! (MOCKED FEEDBACK)')).toBeVisible();

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/data-mock-scoring.png', fullPage: true });
        });
    });
});
