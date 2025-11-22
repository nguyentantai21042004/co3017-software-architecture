# Client-API Integration Gap Analysis

This document outlines the key integration issues between the Web Client (frontend) and backend microservices, and proposes solutions to achieve seamless and full integration.

---

## 1. Key Blocking Factors for Integration

### 🔴 1.1. Client Uses Mock Data by Default

- **Current State:**  
  In the file `sources/client/services/api.ts`, the line `const USE_MOCK_DATA = true` is present.  
  **Impact:** The client app does not communicate with the backend.
- **Required Action:**  
  Change:
  ```ts
  const USE_MOCK_DATA = true;
  ```
  to
  ```ts
  const USE_MOCK_DATA = false;
  ```
  in `sources/client/services/api.ts`.

---

### 🔴 1.2. Content Service Mismatch (“Options” Problem)

**Discovered Differences:**

|                  | Frontend Mock                      | Real Backend              | Status             |
|------------------|------------------------------------|---------------------------|--------------------|
| id               | `1` (number)                       | `26` (number)             | ✅ Match           |
| content          | `"Select..."`                      | `"What is 2 + 2?"`        | ✅ Match           |
| options          | `["A", "B"]`                       | *absent*                  | ❌ Critical error  |
| correct_answer   | `"A"`                              | `"4"`                     | ⚠️ Logic error     |
| difficulty_level | *absent*                           | `1`                       | ℹ️ Minor (unused)  |

- **Frontend expects:**  
  - For multiple-choice questions: loops through `question.options` (e.g., `["A", "B", "C", "D"]`)
  - `correct_answer` is a letter, e.g., `"A"`
- **Backend provides:**  
  - Essay/open-ended questions (no `options` field)
  - `correct_answer` is the answer value (e.g., `"4"`)

**Impact:**  
Questions are rendered but **answer buttons do not appear**. Users cannot enter or submit answers.

- **Proposed Solution (Frontend):**
  - **Hybrid Support:** Build UI to flexibly handle both types at the same time.
  - Check for `question.options`:
    - If present and non-empty: show **Multiple Choice** UI (Radio buttons).
    - If not present: show **Open-ended** UI (Text input).
  - Ensure user's submitted answer is trimmed of leading/trailing whitespace.

---

## 2. Working Services (Good News)

These services have been tested and match client expectations:

- ✅ **2.1. Scoring Service**  
  - Endpoint: `POST /api/scoring/submit`  
  - Payload: `user_id`, `question_id`, `answer`  
  - Response: `correct` (bool), `score` (int), `feedback` (string)
- ✅ **2.2. Learner Model**  
  - Endpoint: `GET /internal/learner/{id}/mastery`  
  - Response: `mastery_score`, `skill_tag`, `last_updated`
- ✅ **2.3. Adaptive Engine**
  - Endpoint: `POST /api/adaptive/next-lesson`
  - Response: `next_lesson_id`, `content_type`, `reason`

---

## 3. Action Plan

To re-enable end-to-end integration:

1. **Enable real API:** Set `USE_MOCK_DATA = false` in `api.ts`.
2. **Update UI:** Refactor `page.tsx` to support **Hybrid UI** (auto-switch between multiple choice and open-ended according to backend data).
3. **Restart all services:** Ensure all backend services (Scoring, Adaptive, etc.) are running.

---

## 4. Technical Notes & Conventions (New)

- **IDs:** Always use **integers** for both frontend and backend.
- **Dates:** Backend formats as `YYYY-MM-DD` (or `YYYY-MM-DD HH:mm:ss` if a time is needed).
- **Scores:** Backend uses integers (0–100); frontend receives as number.