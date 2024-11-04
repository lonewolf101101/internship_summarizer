import path from 'path';
import fs from 'fs';
import { defineNuxtConfig } from 'nuxt/config'; // Correct import

console.log("Server type:", process.env.SERVER_TYPE);

export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },

  devServer: {
    https: {
      key: fs.readFileSync(path.resolve("D:/My_Code/Chimege/internship_summarizer/frontend/tls", 'key.pem')).toString(),
      cert: fs.readFileSync(path.resolve("D:/My_Code/Chimege/internship_summarizer/frontend/tls", 'cert.pem')).toString(),
    },
  },
  plugins: ['~/plugins/vuetify.js'],
  serverHandlers: [
    { route: '/api', handler: '~/server/middleware/proxy.js' }
  ]
});