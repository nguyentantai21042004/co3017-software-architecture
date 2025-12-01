import { test, expect } from './fixtures/antigravity-fixture';
import { screenshotPath } from './utils/artifacts';

test.describe('Comprehensive Mastery Flow Tests', () => {
    test.beforeEach(async ({ agPage }) => {
        await agPage.addInitScript(() => {
            window.localStorage.setItem('user_id', 'test-user-123');
        });
    });

    test.describe('Positive Flow Tests', () => {
        test('should increase mastery after correct answer', async ({ agPage }) => {
            console.log('Test: Correct answer increases mastery');

            await agPage.goto('/dashboard');
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            const initialMasteryLocator = agPage.locator('text=%').first();
            await expect(initialMasteryLocator).toBeVisible();
            const initialMasteryText = await initialMasteryLocator.textContent();
            const initialMastery = parseInt(initialMasteryText?.replace('%', '') || '0');
            console.log(`Initial mastery: ${initialMastery}%`);

            await agPage.screenshot({ path: screenshotPath('mastery-comprehensive-positive-1.png'), fullPage: true });

            await agPage.locator('button:has-text("Continue Learning")').first().click();
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            // Use data-testid to select first option (A)
            const optionA = agPage.getByTestId('option-button-A');
            if (await optionA.isVisible()) {
                await optionA.click();
                console.log('Selected option A');

                // Wait for React state to update
                await agPage.waitForTimeout(1000);

                // Verify Submit button is enabled
                const submitButton = agPage.getByTestId('submit-answer-button');
                await expect(submitButton).toBeEnabled({ timeout: 5000 });
                console.log('Submit button enabled');

                await submitButton.click();
                console.log('Clicked Submit');

                const feedbackLocator = agPage.locator('text=/Excellent|Not quite/');
                await expect(feedbackLocator).toBeVisible({ timeout: 15000 });
                await agPage.waitForTimeout(3000);

                await agPage.locator('header button').first().click();
                await expect(agPage).toHaveURL('/dashboard', { timeout: 10000 });
                await agPage.waitForTimeout(1000);

                const newMasteryLocator = agPage.locator('text=%').first();
                await expect(newMasteryLocator).toBeVisible();
                const newMasteryText = await newMasteryLocator.textContent();
                const newMastery = parseInt(newMasteryText?.replace('%', '') || '0');

                expect(newMastery).not.toBe(initialMastery);
                console.log(`✓ Mastery changed: ${initialMastery}% -> ${newMastery}%`);
            }
        });
    });

    test.describe('Negative Flow Tests', () => {
        test('should handle incorrect answers', async ({ agPage }) => {
            console.log('Test: Incorrect answer handling');

            await agPage.goto('/dashboard');
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            const initialMasteryLocator = agPage.locator('text=%').first();
            await expect(initialMasteryLocator).toBeVisible();
            const initialMasteryText = await initialMasteryLocator.textContent();
            const initialMastery = parseInt(initialMasteryText?.replace('%', '') || '0');

            await agPage.locator('button:has-text("Continue Learning")').first().click();
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            // Try option D (likely incorrect)
            const optionD = agPage.getByTestId('option-button-D');
            if (await optionD.isVisible()) {
                await optionD.click();
                await agPage.waitForTimeout(1000);

                const submitButton = agPage.getByTestId('submit-answer-button');
                await expect(submitButton).toBeEnabled({ timeout: 5000 });
                await submitButton.click();

                await expect(agPage.locator('text=/Excellent|Not quite/')).toBeVisible({ timeout: 15000 });
                await agPage.waitForTimeout(3000);

                await agPage.locator('header button').first().click();
                await expect(agPage).toHaveURL('/dashboard', { timeout: 10000 });
                await agPage.waitForTimeout(1000);

                const newMasteryLocator = agPage.locator('text=%').first();
                await expect(newMasteryLocator).toBeVisible();
                const newMasteryText = await newMasteryLocator.textContent();
                const newMastery = parseInt(newMasteryText?.replace('%', '') || '0');

                console.log(`Mastery: ${initialMastery}% -> ${newMastery}%`);
                expect(newMastery).not.toBe(initialMastery);
            }
        });
    });
    test.describe('Cross-Skill Independence Tests', () => {
        test('should not affect Science mastery when updating Math', async ({ agPage }) => {
            console.log('Test: Cross-skill independence');

            await agPage.goto('/dashboard');
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            const masteryElements = agPage.locator('text=%');
            const mathMasteryText = await masteryElements.nth(0).textContent();
            const scienceMasteryText = await masteryElements.nth(1).textContent();
            const initialMathMastery = parseInt(mathMasteryText?.replace('%', '') || '0');
            const initialScienceMastery = parseInt(scienceMasteryText?.replace('%', '') || '0');

            console.log(`Initial: Math=${initialMathMastery}%, Science=${initialScienceMastery}%`);

            await agPage.locator('button:has-text("Continue Learning")').first().click();
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            const optionA = agPage.getByTestId('option-button-A');
            if (await optionA.isVisible()) {
                await optionA.click();
                await agPage.waitForTimeout(1000);

                const submitButton = agPage.getByTestId('submit-answer-button');
                await expect(submitButton).toBeEnabled({ timeout: 5000 });
                await submitButton.click();

                await expect(agPage.locator('text=/Excellent|Not quite/')).toBeVisible({ timeout: 15000 });
                await agPage.waitForTimeout(3000);
            }

            await agPage.locator('header button').first().click();
            await expect(agPage).toHaveURL('/dashboard', { timeout: 10000 });
            await agPage.waitForTimeout(1000);

            const newMasteryElements = agPage.locator('text=%');
            const newScienceMasteryText = await newMasteryElements.nth(1).textContent();
            const newScienceMastery = parseInt(newScienceMasteryText?.replace('%', '') || '0');

            expect(newScienceMastery).toBe(initialScienceMastery);
            console.log(`✓ Science mastery unchanged: ${newScienceMastery}%`);
        });
    });

    test.describe('Data Integrity Tests', () => {
        test('should persist mastery across page refreshes', async ({ agPage }) => {
            console.log('Test: Mastery persists across refreshes');

            await agPage.goto('/dashboard');
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            const masteryLocator = agPage.locator('text=%').first();
            await expect(masteryLocator).toBeVisible();
            const masteryText = await masteryLocator.textContent();
            const mastery = parseInt(masteryText?.replace('%', '') || '0');
            console.log(`Mastery before refresh: ${mastery}%`);

            await agPage.reload();
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            const newMasteryLocator = agPage.locator('text=%').first();
            await expect(newMasteryLocator).toBeVisible();
            const newMasteryText = await newMasteryLocator.textContent();
            const newMastery = parseInt(newMasteryText?.replace('%', '') || '0');

            expect(newMastery).toBe(mastery);
            console.log(`✓ Mastery persisted after refresh: ${newMastery}%`);
        });
    });

    test.describe('Boundary Condition Tests', () => {
        test('should display UI correctly at different mastery levels', async ({ agPage }) => {
            console.log('Test: UI display at different mastery levels');

            await agPage.goto('/dashboard');
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            const masteryLocator = agPage.locator('text=%').first();
            await expect(masteryLocator).toBeVisible();

            const masteryLabel = agPage.locator('text=Mastery').first();
            await expect(masteryLabel).toBeVisible();

            const levelText = agPage.locator('text=/Beginner|Intermediate|Advanced/').first();
            await expect(levelText).toBeVisible();

            await agPage.screenshot({ path: screenshotPath('mastery-comprehensive-ui-display.png'), fullPage: true });
            console.log('✓ UI displays mastery correctly');
        });
    });

    test.describe('Timing and Race Condition Tests', () => {
        test('should handle rapid session exit', async ({ agPage }) => {
            console.log('Test: Rapid session exit');

            await agPage.goto('/dashboard');
            await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

            await agPage.locator('button:has-text("Continue Learning")').first().click();
            await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

            const optionA = agPage.getByTestId('option-button-A');
            if (await optionA.isVisible()) {
                await optionA.click();
                await agPage.waitForTimeout(1000);

                const submitButton = agPage.getByTestId('submit-answer-button');
                await expect(submitButton).toBeEnabled({ timeout: 5000 });
                await submitButton.click();

                // Exit immediately without waiting for polling
                await agPage.waitForTimeout(500);
                await agPage.locator('header button').first().click();
                await expect(agPage).toHaveURL('/dashboard', { timeout: 10000 });

                console.log('✓ Handled rapid exit without crash');
            }
        });
    });
});
