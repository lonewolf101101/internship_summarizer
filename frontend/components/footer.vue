<template>
    <footer class="footer">
      <div class="container">
        <p>&copy; 2024 Summarizer. All rights reserved.</p>
        <nav>
          <a href="/about">About</a> | <a href="/contact">Contact</a>
        </nav>
      </div>
    </footer>
  </template>
  
  <script>
  export default {};
  </script>
  
  <style scoped>
  .footer {
    padding: 20px;
    background-color: #333;
    color: white;
    text-align: center;
  }
  
  .container {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  nav a {
    color: #ccc;
    text-decoration: none;
    margin: 0 5px;
  }
  
  nav a:hover {
    color: white;
  }
  </style>
  