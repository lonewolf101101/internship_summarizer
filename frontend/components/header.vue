<template>
  <header class="header">
    <div class="container">
      <h1 @click="navigateToHome" class="logo">Summarizer</h1>
      <nav class="nav">
        <!-- <button  @click="navigateToLogin">Login</button>
        <button  @click="navigateToSignup">Sign Up</button> -->
        <!-- <button v-else @click="logout">Logout</button> -->
      </nav>
    </div>
  </header>
</template>

<script>
export default {
  // computed: {
  //   // Assuming the user login status is stored in Vuex or can be accessed globally
  //   isLoggedIn() {
  //     return this.$store.state.user.isLoggedIn; // Adjust based on your store structure
  //   },
  // },
  methods: {
    navigateToHome() {
      this.$router.push('/');
    },
    // navigateToLogin() {
    //   this.$router.push('/login');
    // },
    // navigateToSignup() {
    //   this.$router.push('/signup');
    // },
    // logout() {
    //   // Perform any necessary logout actions (e.g., clear session, call logout API, etc.)
    //   // this.$store.commit('logoutUser'); // Assuming you have a Vuex mutation to log out
    //   this.$router.push('/');
    // },
  },
};
</script>

<style scoped>
.header {
  display: flex;
  align-items: center;
  padding: 1rem 1.5rem;
  background-color: #333;
  color: white;
}

.container {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

.logo {
  font-size: 24px;
  font-weight: bold;
}

.nav button {
  margin-left: 1rem;
  margin-top: 0.2rem;
  padding: 0.8rem 1rem;
  background-color: #555;
  color: white;
  border: none;
  border-radius: 10px;
  cursor: pointer;
}

.nav button:hover {
  background-color: #777;
}
</style>
