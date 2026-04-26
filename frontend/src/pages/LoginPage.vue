<template>
  <div class="container">
    <h2>Login</h2>

    <form @submit.prevent="login">
      <input v-model="email" placeholder="Email" />
      <input v-model="password" type="password" placeholder="Password" />

      <button>Login</button>
    </form>

    <p v-if="error">{{ error }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const email = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()

const login = async () => {
  try {
    const res = await axios.post('http://localhost:8080/dashboard/v1/auth/login', {
      email: email.value,
      password: password.value
    })

    const data = res.data.data

    localStorage.setItem('token', data.token)
    localStorage.setItem('role', data.role)

    router.push('/')
  } catch (err) {
    error.value = 'Login gagal'
  }
}
</script>