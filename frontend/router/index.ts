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
    path: '/users',
    component: () => import('../feature/users/views/Users.vue'),
    meta: { auth: true, admin: true },
  },
  {
    name: 'user-detail',
    path: '/users/:user_id',
    component: () => import('../feature/users/views/UserDetail.vue'),
    meta: { auth: true, admin: true },
  },
  {
    name: 'groups',
    path: '/groups',
    component: () => import('../feature/groups/views/Groups.vue'),
    meta: { auth: true, admin: true },
  },
  {
    name: 'group-detail',
    path: '/groups/:group_id',
    component: () => import('../feature/groups/views/GroupDetail.vue'),
    meta: { auth: true, admin: true },
  },
  {
    name: 'packs',
    path: '/packs',
    component: () => import('../feature/packs/views/Packs.vue'),
    meta: { auth: true },
  },
  {
    name: 'pack-detail',
    path: '/packs/:pack_id',
    component: () => import('../feature/packs/views/PackDetail.vue'),
    meta: { auth: true },
  },
  {
    name: 'pack-build-detail',
    path: '/packs/:pack_id/builds/:build_id',
    component: () => import('../feature/packs/views/BuildDetail.vue'),
    meta: { auth: true },
  },
  {
    name: 'mods',
    path: '/mods',
    component: () => import('../feature/mods/views/Mods.vue'),
    meta: { auth: true },
  },
  {
    name: 'mod-detail',
    path: '/mods/:mod_id',
    component: () => import('../feature/mods/views/ModDetail.vue'),
    meta: { auth: true },
  },
  {
    name: 'mod-version-detail',
    path: '/mods/:mod_id/versions/:version_id',
    component: () => import('../feature/mods/views/VersionDetail.vue'),
    meta: { auth: true },
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
