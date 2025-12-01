import { test as base, type Page, type BrowserContext } from '@playwright/test';

export type AntigravityFixture = {
    agPage: Page;
    agContext: BrowserContext;
};

export const test = base.extend<AntigravityFixture>({
    agContext: async ({ browser }, use) => {
        const context = await browser.newContext({
            userAgent: 'Antigravity-Browser/1.0 (Test-Automation)',
            viewport: { width: 1280, height: 720 },
            recordVideo: {
                dir: 'test-results/videos',
                size: { width: 1280, height: 720 },
            },
        });
        await use(context);
        await context.close();
    },

    agPage: async ({ agContext }, use) => {
        const page = await agContext.newPage();

        // Antigravity specific setup
        await page.addInitScript(() => {
            console.log('Antigravity Browser Initialized');
            (window as any)['__ANTIGRAVITY_ENABLED__'] = true;
        });

        // Enable console log capturing
        page.on('console', msg => {
            console.log(`[Browser Console] ${msg.type()}: ${msg.text()}`);
        });

        await use(page);

        // Cleanup if needed
        await page.close();
    },
});

export { expect } from '@playwright/test';
