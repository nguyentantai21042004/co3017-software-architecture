import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function getMasteryColor(score: number) {
  if (score < 50) return "#EF4444" // Red
  if (score < 80) return "#F59E0B" // Yellow/Orange
  return "#22C55E" // Green
}

export const delay = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms))
