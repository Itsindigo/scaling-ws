import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    hmr: {
      overlay: true,
      host: "localhost",
      port: 3000,
    },
    host: "0.0.0.0",
    port: 3000,
    watch: {
      usePolling: true,
    },
  },
});
