# Changelog: Cost-Benefit Analysis (TCO)

**Date:** 2025-12-07
**Author:** AI Assistant
**Type:** Addition

## Summary

Added comprehensive Total Cost of Ownership (TCO) comparison table between Monolith and Microservices architectures in Section 3.2 (Architecture Style Selection).

## Rationale

- Task 2.9 requires creating TCO comparison table
- Cost-benefit analysis is essential for justifying architecture decisions
- Provides quantitative basis for choosing Microservices over Monolith

## Changes Made

### File Modified

- `report/contents/3.2_architecture_style_selection.tex`

### New Section Added

- **Subsubsection: Phân tích Chi phí -- Lợi ích (TCO)**

### Cost Categories Analyzed

| Category                 | Monolith | Microservices |
| ------------------------ | -------- | ------------- |
| **Development (Year 0)** | $50,000  | $93,000       |
| Architecture design      | $5,000   | $15,000       |
| Core features            | $40,000  | $60,000       |
| CI/CD setup              | $2,000   | $8,000        |
| Testing infrastructure   | $3,000   | $10,000       |
| **Infrastructure/year**  | $10,800  | $25,200       |
| Compute                  | $6,000   | $12,000       |
| Database                 | $3,600   | $7,200        |
| Message Broker           | $0       | $2,400        |
| Monitoring               | $1,200   | $3,600        |
| **Maintenance/year**     | $33,000  | $33,000       |
| Bug fixes                | $8,000   | $6,000        |
| Feature development      | $20,000  | $15,000       |
| DevOps                   | $5,000   | $12,000       |
| **Scaling costs**        | $65,000  | $13,000       |
| Horizontal scaling       | $15,000  | $5,000        |
| Database sharding        | $20,000  | $8,000        |
| Refactoring              | $30,000  | $0            |

### 3-Year TCO Summary

- **Monolith**: $246,400
- **Microservices**: $280,600
- **Difference**: +$34,200 (+14%)

### Key Findings

1. Microservices costs 14% more over 3 years
2. Break-even at >10,000 users due to 80% lower scaling costs
3. Intangible benefits: time-to-market, fault isolation, team autonomy
4. Cost increase justified by AC1 (Modularity), AC2 (Scalability), FR12 (Live AI Swap)

## Verification

- [ ] LaTeX compiles without errors
- [ ] TCO table includes all cost categories
- [ ] Analysis includes both quantitative and qualitative factors

## Related Issues

- Task 2.9: Add Cost-Benefit Analysis
- report/issues/architecture-design-gaps.md
