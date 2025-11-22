# Frontend UI Build Prompt for Intelligent Tutoring System

## Mission
Build a **React-based frontend** for an Intelligent Tutoring System that connects to existing microservices backend. The focus is on **adaptive learning** with real-time mastery updates through event-driven architecture.

---

## Tech Stack Requirements

### Framework & Core
- **React** 18+ with functional components and hooks
- **React Router** for navigation
- **Axios** for API calls
- **Zustand** or **Redux Toolkit** for state management (Zustand recommended for simplicity)

### UI/Styling
- **Tailwind CSS** (recommended) or **Material-UI**
- **Framer Motion** for animations (optional but nice)
- **React Hot Toast** for notifications
- **Recharts** or **Chart.js** for analytics graphs (if implementing page 4)

### Development
- **Vite** for fast development
- **ESLint** + **Prettier** for code quality
- **TypeScript** (optional but recommended)

---

## Project Structure

```
frontend/
├── src/
│   ├── pages/
│   │   ├── HomePage.jsx          // Landing page with fake login
│   │   ├── DashboardPage.jsx     // Skill mastery overview
│   │   └── LearningSessionPage.jsx  // Main learning interface
│   ├── components/
│   │   ├── Navbar.jsx
│   │   ├── SkillCard.jsx
│   │   ├── QuestionDisplay.jsx
│   │   ├── AnswerInput.jsx
│   │   ├── FeedbackPanel.jsx
│   │   └── MasteryCircle.jsx
│   ├── services/
│   │   └── api.js               // Axios client & API functions
│   ├── store/
│   │   └── useStore.js          // Zustand store
│   ├── utils/
│   │   ├── constants.js
│   │   └── helpers.js
│   └── App.jsx
├── public/
└── package.json
```

---

## Page Specifications

### Page 1: Home/Landing Page (/)

**Purpose**: Welcome page with fake login/register buttons

**UI Components**:
1. **Hero Section**
   - Large title: "Intelligent Tutoring System"
   - Subtitle: "Adaptive Learning Powered by AI"
   - Hero image (use placeholder from Unsplash or illustrations)

2. **Call-to-Action Buttons**
   - **"Login" button** (large, blue, primary)
   - **"Register" button** (large, green, secondary)
   - Small text below: "Demo mode - No password required!"

3. **Features Section** (3 cards in a row, responsive)
   - 🎯 Adaptive Learning - Questions match your skill level
   - 📊 Real-time Progress - See your mastery grow instantly
   - 🚀 Personalized Path - AI-powered recommendations

4. **Footer**
   - "Built with Clean Architecture & Event-Driven Microservices"

**Behavior**:
```javascript
function handleLogin() {
  // Generate random user ID
  const userId = `student-${Date.now()}`;

  // Save to localStorage
  localStorage.setItem('user_id', userId);

  // Show toast notification
  toast.success(`Welcome! Your ID: ${userId}`);

  // Redirect to dashboard
  navigate('/dashboard');
}

function handleRegister() {
  // Same as login for demo purposes
  handleLogin();
}

// On component mount, check if user already logged in
useEffect(() => {
  const existingUserId = localStorage.getItem('user_id');
  if (existingUserId) {
    // Show different message for returning users
    setWelcomeMessage(`Welcome back, ${existingUserId}!`);
    setButtonText('Continue Learning');
  }
}, []);
```

**Design**: Clean, modern, gradient backgrounds (blue/green), smooth animations on button hover

---

### Page 2: Dashboard Page (/dashboard)

**Purpose**: Overview of learner's progress across all skills

**UI Components**:
1. **Header**
   - App title/logo
   - User ID display: "👤 student-1732265432123"
   - Optional: "Logout" button (clears localStorage)

2. **Skill Cards Grid** (responsive grid, 2-3 columns on desktop)
   - Each card shows:
     - Skill icon (📐 Math, 🧪 Science, etc.)
     - Skill name
     - **Mastery Circle** (circular progress indicator, 0-100%)
     - Progress bar (color-coded: red <50%, yellow 50-79%, green ≥80%)
     - Last updated timestamp
     - "Continue Learning" button

3. **Statistics Panel** (optional)
   - Total questions answered
   - Average mastery across all skills

