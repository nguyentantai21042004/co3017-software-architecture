# Client Web Application

**Ứng dụng web frontend cho Hệ thống Gia sư Thông minh**

---

## Mục lục

- [Tổng quan](#tổng-quan)
- [Công nghệ](#công-nghệ)
- [Kiến trúc](#kiến-trúc)
- [Cấu trúc Thư mục](#cấu-trúc-thư-mục)
- [Pages](#pages)
- [Components](#components)
- [State Management](#state-management)
- [API Integration](#api-integration)
- [Cấu hình](#cấu-hình)
- [Phát triển Local](#phát-triển-local)
- [Testing](#testing)
- [Deployment](#deployment)

---

## Tổng quan

Client là ứng dụng web frontend của hệ thống ITS (Intelligent Tutoring System - Hệ thống Gia sư Thông minh). Ứng dụng cung cấp giao diện người dùng cho học sinh để:

- **Đăng nhập/Đăng ký**: Xác thực và quản lý tài khoản học sinh
- **Dashboard**: Xem tổng quan tiến độ học tập và mức độ thành thạo (Mastery)
- **Học tập thích ứng**: Nhận câu hỏi được cá nhân hóa dựa trên trình độ
- **Làm bài kiểm tra**: Trả lời câu hỏi và nhận phản hồi ngay lập tức
- **Theo dõi tiến độ**: Xem biểu đồ và thống kê học tập

Client giao tiếp với các backend microservices thông qua REST API để cung cấp trải nghiệm học tập liền mạch và cá nhân hóa.

---

## Công nghệ

### Core Technologies

- **Next.js 15**: React framework với App Router
- **React 19**: UI library
- **TypeScript**: Type-safe JavaScript
- **Tailwind CSS 4**: Utility-first CSS framework
- **Node.js 18.17+** hoặc **20.3+**: JavaScript runtime

### UI Libraries

- **Radix UI**: Headless UI components (Dialog, Dropdown, Toast, etc.)
- **Lucide React**: Icon library
- **Framer Motion**: Animation library
- **Recharts**: Charting library cho biểu đồ tiến độ

### State Management & Data Fetching

- **Zustand**: Lightweight state management
- **Axios**: HTTP client cho API calls
- **React Hook Form**: Form handling với Zod validation

### Development & Testing

- **Jest**: Unit testing framework
- **React Testing Library**: Component testing
- **Playwright**: End-to-end (E2E) testing
- **ESLint**: Code linting

### Build & Deployment

- **Docker**: Containerization
- **Standalone Output**: Optimized production build

---

## Kiến trúc

### Next.js App Router

Client sử dụng **Next.js App Router** (thư mục `app/`) - hệ thống routing mới của Next.js 15 với các tính năng:

- **File-based Routing**: Mỗi thư mục trong `app/` tương ứng với một route
- **Server Components**: Mặc định components render trên server
- **Client Components**: Sử dụng `"use client"` directive cho interactive components
- **Layouts**: Shared layouts giữa các pages
- **Loading States**: Built-in loading UI

```
┌─────────────────────────────────────────────────────────┐
│                    App Router Structure                 │
├─────────────────────────────────────────────────────────┤
│  app/                                                   │
│  ├── layout.tsx          # Root layout (HTML, fonts)    │
│  ├── page.tsx            # Homepage (/)                 │
│  ├── globals.css         # Global styles                │
│  ├── dashboard/                                         │
│  │   └── page.tsx        # Dashboard (/dashboard)       │
│  ├── learn/                                             │
│  │   └── [skill]/        # Dynamic route                │
│  │       └── page.tsx    # Learning page (/learn/:skill)│
│  ├── quizzes/                                           │
│  │   ├── page.tsx        # Quiz list (/quizzes)         │
│  │   └── [quizId]/       # Dynamic route                │
│  │       └── page.tsx    # Quiz detail (/quizzes/:id)   │
│  ├── login/                                             │
│  │   └── page.tsx        # Login (/login)               │
│  ├── register/                                          │
│  │   └── page.tsx        # Register (/register)         │
│  └── profile/                                           │
│      └── page.tsx        # Profile (/profile)           │
└─────────────────────────────────────────────────────────┘
```

### Component Architecture

```
┌─────────────────────────────────────────────────────────┐
│                      Pages (app/)                       │
│  - Route handlers                                       │
│  - Data fetching                                        │
│  - Layout composition                                   │
└─────────────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────────────┐
│                Feature Components                       │
│  (components/learning/, components/quiz/)               │
│  - Business logic                                       │
│  - Feature-specific UI                                  │
└─────────────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────────────┐
│                   UI Components                         │
│  (components/ui/)                                       │
│  - Reusable primitives                                  │
│  - Radix UI wrappers                                    │
│  - Styled with Tailwind                                 │
└─────────────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────────────┐
│                   Services & Store                      │
│  (services/, store/)                                    │
│  - API integration                                      │
│  - Global state management                              │
└─────────────────────────────────────────────────────────┘
```

---

## Cấu trúc Thư mục

```
sources/client/
├── app/                           # Next.js App Router
│   ├── layout.tsx                 # Root layout
│   ├── page.tsx                   # Homepage
│   ├── globals.css                # Global styles
│   ├── dashboard/                 # Dashboard page
│   ├── learn/                     # Learning pages
│   │   └── [skill]/               # Dynamic skill route
│   ├── quizzes/                   # Quiz pages
│   │   └── [quizId]/              # Dynamic quiz route
│   ├── login/                     # Login page
│   ├── register/                  # Register page
│   └── profile/                   # Profile page
│
├── components/                    # React components
│   ├── ui/                        # Reusable UI primitives
│   │   ├── button.tsx             # Button component
│   │   ├── card.tsx               # Card component
│   │   ├── dialog.tsx             # Dialog/Modal
│   │   ├── progress.tsx           # Progress bar
│   │   ├── toast.tsx              # Toast notifications
│   │   └── ...                    # 50+ UI components
│   ├── learning/                  # Learning feature components
│   │   └── mastery-circle.tsx     # Mastery visualization
│   ├── quiz/                      # Quiz feature components
│   │   ├── quiz-interface.tsx     # Quiz UI
│   │   ├── quiz-list.tsx          # Quiz listing
│   │   ├── quiz-timer.tsx         # Timer component
│   │   ├── results-dialog.tsx     # Results modal
│   │   └── start-quiz-dialog.tsx  # Start quiz modal
│   ├── navbar.tsx                 # Navigation bar
│   └── theme-provider.tsx         # Theme context
│
├── services/                      # API integration
│   └── api.ts                     # API client & endpoints
│
├── store/                         # State management
│   └── useStore.ts                # Zustand store
│
├── lib/                           # Utilities
│   ├── api-helpers.ts             # API helper functions
│   ├── env-config.ts              # Environment configuration
│   └── utils.ts                   # General utilities
│
├── hooks/                         # Custom React hooks
│   ├── use-mobile.ts              # Mobile detection
│   └── use-toast.ts               # Toast hook
│
├── types/                         # TypeScript types
│   └── api.ts                     # API response types
│
├── e2e/                           # Playwright E2E tests
│   ├── dashboard.spec.ts          # Dashboard tests
│   ├── learning-flow.spec.ts      # Learning flow tests
│   ├── api-integration.spec.ts    # API integration tests
│   └── ...                        # More E2E tests
│
├── __tests__/                     # Jest unit tests
│   ├── components/                # Component tests
│   ├── lib/                       # Utility tests
│   └── services/                  # Service tests
│
├── public/                        # Static assets
│   ├── icon.svg                   # App icon
│   └── placeholder.svg            # Placeholder images
│
├── scripts/                       # Utility scripts
│   ├── run-e2e-local.sh           # Run E2E tests locally
│   ├── start-services.sh          # Start backend services
│   └── ...                        # More scripts
│
├── package.json                   # Dependencies & scripts
├── next.config.mjs                # Next.js configuration
├── tailwind.config.ts             # Tailwind configuration
├── tsconfig.json                  # TypeScript configuration
├── jest.config.js                 # Jest configuration
├── playwright.config.ts           # Playwright configuration
└── Dockerfile                     # Docker image
```

---

## Pages

### 1. Homepage (`/`)

**File:** `app/page.tsx`

Trang chủ giới thiệu hệ thống ITS với:

- Hero section với mô tả hệ thống
- Features section (Adaptive Learning, Real-time Progress, Personalized Path)
- Login/Register buttons cho người dùng mới
- Continue Learning button cho người dùng đã đăng nhập

**Key Features:**

- Kiểm tra `localStorage` để xác định returning user
- Demo mode - không cần password
- Redirect đến Dashboard sau khi login

### 2. Dashboard (`/dashboard`)

**File:** `app/dashboard/page.tsx`

Bảng điều khiển chính hiển thị:

- Danh sách các kỹ năng (Skills) có sẵn
- Mức độ thành thạo (Mastery) cho mỗi kỹ năng
- Circular progress indicators
- "Continue Learning" buttons

**Data Flow:**

```
1. Fetch available skills từ Content Service
2. Fetch mastery scores từ Learner Model Service
3. Display skill cards với mastery visualization
4. Navigate đến /learn/:skill khi click "Continue Learning"
```

### 3. Learning Page (`/learn/[skill]`)

**File:** `app/learn/[skill]/page.tsx`

Trang học tập thích ứng với:

- Hiển thị câu hỏi được đề xuất bởi Adaptive Engine
- Multiple choice answers
- Submit và nhận feedback ngay lập tức
- Cập nhật mastery sau mỗi câu trả lời
- Remedial questions cho học sinh cần ôn tập

**Adaptive Learning Flow:**

```
1. Call Adaptive Engine để lấy next lesson recommendation
2. Fetch question từ Content Service
3. Display question và options
4. Submit answer đến Scoring Service
5. Show feedback (correct/incorrect)
6. Update mastery display
7. Repeat với câu hỏi tiếp theo
```

### 4. Quizzes (`/quizzes`)

**File:** `app/quizzes/page.tsx`

Danh sách các bài kiểm tra với:

- Quiz cards với thông tin (skill, difficulty, question count)
- Start quiz dialog
- Quiz timer
- Results summary

### 5. Login (`/login`) & Register (`/register`)

**Files:** `app/login/page.tsx`, `app/register/page.tsx`

Trang xác thực với:

- Form validation với React Hook Form + Zod
- Demo mode (tự động tạo user ID)
- Redirect đến Dashboard sau khi thành công

### 6. Profile (`/profile`)

**File:** `app/profile/page.tsx`

Trang hồ sơ người dùng với:

- User information
- Learning statistics
- Mastery overview
- Logout functionality

---

## Components

### UI Components (`components/ui/`)

Thư viện 50+ UI components được xây dựng trên Radix UI và styled với Tailwind CSS:

| Component       | Description          | Radix UI Base         |
| --------------- | -------------------- | --------------------- |
| `button.tsx`    | Button với variants  | -                     |
| `card.tsx`      | Card container       | -                     |
| `dialog.tsx`    | Modal dialog         | `@radix-ui/dialog`    |
| `dropdown.tsx`  | Dropdown menu        | `@radix-ui/dropdown`  |
| `progress.tsx`  | Progress bar         | `@radix-ui/progress`  |
| `toast.tsx`     | Toast notifications  | `@radix-ui/toast`     |
| `tabs.tsx`      | Tab navigation       | `@radix-ui/tabs`      |
| `select.tsx`    | Select dropdown      | `@radix-ui/select`    |
| `checkbox.tsx`  | Checkbox input       | `@radix-ui/checkbox`  |
| `avatar.tsx`    | User avatar          | `@radix-ui/avatar`    |
| `tooltip.tsx`   | Tooltip              | `@radix-ui/tooltip`   |
| `accordion.tsx` | Collapsible sections | `@radix-ui/accordion` |

### Feature Components

#### Learning Components (`components/learning/`)

**`mastery-circle.tsx`**: Circular progress indicator cho mastery score

```tsx
<MasteryCircle score={75} skill="math" />
```

#### Quiz Components (`components/quiz/`)

| Component               | Description                        |
| ----------------------- | ---------------------------------- |
| `quiz-interface.tsx`    | Main quiz UI với questions/answers |
| `quiz-list.tsx`         | List of available quizzes          |
| `quiz-timer.tsx`        | Countdown timer cho quiz           |
| `results-dialog.tsx`    | Quiz results modal                 |
| `start-quiz-dialog.tsx` | Start quiz confirmation            |

---

## State Management

### Zustand Store

Client sử dụng **Zustand** - một lightweight state management library cho React. Zustand được chọn vì:

- **Đơn giản**: API minimal, không cần boilerplate
- **Hiệu năng**: Chỉ re-render components khi state thay đổi
- **TypeScript**: Full type support
- **Persistence**: Dễ dàng persist state vào localStorage

**File:** `store/useStore.ts`

```typescript
interface StoreState {
  // User state
  userId: string | null;
  setUserId: (id: string | null) => void;

  // Mastery data
  masteryData: { [skill: string]: number };
  setMastery: (skill: string, score: number) => void;

  // Current learning state
  currentQuestion: unknown;
  setCurrentQuestion: (question: unknown) => void;

  currentMastery: number;
  setCurrentMastery: (score: number) => void;

  // Loading state
  isLoading: boolean;
  setIsLoading: (loading: boolean) => void;
}
```

### Store Usage

```tsx
// Trong component
import { useStore } from "@/store/useStore";

function Dashboard() {
  const { userId, masteryData, setMastery } = useStore();

  // Sử dụng state
  console.log(`User: ${userId}`);
  console.log(`Math mastery: ${masteryData.math}%`);

  // Update state
  setMastery("math", 85);
}
```

### Persistence

User ID được persist vào `localStorage` để duy trì session:

```typescript
setUserId: (id) => {
  if (id) {
    localStorage.setItem("user_id", id);
  } else {
    localStorage.removeItem("user_id");
  }
  set({ userId: id });
};
```

---

## API Integration

### API Client

**File:** `services/api.ts`

Client sử dụng **Axios** để giao tiếp với backend services:

```typescript
const apiClient = axios.create({
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
});
```

### API Endpoints

| Method | Endpoint                             | Service         | Description                   |
| ------ | ------------------------------------ | --------------- | ----------------------------- |
| GET    | `/api/content/skills`                | Content         | Lấy danh sách skills          |
| GET    | `/api/content/{id}`                  | Content         | Lấy chi tiết câu hỏi          |
| GET    | `/internal/learner/{userId}/mastery` | Learner Model   | Lấy mastery score             |
| POST   | `/api/adaptive/next-lesson`          | Adaptive Engine | Lấy đề xuất bài học tiếp theo |
| POST   | `/api/scoring/submit`                | Scoring         | Nộp câu trả lời               |
| GET    | `/api/scoring/answered-questions`    | Scoring         | Lấy câu hỏi đã trả lời        |

### API Functions

```typescript
export const api = {
  // Lấy danh sách skills
  getAvailableSkills: async () => {
    return apiClient.get(`${API_URLS.content}/api/content/skills`);
  },

  // Lấy mastery score
  getMastery: async (userId: string, skill: string) => {
    return apiClient.get(
      `${API_URLS.learner}/internal/learner/${userId}/mastery?skill=${skill}`
    );
  },

  // Lấy đề xuất bài học
  getNextLesson: async (userId: string, skill: string) => {
    return apiClient.post(`${API_URLS.adaptive}/next-lesson`, {
      user_id: userId,
      current_skill: skill,
    });
  },

  // Lấy chi tiết câu hỏi
  getQuestion: async (questionId: number) => {
    return apiClient.get(`${API_URLS.content}/api/content/${questionId}`);
  },

  // Nộp câu trả lời
  submitAnswer: async (userId: string, questionId: number, answer: string) => {
    return apiClient.post(`${API_URLS.scoring}/api/scoring/submit`, {
      user_id: userId,
      question_id: questionId,
      answer: answer,
    });
  },
};
```

### Error Handling

API client có interceptor để xử lý errors:

```typescript
apiClient.interceptors.response.use(
  (response) => {
    // Check for error_code in response
    if (response.data?.error_code && response.data.error_code !== 0) {
      return Promise.reject({
        userMessage: response.data.message || "An error occurred",
      });
    }
    return response;
  },
  (error) => {
    let message = "An unexpected error occurred";
    if (error.response) {
      message = error.response.data?.message || "Server Error";
    } else if (error.request) {
      message = "Network Error: Unable to reach the server";
    }
    return Promise.reject({ ...error, userMessage: message });
  }
);
```

---

## Cấu hình

### Environment Variables

Client sử dụng các biến môi trường sau:

| Variable                       | Description               | Default                              |
| ------------------------------ | ------------------------- | ------------------------------------ |
| `NEXT_PUBLIC_CONTENT_API_URL`  | Content Service URL       | `http://localhost:8081`              |
| `NEXT_PUBLIC_SCORING_API_URL`  | Scoring Service URL       | `http://localhost:8082`              |
| `NEXT_PUBLIC_LEARNER_API_URL`  | Learner Model Service URL | `http://localhost:8083`              |
| `NEXT_PUBLIC_ADAPTIVE_API_URL` | Adaptive Engine URL       | `http://localhost:8084/api/adaptive` |
| `NEXT_PUBLIC_CLIENT_URL`       | Client URL (for E2E)      | `http://localhost:3001`              |
| `NODE_ENV`                     | Environment mode          | `development`                        |
| `PORT`                         | Server port               | `3000`                               |

### Environment Files

```bash
# Local development
.env.local

# Staging environment
.env.staging

# Test environment
.env.test
```

**Example `.env.local`:**

```env
# API Service URLs (Local Development)
NEXT_PUBLIC_CONTENT_API_URL=http://localhost:8081
NEXT_PUBLIC_SCORING_API_URL=http://localhost:8082
NEXT_PUBLIC_LEARNER_API_URL=http://localhost:8083
NEXT_PUBLIC_ADAPTIVE_API_URL=http://localhost:8084/api/adaptive

# Client Application URL (E2E tests sử dụng port 3001)
NEXT_PUBLIC_CLIENT_URL=http://localhost:3001
```

### Environment Configuration Module

**File:** `lib/env-config.ts`

Module này cung cấp cấu hình API URLs với fallback values:

```typescript
export const API_URLS = {
  content: getEnvUrl(
    process.env.NEXT_PUBLIC_CONTENT_API_URL,
    "http://localhost:8081"
  ),
  scoring: getEnvUrl(
    process.env.NEXT_PUBLIC_SCORING_API_URL,
    "http://localhost:8082"
  ),
  learner: getEnvUrl(
    process.env.NEXT_PUBLIC_LEARNER_API_URL,
    "http://localhost:8083"
  ),
  adaptive: getEnvUrl(
    process.env.NEXT_PUBLIC_ADAPTIVE_API_URL,
    "http://localhost:8084/api/adaptive"
  ),
  client: getEnvUrl(
    process.env.NEXT_PUBLIC_CLIENT_URL,
    "http://localhost:3000"
  ),
};
```

Module cũng cung cấp các utility functions:

- `validateEnvConfig()`: Kiểm tra cấu hình môi trường
- `getEnvironment()`: Lấy tên môi trường hiện tại (local, test, staging, production)

---

## Phát triển Local

### Prerequisites

- **Node.js 18.17+** hoặc **20.3+**: JavaScript runtime (Next.js 15 yêu cầu)
- **npm 9+** hoặc **pnpm**
- **Backend services** đang chạy (Content, Scoring, Learner Model, Adaptive Engine)

### Kiểm tra Prerequisites

```bash
# Kiểm tra Node.js version
node --version
# Expected: v18.17.x hoặc v20.3.x trở lên

# Kiểm tra npm version
npm --version
# Expected: 9.x.x hoặc cao hơn
```

### Setup Project

```bash
# Navigate đến client directory
cd sources/client

# Install dependencies
npm install --legacy-peer-deps

# Copy environment file
cp .env.local.example .env.local

# Edit .env.local nếu cần
```

### Run Development Server

```bash
# Start development server
npm run dev

# Server chạy tại http://localhost:3000
```

### Build Production

```bash
# Build production bundle
npm run build

# Start production server
npm run start
```

### Available Scripts

| Script                  | Description                            |
| ----------------------- | -------------------------------------- |
| `npm run dev`           | Start development server (port 3000)   |
| `npm run build`         | Build production bundle                |
| `npm run start`         | Start production server                |
| `npm run lint`          | Run ESLint (cần cài đặt eslint global) |
| `npm run test`          | Run Jest unit tests                    |
| `npm run test:watch`    | Run Jest in watch mode                 |
| `npm run test:coverage` | Run Jest with coverage                 |
| `npm run test:e2e`      | Run Playwright E2E tests (port 3001)   |
| `npm run test:e2e:ui`   | Run Playwright with UI                 |

### Verify Setup

```bash
# Kiểm tra client đang chạy
curl http://localhost:3000

# Kiểm tra API connectivity (cần backend services)
curl http://localhost:8081/health  # Content Service
curl http://localhost:8082/health  # Scoring Service
curl http://localhost:8083/health  # Learner Model Service
curl http://localhost:8084/health  # Adaptive Engine
```

---

## Testing

### Test Structure

```
sources/client/
├── __tests__/                     # Jest unit tests
│   ├── components/                # Component tests
│   │   └── *.test.tsx
│   ├── lib/                       # Utility tests
│   │   └── *.test.ts
│   └── services/                  # Service tests
│       └── *.test.ts
│
└── e2e/                           # Playwright E2E tests
    ├── fixtures/                  # Test fixtures
    ├── utils/                     # Test utilities
    ├── dashboard.spec.ts          # Dashboard tests
    ├── learning-flow.spec.ts      # Learning flow tests
    ├── api-integration.spec.ts    # API integration tests
    ├── mastery-flow-comprehensive.spec.ts
    ├── mastery-persistence.spec.ts
    ├── error-handling.spec.ts
    ├── ui-components.spec.ts
    └── global-setup.ts            # Global test setup
```

### Unit Testing (Jest)

**Configuration:** `jest.config.js`

```javascript
module.exports = {
  testEnvironment: "jest-environment-jsdom",
  setupFilesAfterEnv: ["<rootDir>/jest.setup.js"],
  moduleNameMapper: {
    "^@/(.*)$": "<rootDir>/$1",
  },
  testMatch: ["**/__tests__/**/*.[jt]s?(x)", "**/?(*.)+(spec|test).[jt]s?(x)"],
  testPathIgnorePatterns: ["/node_modules/", "/.next/", "/e2e/"],
};
```

**Running Unit Tests:**

```bash
# Run all unit tests
npm run test

# Run in watch mode
npm run test:watch

# Run with coverage
npm run test:coverage

# Run specific test file
npm run test -- __tests__/components/button.test.tsx
```

**Example Unit Test:**

```typescript
// __tests__/lib/utils.test.ts
import { getMasteryColor } from "@/lib/utils";

describe("getMasteryColor", () => {
  it("returns red for low mastery", () => {
    expect(getMasteryColor(25)).toBe("#ef4444");
  });

  it("returns yellow for medium mastery", () => {
    expect(getMasteryColor(50)).toBe("#eab308");
  });

  it("returns green for high mastery", () => {
    expect(getMasteryColor(80)).toBe("#22c55e");
  });
});
```

### E2E Testing (Playwright)

**Configuration:** `playwright.config.ts`

```typescript
export default defineConfig({
  testDir: "./e2e",
  globalSetup: require.resolve("./e2e/global-setup.ts"),
  globalTeardown: require.resolve("./e2e/global-teardown.ts"),
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  workers: process.env.CI ? 1 : undefined,
  reporter: [
    ["html", { outputFolder: reportBaseDir, open: "never" }],
    ["json", { outputFile: `${artifactsBaseDir}/results.json` }],
  ],
  use: {
    baseURL: process.env.NEXT_PUBLIC_CLIENT_URL || "http://localhost:3001",
    trace: "on-first-retry",
    screenshot: { mode: "only-on-failure", fullPage: true },
    video: { mode: "retain-on-failure", size: { width: 1280, height: 720 } },
  },
  projects: [
    { name: "chromium", use: { ...devices["Desktop Chrome"] } },
    { name: "firefox", use: { ...devices["Desktop Firefox"] } },
    { name: "webkit", use: { ...devices["Desktop Safari"] } },
  ],
  // Web server chỉ khởi động cho local environment
  webServer:
    environment === "local"
      ? {
          command: "npm run dev -- -p 3001",
          url: baseURL,
          reuseExistingServer: false,
          timeout: 120 * 1000,
        }
      : undefined,
});
```

**Running E2E Tests:**

```bash
# Run all E2E tests
npm run test:e2e

# Run with UI mode
npm run test:e2e:ui

# Run specific test file
npx playwright test e2e/dashboard.spec.ts

# Run with specific browser
npx playwright test --project=chromium

# Generate test report
npx playwright show-report
```

**Example E2E Test:**

```typescript
// e2e/dashboard.spec.ts
import { test, expect } from "@playwright/test";

test.describe("Dashboard", () => {
  test.beforeEach(async ({ page }) => {
    // Setup: Login and navigate to dashboard
    await page.goto("/");
    await page.click("text=Login");
    await page.waitForURL("/dashboard");
  });

  test("displays skill cards", async ({ page }) => {
    // Wait for skills to load
    await expect(page.locator('[data-testid="skill-card"]')).toBeVisible();

    // Verify mastery is displayed
    await expect(page.locator("text=Mastery")).toBeVisible();
  });

  test("navigates to learning page", async ({ page }) => {
    // Click continue learning
    await page.click("text=Continue Learning");

    // Verify navigation
    await expect(page).toHaveURL(/\/learn\//);
  });
});
```

### E2E Test Categories

| Test File                            | Description               |
| ------------------------------------ | ------------------------- |
| `dashboard.spec.ts`                  | Dashboard functionality   |
| `learning-flow.spec.ts`              | Complete learning flow    |
| `api-integration.spec.ts`            | API integration tests     |
| `mastery-flow-comprehensive.spec.ts` | Mastery update flow       |
| `mastery-persistence.spec.ts`        | Mastery data persistence  |
| `error-handling.spec.ts`             | Error handling scenarios  |
| `ui-components.spec.ts`              | UI component interactions |

### Test Scripts

```bash
# Start backend services for E2E
npm run e2e:start-services

# Stop backend services
npm run e2e:stop-services

# Verify services are running
npm run e2e:verify-services

# Setup test data
npm run e2e:setup-data

# Cleanup test data
npm run e2e:cleanup-data

# Run E2E with services
npm run test:e2e:with-services
```

---

## Deployment

### Docker Deployment

**Dockerfile:**

```dockerfile
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install --legacy-peer-deps
COPY . .
RUN npm run build

FROM node:20-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public
EXPOSE 3000
CMD ["node", "server.js"]
```

> **Lưu ý:** Dockerfile expose port 3000 nội bộ, nhưng docker-compose.yml map sang port 3001 bên ngoài.

**Build và Run với Docker:**

```bash
# Build Docker image
docker build -t its-client:latest .

# Run container
docker run -d \
  --name its-client \
  -p 3001:3001 \
  -e NEXT_PUBLIC_CONTENT_API_URL=http://content-service:8081 \
  -e NEXT_PUBLIC_SCORING_API_URL=http://scoring-service:8082 \
  -e NEXT_PUBLIC_LEARNER_API_URL=http://learner-model-api:8083 \
  -e NEXT_PUBLIC_ADAPTIVE_API_URL=http://adaptive-engine:8084/api/adaptive \
  its-client:latest
```

### Docker Compose

**Từ `sources/docker-compose.yml`:**

```yaml
client:
  build:
    context: ./client
    dockerfile: Dockerfile
  image: its-client:latest
  container_name: its-client
  restart: unless-stopped
  environment:
    NODE_ENV: production
    PORT: 3001
    NEXT_PUBLIC_API_BASE_URL: http://localhost:8084
    NEXT_PUBLIC_CONTENT_API: http://localhost:8081
    NEXT_PUBLIC_ADAPTIVE_API: http://localhost:8084
  ports:
    - "3001:3001"
  depends_on:
    adaptive-engine:
      condition: service_started
    content-service:
      condition: service_started
  networks:
    - its-network
```

**Deploy với Docker Compose:**

```bash
# Từ thư mục sources/
cd sources

# Build và start client
docker-compose up -d client

# Hoặc build lại
docker-compose up -d --build client

# Xem logs
docker-compose logs -f client
```

### Production Build

```bash
# Build production bundle
npm run build

# Output:
# - .next/standalone/  # Standalone server
# - .next/static/      # Static assets
# - public/            # Public files

# Start production server
npm run start
```

### Standalone Output

Next.js được cấu hình với `output: 'standalone'` để tạo minimal production build:

```javascript
// next.config.mjs
const nextConfig = {
  output: "standalone",
  // ...
};
```

**Benefits:**

- Smaller Docker image size
- Faster startup time
- Only includes necessary dependencies

---

## Tài liệu tham khảo

- [Next.js Documentation](https://nextjs.org/docs)
- [React Documentation](https://react.dev)
- [Tailwind CSS](https://tailwindcss.com/docs)
- [Radix UI](https://www.radix-ui.com/docs/primitives)
- [Zustand](https://github.com/pmndrs/zustand)
- [Playwright](https://playwright.dev/docs/intro)
- [Jest](https://jestjs.io/docs/getting-started)

---

## Liên kết Liên quan

### Tài liệu Dự án

| Tài liệu                | Đường dẫn                                                    | Mô tả                                |
| ----------------------- | ------------------------------------------------------------ | ------------------------------------ |
| **Root README**         | [../../README.md](../../README.md)                           | Tổng quan dự án, cấu trúc repository |
| **Sources README**      | [../README.md](../README.md)                                 | Hướng dẫn microservices và Docker    |
| **Report README**       | [../../report/README.md](../../report/README.md)             | Hướng dẫn build báo cáo LaTeX        |
| **Presentation README** | [../../presentation/README.md](../../presentation/README.md) | Hướng dẫn xem/build slides           |
| **Báo cáo PDF**         | [../../report/main.pdf](../../report/main.pdf)               | Báo cáo kiến trúc phần mềm           |

### Backend Service READMEs

| Service                   | Đường dẫn                                                    | Port | Mô tả                    |
| ------------------------- | ------------------------------------------------------------ | ---- | ------------------------ |
| **Content Service**       | [../content/README.md](../content/README.md)                 | 8081 | Quản lý nội dung học tập |
| **Scoring Service**       | [../scoring/README.md](../scoring/README.md)                 | 8082 | Chấm điểm, feedback      |
| **Learner Model Service** | [../learner-model/README.md](../learner-model/README.md)     | 8083 | Skill mastery tracking   |
| **Adaptive Engine**       | [../adaptive-engine/README.md](../adaptive-engine/README.md) | 8084 | Recommendation algorithm |

### Tài liệu Kiến trúc

| Tài liệu                   | Đường dẫn                                                                                              | Nội dung              |
| -------------------------- | ------------------------------------------------------------------------------------------------------ | --------------------- |
| **SOLID Principles**       | [../../markdown/report/6-SOLID-principles.md](../../markdown/report/6-SOLID-principles.md)             | Ví dụ áp dụng SOLID   |
| **Architecture Decisions** | [../../markdown/report/5-architecture-decisions.md](../../markdown/report/5-architecture-decisions.md) | ADRs                  |
| **Microservices Analysis** | [../../markdown/microservices.md](../../markdown/microservices.md)                                     | Chi tiết domain model |

### Testing Documentation

| Tài liệu              | Đường dẫn                                                          | Mô tả                          |
| --------------------- | ------------------------------------------------------------------ | ------------------------------ |
| **E2E Tests**         | [e2e/](./e2e/)                                                     | Playwright E2E test specs      |
| **Unit Tests**        | [**tests**/](./__tests__/)                                         | Jest unit tests                |
| **Testing Guide**     | [TESTING.md](./TESTING.md)                                         | Hướng dẫn testing              |
| **Environment Setup** | [docs/TEST_ENVIRONMENT_SETUP.md](./docs/TEST_ENVIRONMENT_SETUP.md) | Test environment configuration |

### Configuration Files

| File                  | Đường dẫn                                      | Mô tả                   |
| --------------------- | ---------------------------------------------- | ----------------------- |
| **Docker Compose**    | [../docker-compose.yml](../docker-compose.yml) | Cấu hình deployment     |
| **Playwright Config** | [playwright.config.ts](./playwright.config.ts) | E2E test configuration  |
| **Jest Config**       | [jest.config.js](./jest.config.js)             | Unit test configuration |
| **Next.js Config**    | [next.config.mjs](./next.config.mjs)           | Next.js configuration   |

---

**Client Web Application** - Phần của Intelligent Tutoring System (ITS)  
CO3017 - Kiến Trúc Phần Mềm - HCMUT
