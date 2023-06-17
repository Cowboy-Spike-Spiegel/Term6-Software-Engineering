import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      // component: HomeView
      component: () => import('../views/LoginRegister.vue')
    },
    {
      path: '/home',
      name: 'home',
      // component: HomeView
      component: () => import('../views/HomeView.vue')
    },

  ]
})

export default router
