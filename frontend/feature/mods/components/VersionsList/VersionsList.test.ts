import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const listVersionsMock = vi.fn()

vi.mock('../../../../client', () => ({
  listVersions: (...args: unknown[]) => listVersionsMock(...args),
}))

const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRouter: () => ({ push: routerPushMock }),
}))

import VersionsList from './VersionsList.vue'

function makeVersion(id: string, name: string) {
  return {
    id,
    name,
    public: true,
    created_at: '2026-01-01T00:00:00Z',
    updated_at: '2026-01-01T00:00:00Z',
  }
}

describe('VersionsList', () => {
  beforeEach(() => {
    listVersionsMock.mockReset()
    routerPushMock.mockReset()
    listVersionsMock.mockResolvedValue({
      data: {
        versions: [makeVersion('1', '1.0.0'), makeVersion('2', '1.1.0')],
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

  it('loads versions for the given mod on mount', async () => {
    mount(VersionsList, { props: { modId: 'mod-1' } })
    await flushPromises()

    expect(listVersionsMock).toHaveBeenCalledWith({
      path: { mod_id: 'mod-1' },
    })
  })

  it('renders the loaded versions', async () => {
    const wrapper = mount(VersionsList, { props: { modId: 'mod-1' } })
    await flushPromises()

    expect(wrapper.text()).toContain('1.0.0')
    expect(wrapper.text()).toContain('1.1.0')
    expect(wrapper.text()).toContain('2 version(s) total')
  })

  it('navigates to the version detail page when "view" is clicked', async () => {
    const wrapper = mount(VersionsList, { props: { modId: 'mod-1' } })
    await flushPromises()

    await wrapper
      .findAll('button[aria-label="View version details"]')[0]
      .trigger('click')

    expect(routerPushMock).toHaveBeenCalledWith({
      name: 'mod-version-detail',
      params: { mod_id: 'mod-1', version_id: '1' },
    })
  })
})
