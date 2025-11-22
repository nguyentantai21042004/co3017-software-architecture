import QuizList from "@/components/quiz/quiz-list"
import Link from "next/link"
import { Button } from "@/components/ui/button"

// Mock data - sẽ được cung cấp từ các file JSON
const mockQuizzes = [
  {
    quizId: "quiz-aws-fundamentals-001",
    title: "Bài kiểm tra kiến thức cơ bản về AWS",
    description:
      "Bài kiểm tra này đánh giá hiểu biết cơ bản về các dịch vụ, khái niệm và best practices trên nền tảng Amazon Web Services (AWS).",
    timeLimitInSeconds: 1800,
    questions: Array(25).fill({ questionId: "" }),
  },
  {
    quizId: "quiz-javascript-basics",
    title: "Bài kiểm tra JavaScript cơ bản",
    description: "Kiểm tra kiến thức về các khái niệm cơ bản của JavaScript như variables, functions, async/await.",
    timeLimitInSeconds: 1200,
    questions: Array(20).fill({ questionId: "" }),
  },
  {
    quizId: "quiz-react-advanced",
    title: "Bài kiểm tra React nâng cao",
    description: "Đánh giá sâu về React hooks, context API, performance optimization và state management.",
    timeLimitInSeconds: 1500,
    questions: Array(15).fill({ questionId: "" }),
  },
]

export default function QuizzesPage() {
  return (
    <div className="min-h-screen bg-background">
      {/* Header */}
      <div className="border-b border-border">
        <div className="max-w-6xl mx-auto px-4 py-8">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-bold">Bài Kiểm Tra</h1>
              <p className="text-muted-foreground mt-2">Danh sách các bài kiểm tra online có sẵn</p>
            </div>
            <Link href="/dashboard">
              <Button variant="outline">← Quay lại</Button>
            </Link>
          </div>
        </div>
      </div>

      {/* Content */}
      <div className="max-w-6xl mx-auto px-4 py-8">
        <QuizList quizzes={mockQuizzes} />
      </div>
    </div>
  )
}
