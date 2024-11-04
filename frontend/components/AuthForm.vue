<template>
    <form @submit.prevent="submit">
      <input v-model="email" type="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit">{{ isRegister ? 'Register' : 'Login' }}</button>
    </form>
  </template>
  
  <script setup>
  import { ref, defineEmits, defineProps } from 'vue'
  
  const emit = defineEmits()
  const props = defineProps({
    isRegister: { type: Boolean, default: false }
  })
  
  const email = ref('')
  const password = ref('')
  
  const submit = () => {
    emit('submit', { email: email.value, password: password.value })
  }
  </script>
  