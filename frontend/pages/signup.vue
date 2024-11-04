<template>
  <form @submit.prevent="handleSubmit" class="form">
    <h1 class="form-title">Шинээр бүртгүүлэх</h1>

    <div class="form-field">
      <label class="form-label" for="username">Username</label>
      <input
        type="text"
        id="username"
        v-model="formData.username" 
        class="form-input"
      />
    </div>

    <div class="form-field">
      <label class="form-label" for="email">Email</label>
      <input
        type="email"
        id="email"
        v-model="formData.email"
        class="form-input"
      />
    </div>

    <button type="submit" class="form-submit">register</button>
    <p>or</p>
    <div class="login-buttons">
      <button @click="loginWithGoogle" class="form-button">register with Google</button>
    </div>
  </form>
 
</template>

  
<script>
export default {
  data() {
    return {
      formData: {
        username: '',
        email: '',
        urgainOvog: '',
        description: ''
      }
    };
  },
  methods: {
    handleSubmit() {
      // Handle form submission logic here
      console.log('Form submitted:', this.formData);
    },
    loginWithGoogle() {
      window.location.href = "http://localhost:3300/pub/auth/google/login";
    },
  }
};
</script>

<style scoped>
p {
  margin: 0; /* Remove default margin */
  text-align: center; /* Center text */
}
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
.form {
  max-width: 50rem;
  margin: 0 auto; /* Center the form */
  padding: 2rem; /* Padding for the container */
  border: 2px solid #e5e7eb; /* Border color */
  border-radius: 0.75rem; /* Rounded corners */
  display: flex;
  flex-direction: column;
  gap: 1rem; /* Space between elements */
  background-color: #ffffff; /* Background color */
}

.form-title {
  font-size: 1.875rem; /* Tailwind's text-3xl */
  font-weight: bold; /* Tailwind's font-bold */
  color: #1f2937; /* Tailwind's text-gray-800 */
  margin-bottom: 1rem; /* Tailwind's mb-4 */
}

.form-label {
  display: block;
  color: #4b5563; /* Tailwind's text-gray-600 */
  margin-bottom: 0.25rem; /* Tailwind's mb-1 */
}

.form-input {
  width: 100%;
  padding: 0.5rem; /* Tailwind's p-2 */
  border: 1px solid #d1d5db; /* Tailwind's border-gray-300 */
  border-radius: 0.375rem; /* Tailwind's rounded-lg */
}

.form-submit {
  background-color: #3498db; /* Tailwind's bg-blue-500 */
  color: #ffffff; /* Tailwind's text-white */
  padding: 0.5rem; /* Tailwind's py-2 */
  border-radius: 0.375rem; /* Tailwind's rounded */
  cursor: pointer;
}
</style>