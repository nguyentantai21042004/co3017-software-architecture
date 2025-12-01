import { test, expect } from './fixtures/antigravity-fixture';
import { screenshotPath } from './utils/artifacts';

test.describe('Learning Session Flow', () => {
  test.beforeEach(async ({ agPage }) => {
    // Set up user in localStorage before navigation
    await agPage.addInitScript(() => {
      window.localStorage.setItem('user_id', 'test-user-123');
    });
  });

  test('should initialize learning session correctly', async ({ agPage }) => {
    // Navigate to learning session for math skill
    await agPage.goto('/learn/math');

    // Wait for session to load (either question or loading indicator)
    await agPage.waitForLoadState('networkidle');

    // Verify we're on the learning page (skill badge should be visible)
    const skillBadge = agPage.locator('text=math').first();
    await expect(skillBadge).toBeVisible({ timeout: 10000 });

    // Verify mastery label is displayed in header
    await expect(agPage.locator('text=Mastery')).toBeVisible();

    // Wait for question to load (h2 contains question text)
    await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

    // Capture screenshot of initialized session
    await agPage.screenshot({ path: screenshotPath('learning-session-initialized.png'), fullPage: true });
  });

  test('should display question with submit button', async ({ agPage }) => {
    await agPage.goto('/learn/math');
    await agPage.waitForLoadState('networkidle');

    // Wait for question to load
    await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

    // Verify question content exists
    const questionText = await agPage.locator('h2').first().textContent();
    expect(questionText).toBeTruthy();

    // Verify submit button exists (initially disabled)
    const submitButton = agPage.locator('button:has-text("Submit Answer")');
    await expect(submitButton).toBeVisible();

    // Capture screenshot
    await agPage.screenshot({ path: screenshotPath('learning-question-display.png'), fullPage: true });
  });

  test('should allow answer selection and submission', async ({ agPage }) => {
    await agPage.goto('/learn/math');
    await agPage.waitForLoadState('networkidle');

    // Wait for question
    await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

    // Try to find and click an option button
    const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });
    const optionCount = await optionButtons.count();

    if (optionCount > 0) {
      // Click first option
      await optionButtons.first().click();

      // Wait a bit for state update
      await agPage.waitForTimeout(500);

      // Capture screenshot with selected answer
      await agPage.screenshot({ path: screenshotPath('learning-answer-selected.png'), fullPage: true });

      // Submit button should now be enabled
      const submitButton = agPage.locator('button:has-text("Submit Answer")');
      await expect(submitButton).toBeEnabled({ timeout: 5000 });

      // Click submit
      await submitButton.click();

      // Wait for feedback (either "Excellent" or "Not quite right")
      await expect(agPage.locator('text=/Excellent|Not quite/')).toBeVisible({ timeout: 15000 });

      // Capture screenshot with feedback
      await agPage.screenshot({ path: screenshotPath('learning-feedback-shown.png'), fullPage: true });
    }
  });

  test('should show feedback panel after submission', async ({ agPage }) => {
    await agPage.goto('/learn/math');
    await agPage.waitForLoadState('networkidle');

    // Wait for question and select answer
    await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

    const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });
    if (await optionButtons.count() > 0) {
      await optionButtons.first().click();
      await agPage.waitForTimeout(500);

      await agPage.locator('button:has-text("Submit Answer")').click();

      // Wait for feedback panel
      const feedbackPanel = agPage.locator('text=/Excellent|Not quite/');
      await expect(feedbackPanel).toBeVisible({ timeout: 15000 });

      // Verify Next Question button appears (may take time for mastery update)
      const nextButton = agPage.locator('button:has-text("Next Question"), button:has-text("Updating Progress")');
      await expect(nextButton).toBeVisible({ timeout: 15000 });

      // Capture screenshot
      await agPage.screenshot({ path: screenshotPath('learning-feedback-panel.png'), fullPage: true });
    }
  });

  test('should allow continuing to next question', async ({ agPage }) => {
    await agPage.goto('/learn/math');
    await agPage.waitForLoadState('networkidle');

    // Complete first question
    await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

    const optionButtons = agPage.locator('button').filter({ hasText: /^[A-D]/ });
    if (await optionButtons.count() > 0) {
      // Get first question ID
      const firstQuestionId = await agPage.locator('text=/ID:/').first().textContent();

      // Submit answer
      await optionButtons.first().click();
      await agPage.waitForTimeout(500);
      await agPage.locator('button:has-text("Submit Answer")').click();

      // Wait for Next Question button to be enabled
      const nextButton = agPage.locator('button:has-text("Next Question")');
      await expect(nextButton).toBeEnabled({ timeout: 15000 });

      // Click Next
      await nextButton.click();

      // Wait for new question to load
      await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

      // Verify it's a different question
      const secondQuestionId = await agPage.locator('text=/ID:/').first().textContent();
      expect(secondQuestionId).not.toBe(firstQuestionId);

      // Capture screenshot
      await agPage.screenshot({ path: screenshotPath('learning-next-question.png'), fullPage: true });
    }
  });

  test('should allow exiting session and return to dashboard', async ({ agPage }) => {
    await agPage.goto('/learn/math');
    await agPage.waitForLoadState('networkidle');

    // Wait for session to load
    await expect(agPage.locator('h2').first()).toBeVisible({ timeout: 15000 });

    // Find and click exit button (X icon in header)
    const exitButton = agPage.locator('header button').first();
    await exitButton.click();

    // Should navigate back to dashboard
    await expect(agPage).toHaveURL('/dashboard', { timeout: 10000 });

    // Verify we're on dashboard
    await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

    // Capture screenshot
    await agPage.screenshot({ path: screenshotPath('learning-after-exit.png'), fullPage: true });
  });
});
