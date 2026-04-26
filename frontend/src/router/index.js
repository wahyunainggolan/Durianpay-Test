import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '../pages/LoginPage.vue'
import DashboardPage from '../pages/DashboardPage.vue'

const routes = [
  { path: '/login', component: LoginPage },
  { path: '/', component: DashboardPage }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 🔐 Proteksi route
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')

  if (to.path !== '/login' && !token) {
    return next('/login')
  }

  if (to.path === '/login' && token) {
    return next('/')
  }

  next()
})

export default router