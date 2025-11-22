# Use lightweight Node.js image
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy package files first to leverage cache
COPY package.json package-lock.json* ./

# Install dependencies with legacy-peer-deps to handle React 19 conflicts
RUN npm install --legacy-peer-deps

# Copy the rest of the application code
COPY . .

# Expose the default Next.js port
EXPOSE 3000

# Start the application in development mode
# For hot reload to work with Docker, ensure you mount the volume:
# -v $(pwd):/app -v /app/node_modules
CMD ["npm", "run", "dev"]
