import { getMasteryColor, cn, delay } from '@/lib/utils'

describe('lib/utils', () => {
  describe('getMasteryColor', () => {
    it('should return red for scores below 50', () => {
      expect(getMasteryColor(0)).toBe('#EF4444')
      expect(getMasteryColor(25)).toBe('#EF4444')
      expect(getMasteryColor(49)).toBe('#EF4444')
    })

    it('should return yellow/orange for scores between 50 and 79', () => {
      expect(getMasteryColor(50)).toBe('#F59E0B')
      expect(getMasteryColor(65)).toBe('#F59E0B')
      expect(getMasteryColor(79)).toBe('#F59E0B')
    })

    it('should return green for scores 80 and above', () => {
      expect(getMasteryColor(80)).toBe('#22C55E')
      expect(getMasteryColor(90)).toBe('#22C55E')
      expect(getMasteryColor(100)).toBe('#22C55E')
    })
  })

  describe('cn', () => {
    it('should merge class names correctly', () => {
      expect(cn('foo', 'bar')).toBe('foo bar')
      expect(cn('foo', false && 'bar', 'baz')).toBe('foo baz')
    })

    it('should handle conditional classes', () => {
      expect(cn('base', true && 'conditional')).toBe('base conditional')
      expect(cn('base', false && 'conditional')).toBe('base')
    })
  })

  describe('delay', () => {
    it('should return a promise that resolves after the specified time', async () => {
      const start = Date.now()
      await delay(100)
      const end = Date.now()
      const elapsed = end - start
      
      // Allow some tolerance for timing
      expect(elapsed).toBeGreaterThanOrEqual(90)
      expect(elapsed).toBeLessThan(150)
    })
  })
})

