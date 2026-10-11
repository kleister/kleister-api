import { mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import { defineComponent, h } from 'vue'

const listModsMock = vi.fn()

vi.mock('../../../client', () => ({
  listMods: (...args: unknown[]) => listModsMock(...args),
}))

import ModsProvider from './ModsProvider.vue'
import { useMods } from './context'

function mountWithContext() {
  let context!: NonNullable<ReturnType<typeof useMods>>

  const Consumer = defineComponent({
    setup() {
      context = useMods()
      return () => h('div')
    },
  })

  mount(ModsProvider, {
    slots: {
      default: () => h(Consumer),
    },
  })

  return context
}

describe('ModsProvider', () => {
  beforeEach(() => {
    listModsMock.mockReset()
  })

  it('starts with an empty, non-loading state', () => {
    const { mods, total, isLoading } = mountWithContext()

    expect(mods.value).toEqual([])
    expect(total.value).toBe(0)
    expect(isLoading.value).toBe(false)
  })

  it('loads mods and stores the total', async () => {
    listModsMock.mockResolvedValue({
      data: {
        mods: [{ id: '1', name: 'JEI' }],
        total: 1,
        limit: 10,
        offset: 0,
      },
      error: undefined,
    })

    const { mods, total, isLoading, loadMods } = mountWithContext()

    await loadMods({ limit: 10, offset: 0 })

    expect(mods.value).toHaveLength(1)
    expect(total.value).toBe(1)
    expect(isLoading.value).toBe(false)
  })

  it('throws and stops loading on error', async () => {
    listModsMock.mockResolvedValue({
      data: undefined,
      error: { status: 500 },
    })

    const { isLoading, loadMods } = mountWithContext()

    await expect(loadMods()).rejects.toEqual({ status: 500 })
    expect(isLoading.value).toBe(false)
  })

  it('reuses the last query when called without arguments', async () => {
    listModsMock.mockResolvedValue({
      data: { mods: [], total: 0, limit: 10, offset: 0 },
      error: undefined,
    })

    const { loadMods } = mountWithContext()

    await loadMods({ search: 'jei' })
    await loadMods()

    expect(listModsMock).toHaveBeenNthCalledWith(1, {
      query: { search: 'jei' },
    })
    expect(listModsMock).toHaveBeenNthCalledWith(2, {
      query: { search: 'jei' },
    })
  })
})
