<template>
    <div class="container">
      <h1>Submit Text and PDF</h1>
      <form @submit.prevent>
        <div>
          <label>Enter Text:</label>
          <input
            v-model="textInput"
            type="text"
            placeholder="Enter some text"
          />
        </div>
        <div>
          <label>Upload PDF:</label>
          <input
            type="file"
            accept="application/pdf"
            @change="handlePdfChange"
          />
        </div>
        <button type="submit">Submit</button>
      </form>
  
      <div v-if="pdfFile">
        <h2>Uploaded Data:</h2>
        <p><strong>Text:</strong> {{ textInput }}</p>
        <p><strong>PDF File:</strong> {{ pdfFile.name }}</p>
      </div>
    </div>
</template>
  
  <script>
  import '~/assets/css/first.css';
  
  export default {
    data() {
      return {
        textInput: '',
        pdfFile: null,
      };
    },
    methods: {
      handlePdfChange(event) {
        this.pdfFile = event.target.files[0];
      },
    },
  };
  </script>
  