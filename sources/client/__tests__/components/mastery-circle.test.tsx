import { render, screen } from '@testing-library/react'
import { MasteryCircle } from '@/components/learning/mastery-circle'

describe('MasteryCircle', () => {
  it('should render the score percentage', () => {
    render(<MasteryCircle score={75} />)
    expect(screen.getByText('75%')).toBeInTheDocument()
  })

  it('should render with default size', () => {
    const { container } = render(<MasteryCircle score={50} />)
    const svg = container.querySelector('svg')
    expect(svg).toBeInTheDocument()
  })

  it('should render with custom size', () => {
    const { container } = render(<MasteryCircle score={80} size={100} />)
    const wrapper = container.firstChild as HTMLElement
    expect(wrapper).toHaveStyle({ width: '100px', height: '100px' })
  })

  it('should display correct score for different values', () => {
    const { rerender } = render(<MasteryCircle score={0} />)
    expect(screen.getByText('0%')).toBeInTheDocument()

    rerender(<MasteryCircle score={100} />)
    expect(screen.getByText('100%')).toBeInTheDocument()

    rerender(<MasteryCircle score={45} />)
    expect(screen.getByText('45%')).toBeInTheDocument()
  })

  it('should render SVG circle elements', () => {
    const { container } = render(<MasteryCircle score={60} />)
    const circles = container.querySelectorAll('circle')
    expect(circles.length).toBeGreaterThan(0)
  })
})

