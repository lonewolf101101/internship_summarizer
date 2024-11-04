<template>
  <div class="container">
    <h1>Text Summarizer</h1>
    <textarea 
      v-model="text" 
      placeholder="Enter your text here..." 
      :maxlength="maxChars"
    ></textarea>
    <div class="char-count">
      {{ remainingChars }} characters remaining
    </div>
    <button @click="submitText" :disabled="isLoading">
      <span v-if="isLoading" class="loader"></span>
      <span v-else>Submit</span>
    </button>

    <div v-if="errorMessage" class="error">{{ errorMessage }}</div>

    <div>
      <input type="file" @change="handleFileSelection" accept=".pdf" />
      <button @click="uploadPdf">Upload PDF to Backend</button>
    </div>

    <div v-if="response && response.summary">
      <h2>Summary</h2>
      <textarea v-model="response.summary" rows="5"></textarea>
      <button @click="copyToClipboard">Copy to Clipboard</button>
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
        maxChars: 2000,
        pdfFile: null,
        isLoading: false, // Add this line
      }
    },
    computed: {
      remainingChars() {
        return this.maxChars - this.text.length;
      }
    },
    methods: {
      copyToClipboard() {
        const textarea = document.createElement('textarea');
        textarea.value = this.response.summary;
        document.body.appendChild(textarea);
        textarea.select();
        document.execCommand('copy');
        document.body.removeChild(textarea);
        this.errorMessage = 'Summary copied to clipboard!';
      },
      async submitText() {
        this.isLoading = true; // Set loading to true
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
        } finally {
          this.isLoading = false; // Set loading to false after request completes
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

  const formData = new FormData(); // Create a FormData object
  formData.append('file', this.pdfFile); // Append the PDF file

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
}

    }
  }

  </script>
  
  <style scoped>
  .loader {
    border: 2px solid #f3f3f3; /* Light grey */
    border-top: 2px solid #3498db; /* Blue */
    border-radius: 50%;
    width: 16px;
    height: 16px;
    animation: spin 1s linear infinite;
    display: inline-block; /* To align it with the text */
    margin-right: 5px; /* Space between loader and text */
  }

  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }


  h1 {
    text-align: center;
    color: #4a90e2;
    margin-bottom: 20px;
  }
  
  .container {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
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
  