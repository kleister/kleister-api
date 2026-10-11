import { flushPromises, mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'

const showVersionMock = vi.fn()
const listVersionBuildsMock = vi.fn()
const attachVersionToBuildMock = vi.fn()
const deleteVersionFromBuildMock = vi.fn()

vi.mock('../../../client', () => ({
  showVersion: (...args: unknown[]) => showVersionMock(...args),
  listVersionBuilds: (...args: unknown[]) => listVersionBuildsMock(...args),
  attachVersionToBuild: (...args: unknown[]) =>
    attachVersionToBuildMock(...args),
  deleteVersionFromBuild: (...args: unknown[]) =>
    deleteVersionFromBuildMock(...args),
}))

const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => ({ params: { mod_id: 'mod-1', version_id: 'version-1' } }),
  useRouter: () => ({ push: routerPushMock }),
}))

import VersionDetail from './VersionDetail.vue'

describe('VersionDetail', () => {
  beforeEach(() => {
    showVersionMock.mockReset()
    listVersionBuildsMock.mockReset()
    attachVersionToBuildMock.mockReset()
    deleteVersionFromBuildMock.mockReset()

    showVersionMock.mockResolvedValue({
      data: { id: 'version-1', name: '1.2.3' },
      error: undefined,
    })
    listVersionBuildsMock.mockResolvedValue({
      data: {
        total: 1,
        limit: 50,
        offset: 0,
        builds: [
          {
            build_id: 'build-1',
            build: { id: 'build-1', name: '1.0.0', pack: { id: 'pack-1' } },
            version_id: 'version-1',
          },
        ],
      },
      error: undefined,
    })
  })

  it('loads the version and its attached builds', async () => {
    const wrapper = mount(VersionDetail)
    await flushPromises()

    expect(showVersionMock).toHaveBeenCalledWith({
      path: { mod_id: 'mod-1', version_id: 'version-1' },
    })
    expect(listVersionBuildsMock).toHaveBeenCalledWith({
      path: { mod_id: 'mod-1', version_id: 'version-1' },
    })
    expect(wrapper.text()).toContain('1.0.0')
  })

  it('attaches a build by pack and build identifier', async () => {
    attachVersionToBuildMock.mockResolvedValue({
      data: { message: 'ok' },
      error: undefined,
    })

    const wrapper = mount(VersionDetail, { attachTo: document.body })
    await flushPromises()

    await wrapper.find('#attach-pack').setValue('pack-2')
    await wrapper.find('#attach-build').setValue('build-2')
    await wrapper.find('form').trigger('submit')
    await flushPromises()

    expect(attachVersionToBuildMock).toHaveBeenCalledWith({
      path: { mod_id: 'mod-1', version_id: 'version-1' },
      body: { pack: 'pack-2', build: 'build-2' },
    })

    wrapper.unmount()
  })

  it('removes an attached build', async () => {
    deleteVersionFromBuildMock.mockResolvedValue({
      data: { message: 'ok' },
      error: undefined,
    })

    const wrapper = mount(VersionDetail, { attachTo: document.body })
    await flushPromises()

    await wrapper.find('button[aria-label="Remove 1.0.0"]').trigger('click')
    await flushPromises()

    expect(deleteVersionFromBuildMock).toHaveBeenCalledWith({
      path: { mod_id: 'mod-1', version_id: 'version-1' },
      body: { pack: 'pack-1', build: 'build-1' },
    })

    wrapper.unmount()
  })
})
