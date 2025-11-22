"use client"

import { useEffect, useState } from "react"
import { useRouter, useParams } from "next/navigation"
import { useStore } from "@/store/useStore"
import { api } from "@/services/api"
import { toast, Toaster } from "react-hot-toast"
import { Button } from "@/components/ui/button"
import { Card } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { X, ArrowRight, CheckCircle, XCircle, Loader2, TrendingUp } from "lucide-react"
import { motion, AnimatePresence } from "framer-motion"
import { delay } from "@/lib/utils"
import { MasteryCircle } from "@/components/learning/mastery-circle"
import confetti from "canvas-confetti"

/**
 * Parse option string format "A. Answer text" into {key, text}
 * Handles both formats:
 * - "A. Answer text" -> {key: "A", text: "Answer text"}
 * - "Answer text only" -> {key: "Answer text only", text: "Answer text only"}
 */
function parseOption(optionStr: string): { key: string; text: string } {
  const match = optionStr.match(/^([A-Z])\.\s*(.+)$/)
  if (match) {
    return { key: match[1], text: match[2] }
  }
  // Fallback for options without prefix
  return { key: optionStr, text: optionStr }
}

export default function LearningSessionPage() {
  const router = useRouter()
  const params = useParams()
  const skillId = params.skill as string
  const { userId, currentMastery, setCurrentMastery } = useStore()

  // Local state for session flow
  const [loading, setLoading] = useState(true)
  const [submitting, setSubmitting] = useState(false)
  const [question, setQuestion] = useState<any>(null)
  const [contentType, setContentType] = useState<"standard" | "remedial">("standard")
  const [userAnswer, setUserAnswer] = useState<string | null>(null)
  const [feedback, setFeedback] = useState<any>(null)
  const [showFeedback, setShowFeedback] = useState(false)
  const [canContinue, setCanContinue] = useState(false)

  // Load initial data
  useEffect(() => {
    if (!userId) {
      router.push("/")
      return
    }
    loadSession()
  }, [userId, skillId])

  const loadSession = async () => {
    try {
      setLoading(true)
      // Step 1: Get Mastery
      const masteryRes = await api.getMastery(userId!, skillId)
      const currentScore = masteryRes.data.data.mastery_score
      setCurrentMastery(currentScore)

      // Step 2 & 3: Load next question
      await loadNextQuestion(currentScore)
    } catch (error) {
      toast.error("Failed to start learning session")
    } finally {
      setLoading(false)
    }
  }

  const loadNextQuestion = async (currentScore: number) => {
    try {
      // Step 2: Adaptive Engine Recommendation
      const adaptiveRes = await api.getNextLesson(userId!, skillId)
      const { next_lesson_id, content_type } = adaptiveRes.data.data

      setContentType(content_type)

      // Step 3: Fetch Question Content
      const questionRes = await api.getQuestion(next_lesson_id)
      setQuestion(questionRes.data.data)

      // Reset state for new question
      setUserAnswer(null)
      setFeedback(null)
      setShowFeedback(false)
      setCanContinue(false)
    } catch (error) {
      toast.error("Error loading next question")
    }
  }

  const parseOption = (opt: string) => {
    // Matches "A. Content" or "1. Content"
    const match = opt.match(/^([A-Z0-9]+)\.\s*(.+)$/)
    if (match) {
      return { key: match[1], text: match[2] }
    }
    // Fallback if format doesn't match (e.g. just "A")
    return { key: opt, text: opt }
  }

  const handleSubmit = async () => {
    if (!userAnswer || !question) return

    setSubmitting(true)
    try {
      // Step 4: Submit Answer
      // If it's a multiple choice question (has options), we need to parse the key (e.g. "A")
      // If it's open text, we just trim the input
      let finalAnswer = userAnswer.trim()

      if (question.options && question.options.length > 0) {
        const parsed = parseOption(userAnswer)
        finalAnswer = parsed.key
      }

      const response = await api.submitAnswer(userId!, question.id, finalAnswer)
      const result = response.data.data

      setFeedback(result)
      setShowFeedback(true)

      if (result.correct) {
        confetti({
          particleCount: 100,
          spread: 70,
          origin: { y: 0.6 },
        })
      }

      // Step 5: Poll for Mastery Update
      // Calculate expected mastery (simple mock logic for demo: +5 if correct)
      const expectedMastery = result.correct ? Math.min(100, currentMastery + 5) : Math.max(0, currentMastery - 2)

      await pollForMasteryUpdate(expectedMastery)

      // Enable next button
      setTimeout(() => setCanContinue(true), 1000)
    } catch (error) {
      toast.error("Failed to submit answer")
    } finally {
      setSubmitting(false)
    }
  }

  const pollForMasteryUpdate = async (expectedScore: number) => {
    const startTime = Date.now()
    const timeout = 5000 // 5s timeout for demo

    // Mock polling simulation
    while (Date.now() - startTime < timeout) {
      await delay(500)
      // In real app, we'd fetch api.getMastery() here.
      // For demo, we'll just simulate the update happening after delay
      if (Date.now() - startTime > 1500) {
        setCurrentMastery(expectedScore)
        return
      }
    }
  }

  const handleNext = async () => {
    setLoading(true)
    await loadNextQuestion(currentMastery)
    setLoading(false)
  }

  const handleExit = () => {
    router.push("/dashboard")
  }

  if (loading && !question) {
    return (
      <div className="flex flex-col items-center justify-center h-screen gap-4">
        <Loader2 className="h-10 w-10 animate-spin text-primary" />
        <p className="text-muted-foreground animate-pulse">Consulting AI Tutor...</p>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-slate-50 dark:bg-zinc-950 flex flex-col">
      <Toaster />

      {/* Top Bar */}
      <header className="bg-background border-b h-16 px-4 flex items-center justify-between sticky top-0 z-10 shadow-sm">
        <div className="flex items-center gap-4">
          <Badge variant="outline" className="text-base py-1 px-3 uppercase tracking-wide">
            {skillId}
          </Badge>
          <div className="flex items-center gap-2 bg-secondary/10 px-3 py-1 rounded-full">
            <span className="text-sm text-muted-foreground font-medium">Mastery:</span>
            <MasteryCircle score={currentMastery} size={32} />
          </div>
        </div>
        <Button variant="ghost" size="icon" onClick={handleExit}>
          <X className="h-5 w-5" />
        </Button>
      </header>

      {/* Main Content */}
      <main className="flex-1 container max-w-3xl py-8 flex flex-col justify-center">
        <AnimatePresence mode="wait">
          {question && (
            <motion.div
              key={question.id}
              initial={{ opacity: 0, x: 20 }}
              animate={{ opacity: 1, x: 0 }}
              exit={{ opacity: 0, x: -20 }}
              transition={{ duration: 0.3 }}
            >
              <Card className="p-6 md:p-8 shadow-md border-t-4 border-t-primary">
                {/* Question Header */}
                <div className="flex justify-between items-start mb-6">
                  <Badge
                    variant={contentType === "remedial" ? "destructive" : "default"}
                    className={
                      contentType === "remedial" ? "bg-orange-500 hover:bg-orange-600" : "bg-blue-600 hover:bg-blue-700"
                    }
                  >
                    {contentType === "remedial" ? "Remedial • Foundation" : "Standard • Challenge"}
                  </Badge>
                  <span className="text-xs text-muted-foreground font-mono">ID: {question.id}</span>
                </div>

                {/* Question Text */}
                <h2 className="text-xl md:text-2xl font-bold mb-8 leading-relaxed text-balance">{question.content}</h2>

                {/* Answer Options or Text Input */}
                <div className="mb-8">
                  {question.options && question.options.length > 0 ? (
                    <div className="grid gap-3">
                      {question.options.map((opt: string) => {
                        const { key, text } = parseOption(opt)
                        const isSelected = userAnswer === key
                        return (
                          <button
                            key={opt}
                            onClick={() => !showFeedback && setUserAnswer(key)}
                            disabled={showFeedback || submitting}
                            className={`
                              w-full p-4 rounded-lg border-2 text-left transition-all font-medium text-lg flex items-center
                              ${isSelected
                                ? "border-primary bg-primary/5 ring-1 ring-primary"
                                : "border-muted hover:border-primary/50 hover:bg-muted/30"
                              }
                              ${showFeedback && isSelected && feedback?.correct ? "border-green-500 bg-green-50 text-green-700" : ""}
                              ${showFeedback && isSelected && !feedback?.correct ? "border-red-500 bg-red-50 text-red-700" : ""}
                              disabled:cursor-not-allowed
                            `}
                          >
                            <div className={`
                              h-8 w-8 rounded-full flex items-center justify-center mr-4 font-bold shrink-0 border
                              ${isSelected ? "bg-primary text-primary-foreground border-primary" : "bg-muted text-muted-foreground border-transparent"}
                              ${showFeedback && isSelected && feedback?.correct ? "bg-green-600 text-white border-green-600" : ""}
                              ${showFeedback && isSelected && !feedback?.correct ? "bg-red-600 text-white border-red-600" : ""}
                            `}>
                              {key}
                            </div>
                            <span>{text}</span>
                          </button>
                        )
                      })}
                    </div>
                  ) : (
                    <div className="space-y-4">
                      <label className="block text-sm font-medium text-muted-foreground">
                        Your Answer:
                      </label>
                      <input
                        type="text"
                        value={userAnswer || ""}
                        onChange={(e) => setUserAnswer(e.target.value)}
                        disabled={showFeedback || submitting}
                        className="w-full p-4 rounded-lg border-2 border-muted bg-background text-lg focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all"
                        placeholder="Type your answer here..."
                        onKeyDown={(e) => {
                          if (e.key === "Enter" && userAnswer && !submitting) {
                            handleSubmit()
                          }
                        }}
                      />
                    </div>
                  )}
                </div>

                {/* Submit Area */}
                {!showFeedback && (
                  <Button
                    size="lg"
                    className="w-full text-lg py-6"
                    onClick={handleSubmit}
                    disabled={!userAnswer || submitting}
                  >
                    {submitting && <Loader2 className="mr-2 h-5 w-5 animate-spin" />}
                    {submitting ? "Submitting..." : "Submit Answer"}
                  </Button>
                )}
              </Card>
            </motion.div>
          )}
        </AnimatePresence>
      </main>

      {/* Feedback Panel (Fixed Bottom) */}
      <AnimatePresence>
        {showFeedback && feedback && (
          <motion.div
            initial={{ y: "100%" }}
            animate={{ y: 0 }}
            exit={{ y: "100%" }}
            transition={{ type: "spring", stiffness: 300, damping: 30 }}
            className={`
              fixed bottom-0 left-0 right-0 p-6 border-t shadow-2xl z-20
              ${feedback.correct ? "bg-green-50 dark:bg-green-950/30 border-green-200" : "bg-red-50 dark:bg-red-950/30 border-red-200"}
            `}
          >
            <div className="container max-w-3xl flex flex-col md:flex-row items-center justify-between gap-4">
              <div className="flex items-center gap-4">
                <div
                  className={`
                  h-12 w-12 rounded-full flex items-center justify-center shrink-0
                  ${feedback.correct ? "bg-green-100 text-green-600" : "bg-red-100 text-red-600"}
                `}
                >
                  {feedback.correct ? <CheckCircle className="h-6 w-6" /> : <XCircle className="h-6 w-6" />}
                </div>
                <div>
                  <h3 className={`font-bold text-lg ${feedback.correct ? "text-green-800" : "text-red-800"}`}>
                    {feedback.correct ? "Excellent!" : "Not quite right..."}
                  </h3>
                  <p className="text-muted-foreground">{feedback.feedback}</p>
                </div>
              </div>

              <div className="flex items-center gap-4 w-full md:w-auto">
                <div className="hidden md:flex flex-col items-end mr-4">
                  <span className="text-xs uppercase text-muted-foreground font-bold">New Mastery</span>
                  <span className="text-2xl font-bold font-mono text-primary flex items-center gap-1">
                    {currentMastery}%
                    <TrendingUp className="h-4 w-4" />
                  </span>
                </div>
                <Button
                  size="lg"
                  onClick={handleNext}
                  disabled={!canContinue}
                  className={`w-full md:w-auto ${feedback.correct ? "bg-green-600 hover:bg-green-700" : "bg-primary"}`}
                >
                  {canContinue ? (
                    <>
                      Next Question <ArrowRight className="ml-2 h-5 w-5" />
                    </>
                  ) : (
                    <>
                      Updating Progress <Loader2 className="ml-2 h-4 w-4 animate-spin" />
                    </>
                  )}
                </Button>
              </div>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  )
}
