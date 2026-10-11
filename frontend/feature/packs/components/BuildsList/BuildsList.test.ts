import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const listBuildsMock = vi.fn()

vi.mock('../../../../client', () => ({
  listBuilds: (...args: unknown[]) => listBuildsMock(...args),
}))

const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRouter: () => ({ push: routerPushMock }),
}))

import BuildsList from './BuildsList.vue'

function makeBuild(id: string, name: string) {
  return {
    id,
    name,
    latest: true,
    recommended: false,
    public: true,
    created_at: '2026-01-01T00:00:00Z',
    updated_at: '2026-01-01T00:00:00Z',
  }
}

describe('BuildsList', () => {
  beforeEach(() => {
    listBuildsMock.mockReset()
    routerPushMock.mockReset()
    listBuildsMock.mockResolvedValue({
      data: {
        builds: [makeBuild('1', '1.0.0'), makeBuild('2', '1.1.0')],
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

  it('loads builds for the given pack on mount', async () => {
    mount(BuildsList, { props: { packId: 'pack-1' } })
    await flushPromises()

    expect(listBuildsMock).toHaveBeenCalledWith({
      path: { pack_id: 'pack-1' },
    })
  })

  it('renders the loaded builds', async () => {
    const wrapper = mount(BuildsList, { props: { packId: 'pack-1' } })
    await flushPromises()

    expect(wrapper.text()).toContain('1.0.0')
    expect(wrapper.text()).toContain('1.1.0')
    expect(wrapper.text()).toContain('2 build(s) total')
  })

  it('navigates to the build detail page when "view" is clicked', async () => {
    const wrapper = mount(BuildsList, { props: { packId: 'pack-1' } })
    await flushPromises()

    await wrapper
      .findAll('button[aria-label="View build details"]')[0]
      .trigger('click')

    expect(routerPushMock).toHaveBeenCalledWith({
      name: 'pack-build-detail',
      params: { pack_id: 'pack-1', build_id: '1' },
    })
  })
})
