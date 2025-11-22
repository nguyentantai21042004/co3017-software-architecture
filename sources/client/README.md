# Client Web Application

This is the React-based frontend for the Intelligent Tutoring System, built with Next.js.

## 🚀 Quick Start with Docker (Recommended)

You can run the application in a Docker container with hot-reload enabled without installing Node.js locally.

### 1. Build the Image
```bash
docker build -t client-app .
```

### 2. Run with Hot Reload
Run the following command to start the app and map your local files to the container (enabling hot reload):

```bash
docker run -p 3000:3000 \
  -v $(pwd):/app \
  -v /app/node_modules \
  client-app
```

- Access the app at: [http://localhost:3000](http://localhost:3000)
- Edit files in `src/` and see changes instantly!

---

## 🛠️ Local Development

If you prefer running locally:

1. **Install Dependencies**
   ```bash
   npm install --legacy-peer-deps
   ```
   *(Note: `--legacy-peer-deps` is required due to React 19 peer dependency conflicts)*

2. **Start Dev Server**
   ```bash
   npm run dev
   ```

3. **Build for Production**
   ```bash
   npm run build
   npm start
   ```

## 📂 Project Structure

- **`src/app`**: Next.js App Router pages
- **`src/components`**: Reusable UI components
- **`src/services`**: API client and service calls
- **`src/store`**: State management (Zustand)
- **`documents/`**: Project specifications and prompts

## 🔧 Tech Stack

- **Framework**: Next.js 15 (React 19)
- **Styling**: Tailwind CSS
- **State**: Zustand
- **HTTP Client**: Axios
- **UI Components**: Radix UI + Lucide Icons
