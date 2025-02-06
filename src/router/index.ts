import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import Dashboard from '../views/Dashboard.vue'
import MyVideos from '../views/MyVideos.vue'
import Analytics from '../views/Analytics.vue'
import Settings from '../views/Settings.vue'
import Help from '../views/Help.vue'
import Login from '../views/auth/Login.vue'
import Register from '../views/auth/Register.vue'
import SubscriptionPage from '../views/SubscriptionPage.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Dashboard',
      component: Dashboard,
      meta: { requiresAuth: true }
    },
    {
      path: '/videos',
      name: 'MyVideos',
      component: MyVideos,
      meta: { requiresAuth: true }
    },
    {
      path: '/subscriptions',
      name: 'Subscriptions',
      component: SubscriptionPage,
      meta: { requiresAuth: true }
    },
    {
      path: '/settings',
      name: 'Settings',
      component: Settings,
      meta: { requiresAuth: true }
    },
    {
      path: '/help',
      name: 'Help',
      component: Help,
      meta: { requiresAuth: true }
    },
    {
      path: '/auth/login',
      name: 'Login',
      component: Login,
      meta: { requiresAuth: false }
    },
    {
      path: '/auth/register',
      name: 'Register',
      component: Register,
      meta: { requiresAuth: false }
    }
  ]
})

// Navigation guard

export default router