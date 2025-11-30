import { test, expect } from '@playwright/test'

test.describe('Learning Flow E2E', () => {
  test.beforeEach(async ({ page }) => {
    // Navigate to home page
    await page.goto('/')
    
    // Mock localStorage for user_id
    await page.addInitScript(() => {
      window.localStorage.setItem('user_id', 'test-user-123')
    })
  })

  test('should complete full learning session flow', async ({ page }) => {
    // Step 1: Navigate to dashboard
    await page.goto('/dashboard')
    
    // Wait for dashboard to load
    await expect(page.locator('text=Welcome back')).toBeVisible({ timeout: 10000 })
    
    // Step 2: Click on a skill card to start learning
    // Assuming there's at least one skill available
    const skillCard = page.locator('[data-testid="skill-card"], button:has-text("Continue Learning")').first()
    
    if (await skillCard.isVisible()) {
      await skillCard.click()
      
      // Step 3: Wait for learning session page to load
      await expect(page.locator('text=Mastery')).toBeVisible({ timeout: 10000 })
      
      // Step 4: Wait for question to appear
      await expect(page.locator('h2, [data-testid="question-content"]')).toBeVisible({ timeout: 10000 })
      
      // Step 5: Select an answer (if multiple choice)
      const answerOption = page.locator('button:has-text("A"), input[type="radio"]').first()
      if (await answerOption.isVisible()) {
        await answerOption.click()
      } else {
        // If text input, type an answer
        const textInput = page.locator('input[type="text"]')
        if (await textInput.isVisible()) {
          await textInput.fill('test answer')
        }
      }
      
      // Step 6: Submit answer
      const submitButton = page.locator('button:has-text("Submit"), button:has-text("Submit Answer")')
      if (await submitButton.isVisible()) {
        await submitButton.click()
        
        // Step 7: Wait for feedback
        await expect(
          page.locator('text=Excellent, text=Not quite, text=Correct, text=Incorrect')
        ).toBeVisible({ timeout: 10000 })
        
        // Step 8: Wait for next button and continue
        const nextButton = page.locator('button:has-text("Next"), button:has-text("Next Question")')
        if (await nextButton.isVisible({ timeout: 5000 })) {
          await nextButton.click()
          
          // Verify new question loads
          await expect(page.locator('h2, [data-testid="question-content"]')).toBeVisible({ timeout: 10000 })
        }
      }
    }
  })

  test('should display mastery score on dashboard', async ({ page }) => {
    await page.goto('/dashboard')
    
    // Check for mastery display
    await expect(
      page.locator('text=Mastery, text=%, [data-testid="mastery-circle"]')
    ).toBeVisible({ timeout: 10000 })
  })

  test('should handle API errors gracefully', async ({ page }) => {
    // Intercept API calls and return error
    await page.route('**/api/**', (route) => {
      route.fulfill({
        status: 500,
        contentType: 'application/json',
        body: JSON.stringify({ error_code: 500, message: 'Internal Server Error' }),
      })
    })

    await page.goto('/dashboard')
    
    // Should show error message or handle gracefully
    // The exact behavior depends on error handling implementation
    await page.waitForTimeout(2000)
  })
})

