"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Loader2 } from "lucide-react"

interface ResultsDialogProps {
  quizId: string
  title: string
  totalQuestions: number
  answers: Record<string, string>
  onClose?: () => void
}

interface SubmitResponse {
  success: boolean
  score: number
  percentage: number
  message: string
  skillsUpdated: Array<{
    skillName: string
    pointsEarned: number
    totalPoints: number
  }>
}

export default function ResultsDialog({ quizId, title, totalQuestions, answers, onClose }: ResultsDialogProps) {
  const router = useRouter()
  const [isLoading, setIsLoading] = useState(true)
  const [result, setResult] = useState<SubmitResponse | null>(null)
  const [error, setError] = useState<string | null>(null)

  // Simulate API call to submit quiz results
  const submitQuizResults = async () => {
    try {
      setIsLoading(true)

      // Simulate API delay
      await new Promise((resolve) => setTimeout(resolve, 2000))

      // Mock API response
      const mockResponse: SubmitResponse = {
        success: true,
        score: Math.floor(Math.random() * (20 - 10 + 1)) + 10,
        percentage: Math.floor(Math.random() * (100 - 60 + 1)) + 60,
        message: "Bài kiểm tra của bạn đã được xử lý thành công!",
        skillsUpdated: [
          {
            skillName: "AWS Fundamentals",
            pointsEarned: Math.floor(Math.random() * (50 - 20 + 1)) + 20,
            totalPoints: 100,
          },
          {
            skillName: "Cloud Architecture",
            pointsEarned: Math.floor(Math.random() * (40 - 15 + 1)) + 15,
            totalPoints: 80,
          },
        ],
      }

      setResult(mockResponse)
    } catch (err) {
      setError("Có lỗi xảy ra khi nộp bài kiểm tra. Vui lòng thử lại.")
      console.error("Error submitting quiz:", err)
    } finally {
      setIsLoading(false)
    }
  }

  // Trigger API call on mount
  const [hasInitiated, setHasInitiated] = useState(false)
  if (!hasInitiated) {
    setHasInitiated(true)
    submitQuizResults()
  }

  if (isLoading) {
    return (
      <div className="fixed inset-0 bg-black/50 flex items-center justify-center p-4 z-50">
        <Card className="w-full max-w-md">
          <CardHeader>
            <CardTitle className="text-center">Đang xử lý kết quả...</CardTitle>
          </CardHeader>
          <CardContent className="flex flex-col items-center gap-6">
            <div className="space-y-3 w-full">
              <div className="h-2 bg-muted rounded-full overflow-hidden">
                <div className="h-full bg-foreground animate-pulse" style={{ width: "60%" }} />
              </div>
              <p className="text-center text-sm text-muted-foreground">Đang cập nhật điểm của bạn...</p>
            </div>
            <Loader2 className="w-8 h-8 animate-spin" />
          </CardContent>
        </Card>
      </div>
    )
  }

  if (error) {
    return (
      <div className="fixed inset-0 bg-black/50 flex items-center justify-center p-4 z-50">
        <Card className="w-full max-w-md border-destructive">
          <CardHeader>
            <CardTitle className="text-center text-destructive">Lỗi</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-center text-muted-foreground">{error}</p>
            <Button
              onClick={() => router.push("/quizzes")}
              className="w-full bg-foreground text-background hover:bg-foreground/90"
            >
              Quay lại danh sách bài kiểm tra
            </Button>
          </CardContent>
        </Card>
      </div>
    )
  }

  if (!result) return null

  return (
    <div className="fixed inset-0 bg-black/50 flex items-center justify-center p-4 z-50">
      <Card className="w-full max-w-md">
        <CardHeader>
          <CardTitle className="text-center text-lg">Kết quả bài kiểm tra</CardTitle>
          <p className="text-center text-sm text-muted-foreground mt-1">{title}</p>
        </CardHeader>
        <CardContent className="space-y-6">
          {/* Score Display */}
          <div className="text-center">
            <div className="text-4xl font-bold mb-2">{result.percentage}%</div>
            <div className="text-sm text-muted-foreground mb-4">
              {result.score} / {totalQuestions} câu đúng
            </div>
            <p className="text-sm font-medium text-foreground/80">{result.message}</p>
          </div>

          {/* Skills Updated */}
          <div className="space-y-3 border-t pt-4">
            <h3 className="text-sm font-semibold">Kỹ năng được cập nhật:</h3>
            {result.skillsUpdated.map((skill) => (
              <div key={skill.skillName} className="space-y-2">
                <div className="flex justify-between text-sm">
                  <span className="font-medium">{skill.skillName}</span>
                  <span className="text-foreground/70">+{skill.pointsEarned} điểm</span>
                </div>
                <div className="w-full h-2 bg-muted rounded-full overflow-hidden">
                  <div
                    className="h-full bg-foreground"
                    style={{
                      width: `${(skill.totalPoints / 100) * 100}%`,
                    }}
                  />
                </div>
                <div className="text-xs text-muted-foreground">{skill.totalPoints} / 100 điểm</div>
              </div>
            ))}
          </div>

          {/* Action Buttons */}
          <div className="flex flex-col gap-2">
            <Button
              onClick={() => router.push("/profile")}
              className="w-full bg-foreground text-background hover:bg-foreground/90"
            >
              Xem hồ sơ kỹ năng
            </Button>
            <Button onClick={() => router.push("/quizzes")} variant="outline" className="w-full bg-transparent">
              Quay lại danh sách
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
