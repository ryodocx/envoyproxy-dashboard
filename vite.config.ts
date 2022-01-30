import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  base: './',
  build: {
    outDir: '.tmp/dist',
    minify: false,
  },
  server: {
    port: 8000,
    open: '/index.html',
  },
})
