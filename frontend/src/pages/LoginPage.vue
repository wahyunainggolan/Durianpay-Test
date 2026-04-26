<template>
  <div class="login-page">
    <div class="login-card">
      <div class="login-header">
        <h1>Welcome Back</h1>
        <p>Sign in to your dashboard</p>
      </div>
      <form class="login-form" @submit.prevent="login">
        <div class="input-group">
          <label>Email</label>
          <input
            v-model="email"
            type="email"
            placeholder="you@company.com"
          />
        </div>
        <div class="input-group">
          <label>Password</label>
          <input
            v-model="password"
            type="password"
            placeholder="••••••••"
          />
        </div>
        <p v-if="error" class="error-msg">
          {{ error }}
        </p>
        <button
          class="login-btn"
          :disabled="loading"
        >
          <span v-if="loading">Signing in...</span>
          <span v-else>Login</span>
        </button>

      </form>

      <div class="login-footer">
        <small>© 2026 DurianPay</small>
      </div>

    </div>

  </div>
</template>

<script setup>
import "../styles/pages/login.css"
import { ref } from "vue"
import { useRouter } from "vue-router"
import axios from "axios"

const email = ref("")
const password = ref("")
const error = ref("")
const loading = ref(false)
const router = useRouter()

const login = async () => {
  try {
    loading.value = true
    error.value = ""

    const res = await axios.post(
      "http://localhost:8080/dashboard/v1/auth/login",
      {
        email: email.value,
        password: password.value
      }
    )

    const data = res.data.data

    localStorage.setItem("token", data.token)
    localStorage.setItem("role", data.role)

    router.push("/")

  } catch (err) {
    error.value = "Invalid email or password"
  } finally {
    loading.value = false
  }
}
</script>