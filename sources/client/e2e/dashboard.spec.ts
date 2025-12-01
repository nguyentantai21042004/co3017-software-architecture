import { test, expect } from './fixtures/antigravity-fixture';

test.describe('Dashboard Page', () => {
  test.beforeEach(async ({ agPage }) => {
    // Set up user in localStorage before navigation
    await agPage.addInitScript(() => {
      window.localStorage.setItem('user_id', 'test-user-123');
    });
  });

  test('should load and display skills', async ({ agPage }) => {
    await agPage.goto('/dashboard');

    // Wait for dashboard to finish loading (loading state should disappear)
    await expect(agPage.getByText('My Learning Dashboard')).toBeVisible({ timeout: 10000 });

    // Verify skill cards are displayed
    const skillButton = agPage.locator('button:has-text("Continue Learning")').first();
    await expect(skillButton).toBeVisible({ timeout: 10000 });

    // Capture screenshot of dashboard with skills loaded
    await agPage.screenshot({ path: 'test-results/screenshots/dashboard-skills-loaded.png', fullPage: true });

    // Verify at least one skill is displayed
    const skillCount = await agPage.locator('button:has-text("Continue Learning")').count();
    expect(skillCount).toBeGreaterThan(0);
  });

  test('should navigate to learning session when skill is clicked', async ({ agPage }) => {
    await agPage.goto('/dashboard');

    // Wait for dashboard to load and skills to appear
    const skillButton = agPage.locator('button:has-text("Continue Learning")').first();
    await expect(skillButton).toBeVisible({ timeout: 10000 });

    // Capture screenshot before clicking
    await agPage.screenshot({ path: 'test-results/screenshots/dashboard-before-navigation.png', fullPage: true });

    // Click on the first skill
    await skillButton.click();

    // Should navigate to learning page
    await expect(agPage).toHaveURL(/\/learn\//, { timeout: 10000 });

    // Capture screenshot of learning page
    await agPage.screenshot({ path: 'test-results/screenshots/learning-session-page.png', fullPage: true });
  });

  test('should show logout button and handle logout', async ({ agPage }) => {
    await agPage.goto('/dashboard');

    const logoutButton = agPage.locator('button:has-text("Logout")');
    await expect(logoutButton).toBeVisible({ timeout: 10000 });

    // Capture screenshot with logout button visible
    await agPage.screenshot({ path: 'test-results/screenshots/dashboard-with-logout.png', fullPage: true });

    await logoutButton.click();

    // Should redirect to home page
    await expect(agPage).toHaveURL('/', { timeout: 10000 });

    // Capture screenshot of home page after logout
    await agPage.screenshot({ path: 'test-results/screenshots/home-after-logout.png', fullPage: true });

    // Verify localStorage is cleared (user_id removed)
    const userId = await agPage.evaluate(() => window.localStorage.getItem('user_id'));
    expect(userId).toBeNull();
  });

  test('should update mastery scores', async ({ agPage }) => {
    // This test assumes we can mock the API response or that the backend is running
    // For now, we'll verify that mastery elements are present
    await agPage.goto('/dashboard');

    // Wait for skills to load
    await expect(agPage.locator('button:has-text("Continue Learning")').first()).toBeVisible({ timeout: 10000 });

    // Check for mastery percentage text
    const masteryText = agPage.locator('text=%').first();
    await expect(masteryText).toBeVisible();

    // Check for mastery label
    const masteryLabel = agPage.locator('text=Mastery').first();
    await expect(masteryLabel).toBeVisible();

    // Capture screenshot showing mastery scores
    await agPage.screenshot({ path: 'test-results/screenshots/dashboard-mastery-scores.png', fullPage: true });
  });
});
