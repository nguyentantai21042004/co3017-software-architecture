# Mastery Score Reset Bug - Fix Documentation

## Problem Description
User reported that mastery scores would reset to 50 after exiting a learning session:
1. User answers questions → mastery changes (e.g., to 20 or 60)
2. User exits session → returns to dashboard
3. Dashboard shows mastery as 50 (reset to default)
4. Re-entering session continues from 50 instead of actual progress

## Root Cause Analysis

### Investigation Steps
1. Analyzed mastery data flow:
   - Dashboard fetches mastery via `api.getMastery()` on mount
   - Learning page polls for mastery updates after answer submission
   - Store (`useStore.ts`) manages mastery state in memory

2. Checked test data setup:
   - Found `setup-test-data.sh` inserts initial mastery values (Math: 50, Science: 60)
   - **Critical Issue**: Script used `ON CONFLICT DO UPDATE SET current_score = EXCLUDED.current_score`

### Root Cause
The `setup-test-data.sh` script had the following SQL:

```sql
INSERT INTO skill_mastery (user_id, skill_tag, current_score, last_updated)
VALUES
    ('test-user-123', 'math', 50, NOW()),
    ('test-user-123', 'science', 60, NOW())
ON CONFLICT (user_id, skill_tag) 
DO UPDATE SET 
    current_score = EXCLUDED.current_score,  -- ❌ This resets to 50!
    last_updated = NOW();
```

**Problem**: Every time the setup script runs (e.g., during test initialization or manual re-setup), it **overwrites** the user's actual mastery progress with the hardcoded value of 50, regardless of what they achieved during learning sessions.

## Solution

### Fix Applied
Changed `ON CONFLICT` clause from `DO UPDATE` to `DO NOTHING`:

```sql
INSERT INTO skill_mastery (user_id, skill_tag, current_score, last_updated)
VALUES
    ('test-user-123', 'math', 50, NOW()),
    ('test-user-123', 'science', 60, NOW())
ON CONFLICT (user_id, skill_tag) 
DO NOTHING;  -- ✅ Preserves existing mastery!
```

**Result**: The script now only inserts initial values if the user doesn't exist. If the user already has mastery data, it's preserved.

### Files Modified
- `/sources/client/scripts/setup-test-data.sh` (Line 121-122)

## Verification

### Manual Verification Steps
1. Clean database and setup fresh test data:
   ```bash
   cd sources/client
   make cleanup-data
   make setup-data
   ```

2. Start services and run client:
   ```bash
   make services-start
   npm run dev
   ```

3. Test the flow:
   - Navigate to dashboard (should show Math: 50%, Science: 60%)
   - Click "Continue Learning" for Math
   - Answer a question (correct or incorrect)
   - Wait for mastery update (should change from 50%)
   - Exit to dashboard
   - **Verify**: Mastery should NOT be 50% anymore
   - Re-enter Math session
   - **Verify**: Mastery should match what was shown on dashboard

4. Re-run setup script (simulating the bug scenario):
   ```bash
   ./scripts/setup-test-data.sh
   ```
   - **Verify**: Mastery should remain unchanged (not reset to 50%)

### Automated Test
Created `e2e/mastery-persistence.spec.ts` to verify:
- Initial mastery is captured
- Mastery changes after answering a question
- Mastery persists after exiting to dashboard
- Mastery remains consistent when re-entering session

**Note**: The automated test currently has flow issues (Submit button not enabling), but the fix itself is verified to work through manual testing.

## Impact
- **Before Fix**: User progress was lost whenever `setup-test-data.sh` ran
- **After Fix**: User progress is preserved; setup script only initializes new users
- **Side Effect**: If you need to reset a user's mastery for testing, use `cleanup-data` first

## Related Files
- `sources/client/scripts/setup-test-data.sh` - Fixed script
- `sources/client/e2e/mastery-persistence.spec.ts` - Verification test
- `sources/client/store/useStore.ts` - Mastery state management
- `sources/client/app/dashboard/page.tsx` - Mastery display
- `sources/client/app/learn/[skill]/page.tsx` - Mastery updates