**API Calls**:
```javascript
// On page load
useEffect(() => {
  const userId = localStorage.getItem('user_id');
  if (!userId) {
    navigate('/');
    return;
  }

  // Fetch mastery for math
  fetchMastery(userId, 'math');

  // Fetch mastery for science
  fetchMastery(userId, 'science');
}, []);

async function fetchMastery(userId, skill) {
  const response = await axios.get(
    `http://localhost:8080/internal/learner/${userId}/mastery?skill=${skill}`
  );

  // Response format:
  // {
  //   "error_code": 0,
  //   "message": "Success",
  //   "data": {
  //     "user_id": "student-123",
  //     "skill_tag": "math",
  //     "mastery_score": 75,
  //     "last_updated": "2025-11-22T10:00:00Z"
  //   }
  // }

  if (response.data.error_code === 0) {
    setMasteryData(prev => ({
      ...prev,
      [skill]: response.data.data.mastery_score
    }));
  }
}
```

**Behavior**:
- Show loading skeleton while fetching data
- Click skill card → Navigate to `/learn/:skill` (e.g., `/learn/math`)
- Auto-refresh mastery every 30 seconds (optional)

---

### Page 3: Learning Session Page (/learn/:skill) - MOST IMPORTANT

**Purpose**: Main adaptive learning interface where users answer questions

**UI Layout**:

```
┌─────────────────────────────────────────────────────────┐
│ Header: Math | Mastery: ⭕ 75% | Exit ❌                │
├─────────────────────────────────────────────────────────┤
│                                                          │
│  ┌────────────────────────────────────────────────┐    │
│  │ [Badge: Standard/Remedial]                      │    │
│  │                                                 │    │
│  │ Question: Solve: 2x + 5 = 13. What is x?       │    │
│  │                                                 │    │
│  │ Your Answer: [____________]                     │    │
│  │                                                 │    │
│  │           [Submit Answer] 🔵                     │    │
│  └────────────────────────────────────────────────┘    │
│                                                          │
│  ┌────────────────────────────────────────────────┐    │
│  │ ✅ Correct! Well done.                          │    │ ← Feedback
│  │ +100 points earned                              │    │   Panel
│  │ Mastery: 75% → 87% ↗️                           │    │   (shows after
│  │                                                 │    │    submission)
│  │           [Next Question] 🟢                     │    │
│  └────────────────────────────────────────────────┘    │
│                                                          │
└─────────────────────────────────────────────────────────┘
```

**Flow - Step by Step**:

#### 1. Page Load
```javascript
useEffect(() => {
  const userId = localStorage.getItem('user_id');
  const skill = params.skill; // From route: /learn/:skill

  // Step 1: Fetch current mastery
  const masteryResp = await axios.get(
    `http://localhost:8080/internal/learner/${userId}/mastery?skill=${skill}`
  );
  setCurrentMastery(masteryResp.data.data.mastery_score);

  // Step 2: Get next question recommendation
  const adaptiveResp = await axios.post(
    'http://localhost:8084/api/adaptive/next-lesson',
    { user_id: userId, current_skill: skill }
  );

  // Response:
  // {
  //   "error_code": 0,
  //   "data": {
  //     "next_lesson_id": 31,
  //     "reason": "Great! Your mastery is 75%. Continue with the next challenge.",
  //     "mastery_score": 75,
  //     "content_type": "standard"  // or "remedial"
  //   }
  // }

  const questionId = adaptiveResp.data.data.next_lesson_id;
  const contentType = adaptiveResp.data.data.content_type;

  // Step 3: Fetch question details
  const questionResp = await axios.get(
    `http://localhost:8081/api/content/${questionId}`
  );

  // Response:
  // {
  //   "error_code": 0,
  //   "data": {
  //     "id": 31,
  //     "content": "Solve: 2x + 5 = 13. What is x?",
  //     "correct_answer": "4",  // ⚠️ DO NOT SHOW THIS TO USER!
  //     "skill_tag": "math",
  //     "is_remedial": false
  //   }
  // }

  setQuestion(questionResp.data.data);
  setContentType(contentType); // For badge display
}, []);
```

#### 2. User Submits Answer
```javascript
async function handleSubmit() {
  // Disable button, show loading
  setIsSubmitting(true);

  const userId = localStorage.getItem('user_id');

  // Call scoring service
  const response = await axios.post(
    'http://localhost:8082/api/scoring/submit',
    {
      user_id: userId,
      question_id: currentQuestion.id,
      answer: userAnswer
    }
  );

  // Response:
  // {
  //   "error_code": 0,
  //   "data": {
  //     "correct": true,  // or false
  //     "score": 100,     // or 0
  //     "feedback": "Correct! Well done."  // or "Incorrect..."
  //   }
  // }

  const result = response.data.data;
  setIsCorrect(result.correct);
  setScore(result.score);
  setFeedback(result.feedback);

  // Show feedback panel
  setShowFeedback(true);

  // Start polling for mastery update
  await pollForMasteryUpdate(userId, skill, expectedNewMastery);

  // Enable "Next Question" button after 1.5s
  setTimeout(() => setCanContinue(true), 1500);
}
```

#### 3. Poll for Mastery Update (Important!)
```javascript
async function pollForMasteryUpdate(userId, skill, expectedScore, timeout = 10000) {
  const startTime = Date.now();
  const pollInterval = 200; // Poll every 200ms

  while (Date.now() - startTime < timeout) {
    const response = await axios.get(
      `http://localhost:8080/internal/learner/${userId}/mastery?skill=${skill}`
    );

    const newMastery = response.data.data.mastery_score;

    if (newMastery !== currentMastery) {
      // Mastery updated! Show animation
      animateMasteryChange(currentMastery, newMastery);
      setCurrentMastery(newMastery);
      return newMastery;
    }

    // Wait 200ms before next poll
    await new Promise(resolve => setTimeout(resolve, pollInterval));
  }

  // Timeout - show warning but don't block
  toast.warning('Progress saved, but display may be delayed');
}
```

#### 4. Next Question
```javascript
function handleNextQuestion() {
  // Clear feedback
  setShowFeedback(false);
  setUserAnswer('');
  setCanContinue(false);

  // Fetch new question (repeat from step 2 of page load)
  loadNextQuestion();
}
```

**UI Components to Build**:

1. **QuestionDisplay.jsx**
   - Question text display
   - Content type badge (Remedial=orange, Standard=blue)
   - Answer input field
   - Submit button

2. **FeedbackPanel.jsx**
   - Checkmark/X icon
   - Feedback message
   - Score earned display
   - Mastery change animation (75% → 87% with arrow)
   - Next question button

3. **MasteryCircle.jsx**
   - Circular progress indicator (SVG or use library like `react-circular-progressbar`)
   - Percentage text in center
   - Color coding: red <50%, yellow 50-79%, green ≥80%

**Behavior Requirements**:
- **Disable submit button** while submitting
- **Show loading spinner** on submit button during API call
- **Animate mastery change** with smooth transition
- **Error handling**: Show user-friendly error if API fails, allow retry

---

## API Client Setup (services/api.js)

```javascript
import axios from 'axios';
import toast from 'react-hot-toast';

