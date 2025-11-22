"use client"

import { useState, useEffect } from "react"
import { useParams, useRouter } from "next/navigation"
import QuizInterface from "@/components/quiz/quiz-interface"
import ResultsDialog from "@/components/quiz/results-dialog"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Skeleton } from "@/components/ui/skeleton"

interface Question {
  questionId: string
  text: string
  questionType: string
  options: Array<{ value: string; text: string }>
}

interface Quiz {
  quizId: string
  title: string
  timeLimitInSeconds: number
  questions: Question[]
}

export default function TakeQuizPage() {
  const params = useParams()
  const router = useRouter()
  const quizId = params.quizId as string
  const [quiz, setQuiz] = useState<Quiz | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const [showResults, setShowResults] = useState(false)
  const [answers, setAnswers] = useState<Record<string, string>>({})

  useEffect(() => {
    const loadQuiz = async () => {
      try {
        const response = await fetch(`/data/quizzes/${quizId}.json`)
        if (response.ok) {
          const data = await response.json()
          setQuiz(data)
        }
      } catch (error) {
        console.error("Error loading quiz:", error)
      } finally {
        setIsLoading(false)
      }
    }

    loadQuiz()
  }, [quizId])

  const handleSubmit = (submittedAnswers: Record<string, string>) => {
    setAnswers(submittedAnswers)
    setShowResults(true)
  }

  if (isLoading) {
    return (
      <div className="min-h-screen bg-background p-8">
        <div className="max-w-4xl mx-auto space-y-4">
          <Skeleton className="h-8 w-1/4" />
          <Skeleton className="h-4 w-full" />
          <Skeleton className="h-64 w-full" />
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
      <QuizInterface quiz={quiz} onSubmit={handleSubmit} />
      {showResults && (
        <ResultsDialog quizId={quizId} title={quiz.title} totalQuestions={quiz.questions.length} answers={answers} />
      )}
    </>
  )
}
