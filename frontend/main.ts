import { createApp } from 'vue'

import './assets/index.css'

import App from './App.vue'
import router from './router'

import { createPinia } from 'pinia'
import { useAuthStore } from './feature/auth/store/auth'

import { client } from './client/client.gen'
client.setConfig({ baseUrl: '/api/v1' })

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)

const { init: initAuth } = useAuthStore()
initAuth()

app.mount('#app')
