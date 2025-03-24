import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import autoprefixer from 'autoprefixer'
import tailwind from 'tailwindcss'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  root: process.cwd() + '/frontend',
  css: {
    postcss: {
      plugins: [tailwind(), autoprefixer()],
    },
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./frontend', import.meta.url)),
    },
  },
  server: {
    proxy: {
      '/api': 'http://localhost:8080',
    },
  },
  build: {
    outDir: '../pkg/frontend/files',
    sourcemap: true,
    manifest: 'manifest.json',
    emptyOutDir: true,
  },
});
