"use client"

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Progress } from "@/components/ui/progress"

interface Skill {
  skillId: string
  skillName: string
  category: string
  currentPoints: number
  maxPoints: number
  level: "Beginner" | "Intermediate" | "Advanced" | "Expert"
  lastUpdated: string
}

const mockUserSkills: Skill[] = [
  {
    skillId: "skill-001",
    skillName: "AWS Fundamentals",
    category: "Cloud",
    currentPoints: 75,
    maxPoints: 100,
    level: "Intermediate",
    lastUpdated: "2024-01-15",
  },
  {
    skillId: "skill-002",
    skillName: "Cloud Architecture",
    category: "Cloud",
    currentPoints: 62,
    maxPoints: 100,
    level: "Intermediate",
    lastUpdated: "2024-01-14",
  },
  {
    skillId: "skill-003",
    skillName: "Python Programming",
    category: "Programming",
    currentPoints: 88,
    maxPoints: 100,
    level: "Advanced",
    lastUpdated: "2024-01-10",
  },
  {
    skillId: "skill-004",
    skillName: "Data Analytics",
    category: "Data",
    currentPoints: 45,
    maxPoints: 100,
    level: "Beginner",
    lastUpdated: "2024-01-05",
  },
  {
    skillId: "skill-005",
    skillName: "JavaScript & React",
    category: "Programming",
    currentPoints: 92,
    maxPoints: 100,
    level: "Advanced",
    lastUpdated: "2024-01-12",
  },
  {
    skillId: "skill-006",
    skillName: "SQL Database Design",
    category: "Database",
    currentPoints: 58,
    maxPoints: 100,
    level: "Intermediate",
    lastUpdated: "2024-01-08",
  },
]

const getLevelColor = (level: string) => {
  switch (level) {
    case "Beginner":
      return "text-muted-foreground"
    case "Intermediate":
      return "text-foreground"
    case "Advanced":
      return "text-foreground font-semibold"
    case "Expert":
      return "text-foreground font-bold"
    default:
      return "text-muted-foreground"
  }
}

const getLevelBgColor = (level: string) => {
  switch (level) {
    case "Beginner":
      return "bg-muted"
    case "Intermediate":
      return "bg-muted"
    case "Advanced":
      return "bg-foreground/10"
    case "Expert":
      return "bg-foreground/20"
    default:
      return "bg-muted"
  }
}

export default function ProfilePage() {
  // Group skills by category
  const skillsByCategory = mockUserSkills.reduce(
    (acc, skill) => {
      if (!acc[skill.category]) {
        acc[skill.category] = []
      }
      acc[skill.category].push(skill)
      return acc
    },
    {} as Record<string, Skill[]>,
  )

  const totalPoints = mockUserSkills.reduce((sum, skill) => sum + skill.currentPoints, 0)
  const maxTotalPoints = mockUserSkills.reduce((sum, skill) => sum + skill.maxPoints, 0)

  return (
    <div className="min-h-screen bg-background p-4 md:p-8">
      <div className="max-w-4xl mx-auto">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold mb-2">Hồ sơ kỹ năng</h1>
          <p className="text-muted-foreground">Theo dõi tiến độ và phát triển của các kỹ năng của bạn</p>
        </div>

        {/* Overall Stats */}
        <Card className="mb-8">
          <CardHeader>
            <CardTitle>Tổng quan</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div>
                <p className="text-sm text-muted-foreground mb-2">Tổng điểm</p>
                <p className="text-3xl font-bold">
                  {totalPoints}/{maxTotalPoints}
                </p>
              </div>
              <div>
                <p className="text-sm text-muted-foreground mb-2">Số kỹ năng</p>
                <p className="text-3xl font-bold">{mockUserSkills.length}</p>
              </div>
              <div>
                <p className="text-sm text-muted-foreground mb-2">Mức trung bình</p>
                <p className="text-3xl font-bold">{Math.round((totalPoints / maxTotalPoints) * 100)}%</p>
              </div>
            </div>
            <div className="pt-4 border-t">
              <div className="flex justify-between text-sm mb-2">
                <span className="font-medium">Tiến độ tổng thể</span>
                <span className="text-muted-foreground">{Math.round((totalPoints / maxTotalPoints) * 100)}%</span>
              </div>
              <Progress value={(totalPoints / maxTotalPoints) * 100} className="h-3" />
            </div>
          </CardContent>
        </Card>

        {/* Skills by Category */}
        {Object.entries(skillsByCategory).map(([category, skills]) => (
          <Card key={category} className="mb-6">
            <CardHeader>
              <CardTitle className="text-lg">{category}</CardTitle>
            </CardHeader>
            <CardContent className="space-y-6">
              {skills.map((skill) => (
                <div key={skill.skillId} className="space-y-3">
                  <div className="flex items-start justify-between">
                    <div>
                      <h3 className="font-semibold text-foreground">{skill.skillName}</h3>
                      <p className="text-xs text-muted-foreground mt-1">
                        Cập nhật lần cuối: {new Date(skill.lastUpdated).toLocaleDateString("vi-VN")}
                      </p>
                    </div>
                    <span
                      className={`px-3 py-1 rounded text-xs font-medium ${getLevelBgColor(
                        skill.level,
                      )} ${getLevelColor(skill.level)}`}
                    >
                      {skill.level === "Beginner" && "Sơ cấp"}
                      {skill.level === "Intermediate" && "Trung cấp"}
                      {skill.level === "Advanced" && "Nâng cao"}
                      {skill.level === "Expert" && "Chuyên gia"}
                    </span>
                  </div>

                  {/* Progress Bar */}
                  <div className="space-y-1">
                    <div className="flex justify-between text-sm">
                      <span className="text-muted-foreground">Điểm</span>
                      <span className="font-medium">
                        {skill.currentPoints} / {skill.maxPoints}
                      </span>
                    </div>
                    <Progress value={(skill.currentPoints / skill.maxPoints) * 100} className="h-2" />
                  </div>
                </div>
              ))}
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  )
}
