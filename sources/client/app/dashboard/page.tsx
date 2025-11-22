"use client"

import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"
import { useStore } from "@/store/useStore"
import { api } from "@/services/api"
import { toast, Toaster } from "react-hot-toast"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Progress } from "@/components/ui/progress"
import { Calculator, FlaskConical, History, LogOut, Loader2, PlayCircle } from "lucide-react"
import { getMasteryColor } from "@/lib/utils"
import { motion } from "framer-motion"

const skills = [
  { id: "math", name: "Mathematics", icon: Calculator, description: "Algebra, Geometry, and Calculus" },
  { id: "science", name: "Science", icon: FlaskConical, description: "Physics, Chemistry, and Biology" },
  { id: "history", name: "History", icon: History, description: "World History and Civilizations" },
]

export default function DashboardPage() {
  const router = useRouter()
  const { userId, setUserId, masteryData, setMastery } = useStore()
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    // Check auth
    const storedUserId = localStorage.getItem("user_id")
    if (!storedUserId && !userId) {
      router.push("/")
      return
    }
    if (storedUserId && !userId) {
      setUserId(storedUserId)
    }

    // Fetch data
    const fetchData = async () => {
      if (!userId && !storedUserId) return
      const uid = userId || (storedUserId as string)

      try {
        setLoading(true)
        // Fetch mastery for all skills parallel
        await Promise.all(
          skills.map(async (skill) => {
            const response = await api.getMastery(uid, skill.id)
            if (response.data.error_code === 0) {
              setMastery(skill.id, response.data.data.mastery_score)
            }
          }),
        )
      } catch (error) {
        toast.error("Failed to load mastery data")
      } finally {
        setLoading(false)
      }
    }

    fetchData()
  }, [userId, router, setMastery, setUserId])

  const handleLogout = () => {
    setUserId(null)
    localStorage.removeItem("user_id")
    router.push("/")
  }

  const handleStartLearning = (skillId: string) => {
    router.push(`/learn/${skillId}`)
  }

  if (!userId && loading)
    return (
      <div className="flex items-center justify-center h-screen">
        <Loader2 className="h-8 w-8 animate-spin" />
      </div>
    )

  return (
    <div className="min-h-screen bg-background">
      <Toaster position="top-right" />

      {/* Header */}
      <header className="border-b bg-card">
        <div className="container flex h-16 items-center justify-between py-4">
          <div className="flex items-center gap-2">
            <div className="h-8 w-8 rounded-full bg-zinc-100 border flex items-center justify-center">
              <span className="font-bold text-black">I</span>
            </div>
            <h1 className="text-xl font-bold">My Learning Dashboard</h1>
          </div>
          <div className="flex items-center gap-4">
            <div className="hidden md:flex flex-col items-end">
              <span className="text-sm font-medium">Student ID</span>
              <span className="text-xs text-muted-foreground font-mono">{userId}</span>
            </div>
            <Button variant="outline" size="sm" onClick={handleLogout}>
              <LogOut className="h-4 w-4 mr-2" />
              Logout
            </Button>
          </div>
        </div>
      </header>

      <main className="container py-8">
        <div className="mb-8">
          <h2 className="text-3xl font-bold tracking-tight mb-2">Welcome back!</h2>
          <p className="text-muted-foreground">
            Here is an overview of your skill mastery. Pick a subject to continue learning.
          </p>
        </div>

        {loading ? (
          <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-[250px] rounded-xl border bg-card/50 animate-pulse" />
            ))}
          </div>
        ) : (
          <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
            {skills.map((skill, index) => {
              const mastery = masteryData[skill.id] || 0
              const masteryColor = getMasteryColor(mastery)

              return (
                <motion.div
                  key={skill.id}
                  initial={{ opacity: 0, y: 20 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ delay: index * 0.1 }}
                >
                  <Card className="overflow-hidden hover:shadow-lg transition-shadow border-2 h-full flex flex-col">
                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                      <CardTitle className="text-xl font-bold">{skill.name}</CardTitle>
                      <skill.icon className="h-6 w-6 text-muted-foreground" />
                    </CardHeader>
                    <CardContent className="pt-6 flex-1">
                      <div className="flex flex-col items-center justify-center mb-6">
                        {/* Simple CSS Circular Progress Placeholder */}
                        <div
                          className="relative h-32 w-32 rounded-full flex items-center justify-center border-8 transition-colors duration-500"
                          style={{ borderColor: masteryColor + "30" }} // Low opacity border
                        >
                          <div
                            className="absolute inset-0 rounded-full border-8 border-transparent transition-all duration-1000"
                            style={{
                              borderColor: masteryColor,
                              clipPath: `polygon(0 0, 100% 0, 100% 100%, 0 100%)`, // Simplified clip for demo
                              borderTopColor: masteryColor,
                              borderRightColor: mastery >= 25 ? masteryColor : "transparent",
                              borderBottomColor: mastery >= 50 ? masteryColor : "transparent",
                              borderLeftColor: mastery >= 75 ? masteryColor : "transparent",
                              transform: `rotate(${(mastery / 100) * 360}deg)`, // Very rough approximation for visual only
                            }}
                          />
                          {/* Better implementation would use SVG circle stroke-dasharray */}
                          <svg className="absolute inset-0 h-full w-full -rotate-90 transform" viewBox="0 0 100 100">
                            <circle
                              className="text-gray-200"
                              strokeWidth="8"
                              stroke="currentColor"
                              fill="transparent"
                              r="42"
                              cx="50"
                              cy="50"
                            />
                            <circle
                              className="transition-all duration-1000 ease-out"
                              strokeWidth="8"
                              strokeDasharray={264}
                              strokeDashoffset={264 - (264 * mastery) / 100}
                              strokeLinecap="round"
                              stroke={masteryColor}
                              fill="transparent"
                              r="42"
                              cx="50"
                              cy="50"
                            />
                          </svg>

                          <div className="flex flex-col items-center">
                            <span className="text-3xl font-bold" style={{ color: masteryColor }}>
                              {mastery}%
                            </span>
                            <span className="text-xs text-muted-foreground uppercase font-semibold">Mastery</span>
                          </div>
                        </div>
                      </div>
                      <div className="space-y-2">
                        <div className="flex justify-between text-xs">
                          <span className="text-muted-foreground">Progress</span>
                          <span className="font-medium" style={{ color: masteryColor }}>
                            {mastery < 50 ? "Beginner" : mastery < 80 ? "Intermediate" : "Advanced"}
                          </span>
                        </div>
                        <Progress value={mastery} className="h-2" indicatorColor={masteryColor} />
                      </div>
                    </CardContent>
                    <CardFooter className="bg-muted/50 pt-4">
                      <Button className="w-full" onClick={() => handleStartLearning(skill.id)}>
                        Continue Learning <PlayCircle className="ml-2 h-4 w-4" />
                      </Button>
                    </CardFooter>
                  </Card>
                </motion.div>
              )
            })}
          </div>
        )}
      </main>
    </div>
  )
}
