import { createRouter, createWebHistory } from 'vue-router'


const router = createRouter({
  history: createWebHistory(),
  routes: [
   
    {
      path: '/history',
      name: 'history',
      component: () => import('../views/History.vue')
    },
  ]
})

export default router