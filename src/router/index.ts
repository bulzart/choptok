import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import SubscriptionPage from '../components/SubscriptionPage.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/upgrade',
      name: 'upgrade',
      component: SubscriptionPage
    },
    {
      path: '/history',
      name: 'history',
      component: () => import('../views/History.vue')
    },
    {
      path: '/favorites',
      name: 'favorites',
      component: () => import('../views/Favorites.vue')
    },
    {
      path: '/settings',
      name: 'settings',
      component: () => import('../views/Settings.vue')
    },
    {
      path: '/help',
      name: 'help',
      component: () => import('../views/Help.vue')
    }
  ]
})

export default router