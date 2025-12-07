# Changelog: Risk Matrix

**Date:** 2025-12-07
**Author:** AI Assistant
**Type:** Addition

## Summary

Added comprehensive Architecture Risk Matrix to Section 3.2 (Architecture Style Selection) with 10 identified risks, probability/impact analysis, and mitigation strategies.

## Rationale

- Task 2.6 requires creating a risk matrix with probability/impact analysis
- Scoring rubric identifies missing risk matrix as -1.5 points gap
- Risk management is essential for architecture documentation

## Changes Made

### File Modified

- `report/contents/3.2_architecture_style_selection.tex`

### New Section Added

- **Subsubsection: Ma trận Rủi ro Kiến trúc** (Architecture Risk Matrix)

### Risks Identified (10 total)

| ID  | Risk                                                  | Probability | Impact |
| --- | ----------------------------------------------------- | ----------- | ------ |
| R1  | Operational complexity (DevOps skills gap)            | High        | High   |
| R2  | Distributed data inconsistency (eventual consistency) | Medium      | High   |
| R3  | Single Point of Failure (Auth/Gateway)                | Medium      | High   |
| R4  | High infrastructure costs                             | High        | Medium |
| R5  | Polyglot complexity (Java + Go)                       | Medium      | Medium |
| R6  | Network latency between services                      | Medium      | Medium |
| R7  | Message broker failure (RabbitMQ)                     | Low         | High   |
| R8  | Security breach (internal attack)                     | Low         | High   |
| R9  | AI Model deployment failure                           | Medium      | High   |
| R10 | Database performance degradation                      | Medium      | Medium |

### Risk Categories

- **High Priority (High Probability + High Impact):** R1
- **Critical (Any Probability + High Impact):** R2, R3, R7, R8, R9
- **Moderate:** R4, R5, R6, R10

### Mitigation Strategies Included

- Phased deployment approach (MVP → Full Microservices)
- Saga Pattern for distributed transactions
- High Availability with Kubernetes replicas
- Blue/Green and Canary deployments for AI models
- Network Policies and VPC isolation for security
- Connection pooling and read replicas for database

### Table Format

- Used `longtable` for multi-page support
- Columns: ID, Description, Probability, Impact, Mitigation Strategy, Status
- Added legend explaining probability/impact levels

## Verification

- [ ] LaTeX compiles without errors
- [ ] Risk matrix contains at least 5-7 key risks (actual: 10 risks)
- [ ] Each risk has probability, impact, and mitigation strategy

## Related Issues

- Task 2.6: Create Risk Matrix
- report/issues/architecture-design-gaps.md (identified missing risk matrix)
- ADR-1 through ADR-10 (source of risk identification)

## Impact on Scoring

- Expected improvement: +1.5 points (from missing risk matrix gap)
