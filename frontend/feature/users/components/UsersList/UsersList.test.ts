import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { h } from 'vue'

const listUsersMock = vi.fn()
const createUserMock = vi.fn()

vi.mock('../../../../client', () => ({
  listUsers: (...args: unknown[]) => listUsersMock(...args),
  createUser: (...args: unknown[]) => createUserMock(...args),
}))

let currentRoute: { query: Record<string, string> }
const routerReplaceMock = vi.fn()
const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => currentRoute,
  useRouter: () => ({ replace: routerReplaceMock, push: routerPushMock }),
}))

import UsersProvider from '../../providers'
import UsersList from './UsersList.vue'

function makeUser(id: string, username: string) {
  return {
    id,
    username,
    email: `${username}@example.com`,
    fullname: username,
    admin: false,
    active: true,
    created_at: '2026-01-01T00:00:00Z',
    updated_at: '2026-01-01T00:00:00Z',
  }
}

function mountUsersList() {
  return mount(UsersProvider, {
    slots: {
      default: () => h(UsersList),
    },
  })
}

describe('UsersList', () => {
  beforeEach(() => {
    currentRoute = { query: {} }
    routerReplaceMock.mockReset()
    routerPushMock.mockReset()
    listUsersMock.mockReset()
    createUserMock.mockReset()
    listUsersMock.mockResolvedValue({
      data: {
        users: [makeUser('1', 'john'), makeUser('2', 'jane')],
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
    mountUsersList()
    await flushPromises()

    expect(listUsersMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'username',
        order: 'asc',
        limit: 50,
        offset: 0,
      },
    })
  })

  it('renders the loaded users', async () => {
    const wrapper = mountUsersList()
    await flushPromises()

    expect(wrapper.text()).toContain('john')
    expect(wrapper.text()).toContain('jane')
  })

  it('shows the total number of users', async () => {
    const wrapper = mountUsersList()
    await flushPromises()

    expect(wrapper.text()).toContain('2 user(s) total')
  })

  it('requests the server with the search term after debouncing', async () => {
    const wrapper = mountUsersList()
    await flushPromises()
    listUsersMock.mockClear()

    await wrapper.find('input[placeholder="Search..."]').setValue('john')
    await new Promise((resolve) => setTimeout(resolve, 350))
    await flushPromises()

    expect(listUsersMock).toHaveBeenCalledWith({
      query: {
        search: 'john',
        sort: 'username',
        order: 'asc',
        limit: 50,
        offset: 0,
      },
    })
  })

  it('requests the server with the new sort order when a column header is clicked', async () => {
    const wrapper = mountUsersList()
    await flushPromises()
    listUsersMock.mockClear()

    const emailHeaderButton = wrapper
      .findAll('button')
      .find((button) => button.text().includes('Email'))

    expect(emailHeaderButton).toBeTruthy()
    await emailHeaderButton!.trigger('click')
    await flushPromises()

    expect(listUsersMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'email',
        order: 'asc',
        limit: 50,
        offset: 0,
      },
    })
  })

  it('requests the next page when "Next" is clicked', async () => {
    listUsersMock.mockResolvedValue({
      data: {
        users: [makeUser('1', 'john'), makeUser('2', 'jane')],
        total: 120,
        limit: 50,
        offset: 0,
      },
      error: undefined,
    })

    const wrapper = mountUsersList()
    await flushPromises()
    listUsersMock.mockClear()

    const nextButton = wrapper
      .findAll('button')
      .find((button) => button.text() === 'Next')

    expect(nextButton).toBeTruthy()
    await nextButton!.trigger('click')
    await flushPromises()

    expect(listUsersMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'username',
        order: 'asc',
        limit: 50,
        offset: 50,
      },
    })
  })

  it('restores search, sort, order, page and page size from the URL on mount', async () => {
    currentRoute = {
      query: {
        search: 'jane',
        sort: 'email',
        order: 'desc',
        page: '2',
        limit: '20',
      },
    }
    listUsersMock.mockResolvedValue({
      data: {
        users: [makeUser('2', 'jane')],
        total: 25,
        limit: 20,
        offset: 20,
      },
      error: undefined,
    })

    mountUsersList()
    await flushPromises()

    expect(listUsersMock).toHaveBeenCalledWith({
      query: {
        search: 'jane',
        sort: 'email',
        order: 'desc',
        limit: 20,
        offset: 20,
      },
    })
  })

  it('falls back to the default state for an invalid URL query', async () => {
    currentRoute = { query: { sort: 'not-a-column', page: '-1', limit: '999' } }

    mountUsersList()
    await flushPromises()

    expect(listUsersMock).toHaveBeenCalledWith({
      query: {
        search: undefined,
        sort: 'username',
        order: 'asc',
        limit: 50,
        offset: 0,
      },
    })
  })

  it('reflects the current state in the URL so the view can be shared', async () => {
    listUsersMock.mockResolvedValue({
      data: {
        users: [makeUser('1', 'john'), makeUser('2', 'jane')],
        total: 120,
        limit: 50,
        offset: 0,
      },
      error: undefined,
    })

    const wrapper = mountUsersList()
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
        sort: 'username',
        order: 'asc',
        page: 2,
        limit: 50,
      },
    })
  })

  it('renders details, edit and delete actions for each row', async () => {
    const wrapper = mountUsersList()
    await flushPromises()

    expect(
      wrapper.findAll('button[aria-label="View user details"]')
    ).toHaveLength(2)
    expect(wrapper.findAll('button[aria-label="Edit user"]')).toHaveLength(2)
    expect(wrapper.findAll('button[aria-label="Delete user"]')).toHaveLength(2)
  })

  it('navigates to the user detail page when "details" is clicked', async () => {
    const wrapper = mountUsersList()
    await flushPromises()

    await wrapper
      .findAll('button[aria-label="View user details"]')[0]
      .trigger('click')

    expect(routerPushMock).toHaveBeenCalledWith({
      name: 'user-detail',
      params: { user_id: '1' },
    })
  })
})
