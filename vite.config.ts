import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  base: './',
  root: './frontend',
  build: {
    outDir: '../.tmp/dist',
    minify: false,
  },
})
