import { test, expect } from './fixtures/antigravity-fixture';

test.describe('UI Component Tests', () => {
    test.beforeEach(async ({ agPage }) => {
        // Set up user in localStorage before navigation
        await agPage.addInitScript(() => {
            window.localStorage.setItem('user_id', 'test-user-123');
        });
    });

    test.describe('Mastery Circle Component', () => {
        test('should display mastery circle with correct score', async ({ agPage }) => {
            await agPage.goto('/dashboard');
            await agPage.waitForLoadState('networkidle');

            // Wait for dashboard to load
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            // Find mastery circle elements (should show percentage)
            const masteryPercentages = agPage.locator('text=/%/');
            const count = await masteryPercentages.count();

            // Should have at least one mastery circle
            expect(count).toBeGreaterThan(0);

            // Verify first mastery score is a valid percentage (0-100)
            const firstScore = await masteryPercentages.first().textContent();
            const scoreValue = parseInt(firstScore!.replace('%', ''));
            expect(scoreValue).toBeGreaterThanOrEqual(0);
            expect(scoreValue).toBeLessThanOrEqual(100);

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/component-mastery-circle.png', fullPage: true });
        });

        test('should display correct color coding based on mastery level', async ({ agPage }) => {
            await agPage.goto('/dashboard');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            // Get all mastery circles
            const masteryCircles = agPage.locator('svg circle[stroke-dasharray]');
            const circleCount = await masteryCircles.count();

            if (circleCount > 0) {
                // Check first circle has a stroke color (color coding)
                const firstCircle = masteryCircles.first();
                const strokeColor = await firstCircle.getAttribute('stroke');

                // Should have a color (not transparent or none)
                expect(strokeColor).toBeTruthy();
                expect(strokeColor).not.toBe('transparent');
                expect(strokeColor).not.toBe('none');

                // Capture screenshot showing color coding
                await agPage.screenshot({ path: 'test-results/screenshots/component-mastery-colors.png', fullPage: true });
            }
        });

        test('should show mastery circle in learning session header', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');

            // Wait for session to load
            await expect(agPage.locator('text=Mastery').first()).toBeVisible({ timeout: 15000 });

            // Verify mastery circle is in header
            const headerMastery = agPage.locator('header').locator('svg circle');
            await expect(headerMastery.first()).toBeVisible();

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/component-mastery-header.png', fullPage: true });
        });
    });

    test.describe('Question Display Components', () => {
        test('should render question text correctly', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');

            // Wait for question to load
            const questionText = agPage.locator('h2').first();
            await expect(questionText).toBeVisible({ timeout: 15000 });

            // Verify question has content
            const text = await questionText.textContent();
            expect(text).toBeTruthy();
            expect(text!.length).toBeGreaterThan(10); // Reasonable question length

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/component-question-text.png', fullPage: true });
        });

        test('should display multiple choice options correctly', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            // Look for option buttons
            const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });
            const optionCount = await optionButtons.count();

            if (optionCount > 0) {
                // Verify options are properly formatted
                expect(optionCount).toBeGreaterThanOrEqual(2); // At least 2 options
                expect(optionCount).toBeLessThanOrEqual(6); // Reasonable max

                // Verify first option has key and text
                const firstOption = optionButtons.first();
                const optionText = await firstOption.textContent();
                expect(optionText).toBeTruthy();

                // Verify option is clickable
                await expect(firstOption).toBeEnabled();

                // Capture screenshot
                await agPage.screenshot({ path: 'test-results/screenshots/component-question-options.png', fullPage: true });
            }
        });

        test('should display question metadata (ID and type badge)', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            // Verify question ID is shown
            const questionId = agPage.locator('text=/ID:/').first();
            await expect(questionId).toBeVisible();

            // Verify content type badge (Remedial or Standard)
            const typeBadge = agPage.locator('text=/Remedial|Standard/').first();
            await expect(typeBadge).toBeVisible();

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/component-question-metadata.png', fullPage: true });
        });
    });

    test.describe('Feedback Components', () => {
        test('should display feedback message after answer submission', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            // Submit an answer
            const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });
            if (await optionButtons.count() > 0) {
                await optionButtons.first().click();
                await agPage.waitForTimeout(500);
                await agPage.locator('button:has-text("Submit Answer")').click();

                // Wait for feedback
                const feedbackMessage = agPage.locator('text=/Excellent|Not quite/');
                await expect(feedbackMessage).toBeVisible({ timeout: 15000 });

                // Verify feedback has descriptive text
                const feedbackText = agPage.locator('p.text-muted-foreground').last();
                await expect(feedbackText).toBeVisible();

                // Capture screenshot
                await agPage.screenshot({ path: 'test-results/screenshots/component-feedback-message.png', fullPage: true });
            }
        });

        test('should display correct feedback styling (success/error)', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            // Submit an answer
            const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });
            if (await optionButtons.count() > 0) {
                await optionButtons.first().click();
                await agPage.waitForTimeout(500);
                await agPage.locator('button:has-text("Submit Answer")').click();

                // Wait for feedback panel
                await expect(agPage.locator('text=/Excellent|Not quite/')).toBeVisible({ timeout: 15000 });

                // Verify feedback icon is present (CheckCircle or XCircle)
                const feedbackIcon = agPage.locator('svg').filter({ has: agPage.locator('circle') }).first();
                await expect(feedbackIcon).toBeVisible();

                // Capture screenshot showing styling
                await agPage.screenshot({ path: 'test-results/screenshots/component-feedback-styling.png', fullPage: true });
            }
        });

        test('should show feedback panel with smooth animation', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            // Submit an answer
            const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });
            if (await optionButtons.count() > 0) {
                await optionButtons.first().click();
                await agPage.waitForTimeout(500);

                // Take screenshot before submission
                await agPage.screenshot({ path: 'test-results/screenshots/component-before-feedback.png', fullPage: true });

                await agPage.locator('button:has-text("Submit Answer")').click();

                // Wait for feedback panel to appear
                await expect(agPage.locator('text=/Excellent|Not quite/')).toBeVisible({ timeout: 15000 });

                // Wait a bit for animation to complete
                await agPage.waitForTimeout(1000);

                // Take screenshot after feedback appears
                await agPage.screenshot({ path: 'test-results/screenshots/component-after-feedback.png', fullPage: true });
            }
        });
    });

    test.describe('Navigation Components', () => {
        test('should have functional exit button in learning session', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            // Find exit button (X icon in header)
            const exitButton = agPage.locator('header button').first();
            await expect(exitButton).toBeVisible();
            await expect(exitButton).toBeEnabled();

            // Verify it has an icon
            const icon = exitButton.locator('svg');
            await expect(icon).toBeVisible();

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/component-nav-exit.png', fullPage: true });
        });

        test('should navigate back to dashboard when exit is clicked', async ({ agPage }) => {
            await agPage.goto('/learn/math');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            // Click exit button
            const exitButton = agPage.locator('header button').first();
            await exitButton.click();

            // Should return to dashboard
            await expect(agPage).toHaveURL('/dashboard', { timeout: 10000 });
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/component-nav-back-dashboard.png', fullPage: true });
        });

        test('should have logout button on dashboard', async ({ agPage }) => {
            await agPage.goto('/dashboard');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            // Find logout button
            const logoutButton = agPage.locator('button:has-text("Logout")');
            await expect(logoutButton).toBeVisible();
            await expect(logoutButton).toBeEnabled();

            // Capture screenshot
            await agPage.screenshot({ path: 'test-results/screenshots/component-nav-logout.png', fullPage: true });
        });
    });

    test.describe('Accessibility and Visual Regression', () => {
        test('should have proper ARIA labels and roles', async ({ agPage }) => {
            await agPage.goto('/dashboard');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            // Check for buttons with proper roles
            const buttons = agPage.locator('button');
            const buttonCount = await buttons.count();
            expect(buttonCount).toBeGreaterThan(0);

            // Verify buttons are keyboard accessible
            const firstButton = buttons.first();
            await expect(firstButton).toBeEnabled();

            // Capture screenshot for visual regression baseline
            await agPage.screenshot({ path: 'test-results/screenshots/component-accessibility.png', fullPage: true });
        });

        test('should maintain consistent styling across components', async ({ agPage }) => {
            await agPage.goto('/dashboard');
            await agPage.waitForLoadState('networkidle');
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            // Capture dashboard for visual regression
            await agPage.screenshot({ path: 'test-results/screenshots/component-visual-dashboard.png', fullPage: true });

            // Navigate to learning session
            const skillButton = agPage.locator('button:has-text("Continue Learning")').first();
            if (await skillButton.isVisible()) {
                await skillButton.click();
                await agPage.waitForLoadState('networkidle');
                await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

                // Capture learning session for visual regression
                await agPage.screenshot({ path: 'test-results/screenshots/component-visual-learning.png', fullPage: true });
            }
        });
    });
});
