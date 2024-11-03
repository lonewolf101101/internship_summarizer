// file: server/api/postpdf.js
import { defineEventHandler, readRawBody } from 'h3'; // Use h3 for Nuxt 3 APIs

export default defineEventHandler(async (event) => {
  // Read raw body for PDF
  const body = await readRawBody(event);

  try {
    // Forward the PDF data to your Go backend
    const response = await $fetch('http://localhost:3300/postpdf', {
      method: 'POST',
      body,
      headers: {
        'Content-Type': 'application/pdf',
      },
    });

    // Return the response from the Go backend  
    return response;
  } catch (error) {
    console.error('Error forwarding PDF to backend:', error);
    // Send a generic error response
    return {
      statusCode: 500,
      body: 'Error uploading PDF to backend.'
    };
  }
});
