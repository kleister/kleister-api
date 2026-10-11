import { mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import { defineComponent, h } from 'vue'

const listGroupsMock = vi.fn()

vi.mock('../../../client', () => ({
  listGroups: (...args: unknown[]) => listGroupsMock(...args),
}))

import GroupsProvider from './GroupsProvider.vue'
import { useGroups } from './context'

function mountWithContext() {
  let context!: NonNullable<ReturnType<typeof useGroups>>

  const Consumer = defineComponent({
    setup() {
      context = useGroups()
      return () => h('div')
    },
  })

  mount(GroupsProvider, {
    slots: {
      default: () => h(Consumer),
    },
  })

  return context
}

describe('GroupsProvider', () => {
  beforeEach(() => {
    listGroupsMock.mockReset()
  })

  it('starts with an empty, non-loading state', () => {
    const { groups, total, isLoading } = mountWithContext()

    expect(groups.value).toEqual([])
    expect(total.value).toBe(0)
    expect(isLoading.value).toBe(false)
  })

  it('loads groups and stores the total', async () => {
    listGroupsMock.mockResolvedValue({
      data: {
        groups: [{ id: '1', name: 'Admins' }],
        total: 1,
        limit: 10,
        offset: 0,
      },
      error: undefined,
    })

    const { groups, total, isLoading, loadGroups } = mountWithContext()

    await loadGroups({ limit: 10, offset: 0 })

    expect(groups.value).toHaveLength(1)
    expect(total.value).toBe(1)
    expect(isLoading.value).toBe(false)
  })

  it('throws and stops loading on error', async () => {
    listGroupsMock.mockResolvedValue({
      data: undefined,
      error: { status: 500 },
    })

    const { isLoading, loadGroups } = mountWithContext()

    await expect(loadGroups()).rejects.toEqual({ status: 500 })
    expect(isLoading.value).toBe(false)
  })

  it('reuses the last query when called without arguments', async () => {
    listGroupsMock.mockResolvedValue({
      data: { groups: [], total: 0, limit: 10, offset: 0 },
      error: undefined,
    })

    const { loadGroups } = mountWithContext()

    await loadGroups({ search: 'admins' })
    await loadGroups()

    expect(listGroupsMock).toHaveBeenNthCalledWith(1, {
      query: { search: 'admins' },
    })
    expect(listGroupsMock).toHaveBeenNthCalledWith(2, {
      query: { search: 'admins' },
    })
  })
})
