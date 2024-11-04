<template>
  <div class="login-container">
    <h1 class="login-title">Welcome to the App!</h1>
    <form @submit.prevent="loginWithEmail" class="login-form">
      <input v-model="email" type="email" placeholder="Email" required class="form-input" />
      <input v-model="password" type="password" placeholder="Password" required class="form-input" />
      <button type="submit" class="form-submit">Login</button>
    </form>
    <div class="login-buttons">
      <button @click="loginWithGoogle" class="form-button">Login with Google</button>
    </div>
  </div>
</template>


<script>
export default {
  data() {
    return {
      email: '',
      password: '',
    };
  },
  methods: {
    loginWithGoogle() {
      window.location.href = "http://localhost:3300/pub/auth/google/login";
    },
    async logout() {
      const response = await this.$axios.get('/pub/auth/logout');
    },
    async loginWithEmail() {
  try {
    const response = await useFetch('/api/pub/auth/basic/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ // Ensure you stringify the body
        email: this.email,
        password: this.password,
      }),
    });

    // Assume the response is JSON formatted
    if (!response.ok) { 
      const errorData = await response.json();
      throw new Error(errorData.message || 'Login failed');
    }

    const data = await response.json(); // Parse the JSON response
    console.log('Login successful:', data);

    // Example of reading a cookie after login
    const myCookieValue = this.getCookie('myCookieName'); // Replace 'myCookieName' with the actual cookie name
    console.log('Cookie Value:', myCookieValue);
    
    // Optionally store user info in Vuex or local storage
    this.$store.commit('setUser', data.user);
    
    // Redirect to another page
    this.$router.push('/profile');

  } catch (error) {
    console.error('Login failed:', error);
    // Optionally show an error message to the user
  }
},

getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);

  if (parts.length === 2) {
    return parts.pop().split(';').shift();
  }
  return null; // Return null if the cookie doesn't exist
},

  },
};
</script>

<style scoped>
.login-container {
  width: 100%;
  max-width: 50rem; /* Optional: Limit the width of the form */
  margin: 0 auto; /* Center the form */
  padding: 2rem; /* Padding for the container */
  border: 2px solid #e5e7eb; /* Border color */
  border-radius: 0.75rem; /* Rounded corners */
  display: flex;
  flex-direction: column;
  gap: 1rem; /* Space between elements */
  background-color: #ffffff; /* Background color */
}

.login-title {
  font-size: 1.875rem; /* Font size for the title */
  font-weight: bold; /* Bold title */
  color: #1f2937; /* Dark text color */
  text-align: center; /* Center the title */
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 1rem; /* Space between inputs */
}

.form-input {
  width: 100%;
  padding: 0.5rem; /* Padding for inputs */
  border: 1px solid #d1d5db; /* Border color */
  border-radius: 0.375rem; /* Rounded corners */
}

.form-submit,
.form-button {
  width: 100%; /* Full width buttons */
  padding: 0.5rem; /* Padding for buttons */
  border: none; /* Remove default border */
  border-radius: 0.375rem; /* Rounded corners */
  background-color: #3498db; /* Button background color */
  color: #ffffff; /* White text color */
  cursor: pointer; /* Pointer cursor on hover */
  font-size: 1rem; /* Font size for buttons */
}

.form-button {
  background-color: #4b5563; /* Darker background for secondary buttons */
}

.form-submit:hover,
.form-button:hover {
  opacity: 0.9; /* Slightly reduce opacity on hover */
}
</style>
