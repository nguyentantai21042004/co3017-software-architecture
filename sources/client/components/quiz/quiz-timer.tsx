"use client"

interface QuizTimerProps {
  timeLeft: number
  onTimeUp: () => void
}

export default function QuizTimer({ timeLeft, onTimeUp }: QuizTimerProps) {
  const minutes = Math.floor(timeLeft / 60)
  const seconds = timeLeft % 60
  const isWarning = timeLeft < 300 // Less than 5 minutes

  const formatTime = (time: number) => {
    return time.toString().padStart(2, "0")
  }

  return (
    <div
      className={`text-2xl font-mono font-bold px-4 py-2 rounded-lg border-2 transition-colors ${
        isWarning
          ? "border-destructive bg-destructive/10 text-destructive"
          : "border-foreground/20 bg-muted text-foreground"
      }`}
    >
      {formatTime(minutes)}:{formatTime(seconds)}
    </div>
  )
}
