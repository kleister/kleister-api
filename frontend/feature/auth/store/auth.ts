import { client } from '../../../client/client.gen'
import { defineStore } from 'pinia'
import { computed, reactive, unref } from 'vue'
import { useRouter } from 'vue-router'

const AUTH_STORAGE_KEYS = Object.freeze({
  accessToken: 'kleister.auth.access_token',
})

interface ParsedBearerToken {
  admin: boolean
  profile: string
  email: string
  exp: number
  iat: number
  ident: string
  iss: string
  login: string
  name: string
}

interface Token {
  accessToken: string
  expires: number
}

interface User {
  displayName: string
  email: string
  profile: string
  isAdmin: boolean
}

function parseJwt(token: string): ParsedBearerToken {
  const base64Url = token.split('.')[1]
  const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
  const jsonPayload = decodeURIComponent(
    atob(base64)
      .split('')
      .map(function (c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
      })
      .join('')
  )

  return JSON.parse(jsonPayload) satisfies ParsedBearerToken
}

export const useAuthStore = defineStore('auth', () => {
  const router = useRouter()

  const user = reactive<User>({
    displayName: '',
    email: '',
    profile: '',
    isAdmin: false,
  })
  const token = reactive<Token>({ accessToken: '', expires: 0 })

  const isAuthed = computed(() => unref(token).accessToken !== '')

  function signInUser(accessToken: string) {
    const parsedToken = parseJwt(accessToken)

    if (parsedToken.exp < Date.now() / 1000) {
      localStorage.removeItem(AUTH_STORAGE_KEYS.accessToken)
      return
    }

    Object.assign(token, {
      accessToken: accessToken,
      expires: parsedToken.exp,
    })
    Object.assign(user, {
      displayName: parsedToken.name,
      email: parsedToken.email,
      profile: parsedToken.profile,
      isAdmin: parsedToken.admin,
    })

    localStorage.setItem(AUTH_STORAGE_KEYS.accessToken, accessToken)
    client.setConfig({ headers: { Authorization: `Bearer ${accessToken}` } })
  }

  function init() {
    const accessToken = localStorage.getItem(AUTH_STORAGE_KEYS.accessToken)

    if (!accessToken) {
      return
    }

    signInUser(accessToken)
  }

  async function signOutUser() {
    Object.assign(user, {
      displayName: '',
      email: '',
      profile: '',
      isAdmin: false,
    })
    Object.assign(token, { accessToken: '', expires: 0 })

    localStorage.removeItem(AUTH_STORAGE_KEYS.accessToken)
    client.setConfig({ headers: { Authorization: undefined } })
    await router.push({ name: 'signin' })
  }

  return { user, token, isAuthed, signInUser, init, signOutUser }
})
