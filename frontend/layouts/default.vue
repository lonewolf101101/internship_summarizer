<template>
    <div class="layout">
      <Header/>
      <main class="content">
        <NuxtPage />
      </main>
      <Footer/>
    </div>
  </template>
  
  <script>
  import Header from '~/components/header.vue';
  import Footer from '~/components/footer.vue';
 
  export default {
    components: {
      Header,
      Footer,
    },
  };
  </script>
  
  <style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

body {
  font-family: Arial, sans-serif;
  background-color: #f0f0f0;
  color: #333;
  line-height: 1.6;
}

.content {
  flex: 1;
  padding: 20px;
  }
  </style>
  