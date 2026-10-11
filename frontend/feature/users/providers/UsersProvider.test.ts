import { mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import { defineComponent, h } from 'vue'

const listUsersMock = vi.fn()

vi.mock('../../../client', () => ({
  listUsers: (...args: unknown[]) => listUsersMock(...args),
}))

import UsersProvider from './UsersProvider.vue'
import { useUsers } from './context'

function mountWithContext() {
  let context!: NonNullable<ReturnType<typeof useUsers>>

  const Consumer = defineComponent({
    setup() {
      context = useUsers()
      return () => h('div')
    },
  })

  mount(UsersProvider, {
    slots: {
      default: () => h(Consumer),
    },
  })

  return context
}

describe('UsersProvider', () => {
  beforeEach(() => {
    listUsersMock.mockReset()
  })

  it('starts with an empty, non-loading state', () => {
    const { users, total, isLoading } = mountWithContext()

    expect(users.value).toEqual([])
    expect(total.value).toBe(0)
    expect(isLoading.value).toBe(false)
  })

  it('loads users and stores the total', async () => {
    listUsersMock.mockResolvedValue({
      data: {
        users: [{ id: '1', username: 'john' }],
        total: 1,
        limit: 10,
        offset: 0,
      },
      error: undefined,
    })

    const { users, total, isLoading, loadUsers } = mountWithContext()

    await loadUsers({ limit: 10, offset: 0 })

    expect(users.value).toHaveLength(1)
    expect(total.value).toBe(1)
    expect(isLoading.value).toBe(false)
  })

  it('throws and stops loading on error', async () => {
    listUsersMock.mockResolvedValue({
      data: undefined,
      error: { status: 500 },
    })

    const { isLoading, loadUsers } = mountWithContext()

    await expect(loadUsers()).rejects.toEqual({ status: 500 })
    expect(isLoading.value).toBe(false)
  })

  it('reuses the last query when called without arguments', async () => {
    listUsersMock.mockResolvedValue({
      data: { users: [], total: 0, limit: 10, offset: 0 },
      error: undefined,
    })

    const { loadUsers } = mountWithContext()

    await loadUsers({ search: 'john' })
    await loadUsers()

    expect(listUsersMock).toHaveBeenNthCalledWith(1, {
      query: { search: 'john' },
    })
    expect(listUsersMock).toHaveBeenNthCalledWith(2, {
      query: { search: 'john' },
    })
  })
})
