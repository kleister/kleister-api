import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { h } from 'vue'

const listPacksMock = vi.fn()
const createPackMock = vi.fn()

vi.mock('../../../../client', () => ({
  listPacks: (...args: unknown[]) => listPacksMock(...args),
  createPack: (...args: unknown[]) => createPackMock(...args),
}))

let currentRoute: { query: Record<string, string> }
const routerReplaceMock = vi.fn()
const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => currentRoute,
  useRouter: () => ({ replace: routerReplaceMock, push: routerPushMock }),
}))

import PacksProvider from '../../providers'
import PacksList from './PacksList.vue'

function makePack(id: string, name: string) {
  return {
    id,
    name,
    slug: name.toLowerCase(),
    website: 'https://example.com',
    public: false,
    created_at: '2026-01-01T00:00:00Z',
    updated_at: '2026-01-01T00:00:00Z',
  }
}

function mountPacksList() {
  return mount(PacksProvider, {
    slots: {
      default: () => h(PacksList),
    },
  })
}

describe('PacksList', () => {
  beforeEach(() => {
    currentRoute = { query: {} }
    routerReplaceMock.mockReset()
    routerPushMock.mockReset()
    listPacksMock.mockReset()
    createPackMock.mockReset()
    listPacksMock.mockResolvedValue({
      data: {
        packs: [makePack('1', 'Vanilla Plus'), makePack('2', 'Tech Heavy')],
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
    mountPacksList()
    await flushPromises()

    expect(listPacksMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'name',
        order: 'asc',
        limit: 50,
        offset: 0,
      },
    })
  })

  it('renders the loaded packs', async () => {
    const wrapper = mountPacksList()
    await flushPromises()

    expect(wrapper.text()).toContain('Vanilla Plus')
    expect(wrapper.text()).toContain('Tech Heavy')
  })

  it('shows the total number of packs', async () => {
    const wrapper = mountPacksList()
    await flushPromises()

    expect(wrapper.text()).toContain('2 pack(s) total')
  })

  it('renders details, edit and delete actions for each row', async () => {
    const wrapper = mountPacksList()
    await flushPromises()

    expect(
      wrapper.findAll('button[aria-label="View pack details"]')
    ).toHaveLength(2)
    expect(wrapper.findAll('button[aria-label="Edit pack"]')).toHaveLength(2)
    expect(wrapper.findAll('button[aria-label="Delete pack"]')).toHaveLength(
      2
    )
  })

  it('navigates to the pack detail page when "details" is clicked', async () => {
    const wrapper = mountPacksList()
    await flushPromises()

    await wrapper
      .findAll('button[aria-label="View pack details"]')[0]
      .trigger('click')

    expect(routerPushMock).toHaveBeenCalledWith({
      name: 'pack-detail',
      params: { pack_id: '1' },
    })
  })
})
