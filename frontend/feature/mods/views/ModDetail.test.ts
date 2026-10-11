import { flushPromises, mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'

const showModMock = vi.fn()
const listModUsersMock = vi.fn()
const listModGroupsMock = vi.fn()
const listVersionsMock = vi.fn()

vi.mock('../../../client', () => ({
  showMod: (...args: unknown[]) => showModMock(...args),
  listModUsers: (...args: unknown[]) => listModUsersMock(...args),
  attachModToUser: vi.fn(),
  permitModUser: vi.fn(),
  deleteModFromUser: vi.fn(),
  listModGroups: (...args: unknown[]) => listModGroupsMock(...args),
  attachModToGroup: vi.fn(),
  permitModGroup: vi.fn(),
  deleteModFromGroup: vi.fn(),
  listVersions: (...args: unknown[]) => listVersionsMock(...args),
}))

const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => ({ params: { mod_id: '1' } }),
  useRouter: () => ({ push: routerPushMock }),
}))

import ModDetail from './ModDetail.vue'

describe('ModDetail', () => {
  beforeEach(() => {
    showModMock.mockReset()
    listModUsersMock.mockReset()
    listModGroupsMock.mockReset()
    listVersionsMock.mockReset()
    listVersionsMock.mockResolvedValue({
      data: { versions: [], total: 0, limit: 50, offset: 0 },
      error: undefined,
    })
  })

  it('loads and renders the mod by id from the route', async () => {
    showModMock.mockResolvedValue({
      data: {
        id: '1',
        name: 'JEI',
        slug: 'jei',
        author: 'mezz',
        created_at: '2026-01-01T00:00:00Z',
        updated_at: '2026-01-02T00:00:00Z',
      },
      error: undefined,
    })

    const wrapper = mount(ModDetail)
    await flushPromises()

    expect(showModMock).toHaveBeenCalledWith({ path: { mod_id: '1' } })
    expect(wrapper.text()).toContain('JEI')
    expect(wrapper.text()).toContain('jei')
  })

  it('navigates back to the mods list', async () => {
    showModMock.mockResolvedValue({
      data: { id: '1', name: 'JEI', slug: 'jei' },
      error: undefined,
    })

    const wrapper = mount(ModDetail)
    await flushPromises()

    await wrapper.find('button').trigger('click')

    expect(routerPushMock).toHaveBeenCalledWith({ name: 'mods' })
  })
})
