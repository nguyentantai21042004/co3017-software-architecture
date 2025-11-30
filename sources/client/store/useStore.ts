import { create } from "zustand"

interface MasteryData {
  [key: string]: number
}

interface StoreState {
  userId: string | null
  setUserId: (id: string | null) => void

  masteryData: MasteryData
  setMastery: (skill: string, score: number) => void

  currentQuestion: unknown
  setCurrentQuestion: (question: unknown) => void

  currentMastery: number
  setCurrentMastery: (score: number) => void

  isLoading: boolean
  setIsLoading: (loading: boolean) => void
}

export const useStore = create<StoreState>((set) => ({
  userId: typeof window !== "undefined" ? localStorage.getItem("user_id") : null,
  setUserId: (id) => {
    if (id) {
      localStorage.setItem("user_id", id)
    } else {
      localStorage.removeItem("user_id")
    }
    set({ userId: id })
  },

  masteryData: {
    math: 0,
    science: 0,
    history: 0,
  },
  setMastery: (skill, score) =>
    set((state) => ({
      masteryData: { ...state.masteryData, [skill]: score },
    })),

  currentQuestion: null,
  setCurrentQuestion: (question) => set({ currentQuestion: question }),

  currentMastery: 0,
  setCurrentMastery: (score) => set({ currentMastery: score }),

  isLoading: false,
  setIsLoading: (loading) => set({ isLoading: loading }),
}))
