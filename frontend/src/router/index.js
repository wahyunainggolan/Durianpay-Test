import { createRouter, createWebHistory } from "vue-router"
import LoginPage from "../pages/LoginPage.vue"
import DashboardPage from "../pages/DashboardPage.vue"

const isTokenValid = (token) => {
  if (!token) return false

  try {
    const payload = JSON.parse(atob(token.split(".")[1]))
    return payload.exp * 1000 > Date.now()
  } catch (e) {
    return false
  }
}

const routes = [
  {
    path: "/login",
    name: "login",
    component: LoginPage,
    meta: { guestOnly: true }
  },
  {
    path: "/",
    name: "dashboard",
    component: DashboardPage,
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token")
  const isAuth = isTokenValid(token)

  if (to.meta.requiresAuth && !isAuth) {
    localStorage.removeItem("token")
    return next({ name: "login" })
  }

  if (to.meta.guestOnly && isAuth) {
    return next({ name: "dashboard" })
  }

  next()
})

export default router