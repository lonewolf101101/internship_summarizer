<template>
    <div>
      <h1>Text Summarizer</h1>
      <textarea 
        v-model="text" 
        placeholder="Enter your text here..." 
        :maxlength="maxChars"
      ></textarea>
      <div class="char-count">
        {{ remainingChars }} characters remaining
      </div>
      <button @click="submitText">Submit</button>
      
      <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
  
      <div>
        <input type="file" @change="handleFileSelection" accept=".pdf" />
        <button @click="uploadPdf">Upload PDF to Backend</button>
        <p v-if="response">Response from Backend: {{ response }}</p>
      </div>
  
      <div v-if="response && response.summary">
        <h2>Summary</h2>
        <p>{{ response.summary }}</p>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios'; // Add this line

  export default {
    data() {
      return {
        text: '',
        response: null,
        errorMessage: '',
        maxChars: 1400,
        pdfFile: null,
      }
    },
    computed: {
      remainingChars() {
        return this.maxChars - this.text.length;
      }
    },
    methods: {
      async submitText() {
        try {
          const res = await fetch('/api/summarize', {
            method: "POST",
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              content: this.text,
              summary: ""
            })
          });
  
          if (res.ok) {
            this.response = await res.json();
            this.errorMessage = ''; // Clear any previous error
            console.log(this.response);
          } else {
            this.errorMessage = 'Failed to get a valid response from the server.';
          }
        } catch (error) {
          console.error('Error sending text to API:', error);
          this.errorMessage = 'Failed to send text to the API. Please try again.';
          this.response = null;
        }
      },
      handleFileSelection(event) {
        this.pdfFile = event.target.files[0];
      },
      async uploadPdf() {
  if (!this.pdfFile) {
    this.errorMessage = 'Please select a PDF file to upload.';
    return;
  }

  const fileData = await this.pdfFile.arrayBuffer(); // Read as binary data

  try {
    const response = await axios.post('/api/postpdf', fileData, {
      headers: {
        'Content-Type': 'application/pdf',
      },
    });

    // Log the response status and data
    console.log('Upload response status:', response.status);
    console.log('Response from backend:', response.data);

    this.response = response.data; // Store the response data
    this.errorMessage = ''; // Clear any previous error
  } catch (error) {
    console.error('Error uploading PDF:', error);
    this.errorMessage = 'Failed to upload PDF. Please try again.';
  }
},

    }
  }

  </script>
  
  <style scoped>
  body {
    font-family: Arial, sans-serif;
    background-color: #f4f7fa;
    color: #333;
    margin: 0;
    padding: 0;
  }
  
  h1 {
    text-align: center;
    color: #4a90e2;
    margin-bottom: 20px;
  }
  
  textarea {
    width: 100%;
    max-width: 600px;
    height: 150px;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
    resize: none;
    font-size: 16px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    margin: 0 auto;
    display: block;
  }
  
  button {
    display: block;
    width: 100%;
    max-width: 200px;
    padding: 10px;
    background-color: #4a90e2;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 16px;
    margin: 20px auto;
    transition: background-color 0.3s;
  }
  
  button:hover {
    background-color: #357ABD;
  }
  
  .error {
    color: red;
    text-align: center;
    margin-top: 10px;
  }
  
  div {
    text-align: center;
  }
  </style>
  