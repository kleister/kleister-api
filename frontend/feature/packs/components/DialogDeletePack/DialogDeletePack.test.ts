import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const deletePackMock = vi.fn()

vi.mock('../../../../client', () => ({
  deletePack: (...args: unknown[]) => deletePackMock(...args),
}))

import DialogDeletePack from './DialogDeletePack.vue'

const pack = {
  id: '1',
  name: 'Vanilla Plus',
  slug: 'vanilla-plus',
  created_at: '2026-01-01T00:00:00Z',
  updated_at: '2026-01-01T00:00:00Z',
}

describe('DialogDeletePack', () => {
  beforeEach(() => {
    deletePackMock.mockReset()
    deletePackMock.mockResolvedValue({
      data: { message: 'ok' },
      error: undefined,
    })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('opens a confirmation modal on a plain click and does not delete yet', async () => {
    const wrapper = mount(DialogDeletePack, {
      props: { pack },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete pack"]').trigger('click')
    await flushPromises()

    expect(deletePackMock).not.toHaveBeenCalled()
    expect(document.body.textContent).toContain('Delete pack')
    expect(document.body.textContent).toContain('Vanilla Plus')

    wrapper.unmount()
  })

  it('deletes the pack when confirming in the modal', async () => {
    const wrapper = mount(DialogDeletePack, {
      props: { pack },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete pack"]').trigger('click')
    await flushPromises()

    const deleteButton = Array.from(document.querySelectorAll('button')).find(
      (button) => button.textContent === 'Delete'
    )
    expect(deleteButton).toBeTruthy()

    deleteButton!.dispatchEvent(new MouseEvent('click', { bubbles: true }))
    await flushPromises()

    expect(deletePackMock).toHaveBeenCalledWith({ path: { pack_id: '1' } })
    expect(wrapper.emitted('deleted')).toBeTruthy()

    wrapper.unmount()
  })

  it('cancels without deleting', async () => {
    const wrapper = mount(DialogDeletePack, {
      props: { pack },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete pack"]').trigger('click')
    await flushPromises()

    const cancelButton = Array.from(document.querySelectorAll('button')).find(
      (button) => button.textContent === 'Cancel'
    )

    cancelButton!.dispatchEvent(new MouseEvent('click', { bubbles: true }))
    await flushPromises()

    expect(deletePackMock).not.toHaveBeenCalled()

    wrapper.unmount()
  })

  it('deletes immediately without the modal when shift-clicking', async () => {
    const wrapper = mount(DialogDeletePack, {
      props: { pack },
      attachTo: document.body,
    })

    await wrapper
      .find('button[aria-label="Delete pack"]')
      .trigger('click', { shiftKey: true })
    await flushPromises()

    expect(deletePackMock).toHaveBeenCalledWith({ path: { pack_id: '1' } })
    expect(wrapper.emitted('deleted')).toBeTruthy()
    expect(document.body.textContent).not.toContain('Are you sure')

    wrapper.unmount()
  })
})
