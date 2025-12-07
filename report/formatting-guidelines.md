# Report Formatting Guidelines

**Document:** Software Architecture Report - Intelligent Tutoring System (ITS)
**Last Updated:** 2025-12-07
**Purpose:** Standardize formatting and writing quality for academic submission

---

## 1. Writing Style

### 1.1 Pronoun Usage

**DO:**

- Use "nhóm" (the team) for first-person plural references
- Use passive voice for objective statements
- Use impersonal constructions

**DON'T:**

- Use "tôi" (I) outside of User Stories
- Use "chúng tôi" (we) - replace with "nhóm"
- Use "thầy cô", "người đọc", "giảng viên"

**Examples:**

```
❌ "Chúng tôi đã thiết kế hệ thống..."
✅ "Nhóm đã thiết kế hệ thống..."

❌ "Người đọc có thể thấy rằng..."
✅ "Có thể thấy rằng..."

❌ "Thầy cô sẽ đánh giá..."
✅ "Hệ thống được đánh giá..."
```

### 1.2 Capitalization Rules

**Capitalize:**

- Proper nouns (names, organizations)
- Technology names: Java, Go, PostgreSQL, RabbitMQ, Spring Boot
- Architecture patterns: Clean Architecture, Microservices
- Design principles: SOLID, SRP, OCP, LSP, ISP, DIP
- System/Service names: Content Service, Scoring Service
- Acronyms: MVP, ADR, API, REST, JWT
- Diagram titles: Figure X.X, Table X.X
- First word of sentences

**Don't Capitalize:**

- Common nouns mid-sentence
- Descriptive phrases

**Examples:**

```
❌ "The Architecture Design uses Clean Architecture"
✅ "The architecture design uses Clean Architecture"

❌ "This Module handles User Authentication"
✅ "This module handles user authentication"
```

### 1.3 Icons & Emojis

**Rule:** No Unicode icons or emojis in academic documents.

**Replacements:**
| Icon | Text Replacement |
|------|------------------|
| ✅ | "Implemented" or "Complete" |
| ❌ | "Not implemented" or "Planned" |
| ⚠️ | "Partial" or "Warning" |
| → | `$\rightarrow$` or `:` |

---

## 2. LaTeX Formatting

### 2.1 Caption Positions

**Tables:** Caption renders ABOVE table (LaTeX handles automatically)

```latex
\begin{table}[H]
    \centering
    \begin{tabular}{...}
        ...
    \end{tabular}
    \caption{Table Title}
    \label{tab:table-name}
\end{table}
```

**Figures:** Caption renders BELOW figure

```latex
\begin{figure}[H]
    \centering
    \includegraphics[width=0.8\textwidth]{images/filename.png}
    \caption{Figure Title}
    \label{fig:figure-name}
\end{figure}
```

### 2.2 Label Naming Convention

- Tables: `\label{tab:descriptive-name}`
- Figures: `\label{fig:descriptive-name}`
- Use kebab-case for label names

### 2.3 Section References

Use manual references with Vietnamese format:

```latex
Mục~3.1  % Section 3.1
Mục~4.2  % Section 4.2
```

### 2.4 Spacing

- Use `\vspace{0.5em}` to `\vspace{1em}` between sections
- Large `\vspace{}` values acceptable for page layout optimization
- Use `\indentpar` for paragraph indentation

---

## 3. File Naming

### 3.1 Image Files

**Convention:** snake_case

```
✅ adaptive_content_delivery_sequence.png
✅ erd_content_service.png
✅ system_decomposition.png
❌ Clean-Architecture-Layers.png
```

### 3.2 LaTeX Content Files

**Convention:** `<chapter>.<section>_<description>.tex`

```
1_executive_summary.tex
2.1_project_scope_and_objectives.tex
3.3_architecture_decision_records.tex
```

---

## 4. Terminology Consistency

### 4.1 Standard Terms

| Term               | Standard Form        | Avoid                        |
| ------------------ | -------------------- | ---------------------------- |
| Microservices      | Microservices        | micro-service, Micro-service |
| API                | API                  | Api, api                     |
| Database           | database (lowercase) | Database (mid-sentence)      |
| PostgreSQL         | PostgreSQL           | Postgresql, postgres         |
| RabbitMQ           | RabbitMQ             | Rabbitmq                     |
| Clean Architecture | Clean Architecture   | clean architecture           |

### 4.2 Service Names

Always use consistent naming:

- Content Service
- Scoring Service
- Learner Model Service
- Adaptive Engine
- User Management Service
- Auth Service
- API Gateway

---

## 5. Quality Checklist

Before submission, verify:

- [ ] No Unicode icons/emojis
- [ ] No improper pronouns
- [ ] Consistent capitalization
- [ ] All captions in correct positions
- [ ] Consistent file naming
- [ ] No broken cross-references
- [ ] LaTeX compiles without errors (run twice)
- [ ] Professional academic tone throughout
- [ ] All terminology consistent

---

## 6. Changelog Files

Create changelog files for significant changes:

```
report/changelog/<feature>-YYYYMMDD.md
```

Example: `report/changelog/fix-pronouns-20251207.md`
