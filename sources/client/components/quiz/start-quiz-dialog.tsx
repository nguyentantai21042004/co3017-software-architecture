"use client"
import { useRouter } from "next/navigation"
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"

interface StartQuizDialogProps {
  isOpen: boolean
  quiz: {
    quizId: string
    title: string
    timeLimitInSeconds: number
    questions: Array<{ questionId: string }>
  }
  onOpenChange: (open: boolean) => void
}

export default function StartQuizDialog({ isOpen, quiz, onOpenChange }: StartQuizDialogProps) {
  const router = useRouter()

  const handleStart = () => {
    router.push(`/quizzes/${quiz.quizId}/take`)
    onOpenChange(false)
  }

  const minutes = Math.floor(quiz.timeLimitInSeconds / 60)

  return (
    <AlertDialog open={isOpen} onOpenChange={onOpenChange}>
      <AlertDialogContent className="max-w-md">
        <AlertDialogHeader>
          <AlertDialogTitle>Xác nhận bắt đầu bài kiểm tra</AlertDialogTitle>
          <AlertDialogDescription asChild>
            <div className="space-y-4 mt-4">
              <p className="font-medium text-foreground">{quiz.title}</p>
              <div className="bg-muted p-4 rounded-lg space-y-2 text-sm">
                <div className="flex justify-between">
                  <span className="text-muted-foreground">Số câu hỏi:</span>
                  <span className="font-medium">{quiz.questions.length} câu</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-muted-foreground">Thời gian:</span>
                  <span className="font-medium">{minutes} phút</span>
                </div>
              </div>
              <p className="text-xs text-muted-foreground">
                Sau khi bắt đầu, bài kiểm tra sẽ tự động nộp khi hết thời gian.
              </p>
            </div>
          </AlertDialogDescription>
        </AlertDialogHeader>
        <div className="flex gap-3 justify-end">
          <AlertDialogCancel>Hủy bỏ</AlertDialogCancel>
          <AlertDialogAction onClick={handleStart} className="bg-foreground text-background hover:bg-foreground/90">
            Bắt đầu
          </AlertDialogAction>
        </div>
      </AlertDialogContent>
    </AlertDialog>
  )
}
