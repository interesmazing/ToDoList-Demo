# Stage 1: Build the Vue.js application
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package.json and package-lock.json (or yarn.lock)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the application for production
RUN npm run build

# Stage 2: Serve the application using Nginx
FROM nginx:1.25-alpine

# Copy the built assets from the 'builder' stage
COPY --from=builder /app/dist /usr/share/nginx/html

# When a request comes in, Nginx will first look for a matching file in its root directory.
# If it doesn't find one (e.g., for a deep link like /todos/123), it will fall back
# to serving /index.html. This allows the Vue router to handle the request.
COPY --from=builder /app/dist/index.html /usr/share/nginx/html/index.html

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
