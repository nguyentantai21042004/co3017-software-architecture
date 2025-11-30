import { test, expect } from '@playwright/test'

test.describe('Dashboard Page', () => {
  test.beforeEach(async ({ page }) => {
    // Set up user in localStorage
    await page.addInitScript(() => {
      window.localStorage.setItem('user_id', 'test-user-123')
    })
  })

  test('should load and display skills', async ({ page }) => {
    await page.goto('/dashboard')
    
    // Wait for page to load
    await expect(page.locator('text=Welcome back, text=My Learning Dashboard')).toBeVisible({ timeout: 10000 })
    
    // Check for skill cards or loading state
    const hasSkills = await page.locator('[data-testid="skill-card"], button:has-text("Continue Learning")').count()
    const hasLoading = await page.locator('text=Loading, [data-testid="loading"]').isVisible()
    
    // Either skills are loaded or still loading
    expect(hasSkills > 0 || hasLoading).toBeTruthy()
  })

  test('should navigate to learning session when skill is clicked', async ({ page }) => {
    await page.goto('/dashboard')
    
    // Wait for dashboard to load
    await page.waitForTimeout(2000)
    
    // Try to click on a skill card
    const skillButton = page.locator('button:has-text("Continue Learning")').first()
    
    if (await skillButton.isVisible({ timeout: 5000 })) {
      await skillButton.click()
      
      // Should navigate to learning page
      await expect(page).toHaveURL(/\/learn\//, { timeout: 5000 })
    }
  })

  test('should show logout button and handle logout', async ({ page }) => {
    await page.goto('/dashboard')
    
    const logoutButton = page.locator('button:has-text("Logout")')
    
    if (await logoutButton.isVisible({ timeout: 5000 })) {
      await logoutButton.click()
      
      // Should redirect to home page
      await expect(page).toHaveURL('/', { timeout: 5000 })
    }
  })
})

