import { test, expect } from '../fixtures/antigravity-fixture';

test('Antigravity Browser loads dashboard', async ({ agPage }) => {
    await agPage.goto('/');

    // Verify title or some element
    // Since we don't know the exact content, we'll just check the URL and maybe title
    await expect(agPage).toHaveURL(/localhost:3000/);

    // Check for the Antigravity init script
    const isAntigravity = await agPage.evaluate(() => (window as any)['__ANTIGRAVITY_ENABLED__']);
    expect(isAntigravity).toBe(true);

    console.log('Test completed successfully');
});
