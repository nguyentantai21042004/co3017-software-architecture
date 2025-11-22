# Client-API Integration Gap Analysis v2

**Status:** Backend `options` support added, but data format mismatches remain.

---

## 1. Resolved Issues (✅)

- **Mock Data Disabled:** Client is now configured to use real backend APIs (`USE_MOCK_DATA = false`).
- **Backend Options Support:** Content Service now has `options` field in `Question` model, Entity, and DTO.
- **Database Seeding:** `content_db` is initialized with questions containing options.
- **Hybrid UI:** Client handles both multiple-choice (if options exist) and open-ended questions.

---

## 2. Remaining Integration Problems (⚠️)

### 🔴 2.1. Option Data Format Mismatch

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

### 🔴 2.2. Submission Value Mismatch

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

## 3. Recommended Fixes

### Option 1: Client-Side Parsing (Recommended)
Update `page.tsx` to parse the option string:
1.  **Display:** Split the string `"A. Text"` -> Label `"A"`, Text `"Text"`.
2.  **Submission:** Send only the Label `"A"` to the backend.

### Option 2: Backend Data Cleanup
Update `01-init-content-db.sql` to store structured JSON:
`[{"key": "A", "text": "..."}]`
*Note: This requires changing Java DTOs and Entity definitions, which is higher risk.*

---

## 4. Action Plan

1.  **Modify `page.tsx`**:
    -   Add a helper function to parse options: `parseOption(optString) -> { key, text }`.
    -   Update UI to display `key` and `text` separately.
    -   Update `handleSubmit` to send only `key`.
2.  **Verify**:
    -   Submit a correct answer and ensure `Scoring Service` returns `correct: true`.
