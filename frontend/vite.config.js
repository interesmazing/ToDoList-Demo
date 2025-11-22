import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    // This is required for Docker container port mapping to work correctly
    host: '0.0.0.0', 
    port: 5173, // This is the port inside the container
    proxy: {
      // Proxy API requests to the backend service.
      // '/api' is the path that the frontend will use to make requests.
      // 'http://backend:8080' is the address of the backend service inside the Docker network.
      // Vite's development server will forward requests from http://localhost:5173/api
      // to http://backend:8080/api.
      // In production, Nginx will handle this proxying.
      '/api': {
        target: 'http://backend:8080',
        changeOrigin: true,
      },
    },
  },
})
