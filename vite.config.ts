import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'node:path';

// https://vite.dev/config/
export default defineConfig({
  root: path.resolve(__dirname, 'web'),
  plugins: [react()],
  build: {
    outDir: "../build/web/",
  },
});
