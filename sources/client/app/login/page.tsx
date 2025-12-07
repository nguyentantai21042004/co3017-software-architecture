"use client"

import { useRouter } from "next/navigation"
import { useStore } from "@/store/useStore"
import { toast, Toaster } from "react-hot-toast"
import { Navbar } from "@/components/navbar"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import Link from "next/link"

export default function LoginPage() {
  const router = useRouter()
  const { setUserId } = useStore()

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // Fake authentication - auto-generate user ID
    const newUserId = `student-${Date.now()}`
    setUserId(newUserId)
    toast.success(`Welcome! Your ID: ${newUserId}`)
    setTimeout(() => router.push("/dashboard"), 1000)
  }

  return (
    <div className="min-h-screen bg-background">
      <Toaster position="top-center" />
      <Navbar />

      <div className="container mx-auto px-4 sm:px-6 lg:px-8 py-16 md:py-24">
        <div className="max-w-md mx-auto">
          <Card>
            <CardHeader className="space-y-1">
              <CardTitle className="text-2xl font-bold">Welcome back</CardTitle>
              <CardDescription>
                Demo mode - Click "Sign In" to auto-login (no password required)
              </CardDescription>
            </CardHeader>
            <CardContent>
              <form className="space-y-4" onSubmit={handleSubmit}>
                <div className="space-y-2">
                  <Label htmlFor="email">Email</Label>
                  <Input id="email" type="email" placeholder="student@example.com" />
                </div>

                <div className="space-y-2">
                  <div className="flex items-center justify-between">
                    <Label htmlFor="password">Password</Label>
                    <Link
                      href="/forgot-password"
                      className="text-sm text-muted-foreground hover:text-foreground transition-colors"
                    >
                      Forgot password?
                    </Link>
                  </div>
                  <Input id="password" type="password" />
                </div>

                <Button type="submit" className="w-full" size="lg">
                  Sign In
                </Button>
              </form>

              <div className="mt-6 text-center text-sm">
                <span className="text-muted-foreground">Don't have an account? </span>
                <Link href="/register" className="font-semibold hover:underline">
                  Sign up
                </Link>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
