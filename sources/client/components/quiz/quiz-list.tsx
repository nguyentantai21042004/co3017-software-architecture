"use client"
import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"

interface Quiz {
  quizId: string
  title: string
  description: string
  timeLimitInSeconds: number
  questions: Array<{ questionId: string }>
}

interface QuizListProps {
  quizzes: Quiz[]
}

export default function QuizList({ quizzes }: QuizListProps) {
  return (
    <div className="space-y-4">
      {quizzes.map((quiz) => (
        <Card key={quiz.quizId} className="hover:border-foreground/50 transition-colors">
          <CardHeader>
            <div className="flex items-start justify-between">
              <div className="flex-1">
                <CardTitle className="text-xl">{quiz.title}</CardTitle>
                <CardDescription className="mt-2">{quiz.description}</CardDescription>
              </div>
            </div>
          </CardHeader>
          <CardContent>
            <div className="flex items-center justify-between">
              <div className="flex gap-4 text-sm text-muted-foreground">
                <span>{quiz.questions.length} câu hỏi</span>
                <span>{Math.floor(quiz.timeLimitInSeconds / 60)} phút</span>
              </div>
              <Link href={`/quizzes/${quiz.quizId}`}>
                <Button className="bg-foreground text-background hover:bg-foreground/90">Bắt đầu làm bài</Button>
              </Link>
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  )
}
