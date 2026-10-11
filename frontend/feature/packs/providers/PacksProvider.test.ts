import { mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import { defineComponent, h } from 'vue'

const listPacksMock = vi.fn()

vi.mock('../../../client', () => ({
  listPacks: (...args: unknown[]) => listPacksMock(...args),
}))

import PacksProvider from './PacksProvider.vue'
import { usePacks } from './context'

function mountWithContext() {
  let context!: NonNullable<ReturnType<typeof usePacks>>

  const Consumer = defineComponent({
    setup() {
      context = usePacks()
      return () => h('div')
    },
  })

  mount(PacksProvider, {
    slots: {
      default: () => h(Consumer),
    },
  })

  return context
}

describe('PacksProvider', () => {
  beforeEach(() => {
    listPacksMock.mockReset()
  })

  it('starts with an empty, non-loading state', () => {
    const { packs, total, isLoading } = mountWithContext()

    expect(packs.value).toEqual([])
    expect(total.value).toBe(0)
    expect(isLoading.value).toBe(false)
  })

  it('loads packs and stores the total', async () => {
    listPacksMock.mockResolvedValue({
      data: {
        packs: [{ id: '1', name: 'Vanilla Plus' }],
        total: 1,
        limit: 10,
        offset: 0,
      },
      error: undefined,
    })

    const { packs, total, isLoading, loadPacks } = mountWithContext()

    await loadPacks({ limit: 10, offset: 0 })

    expect(packs.value).toHaveLength(1)
    expect(total.value).toBe(1)
    expect(isLoading.value).toBe(false)
  })

  it('throws and stops loading on error', async () => {
    listPacksMock.mockResolvedValue({
      data: undefined,
      error: { status: 500 },
    })

    const { isLoading, loadPacks } = mountWithContext()

    await expect(loadPacks()).rejects.toEqual({ status: 500 })
    expect(isLoading.value).toBe(false)
  })

  it('reuses the last query when called without arguments', async () => {
    listPacksMock.mockResolvedValue({
      data: { packs: [], total: 0, limit: 10, offset: 0 },
      error: undefined,
    })

    const { loadPacks } = mountWithContext()

    await loadPacks({ search: 'vanilla' })
    await loadPacks()

    expect(listPacksMock).toHaveBeenNthCalledWith(1, {
      query: { search: 'vanilla' },
    })
    expect(listPacksMock).toHaveBeenNthCalledWith(2, {
      query: { search: 'vanilla' },
    })
  })
})
