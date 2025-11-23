# Client Web Application (Next.js)

**Port:** 3000
**Technology:** Next.js 15, React 19, Node.js 20
**Styling:** Tailwind CSS

## Overview

The frontend application for the Intelligent Tutoring System. It provides a user interface for students to take quizzes, view their progress, and receive adaptive learning recommendations.

## Quick Start (Docker)

```bash
# Start client via Docker Compose (from sources/ root)
make client
```

## Local Setup (Development)

1. **Install Dependencies**:
   ```bash
   npm install --legacy-peer-deps
   ```

2. **Run Development Server**:
   ```bash
   npm run dev
   ```

   Access at [http://localhost:3000](http://localhost:3000).

## Configuration

Environment variables (automatically set in `docker-compose.yml`):

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Service Port | 3000 |
| `NEXT_PUBLIC_API_BASE_URL` | Backend API URL | `http://localhost:8084` |

## Project Structure

- `src/app`: Next.js App Router pages
- `src/components`: Reusable UI components
- `src/services`: API integration
- `src/store`: State management (Zustand)
