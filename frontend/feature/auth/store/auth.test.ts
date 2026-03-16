import { createPinia, setActivePinia } from 'pinia'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

// Mock the generated API client so we don't need a real server.
vi.mock('../../../client/client.gen', () => ({
  client: { setConfig: vi.fn() },
}))

// Mock vue-router so the store can call useRouter() without a real app.
const mockRouterPush = vi.fn()
vi.mock('vue-router', () => ({
  useRouter: () => ({ push: mockRouterPush }),
}))

// Import after mocks are registered.
import { useAuthStore } from './auth'

function makeTestJwt(
  payload: Record<string, unknown>,
  expiresOffsetSeconds = 3600
): string {
  const now = Math.floor(Date.now() / 1000)
  const fullPayload = { iat: now, exp: now + expiresOffsetSeconds, ...payload }
  const encode = (obj: object) =>
    btoa(JSON.stringify(obj))
      .replace(/\+/g, '-')
      .replace(/\//g, '_')
      .replace(/=/g, '')
  const header = encode({ alg: 'HS256', typ: 'JWT' })
  const body = encode(fullPayload)
  return `${header}.${body}.fakesig`
}

const TEST_PAYLOAD = {
  admin: true,
  email: 'admin@test.com',
  ident: '01ABC123',
  iss: 'kleister',
  login: 'admin',
  name: 'Admin User',
}

describe('useAuthStore', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    localStorage.clear()
    vi.clearAllMocks()
  })

  afterEach(() => {
    localStorage.clear()
  })

  it('is not authenticated by default', () => {
    const store = useAuthStore()
    expect(store.isAuthed).toBe(false)
  })

  describe('signInUser', () => {
    it('sets isAuthed to true with a valid token', () => {
      const store = useAuthStore()
      store.signInUser(makeTestJwt(TEST_PAYLOAD))
      expect(store.isAuthed).toBe(true)
    })

    it('populates user fields from the JWT payload', () => {
      const store = useAuthStore()
      store.signInUser(makeTestJwt(TEST_PAYLOAD))
      expect(store.user.displayName).toBe('Admin User')
      expect(store.user.email).toBe('admin@test.com')
      expect(store.user.isAdmin).toBe(true)
    })

    it('persists the token in localStorage', () => {
      const store = useAuthStore()
      const jwt = makeTestJwt(TEST_PAYLOAD)
      store.signInUser(jwt)
      expect(localStorage.getItem('kleister.auth.access_token')).toBe(jwt)
    })

    it('does not authenticate with an already-expired token', () => {
      const store = useAuthStore()
      store.signInUser(makeTestJwt(TEST_PAYLOAD, -10))
      expect(store.isAuthed).toBe(false)
    })
  })

  describe('init', () => {
    it('restores authentication from localStorage', () => {
      const jwt = makeTestJwt(TEST_PAYLOAD)
      localStorage.setItem('kleister.auth.access_token', jwt)

      const store = useAuthStore()
      store.init()

      expect(store.isAuthed).toBe(true)
      expect(store.user.email).toBe('admin@test.com')
    })

    it('does nothing when localStorage is empty', () => {
      const store = useAuthStore()
      store.init()
      expect(store.isAuthed).toBe(false)
    })
  })

  describe('signOutUser', () => {
    it('clears authentication state', async () => {
      const store = useAuthStore()
      store.signInUser(makeTestJwt(TEST_PAYLOAD))
      expect(store.isAuthed).toBe(true)

      await store.signOutUser()

      expect(store.isAuthed).toBe(false)
      expect(store.user.displayName).toBe('')
      expect(store.user.email).toBe('')
      expect(store.user.isAdmin).toBe(false)
    })

    it('removes the token from localStorage', async () => {
      const store = useAuthStore()
      store.signInUser(makeTestJwt(TEST_PAYLOAD))
      await store.signOutUser()
      expect(localStorage.getItem('kleister.auth.access_token')).toBeNull()
    })

    it('redirects to the sign-in page', async () => {
      const store = useAuthStore()
      store.signInUser(makeTestJwt(TEST_PAYLOAD))
      await store.signOutUser()
      expect(mockRouterPush).toHaveBeenCalledOnce()
    })
  })
})
