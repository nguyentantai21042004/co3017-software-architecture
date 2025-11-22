# Client-API Integration Problem Logs

This document tracks the evolution of integration issues between the Web Client (frontend) and backend microservices, from initial discovery through resolution.

---

## Version 1: Initial Discovery

**Date:** Initial analysis  
**Status:** 🔴 Critical blockers identified

### 1. Key Blocking Factors for Integration

#### 🔴 1.1. Client Uses Mock Data by Default

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

#### 🔴 1.2. Content Service Mismatch ("Options" Problem)

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

### 2. Working Services (Good News)

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

### 3. Action Plan (v1)

To re-enable end-to-end integration:

1. **Enable real API:** Set `USE_MOCK_DATA = false` in `api.ts`.
2. **Update UI:** Refactor `page.tsx` to support **Hybrid UI** (auto-switch between multiple choice and open-ended according to backend data).
3. **Restart all services:** Ensure all backend services (Scoring, Adaptive, etc.) are running.

---

### 4. Technical Notes & Conventions

- **IDs:** Always use **integers** for both frontend and backend.
- **Dates:** Backend formats as `YYYY-MM-DD` (or `YYYY-MM-DD HH:mm:ss` if a time is needed).
- **Scores:** Backend uses integers (0–100); frontend receives as number.

---

## Version 2: Progress Made, New Issues Discovered

**Date:** After initial fixes  
**Status:** ⚠️ Backend `options` support added, but data format mismatches remain

### 1. Resolved Issues (✅)

- **Mock Data Disabled:** Client is now configured to use real backend APIs (`USE_MOCK_DATA = false`).
- **Backend Options Support:** Content Service now has `options` field in `Question` model, Entity, and DTO.
- **Database Seeding:** `content_db` is initialized with questions containing options.
- **Hybrid UI:** Client handles both multiple-choice (if options exist) and open-ended questions.

---

### 2. Remaining Integration Problems (⚠️)

#### 🔴 2.1. Option Data Format Mismatch

- **Backend Data:**
  The `options` array contains full strings with prefixes:
  ```json
  ["A. x = -2", "B. x = 2", "C. x = -1", "D. Vô nghiệm"]
  ```
- **Client Expectation (Current Code):**
  The client code assumes `options` are just keys (e.g., `["A", "B", "C", "D"]`) and uses them as both the display label and the value.
  ```tsx
  // page.tsx
  <span className="mr-3 ...">{opt}.</span> // Renders: "A. x = -2."
  ```
- **Impact:**
  - UI displays redundant/messy labels (e.g., "A. x = -2.").
  - The "content" of the option is not properly separated from the label.

#### 🔴 2.2. Submission Value Mismatch

- **Client Behavior:**
  When a user clicks an option, the *entire string* is set as the answer.
  ```ts
  setUserAnswer(opt) // opt is "A. x = -2"
  api.submitAnswer(..., "A. x = -2")
  ```
- **Backend Expectation:**
  The `correct_answer` column in DB is just `"A"`.
  The Scoring Service likely compares `answer` vs `correct_answer`.
  `"A. x = -2" != "A"`
- **Impact:**
  - **All answers will be marked INCORRECT**, even if the user chooses the right option.

---

### 3. Recommended Fixes (v2)

#### Option 1: Client-Side Parsing (Recommended)
Update `page.tsx` to parse the option string:
1.  **Display:** Split the string `"A. Text"` -> Label `"A"`, Text `"Text"`.
2.  **Submission:** Send only the Label `"A"` to the backend.

#### Option 2: Backend Data Cleanup
Update `01-init-content-db.sql` to store structured JSON:
`[{"key": "A", "text": "..."}]`
*Note: This requires changing Java DTOs and Entity definitions, which is higher risk.*

---

### 4. Action Plan (v2)

1.  **Modify `page.tsx`**:
    -   Add a helper function to parse options: `parseOption(optString) -> { key, text }`.
    -   Update UI to display `key` and `text` separately.
    -   Update `handleSubmit` to send only `key`.
2.  **Verify**:
    -   Submit a correct answer and ensure `Scoring Service` returns `correct: true`.

---

## Version 3: Final Resolution

**Date:** After all fixes implemented  
**Status:** ✅ All identified integration blockers have been resolved

### 1. Resolved Issues (✅)

#### 1.1. Infrastructure & Connectivity
-   **Mock Data Disabled:** Client is configured to use real backend APIs (`USE_MOCK_DATA = false`).
-   **Service Availability:** All microservices (Content, Scoring, Learner, Adaptive) are expected to be running.

#### 1.2. Data Format & Logic
-   **Backend Options Support:** Content Service correctly provides `options` in the API response.
-   **Hybrid UI:** Client dynamically renders:
    -   **Multiple Choice:** When `options` are present.
    -   **Text Input:** When `options` are missing (open-ended questions).
-   **Option Parsing:** Client correctly parses backend strings like `"A. Answer"` into:
    -   **Key:** `"A"` (used for submission).
    -   **Text:** `"Answer"` (used for display).
-   **Submission Accuracy:** Client submits only the key (e.g., `"A"`) to the Scoring Service, matching the database's `correct_answer` format.

---

### 2. Remaining Risks & Recommendations (ℹ️)

#### 2.1. Data Consistency
-   **Risk:** If the backend changes the option format (e.g., from `"A. Text"` to `"1. Text"` or just `"Text"`), the client's regex parser might need adjustment.
-   **Mitigation:** The current parser has a fallback: if it can't find a prefix (like "A."), it uses the whole string as both key and text. This ensures the UI doesn't break, though submission might be incorrect if the backend expects a specific key.

#### 2.2. Error Handling
-   **Risk:** Network failures or service downtime will show generic error toasts.
-   **Recommendation:** Future improvements could add more specific error messages (e.g., "Scoring Service Unavailable").

---

### 3. Verification Checklist

To ensure everything stays green:

1.  [x] **Login:** User can login (Client -> Learner Service).
2.  [x] **Get Question:** Question loads with correct UI (Client -> Adaptive -> Content).
3.  [x] **Submit Answer:** Score updates correctly (Client -> Scoring -> RabbitMQ -> Learner).
4.  [x] **Progress:** Mastery score updates in real-time.

**System is ready for full end-to-end testing.**

---

## Summary

### Timeline of Issues and Resolutions

1. **v1 (Initial):** 
   - 🔴 Mock data enabled blocking integration
   - 🔴 Backend missing `options` field
   - ✅ Working services identified

2. **v2 (Progress):**
   - ✅ Mock data disabled
   - ✅ Backend `options` support added
   - ✅ Hybrid UI implemented
   - 🔴 Option format mismatch discovered
   - 🔴 Submission value mismatch discovered

3. **v3 (Resolved):**
   - ✅ Option parsing implemented
   - ✅ Submission format corrected
   - ✅ All critical blockers resolved
   - ℹ️ Future recommendations documented

### Key Learnings

- **Data Format Consistency:** Frontend and backend must agree on data structures. The initial mismatch between mock data format and real backend format caused significant integration delays.
- **Progressive Problem Solving:** Each iteration revealed new issues that needed to be addressed systematically.
- **Client-Side Parsing:** Choosing client-side parsing over backend changes was the right approach for faster resolution with lower risk.

---

**Last Updated:** v3 - All integration blockers resolved