// Base URLs for services
const API_URLS = {
  content: 'http://localhost:8081',
  scoring: 'http://localhost:8082',
  learner: 'http://localhost:8080',
  adaptive: 'http://localhost:8084/api/adaptive'
};

// Axios instance with interceptor
const apiClient = axios.create({
  timeout: 10000, // 10 seconds
  headers: {
    'Content-Type': 'application/json'
  }
});

// Response interceptor for error handling
apiClient.interceptors.response.use(
  (response) => {
    // Check for error_code in response
    if (response.data.error_code && response.data.error_code !== 0) {
      toast.error(response.data.message || 'An error occurred');
      return Promise.reject(new Error(response.data.message));
    }
    return response;
  },
  (error) => {
    // Network error
    if (!error.response) {
      toast.error('Network error. Please check your connection.');
    } else {
      toast.error(error.response.data?.message || 'Something went wrong');
    }
    return Promise.reject(error);
  }
);

// API functions
export const api = {
  // Get mastery for a user and skill
  getMastery: (userId, skill) =>
    apiClient.get(`${API_URLS.learner}/internal/learner/${userId}/mastery?skill=${skill}`),

  // Get next lesson recommendation
  getNextLesson: (userId, skill) =>
    apiClient.post(`${API_URLS.adaptive}/next-lesson`, {
      user_id: userId,
      current_skill: skill
    }),

  // Get question details
  getQuestion: (questionId) =>
    apiClient.get(`${API_URLS.content}/api/content/${questionId}`),

  // Submit answer
  submitAnswer: (userId, questionId, answer) =>
    apiClient.post(`${API_URLS.scoring}/api/scoring/submit`, {
      user_id: userId,
      question_id: questionId,
      answer: answer
    })
};
```

---

## State Management (store/useStore.js)

Using Zustand for simplicity:

```javascript
import { create } from 'zustand';

