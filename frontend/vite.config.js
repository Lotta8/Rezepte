// vite.config.js
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': '/src' // Stelle sicher, dass der Pfad zu deinem "src"-Verzeichnis korrekt ist
    }
  }
})
