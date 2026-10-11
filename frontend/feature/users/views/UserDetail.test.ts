import { flushPromises, mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'

const showUserMock = vi.fn()

vi.mock('../../../client', () => ({
  showUser: (...args: unknown[]) => showUserMock(...args),
}))

const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => ({ params: { user_id: '1' } }),
  useRouter: () => ({ push: routerPushMock }),
}))

import UserDetail from './UserDetail.vue'

describe('UserDetail', () => {
  beforeEach(() => {
    showUserMock.mockReset()
    routerPushMock.mockReset()
  })

  it('loads and renders the user by id from the route', async () => {
    showUserMock.mockResolvedValue({
      data: {
        id: '1',
        username: 'john',
        email: 'john@example.com',
        fullname: 'John Doe',
        admin: true,
        active: true,
        created_at: '2026-01-01T00:00:00Z',
        updated_at: '2026-01-02T00:00:00Z',
      },
      error: undefined,
    })

    const wrapper = mount(UserDetail)
    await flushPromises()

    expect(showUserMock).toHaveBeenCalledWith({ path: { user_id: '1' } })
    expect(wrapper.text()).toContain('john')
    expect(wrapper.text()).toContain('john@example.com')
    expect(wrapper.text()).toContain('John Doe')
  })

  it('navigates back to the users list', async () => {
    showUserMock.mockResolvedValue({
      data: { id: '1', username: 'john' },
      error: undefined,
    })

    const wrapper = mount(UserDetail)
    await flushPromises()

    await wrapper.find('button').trigger('click')

    expect(routerPushMock).toHaveBeenCalledWith({ name: 'users' })
  })
})
