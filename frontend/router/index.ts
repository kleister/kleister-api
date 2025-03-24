import { unref } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'

import NotFound from '@/feature/not-found/views/NotFound.vue'
import { useAuthStore } from '@/feature/auth/store/auth'

const routes = [
  {
    name: 'signin',
    path: '/auth/signin',
    component: () => import('../feature/auth/views/Signin.vue'),
  },
  {
    name: 'callback',
    path: '/auth/callback/:redirect_token',
    component: () => import('../feature/auth/views/Callback.vue'),
  },

  {
    name: 'welcome',
    path: '/',
    component: () => import('../feature/dashboard/views/Dashboard.vue'),
    meta: { auth: true },
  },

  {
    name: 'users',
    path: '/admin/users',
    component: () => import('../feature/users/views/Users.vue'),
    meta: { auth: true, admin: true },
  },
  {
    name: 'groups',
    path: '/admin/groups',
    component: () => import('../feature/groups/views/Groups.vue'),
    meta: { auth: true, admin: true },
  },

  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFound,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to) => {
  const authStore = useAuthStore()

  if (to.meta.auth && !unref(authStore.isAuthed) && to.name !== 'signin') {
    return { name: 'signin', query: { redirect: to.fullPath } }
  }

  if (to.name === 'signin' && unref(authStore.isAuthed)) {
    return { name: 'welcome' }
  }

  if (to.meta.admin && !unref(authStore.user.isAdmin)) {
    return { name: 'welcome' }
  }
})

export default router
