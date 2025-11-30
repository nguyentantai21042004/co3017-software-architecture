// Jest configuration for Next.js
// Using manual config to avoid SWC binary issues in test environment
module.exports = {
    testEnvironment: 'jest-environment-jsdom',
    setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
    moduleNameMapper: {
        '^@/(.*)$': '<rootDir>/$1',
        '^@/components/(.*)$': '<rootDir>/components/$1',
        '^@/lib/(.*)$': '<rootDir>/lib/$1',
        '^@/services/(.*)$': '<rootDir>/services/$1',
        '^@/types/(.*)$': '<rootDir>/types/$1',
        '^@/store/(.*)$': '<rootDir>/store/$1',
    },
    testMatch: [
        '**/__tests__/**/*.[jt]s?(x)',
        '**/?(*.)+(spec|test).[jt]s?(x)',
    ],
    testPathIgnorePatterns: [
        '/node_modules/',
        '/.next/',
        '/e2e/',
    ],
    collectCoverageFrom: [
        'components/**/*.{js,jsx,ts,tsx}',
        'app/**/*.{js,jsx,ts,tsx}',
        'services/**/*.{js,jsx,ts,tsx}',
        'lib/**/*.{js,jsx,ts,tsx}',
        '!**/*.d.ts',
        '!**/node_modules/**',
        '!**/.next/**',
    ],
    transform: {
        '^.+\\.(js|jsx|ts|tsx)$': ['babel-jest', {
            presets: ['next/babel']
        }],
    },
    transformIgnorePatterns: [
        '/node_modules/',
        '^.+\\.module\\.(css|sass|scss)$',
    ],
    moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx', 'json', 'node'],
}