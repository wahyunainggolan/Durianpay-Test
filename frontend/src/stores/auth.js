import { defineStore } from "pinia"
import api from "../api/http"

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: localStorage.getItem("token") || "",
    role: localStorage.getItem("role") || "",
  }),

  actions: {
    async login(email, password) {
      const res = await api.post("/auth/login", { email, password })

      this.token = res.data.data.token
      this.role = res.data.data.role

      localStorage.setItem("token", this.token)
      localStorage.setItem("role", this.role)
    },

    logout() {
      this.token = ""
      localStorage.clear()
    },
  },
})