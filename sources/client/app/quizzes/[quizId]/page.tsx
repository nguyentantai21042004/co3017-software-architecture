"use client"

import { useState, useEffect } from "react"
import { useParams } from "next/navigation"
import StartQuizDialog from "@/components/quiz/start-quiz-dialog"
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Skeleton } from "@/components/ui/skeleton"

interface Question {
  questionId: string
  text: string
  options: Array<{ value: string; text: string }>
}

interface Quiz {
  quizId: string
  title: string
  description: string
  timeLimitInSeconds: number
  questions: Question[]
}

export default function QuizDetailPage() {
  const params = useParams()
  const quizId = params.quizId as string
  const [quiz, setQuiz] = useState<Quiz | null>(null)
  const [isDialogOpen, setIsDialogOpen] = useState(true)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    const loadQuiz = async () => {
      try {
        const response = await fetch(`/data/quizzes/${quizId}.json`)
        if (response.ok) {
          const data = await response.json()
          setQuiz(data)
        } else {
          console.error("Quiz not found")
        }
      } catch (error) {
        console.error("Error loading quiz:", error)
      } finally {
        setIsLoading(false)
      }
    }

    loadQuiz()
  }, [quizId])

  if (isLoading) {
    return (
      <div className="min-h-screen bg-background p-8">
        <div className="max-w-2xl mx-auto space-y-4">
          <Skeleton className="h-8 w-3/4" />
          <Skeleton className="h-4 w-full" />
          <Skeleton className="h-32 w-full" />
        </div>
      </div>
    )
  }

  if (!quiz) {
    return (
      <div className="min-h-screen bg-background p-8">
        <div className="max-w-2xl mx-auto">
          <Card>
            <CardHeader>
              <CardTitle>Không tìm thấy bài kiểm tra</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-muted-foreground">Bài kiểm tra bạn tìm kiếm không tồn tại.</p>
            </CardContent>
          </Card>
        </div>
      </div>
    )
  }

  return (
    <>
      <div className="min-h-screen bg-background p-8">
        <div className="max-w-2xl mx-auto">
          <Card>
            <CardHeader>
              <CardTitle className="text-2xl">{quiz.title}</CardTitle>
              <CardDescription className="text-base mt-4">{quiz.description}</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-6">
                <div className="grid grid-cols-2 gap-4">
                  <div className="bg-muted p-4 rounded-lg">
                    <p className="text-sm text-muted-foreground">Số câu hỏi</p>
                    <p className="text-2xl font-bold">{quiz.questions.length}</p>
                  </div>
                  <div className="bg-muted p-4 rounded-lg">
                    <p className="text-sm text-muted-foreground">Thời gian</p>
                    <p className="text-2xl font-bold">{Math.floor(quiz.timeLimitInSeconds / 60)} phút</p>
                  </div>
                </div>

                <div className="bg-secondary/50 p-4 rounded-lg border border-border">
                  <p className="text-sm font-medium mb-2">Yêu cầu:</p>
                  <ul className="text-sm text-muted-foreground space-y-1">
                    <li>• Bạn phải hoàn thành bài kiểm tra trong thời gian quy định</li>
                    <li>• Mỗi câu hỏi có một đáp án đúng</li>
                    <li>• Không thể quay lại sau khi nộp bài</li>
                    <li>• Kết quả sẽ được hiển thị ngay sau khi nộp</li>
                  </ul>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>

      <StartQuizDialog isOpen={isDialogOpen} quiz={quiz} onOpenChange={setIsDialogOpen} />
    </>
  )
}
