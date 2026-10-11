import { flushPromises, mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'

const showGroupMock = vi.fn()

vi.mock('../../../client', () => ({
  showGroup: (...args: unknown[]) => showGroupMock(...args),
}))

const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => ({ params: { group_id: '1' } }),
  useRouter: () => ({ push: routerPushMock }),
}))

import GroupDetail from './GroupDetail.vue'

describe('GroupDetail', () => {
  beforeEach(() => {
    showGroupMock.mockReset()
    routerPushMock.mockReset()
  })

  it('loads and renders the group by id from the route', async () => {
    showGroupMock.mockResolvedValue({
      data: {
        id: '1',
        name: 'Admins',
        slug: 'admins',
        created_at: '2026-01-01T00:00:00Z',
        updated_at: '2026-01-02T00:00:00Z',
      },
      error: undefined,
    })

    const wrapper = mount(GroupDetail)
    await flushPromises()

    expect(showGroupMock).toHaveBeenCalledWith({ path: { group_id: '1' } })
    expect(wrapper.text()).toContain('Admins')
    expect(wrapper.text()).toContain('admins')
  })

  it('navigates back to the groups list', async () => {
    showGroupMock.mockResolvedValue({
      data: { id: '1', name: 'Admins', slug: 'admins' },
      error: undefined,
    })

    const wrapper = mount(GroupDetail)
    await flushPromises()

    await wrapper.find('button').trigger('click')

    expect(routerPushMock).toHaveBeenCalledWith({ name: 'groups' })
  })
})
