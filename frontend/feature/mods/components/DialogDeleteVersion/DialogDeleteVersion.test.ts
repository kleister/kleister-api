import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const deleteVersionMock = vi.fn()

vi.mock('../../../../client', () => ({
  deleteVersion: (...args: unknown[]) => deleteVersionMock(...args),
}))

import DialogDeleteVersion from './DialogDeleteVersion.vue'

const version = { id: 'version-1', name: '1.0.0' }

describe('DialogDeleteVersion', () => {
  beforeEach(() => {
    deleteVersionMock.mockReset()
    deleteVersionMock.mockResolvedValue({
      data: { message: 'ok' },
      error: undefined,
    })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('deletes the version scoped to the mod when confirming', async () => {
    const wrapper = mount(DialogDeleteVersion, {
      props: { modId: 'mod-1', version },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete version"]').trigger('click')
    await flushPromises()

    const deleteButton = Array.from(document.querySelectorAll('button')).find(
      (button) => button.textContent === 'Delete'
    )
    deleteButton!.dispatchEvent(new MouseEvent('click', { bubbles: true }))
    await flushPromises()

    expect(deleteVersionMock).toHaveBeenCalledWith({
      path: { mod_id: 'mod-1', version_id: 'version-1' },
    })
    expect(wrapper.emitted('deleted')).toBeTruthy()

    wrapper.unmount()
  })

  it('deletes immediately without the modal when shift-clicking', async () => {
    const wrapper = mount(DialogDeleteVersion, {
      props: { modId: 'mod-1', version },
      attachTo: document.body,
    })

    await wrapper
      .find('button[aria-label="Delete version"]')
      .trigger('click', { shiftKey: true })
    await flushPromises()

    expect(deleteVersionMock).toHaveBeenCalledWith({
      path: { mod_id: 'mod-1', version_id: 'version-1' },
    })
    expect(wrapper.emitted('deleted')).toBeTruthy()

    wrapper.unmount()
  })
})
