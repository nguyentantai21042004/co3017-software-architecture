"use client"

import { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group"
import { Label } from "@/components/ui/label"
import { Progress } from "@/components/ui/progress"
import QuizTimer from "./quiz-timer"

interface Option {
  value: string
  text: string
}

interface Question {
  questionId: string
  text: string
  questionType: string
  options: Option[]
}

interface QuizInterfaceProps {
  quiz: {
    quizId: string
    title: string
    timeLimitInSeconds: number
    questions: Question[]
  }
  onSubmit: (answers: Record<string, string>) => void
}

export default function QuizInterface({ quiz, onSubmit }: QuizInterfaceProps) {
  const [currentQuestion, setCurrentQuestion] = useState(0)
  const [answers, setAnswers] = useState<Record<string, string>>({})
  const [timeLeft, setTimeLeft] = useState(quiz.timeLimitInSeconds)
  const [isTimeUp, setIsTimeUp] = useState(false)

  const question = quiz.questions[currentQuestion]
  const progress = ((currentQuestion + 1) / quiz.questions.length) * 100

  useEffect(() => {
    if (isTimeUp) return

    const timer = setInterval(() => {
      setTimeLeft((prev) => {
        if (prev <= 1) {
          setIsTimeUp(true)
          return 0
        }
        return prev - 1
      })
    }, 1000)

    return () => clearInterval(timer)
  }, [isTimeUp])

  const handleAnswerChange = (value: string) => {
    setAnswers((prev) => ({
      ...prev,
      [question.questionId]: value,
    }))
  }

  const handleNext = () => {
    if (currentQuestion < quiz.questions.length - 1) {
      setCurrentQuestion((prev) => prev + 1)
    }
  }

  const handlePrev = () => {
    if (currentQuestion > 0) {
      setCurrentQuestion((prev) => prev - 1)
    }
  }

  const handleSubmit = () => {
    onSubmit(answers)
  }

  const handleTimeUp = () => {
    onSubmit(answers)
  }

  if (isTimeUp) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-background">
        <Card className="w-full max-w-md">
          <CardHeader>
            <CardTitle className="text-center text-destructive">Hết thời gian</CardTitle>
          </CardHeader>
          <CardContent className="text-center">
            <p className="text-muted-foreground mb-6">Bài kiểm tra của bạn đã tự động nộp.</p>
            <Button onClick={handleTimeUp} className="w-full bg-foreground text-background">
              Xem kết quả
            </Button>
          </CardContent>
        </Card>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-background p-4 md:p-8">
      {/* Header */}
      <div className="max-w-4xl mx-auto mb-8">
        <div className="flex items-center justify-between mb-6">
          <h1 className="text-2xl font-bold">{quiz.title}</h1>
          <QuizTimer timeLeft={timeLeft} onTimeUp={handleTimeUp} />
        </div>

        {/* Progress */}
        <div className="space-y-2">
          <div className="flex justify-between text-sm text-muted-foreground">
            <span>
              Câu {currentQuestion + 1} / {quiz.questions.length}
            </span>
            <span>{Math.round(progress)}%</span>
          </div>
          <Progress value={progress} className="h-2" />
        </div>
      </div>

      {/* Question Card */}
      <div className="max-w-4xl mx-auto">
        <Card>
          <CardHeader>
            <CardTitle className="text-lg">{question.text}</CardTitle>
          </CardHeader>
          <CardContent>
            {/* Answer Options */}
            <RadioGroup value={answers[question.questionId] || ""} onValueChange={handleAnswerChange}>
              <div className="space-y-4">
                {question.options.map((option) => (
                  <div
                    key={option.value}
                    className="flex items-center space-x-3 p-4 border rounded-lg hover:bg-muted/50 cursor-pointer transition-colors"
                  >
                    <RadioGroupItem value={option.value} id={option.value} />
                    <Label htmlFor={option.value} className="flex-1 cursor-pointer">
                      {option.text}
                    </Label>
                  </div>
                ))}
              </div>
            </RadioGroup>

            {/* Navigation Buttons */}
            <div className="flex gap-4 mt-8 pt-6 border-t">
              <Button
                variant="outline"
                onClick={handlePrev}
                disabled={currentQuestion === 0}
                className="flex-1 bg-transparent"
              >
                ← Câu trước
              </Button>
              {currentQuestion === quiz.questions.length - 1 ? (
                <Button onClick={handleSubmit} className="flex-1 bg-foreground text-background hover:bg-foreground/90">
                  Nộp bài
                </Button>
              ) : (
                <Button onClick={handleNext} className="flex-1 bg-foreground text-background hover:bg-foreground/90">
                  Câu tiếp →
                </Button>
              )}
            </div>
          </CardContent>
        </Card>

        {/* Question Overview */}
        <Card className="mt-8">
          <CardHeader>
            <CardTitle className="text-sm">Tổng quan câu hỏi</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-4 md:grid-cols-6 lg:grid-cols-10 gap-2">
              {quiz.questions.map((_, index) => (
                <button
                  key={index}
                  onClick={() => setCurrentQuestion(index)}
                  className={`w-8 h-8 rounded text-sm font-medium transition-colors ${
                    index === currentQuestion
                      ? "bg-foreground text-background"
                      : answers[quiz.questions[index].questionId]
                        ? "bg-muted text-foreground border border-foreground"
                        : "bg-muted text-muted-foreground hover:bg-muted/80"
                  }`}
                >
                  {index + 1}
                </button>
              ))}
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
