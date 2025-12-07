"use client"

import { useRouter } from "next/navigation"
import { useStore } from "@/store/useStore"
import { toast, Toaster } from "react-hot-toast"
import { Navbar } from "@/components/navbar"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import Link from "next/link"

export default function RegisterPage() {
  const router = useRouter()
  const { setUserId } = useStore()

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // Fake authentication - auto-generate user ID
    const newUserId = `student-${Date.now()}`
    setUserId(newUserId)
    toast.success(`Account created! Your ID: ${newUserId}`)
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
              <CardTitle className="text-2xl font-bold">Create an account</CardTitle>
              <CardDescription>
                Demo mode - Click "Create Account" to auto-register (no password required)
              </CardDescription>
            </CardHeader>
            <CardContent>
              <form className="space-y-4" onSubmit={handleSubmit}>
                <div className="space-y-2">
                  <Label htmlFor="fullname">Full Name</Label>
                  <Input id="fullname" type="text" placeholder="Nguyen Van A" />
                </div>

                <div className="space-y-2">
                  <Label htmlFor="email">Email</Label>
                  <Input id="email" type="email" placeholder="student@example.com" />
                </div>

                <div className="space-y-2">
                  <Label htmlFor="role">I am a</Label>
                  <Select defaultValue="student">
                    <SelectTrigger id="role">
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="student">Student</SelectItem>
                      <SelectItem value="instructor">Instructor</SelectItem>
                    </SelectContent>
                  </Select>
                </div>

                <div className="space-y-2">
                  <Label htmlFor="password">Password</Label>
                  <Input id="password" type="password" />
                </div>

                <div className="space-y-2">
                  <Label htmlFor="confirm-password">Confirm Password</Label>
                  <Input id="confirm-password" type="password" />
                </div>

                <Button type="submit" className="w-full" size="lg">
                  Create Account
                </Button>
              </form>

              <div className="mt-6 text-center text-sm">
                <span className="text-muted-foreground">Already have an account? </span>
                <Link href="/login" className="font-semibold hover:underline">
                  Sign in
                </Link>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
