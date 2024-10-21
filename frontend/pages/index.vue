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

    <div v-if="response">
      <h2>Summary</h2>
      <p>{{ response.summary }}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      text: '',                 // User's input text
      response: null,           // API response
      errorMessage: '',         // Error message
      maxChars: 1400             // Maximum number of characters allowed
    }
  },
  computed: {
    remainingChars() {
      return this.maxChars - this.text.length; // Calculate remaining characters
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
    }
  }
}
</script>
<style scoped>
body {
  font-family: Arial, sans-serif; /* Use a clean sans-serif font */
  background-color: #f4f7fa;      /* Light background for contrast */
  color: #333;                    /* Darker text for readability */
  margin: 0;
  padding: 0;
}

h1 {
  text-align: center;
  color: #4a90e2;                /* A calming blue for the heading */
  margin-bottom: 20px;
}

h2 {
  color: #333;                    /* Dark color for subheadings */
  margin-top: 20px;
  margin-bottom: 10px;
}

textarea {
  width: 100%;                    /* Full width of the container */
  max-width: 600px;               /* Max width for larger screens */
  height: 150px;                  /* Set a height for the textarea */
  padding: 10px;                  /* Inner padding for better text positioning */
  border: 1px solid #ccc;         /* Light border */
  border-radius: 4px;             /* Rounded corners */
  resize: none;                   /* Disable resizing */
  font-size: 16px;                /* Font size for better readability */
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1); /* Subtle shadow for depth */
  margin: 0 auto;                 /* Centering */
  display: block;                 /* Block element */
}

button {
  display: block;                 /* Make the button a block element */
  width: 100%;                    /* Full width */
  max-width: 200px;              /* Max width for larger screens */
  padding: 10px;                  /* Padding for the button */
  background-color: #4a90e2;      /* Button color */
  color: #fff;                    /* Text color */
  border: none;                   /* Remove border */
  border-radius: 4px;            /* Rounded corners */
  cursor: pointer;                /* Pointer cursor */
  font-size: 16px;                /* Font size */
  margin: 20px auto;              /* Centering */
  transition: background-color 0.3s; /* Smooth transition */
}

button:hover {
  background-color: #357ABD;     /* Darker shade on hover */
}

.error {
  color: red;                    /* Error message color */
  text-align: center;            /* Center the error message */
  margin-top: 10px;              /* Space above the error message */
}

div {
  text-align: center;             /* Centering all divs */
}
</style>