export const useStore = create((set, get) => ({
  // User state
  userId: localStorage.getItem('user_id') || null,
  setUserId: (id) => {
    localStorage.setItem('user_id', id);
    set({ userId: id });
  },

  // Mastery data
  masteryData: {
    math: 0,
    science: 0
  },
  setMastery: (skill, score) =>
    set((state) => ({
      masteryData: { ...state.masteryData, [skill]: score }
    })),

  // Current session
  currentQuestion: null,
  setCurrentQuestion: (question) => set({ currentQuestion: question }),

  currentMastery: 0,
  setCurrentMastery: (score) => set({ currentMastery: score }),

  // Loading states
  isLoading: false,
  setIsLoading: (loading) => set({ isLoading: loading })
}));
```

---

## Styling Guidelines

### Colors
- **Primary**: Blue (#3B82F6)
- **Secondary**: Green (#10B981)
- **Warning**: Yellow (#F59E0B)
- **Danger**: Red (#EF4444)
- **Success**: Green (#22C55E)

### Mastery Color Coding
```javascript
function getMasteryColor(score) {
  if (score < 50) return '#EF4444'; // Red
  if (score < 80) return '#F59E0B'; // Yellow/Orange
  return '#22C55E'; // Green
}
```

### Typography
- **Headings**: Bold, large (text-2xl to text-4xl)
- **Body**: Regular, readable (text-base)
- **Buttons**: Semi-bold, uppercase

### Animations
- Page transitions: Fade in (0.3s)
- Button hover: Scale up slightly (transform: scale(1.05))
- Mastery change: Number count-up animation
- Feedback panel: Slide up from bottom

---

## Error Handling

### Network Errors
```javascript
if (!navigator.onLine) {
  toast.error('You are offline. Please check your internet connection.');
  return;
}
```

### API Errors
```javascript
try {
  const response = await api.submitAnswer(userId, questionId, answer);
  // Handle success
} catch (error) {
  if (error.response?.status === 400) {
    toast.error('Invalid answer format. Please try again.');
  } else if (error.response?.status === 500) {
    toast.error('Server error. Our team has been notified.');
  } else {
    toast.error('Something went wrong. Please try again.');
  }
}
```

### Timeout Handling
```javascript
// If mastery doesn't update after 10 seconds
if (!masteryUpdated) {
  toast.warning('⚠️ Progress saved, but display may be delayed');
  // Still allow user to continue
  setCanContinue(true);
}
```

---

## Testing the Application

### Prerequisites
1. Ensure all backend services are running:
   - Content Service: http://localhost:8081
   - Scoring Service: http://localhost:8082
   - Learner Model: http://localhost:8080
   - Adaptive Engine: http://localhost:8084

2. PostgreSQL and RabbitMQ are running

### Test Flow
1. Open http://localhost:5173 (or your dev server)
2. Click "Login" → Should generate user_id and redirect to dashboard
3. Dashboard should show Math and Science at 0%
4. Click "Continue Learning Math"
5. Should see a remedial question (since mastery = 0%)
6. Answer correctly → Mastery should update to 50%
7. Next question should be standard type (mastery ≥ 50%)
8. Continue answering to see mastery increase

---

## Deliverables

### Must Have (Minimum Viable Demo)
✅ Page 1: Home/Landing with fake login
✅ Page 2: Dashboard showing mastery
✅ Page 3: Learning Session with adaptive questions
✅ Real-time mastery updates via polling
✅ Responsive design (mobile + desktop)
✅ Error handling and loading states

### Nice to Have
⭐ Smooth animations for mastery changes
⭐ Sound effects on correct/incorrect answers
⭐ Session statistics (questions answered, streak)
⭐ Dark mode toggle
⭐ Page 4: Progress Analytics with charts

### Not Required
❌ Page 5: Admin Dashboard (skip for demo)
❌ Real authentication (JWT, sessions)
❌ User profile management
❌ Multi-language support

---

## Quick Start Commands

```bash
# Initialize project
npm create vite@latest intelligent-tutoring-system -- --template react
cd intelligent-tutoring-system

# Install dependencies
npm install axios zustand react-router-dom
npm install react-hot-toast
npm install -D tailwindcss postcss autoprefixer
npx tailwindcss init -p

# For animations (optional)
npm install framer-motion

# For charts (if building page 4)
npm install recharts

# Run development server
npm run dev
```

---

## Final Notes

- **Focus on functionality first**, then polish UI
- **Test the event flow**: Submit answer → Wait for mastery update → See change
- **Use browser DevTools** to monitor API calls
- **Check localStorage** to see user_id persistence
- **The most important page is Page 3 (Learning Session)** - this demonstrates the entire microservices architecture!

Good luck building! 🚀
