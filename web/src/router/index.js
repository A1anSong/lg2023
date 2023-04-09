import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  // {
  //   path: '/init',
  //   name: 'Init',
  //   component: () => import('@/view/init/index.vue')
  // },
  {
    path: '/elogValidate',
    name: 'ElogValidate',
    component: () => import('@/view/lg/elogValidate/elogValidate.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/view/login/index.vue')
  },
  {
    path: '/:catchAll(.*)',
    meta: {
      closeTab: true,
    },
    component: () => import('@/view/error/index.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
