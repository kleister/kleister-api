import { flushPromises, mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'

const showPackMock = vi.fn()
const listPackUsersMock = vi.fn()
const listPackGroupsMock = vi.fn()
const listBuildsMock = vi.fn()

vi.mock('../../../client', () => ({
  showPack: (...args: unknown[]) => showPackMock(...args),
  listPackUsers: (...args: unknown[]) => listPackUsersMock(...args),
  attachPackToUser: vi.fn(),
  permitPackUser: vi.fn(),
  deletePackFromUser: vi.fn(),
  listPackGroups: (...args: unknown[]) => listPackGroupsMock(...args),
  attachPackToGroup: vi.fn(),
  permitPackGroup: vi.fn(),
  deletePackFromGroup: vi.fn(),
  listBuilds: (...args: unknown[]) => listBuildsMock(...args),
}))

const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => ({ params: { pack_id: '1' } }),
  useRouter: () => ({ push: routerPushMock }),
}))

import PackDetail from './PackDetail.vue'

describe('PackDetail', () => {
  beforeEach(() => {
    showPackMock.mockReset()
    listPackUsersMock.mockReset()
    listPackGroupsMock.mockReset()
    listBuildsMock.mockReset()
    listBuildsMock.mockResolvedValue({
      data: { builds: [], total: 0, limit: 50, offset: 0 },
      error: undefined,
    })
  })

  it('loads and renders the pack by id from the route', async () => {
    showPackMock.mockResolvedValue({
      data: {
        id: '1',
        name: 'Vanilla Plus',
        slug: 'vanilla-plus',
        website: 'https://example.com',
        public: true,
        created_at: '2026-01-01T00:00:00Z',
        updated_at: '2026-01-02T00:00:00Z',
      },
      error: undefined,
    })

    const wrapper = mount(PackDetail)
    await flushPromises()

    expect(showPackMock).toHaveBeenCalledWith({ path: { pack_id: '1' } })
    expect(wrapper.text()).toContain('Vanilla Plus')
    expect(wrapper.text()).toContain('vanilla-plus')
  })

  it('navigates back to the packs list', async () => {
    showPackMock.mockResolvedValue({
      data: { id: '1', name: 'Vanilla Plus', slug: 'vanilla-plus' },
      error: undefined,
    })

    const wrapper = mount(PackDetail)
    await flushPromises()

    await wrapper.find('button').trigger('click')

    expect(routerPushMock).toHaveBeenCalledWith({ name: 'packs' })
  })
})
