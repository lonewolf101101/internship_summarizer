// server/api/summarize.js
export default defineEventHandler(async (event) => {
    const body = await readBody(event); // Get the request body
  
    // Forward the request to the external API (your backend)
    const response = await $fetch('http://localhost:3300/summarize', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body)  // Send the request body to the backend API
    });
  
    return response;  // Return the backend API response
  });
  