import { test, expect } from './fixtures/antigravity-fixture';
import { screenshotPath } from './utils/artifacts';

test.describe('Mastery Score Persistence', () => {
    test.beforeEach(async ({ agPage }) => {
        await agPage.addInitScript(() => {
            window.localStorage.setItem('user_id', 'test-user-123');
        });
    });

    test('should persist mastery score after exiting and re-entering session', async ({ agPage }) => {
        // Step 1: Navigate to Dashboard and capture initial mastery
        console.log('Step 1: Navigate to Dashboard');
        await agPage.goto('/dashboard');
        await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

        // Get initial mastery (first skill card, which is Math)
        const initialMasteryLocator = agPage.locator('text=%').first();
        await expect(initialMasteryLocator).toBeVisible();
        const initialMasteryText = await initialMasteryLocator.textContent();
        const initialMastery = parseInt(initialMasteryText?.replace('%', '') || '0');
        console.log(`Initial Math Mastery: ${initialMastery}%`);

        await agPage.screenshot({ path: screenshotPath('mastery-persistence-1-initial-dashboard.png'), fullPage: true });

        // Step 2: Enter Math learning session
        console.log('Step 2: Enter Math Session');
        await agPage.locator('button:has-text("Continue Learning")').first().click();
        await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

        // Step 3: Answer a question
        console.log('Step 3: Answer Question');
        const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });
        const optionCount = await optionButtons.count();

        if (optionCount > 0) {
            // Select first option
            await optionButtons.first().click();
            await agPage.waitForTimeout(500);

            // Submit answer
            await agPage.locator('button:has-text("Submit Answer")').click();

            // Wait for feedback
            await expect(agPage.locator('text=/Excellent|Not quite/')).toBeVisible({ timeout: 15000 });

            // Wait for mastery update to complete (polling happens in background)
            await agPage.waitForTimeout(3000);

            await agPage.screenshot({ path: screenshotPath('mastery-persistence-2-after-answer.png'), fullPage: true });
        }

        // Step 4: Exit to Dashboard
        console.log('Step 4: Exit to Dashboard');
        await agPage.locator('header button').first().click();
        await expect(agPage).toHaveURL('/dashboard', { timeout: 10000 });
        await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

        // Step 5: Verify mastery persisted (should NOT be reset to 50)
        console.log('Step 5: Verify Mastery Persisted');
        await agPage.waitForTimeout(1000); // Wait for dashboard to fetch mastery

        const newMasteryLocator = agPage.locator('text=%').first();
        await expect(newMasteryLocator).toBeVisible();
        const newMasteryText = await newMasteryLocator.textContent();
        const newMastery = parseInt(newMasteryText?.replace('%', '') || '0');
        console.log(`Mastery after session: ${newMastery}%`);

        await agPage.screenshot({ path: screenshotPath('mastery-persistence-3-after-exit.png'), fullPage: true });

        // CRITICAL ASSERTION: Mastery should have changed from initial value
        // If it's still 50, the bug exists (score was reset)
        // If it changed (e.g., 40, 60, etc.), the bug is fixed
        expect(newMastery).not.toBe(initialMastery);
        console.log(`✓ Mastery changed from ${initialMastery}% to ${newMastery}%`);

        // Step 6: Re-enter session to verify consistency
        console.log('Step 6: Re-enter Session');
        await agPage.locator('button:has-text("Continue Learning")').first().click();
        await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

        // Verify mastery display in session header matches dashboard
        const sessionMasteryText = await agPage.locator('text=%').first().textContent();
        const sessionMastery = parseInt(sessionMasteryText?.replace('%', '') || '0');
        console.log(`Mastery in session: ${sessionMastery}%`);

        await agPage.screenshot({ path: screenshotPath('mastery-persistence-4-reenter-session.png'), fullPage: true });

        // Mastery in session should match what we saw on dashboard
        expect(sessionMastery).toBe(newMastery);
        console.log(`✓ Mastery consistent: ${sessionMastery}%`);
    });
});
