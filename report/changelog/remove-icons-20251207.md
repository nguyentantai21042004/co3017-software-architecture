# Changelog: Remove Icons & Emojis

**Date:** 2025-12-07
**Task:** 1.1 - Remove Icons & Emojis
**Proposal:** standardize-report-formatting

## Summary

Removed Unicode arrow symbols (→) from LaTeX files and replaced with appropriate LaTeX equivalents or punctuation.

## Changes Made

### File: `report/contents/2.3_functional_requirements.tex`

- **Line 468-471**: Replaced `→` with `$\rightarrow$` in event chain description
  - Before: `Learner nộp bài → ScoringEngine phát sự kiện`
  - After: `Learner nộp bài $\rightarrow$ ScoringEngine phát sự kiện`

### File: `report/contents/4.3_allocation_view.tex`

- **Line 17-19**: Replaced `→` with `:` in section references
  - Before: `Mục~4.1 Module View → cho biết cấu trúc code`
  - After: `Mục~4.1 Module View: cho biết cấu trúc code`
- **Line 97-100**: Replaced `→` with `:` in network layer descriptions
  - Before: `\textbf{DMZ Layer} → Load Balancer + API Gateway`
  - After: `\textbf{DMZ Layer}: Load Balancer + API Gateway`

## Verification

- No Unicode icons (✅❌⚠️📌✔✨🔥✓×→) remain in LaTeX files
- LaTeX compilation successful

## Files Modified

1. `report/contents/2.3_functional_requirements.tex`
2. `report/contents/4.3_allocation_view.tex`
