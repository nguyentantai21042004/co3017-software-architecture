"use client"

import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"
import { useStore } from "@/store/useStore"
import { toast, Toaster } from "react-hot-toast"
import { TrendingUp, Zap, ArrowRight, LayoutDashboard } from "lucide-react"
import { Button } from "@/components/ui/button"
import Link from "next/link"

export default function HomePage() {
  const router = useRouter()
  const { setUserId } = useStore()
  const [welcomeMessage, setWelcomeMessage] = useState("Welcome back!")
  const [isReturningUser, setIsReturningUser] = useState(false)

  useEffect(() => {
    const existingUserId = localStorage.getItem("user_id")
    if (existingUserId) {
      setIsReturningUser(true)
      setWelcomeMessage(`Welcome back, ${existingUserId}!`)
    }
  }, [])

  const handleLogin = () => {
    const newUserId = `student-${Date.now()}`
    setUserId(newUserId)
    toast.success(`Welcome! Your ID: ${newUserId}`)
    setTimeout(() => router.push("/dashboard"), 1000)
  }

  const handleContinue = () => {
    router.push("/dashboard")
  }

  return (
    <div className="min-h-screen bg-background flex flex-col">
      <Toaster position="top-center" />

      {/* Navbar Placeholder */}
      <header className="border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
        <div className="container flex h-14 items-center">
          <div className="mr-4 flex">
            <Link href="/" className="mr-6 flex items-center space-x-2">
              <span className="font-bold text-xl">ITS Platform</span>
            </Link>
          </div>
          <div className="flex flex-1 items-center justify-between space-x-2 md:justify-end">
            <nav className="flex items-center space-x-4">
              {isReturningUser ? (
                <Button variant="ghost" onClick={handleContinue}>
                  Dashboard
                </Button>
              ) : (
                <Button variant="ghost" onClick={handleLogin}>
                  Login
                </Button>
              )}
            </nav>
          </div>
        </div>
      </header>

      <main className="flex-1">
        {/* Hero Section */}
        <section className="space-y-6 pb-8 pt-6 md:pb-12 md:pt-10 lg:py-32">
          <div className="container flex max-w-[64rem] flex-col items-center gap-4 text-center">
            <h1 className="font-heading text-3xl sm:text-5xl md:text-6xl lg:text-7xl font-bold tracking-tight">
              Intelligent Tutoring System
            </h1>
            <p className="max-w-[42rem] leading-normal text-muted-foreground sm:text-xl sm:leading-8">
              Adaptive Learning Powered by AI. Experience a personalized education path that evolves with your mastery.
            </p>
            <div className="space-x-4 pt-4">
              {isReturningUser ? (
                <Button size="lg" onClick={handleContinue} className="bg-primary text-primary-foreground">
                  Continue Learning <ArrowRight className="ml-2 h-4 w-4" />
                </Button>
              ) : (
                <div className="flex gap-4 justify-center">
                  <Button
                    size="lg"
                    onClick={handleLogin}
                    className="bg-primary text-primary-foreground hover:bg-primary/90 min-w-[150px]"
                  >
                    Login
                  </Button>
                  <Button
                    size="lg"
                    variant="outline"
                    onClick={handleLogin}
                    className="border-primary text-primary hover:bg-primary/5 min-w-[150px] bg-transparent"
                  >
                    Register
                  </Button>
                </div>
              )}
            </div>
            {!isReturningUser && <p className="text-xs text-muted-foreground">Demo mode - No password required!</p>}
          </div>
        </section>

        {/* Features Section */}
        <section className="container space-y-6 py-8 md:py-12 lg:py-24 bg-slate-50 dark:bg-zinc-900/50 rounded-3xl my-8">
          <div className="mx-auto grid justify-center gap-4 sm:grid-cols-2 md:max-w-[64rem] md:grid-cols-3">
            <div className="relative overflow-hidden rounded-lg border bg-background p-2">
              <div className="flex h-[180px] flex-col justify-between rounded-md p-6">
                <Zap className="h-12 w-12 text-primary" />
                <div className="space-y-2">
                  <h3 className="font-bold">Adaptive Learning</h3>
                  <p className="text-sm text-muted-foreground">Questions automatically match your skill level.</p>
                </div>
              </div>
            </div>
            <div className="relative overflow-hidden rounded-lg border bg-background p-2">
              <div className="flex h-[180px] flex-col justify-between rounded-md p-6">
                <TrendingUp className="h-12 w-12 text-primary" />
                <div className="space-y-2">
                  <h3 className="font-bold">Real-time Progress</h3>
                  <p className="text-sm text-muted-foreground">See your mastery grow instantly with live updates.</p>
                </div>
              </div>
            </div>
            <div className="relative overflow-hidden rounded-lg border bg-background p-2">
              <div className="flex h-[180px] flex-col justify-between rounded-md p-6">
                <LayoutDashboard className="h-12 w-12 text-primary" />
                <div className="space-y-2">
                  <h3 className="font-bold">Personalized Path</h3>
                  <p className="text-sm text-muted-foreground">AI-powered recommendations for what to learn next.</p>
                </div>
              </div>
            </div>
          </div>
        </section>
      </main>

      <footer className="py-6 md:px-8 md:py-0">
        <div className="container flex flex-col items-center justify-between gap-4 md:h-24 md:flex-row">
          <p className="text-center text-sm leading-loose text-muted-foreground md:text-left">
            Built with Clean Architecture & Event-Driven Microservices
          </p>
        </div>
      </footer>
    </div>
  )
}
