# Changelog Directory

This directory contains changelog files documenting all significant changes made to the Software Architecture report.

## Purpose

Maintain a detailed audit trail of all modifications, additions, and deletions. Each changelog entry provides context, rationale, and verification for changes made.

## File Naming Convention

`{component-name}-{YYYYMMDD}.md`

Examples:
- `executive-summary-20251201.md`
- `reflection-expansion-20251201.md`
- `risk-matrix-20251201.md`

## Changelog Template

Each changelog file should follow this structure:

```markdown
# Changelog: [Component Name]

**Date:** YYYY-MM-DD
**Author:** [Name]
**Type:** [Addition|Modification|Deletion]

## Summary
Brief description of what changed.

## Rationale
Why this change was made (e.g., scoring rubric requirement, template compliance, verification discrepancy).

## Changes Made
- Detailed list of changes
- File locations
- Line numbers if applicable

## Verification
How the change was verified (compilation, review, comparison with code, etc.)

## Related Issues
Links to issue files or user questions resolved by this change.
```

## Usage

1. **Before Making Changes:** Plan what will be changed
2. **During Changes:** Document as you go
3. **After Changes:** Create changelog file with complete details
4. **Review:** Use changelogs to track progress and understand evolution
