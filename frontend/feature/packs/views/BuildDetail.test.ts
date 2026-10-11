import { flushPromises, mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'

const showBuildMock = vi.fn()
const listBuildVersionsMock = vi.fn()
const attachBuildToVersionMock = vi.fn()
const deleteBuildFromVersionMock = vi.fn()

vi.mock('../../../client', () => ({
  showBuild: (...args: unknown[]) => showBuildMock(...args),
  listBuildVersions: (...args: unknown[]) => listBuildVersionsMock(...args),
  attachBuildToVersion: (...args: unknown[]) =>
    attachBuildToVersionMock(...args),
  deleteBuildFromVersion: (...args: unknown[]) =>
    deleteBuildFromVersionMock(...args),
}))

const routerPushMock = vi.fn()

vi.mock('vue-router', () => ({
  useRoute: () => ({ params: { pack_id: 'pack-1', build_id: 'build-1' } }),
  useRouter: () => ({ push: routerPushMock }),
}))

import BuildDetail from './BuildDetail.vue'

describe('BuildDetail', () => {
  beforeEach(() => {
    showBuildMock.mockReset()
    listBuildVersionsMock.mockReset()
    attachBuildToVersionMock.mockReset()
    deleteBuildFromVersionMock.mockReset()

    showBuildMock.mockResolvedValue({
      data: { id: 'build-1', name: '1.0.0' },
      error: undefined,
    })
    listBuildVersionsMock.mockResolvedValue({
      data: {
        total: 1,
        limit: 50,
        offset: 0,
        versions: [
          {
            version_id: 'version-1',
            version: { id: 'version-1', name: '1.2.3', mod: { id: 'mod-1' } },
            build_id: 'build-1',
          },
        ],
      },
      error: undefined,
    })
  })

  it('loads the build and its attached versions', async () => {
    const wrapper = mount(BuildDetail)
    await flushPromises()

    expect(showBuildMock).toHaveBeenCalledWith({
      path: { pack_id: 'pack-1', build_id: 'build-1' },
    })
    expect(listBuildVersionsMock).toHaveBeenCalledWith({
      path: { pack_id: 'pack-1', build_id: 'build-1' },
    })
    expect(wrapper.text()).toContain('1.2.3')
  })

  it('attaches a version by mod and version identifier', async () => {
    attachBuildToVersionMock.mockResolvedValue({
      data: { message: 'ok' },
      error: undefined,
    })

    const wrapper = mount(BuildDetail, { attachTo: document.body })
    await flushPromises()

    await wrapper.find('#attach-mod').setValue('mod-2')
    await wrapper.find('#attach-version').setValue('version-2')
    await wrapper.find('form').trigger('submit')
    await flushPromises()

    expect(attachBuildToVersionMock).toHaveBeenCalledWith({
      path: { pack_id: 'pack-1', build_id: 'build-1' },
      body: { mod: 'mod-2', version: 'version-2' },
    })

    wrapper.unmount()
  })

  it('removes an attached version', async () => {
    deleteBuildFromVersionMock.mockResolvedValue({
      data: { message: 'ok' },
      error: undefined,
    })

    const wrapper = mount(BuildDetail, { attachTo: document.body })
    await flushPromises()

    await wrapper.find('button[aria-label="Remove 1.2.3"]').trigger('click')
    await flushPromises()

    expect(deleteBuildFromVersionMock).toHaveBeenCalledWith({
      path: { pack_id: 'pack-1', build_id: 'build-1' },
      body: { mod: 'mod-1', version: 'version-1' },
    })

    wrapper.unmount()
  })
})
