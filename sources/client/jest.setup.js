// Learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom'

// Mock Next.js router
jest.mock('next/navigation', () => ({
    useRouter() {
        return {
            push: jest.fn(),
            replace: jest.fn(),
            prefetch: jest.fn(),
            back: jest.fn(),
            pathname: '/',
            query: {},
            asPath: '/',
        }
    },
    useParams() {
        return {}
    },
    usePathname() {
        return '/'
    },
    useSearchParams() {
        return new URLSearchParams()
    },
}))

// Mock framer-motion
jest.mock('framer-motion', () => ({
    motion: {
        div: ({
            children,
            ...props
        }) => < div {
            ...props
        } > {
            children
        } < /div>,
        circle: ({
            ...props
        }) => < circle {
            ...props
        }
        />,
    },
    AnimatePresence: ({
        children
    }) => children,
}))

// Mock canvas-confetti
jest.mock('canvas-confetti', () => ({
    __esModule: true,
    default: jest.fn(),
}))