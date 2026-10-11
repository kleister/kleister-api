import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { h } from 'vue'

const listGroupsMock = vi.fn()
const createGroupMock = vi.fn()

vi.mock('../../../../client', () => ({
  listGroups: (...args: unknown[]) => listGroupsMock(...args),
  createGroup: (...args: unknown[]) => createGroupMock(...args),
}))

let currentRoute: { query: Record<string, string> }
const routerReplaceMock = vi.fn()
const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => currentRoute,
  useRouter: () => ({ replace: routerReplaceMock, push: routerPushMock }),
}))

import GroupsProvider from '../../providers'
import GroupsList from './GroupsList.vue'

function makeGroup(id: string, name: string) {
  return {
    id,
    name,
    slug: name.toLowerCase(),
    created_at: '2026-01-01T00:00:00Z',
    updated_at: '2026-01-01T00:00:00Z',
  }
}

function mountGroupsList() {
  return mount(GroupsProvider, {
    slots: {
      default: () => h(GroupsList),
    },
  })
}

describe('GroupsList', () => {
  beforeEach(() => {
    currentRoute = { query: {} }
    routerReplaceMock.mockReset()
    routerPushMock.mockReset()
    listGroupsMock.mockReset()
    createGroupMock.mockReset()
    listGroupsMock.mockResolvedValue({
      data: {
        groups: [makeGroup('1', 'Admins'), makeGroup('2', 'Editors')],
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
    mountGroupsList()
    await flushPromises()

    expect(listGroupsMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'name',
        order: 'asc',
        limit: 50,
        offset: 0,
      },
    })
  })

  it('renders the loaded groups', async () => {
    const wrapper = mountGroupsList()
    await flushPromises()

    expect(wrapper.text()).toContain('Admins')
    expect(wrapper.text()).toContain('Editors')
  })

  it('shows the total number of groups', async () => {
    const wrapper = mountGroupsList()
    await flushPromises()

    expect(wrapper.text()).toContain('2 group(s) total')
  })

  it('requests the server with the search term after debouncing', async () => {
    const wrapper = mountGroupsList()
    await flushPromises()
    listGroupsMock.mockClear()

    await wrapper.find('input[placeholder="Search..."]').setValue('admins')
    await new Promise((resolve) => setTimeout(resolve, 350))
    await flushPromises()

    expect(listGroupsMock).toHaveBeenCalledWith({
      query: {
        search: 'admins',
        sort: 'name',
        order: 'asc',
        limit: 50,
        offset: 0,
      },
    })
  })

  it('requests the server with the new sort order when a column header is clicked', async () => {
    const wrapper = mountGroupsList()
    await flushPromises()
    listGroupsMock.mockClear()

    const slugHeaderButton = wrapper
      .findAll('button')
      .find((button) => button.text().includes('Slug'))

    expect(slugHeaderButton).toBeTruthy()
    await slugHeaderButton!.trigger('click')
    await flushPromises()

    expect(listGroupsMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'slug',
        order: 'asc',
        limit: 50,
        offset: 0,
      },
    })
  })

  it('requests the next page when "Next" is clicked', async () => {
    listGroupsMock.mockResolvedValue({
      data: {
        groups: [makeGroup('1', 'Admins'), makeGroup('2', 'Editors')],
        total: 120,
        limit: 50,
        offset: 0,
      },
      error: undefined,
    })

    const wrapper = mountGroupsList()
    await flushPromises()
    listGroupsMock.mockClear()

    const nextButton = wrapper
      .findAll('button')
      .find((button) => button.text() === 'Next')

    expect(nextButton).toBeTruthy()
    await nextButton!.trigger('click')
    await flushPromises()

    expect(listGroupsMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'name',
        order: 'asc',
        limit: 50,
        offset: 50,
      },
    })
  })

  it('restores search, sort, order, page and page size from the URL on mount', async () => {
    currentRoute = {
      query: {
        search: 'editors',
        sort: 'slug',
        order: 'desc',
        page: '2',
        limit: '20',
      },
    }
    listGroupsMock.mockResolvedValue({
      data: {
        groups: [makeGroup('2', 'Editors')],
        total: 25,
        limit: 20,
        offset: 20,
      },
      error: undefined,
    })

    mountGroupsList()
    await flushPromises()

    expect(listGroupsMock).toHaveBeenCalledWith({
      query: {
        search: 'editors',
        sort: 'slug',
        order: 'desc',
        limit: 20,
        offset: 20,
      },
    })
  })

  it('falls back to the default state for an invalid URL query', async () => {
    currentRoute = { query: { sort: 'not-a-column', page: '-1', limit: '999' } }

    mountGroupsList()
    await flushPromises()

    expect(listGroupsMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'name',
        order: 'asc',
        limit: 50,
        offset: 0,
      },
    })
  })

  it('reflects the current state in the URL so the view can be shared', async () => {
    listGroupsMock.mockResolvedValue({
      data: {
        groups: [makeGroup('1', 'Admins'), makeGroup('2', 'Editors')],
        total: 120,
        limit: 50,
        offset: 0,
      },
      error: undefined,
    })

    const wrapper = mountGroupsList()
    await flushPromises()
    routerReplaceMock.mockClear()

    const nextButton = wrapper
      .findAll('button')
      .find((button) => button.text() === 'Next')
    await nextButton!.trigger('click')
    await flushPromises()

    expect(routerReplaceMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'name',
        order: 'asc',
        page: 2,
        limit: 50,
      },
    })
  })

  it('renders details, edit and delete actions for each row', async () => {
    const wrapper = mountGroupsList()
    await flushPromises()

    expect(
      wrapper.findAll('button[aria-label="View group details"]')
    ).toHaveLength(2)
    expect(wrapper.findAll('button[aria-label="Edit group"]')).toHaveLength(2)
    expect(wrapper.findAll('button[aria-label="Delete group"]')).toHaveLength(2)
  })

  it('navigates to the group detail page when "details" is clicked', async () => {
    const wrapper = mountGroupsList()
    await flushPromises()

    await wrapper
      .findAll('button[aria-label="View group details"]')[0]
      .trigger('click')

    expect(routerPushMock).toHaveBeenCalledWith({
      name: 'group-detail',
      params: { group_id: '1' },
    })
  })
})
