const fs = require('fs');
const path = require('path');

// Get run ID from env or argument
const runId = process.env.PW_RUN_ID;
const artifactsDir = process.env.PW_ARTIFACTS_DIR || (runId ? `test-results/${runId}` : 'test-results');
const reportPath = path.join(process.cwd(), artifactsDir, 'results.json');

console.log(`\n📊 Aggregating test results from: ${reportPath}`);

if (!fs.existsSync(reportPath)) {
    console.error(`❌ Report file not found: ${reportPath}`);
    process.exit(1);
}

try {
    const report = JSON.parse(fs.readFileSync(reportPath, 'utf8'));
    const total = report.stats.expected + report.stats.unexpected + report.stats.flaky + report.stats.skipped;
    const passed = report.stats.expected;
    const failed = report.stats.unexpected;
    const flaky = report.stats.flaky;
    const skipped = report.stats.skipped;
    const duration = (report.stats.duration / 1000).toFixed(2);

    console.log(`\n=============================================`);
    console.log(`   E2E TEST SUMMARY`);
    console.log(`=============================================`);
    console.log(`Total Tests: ${total}`);
    console.log(`✅ Passed:   ${passed}`);
    console.log(`❌ Failed:   ${failed}`);
    console.log(`⚠️  Flaky:    ${flaky}`);
    console.log(`⏭️  Skipped:  ${skipped}`);
    console.log(`⏱️  Duration: ${duration}s`);
    console.log(`=============================================\n`);

    if (failed > 0) {
        console.log(`❌ FAILED TESTS:`);
        report.suites.forEach(suite => {
            suite.specs.forEach(spec => {
                spec.tests.forEach(test => {
                    if (test.status === 'unexpected') {
                        console.log(`  - ${spec.title}`);
                    }
                });
            });
        });
        process.exit(1);
    } else {
        console.log(`✅ All tests passed!`);
        process.exit(0);
    }

} catch (error) {
    console.error(`❌ Error parsing report: ${error.message}`);
    process.exit(1);
}