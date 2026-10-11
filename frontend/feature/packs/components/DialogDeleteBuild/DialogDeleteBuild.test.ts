import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const deleteBuildMock = vi.fn()

vi.mock('../../../../client', () => ({
  deleteBuild: (...args: unknown[]) => deleteBuildMock(...args),
}))

import DialogDeleteBuild from './DialogDeleteBuild.vue'

const build = { id: 'build-1', name: '1.0.0' }

describe('DialogDeleteBuild', () => {
  beforeEach(() => {
    deleteBuildMock.mockReset()
    deleteBuildMock.mockResolvedValue({
      data: { message: 'ok' },
      error: undefined,
    })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('deletes the build scoped to the pack when confirming', async () => {
    const wrapper = mount(DialogDeleteBuild, {
      props: { packId: 'pack-1', build },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete build"]').trigger('click')
    await flushPromises()

    const deleteButton = Array.from(document.querySelectorAll('button')).find(
      (button) => button.textContent === 'Delete'
    )
    deleteButton!.dispatchEvent(new MouseEvent('click', { bubbles: true }))
    await flushPromises()

    expect(deleteBuildMock).toHaveBeenCalledWith({
      path: { pack_id: 'pack-1', build_id: 'build-1' },
    })
    expect(wrapper.emitted('deleted')).toBeTruthy()

    wrapper.unmount()
  })

  it('deletes immediately without the modal when shift-clicking', async () => {
    const wrapper = mount(DialogDeleteBuild, {
      props: { packId: 'pack-1', build },
      attachTo: document.body,
    })

    await wrapper
      .find('button[aria-label="Delete build"]')
      .trigger('click', { shiftKey: true })
    await flushPromises()

    expect(deleteBuildMock).toHaveBeenCalledWith({
      path: { pack_id: 'pack-1', build_id: 'build-1' },
    })
    expect(wrapper.emitted('deleted')).toBeTruthy()

    wrapper.unmount()
  })
})
