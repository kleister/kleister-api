import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { h } from 'vue'

const listModsMock = vi.fn()
const createModMock = vi.fn()

vi.mock('../../../../client', () => ({
  listMods: (...args: unknown[]) => listModsMock(...args),
  createMod: (...args: unknown[]) => createModMock(...args),
}))

let currentRoute: { query: Record<string, string> }
const routerReplaceMock = vi.fn()
const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => currentRoute,
  useRouter: () => ({ replace: routerReplaceMock, push: routerPushMock }),
}))

import ModsProvider from '../../providers'
import ModsList from './ModsList.vue'

function makeMod(id: string, name: string) {
  return {
    id,
    name,
    slug: name.toLowerCase(),
    author: 'someone',
    public: false,
    created_at: '2026-01-01T00:00:00Z',
    updated_at: '2026-01-01T00:00:00Z',
  }
}

function mountModsList() {
  return mount(ModsProvider, {
    slots: {
      default: () => h(ModsList),
    },
  })
}

describe('ModsList', () => {
  beforeEach(() => {
    currentRoute = { query: {} }
    routerReplaceMock.mockReset()
    routerPushMock.mockReset()
    listModsMock.mockReset()
    createModMock.mockReset()
    listModsMock.mockResolvedValue({
      data: {
        mods: [makeMod('1', 'JEI'), makeMod('2', 'Waila')],
        total: 2,
        limit: 50,
        offset: 0,
      },
      error: undefined,
    })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('loads the first page with the default sorting on mount', async () => {
    mountModsList()
    await flushPromises()

    expect(listModsMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'name',
        order: 'asc',
        limit: 50,
        offset: 0,
      },
    })
  })

  it('renders the loaded mods', async () => {
    const wrapper = mountModsList()
    await flushPromises()

    expect(wrapper.text()).toContain('JEI')
    expect(wrapper.text()).toContain('Waila')
  })

  it('shows the total number of mods', async () => {
    const wrapper = mountModsList()
    await flushPromises()

    expect(wrapper.text()).toContain('2 mod(s) total')
  })

  it('renders details, edit and delete actions for each row', async () => {
    const wrapper = mountModsList()
    await flushPromises()

    expect(
      wrapper.findAll('button[aria-label="View mod details"]')
    ).toHaveLength(2)
    expect(wrapper.findAll('button[aria-label="Edit mod"]')).toHaveLength(2)
    expect(wrapper.findAll('button[aria-label="Delete mod"]')).toHaveLength(2)
  })

  it('navigates to the mod detail page when "details" is clicked', async () => {
    const wrapper = mountModsList()
    await flushPromises()

    await wrapper
      .findAll('button[aria-label="View mod details"]')[0]
      .trigger('click')

    expect(routerPushMock).toHaveBeenCalledWith({
      name: 'mod-detail',
      params: { mod_id: '1' },
    })
  })
})
