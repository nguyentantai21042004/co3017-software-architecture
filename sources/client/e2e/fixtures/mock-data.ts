export const MOCK_USER = {
    id: 'mock-user-999',
    name: 'Mock User',
    email: 'mock@example.com'
};

export const MOCK_QUESTION = {
    id: 999, // Must be number based on api.ts
    content: 'What is 2 + 2? (MOCKED)',
    options: [
        'A. 3',
        'B. 4',
        'C. 5',
        'D. 6'
    ],
    correct_answer: 'B', // api.ts uses snake_case
    skill_tag: 'math', // api.ts uses snake_case
    difficulty_level: 1, // api.ts uses snake_case
    is_remedial: false // api.ts uses snake_case
};

export const MOCK_SCORING_RESPONSE = {
    correct: true,
    score: 100,
    feedback: 'Correct! (MOCKED FEEDBACK)',
    mastery_update: { // This might be nested differently, let's check api.ts mockSubmit
        new_score: 85,
        delta: 5
    }
};

// Based on api.ts mockSubmit:
// data: {
//   correct: isCorrect,
//   score: isCorrect ? 100 : 0,
//   feedback: isCorrect ? "Correct! Well done." : "Incorrect. Let's review this concept.",
// }
// Wait, the frontend expects `masteryRes` polling for update.
// The `submitAnswer` response is just `SubmitAnswerData`.
