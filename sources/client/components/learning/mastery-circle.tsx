"use client"

import { motion } from "framer-motion"
import { getMasteryColor } from "@/lib/utils"

interface MasteryCircleProps {
  score: number
  size?: number
}

export function MasteryCircle({ score, size = 60 }: MasteryCircleProps) {
  const color = getMasteryColor(score)
  const radius = size / 2 - 4
  const circumference = 2 * Math.PI * radius
  const strokeDashoffset = circumference - (score / 100) * circumference

  return (
    <div className="relative flex items-center justify-center" style={{ width: size, height: size }}>
      <svg className="transform -rotate-90 w-full h-full">
        <circle
          className="text-muted/20"
          strokeWidth="4"
          stroke="currentColor"
          fill="transparent"
          r={radius}
          cx={size / 2}
          cy={size / 2}
        />
        <motion.circle
          initial={{ strokeDashoffset: circumference }}
          animate={{ strokeDashoffset }}
          transition={{ duration: 1, ease: "easeOut" }}
          strokeWidth="4"
          strokeDasharray={circumference}
          strokeLinecap="round"
          stroke={color}
          fill="transparent"
          r={radius}
          cx={size / 2}
          cy={size / 2}
        />
      </svg>
      <span className="absolute text-xs font-bold" style={{ color }}>
        {score}%
      </span>
    </div>
  )
}
