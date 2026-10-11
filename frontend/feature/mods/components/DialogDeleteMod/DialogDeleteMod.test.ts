import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const deleteModMock = vi.fn()

vi.mock('../../../../client', () => ({
  deleteMod: (...args: unknown[]) => deleteModMock(...args),
}))

import DialogDeleteMod from './DialogDeleteMod.vue'

const mod = {
  id: '1',
  name: 'JEI',
  slug: 'jei',
  created_at: '2026-01-01T00:00:00Z',
  updated_at: '2026-01-01T00:00:00Z',
}

describe('DialogDeleteMod', () => {
  beforeEach(() => {
    deleteModMock.mockReset()
    deleteModMock.mockResolvedValue({
      data: { message: 'ok' },
      error: undefined,
    })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('opens a confirmation modal on a plain click and does not delete yet', async () => {
    const wrapper = mount(DialogDeleteMod, {
      props: { mod },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete mod"]').trigger('click')
    await flushPromises()

    expect(deleteModMock).not.toHaveBeenCalled()
    expect(document.body.textContent).toContain('Delete mod')
    expect(document.body.textContent).toContain('JEI')

    wrapper.unmount()
  })

  it('deletes the mod when confirming in the modal', async () => {
    const wrapper = mount(DialogDeleteMod, {
      props: { mod },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete mod"]').trigger('click')
    await flushPromises()

    const deleteButton = Array.from(document.querySelectorAll('button')).find(
      (button) => button.textContent === 'Delete'
    )
    expect(deleteButton).toBeTruthy()

    deleteButton!.dispatchEvent(new MouseEvent('click', { bubbles: true }))
    await flushPromises()

    expect(deleteModMock).toHaveBeenCalledWith({ path: { mod_id: '1' } })
    expect(wrapper.emitted('deleted')).toBeTruthy()

    wrapper.unmount()
  })

  it('cancels without deleting', async () => {
    const wrapper = mount(DialogDeleteMod, {
      props: { mod },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete mod"]').trigger('click')
    await flushPromises()

    const cancelButton = Array.from(document.querySelectorAll('button')).find(
      (button) => button.textContent === 'Cancel'
    )

    cancelButton!.dispatchEvent(new MouseEvent('click', { bubbles: true }))
    await flushPromises()

    expect(deleteModMock).not.toHaveBeenCalled()

    wrapper.unmount()
  })

  it('deletes immediately without the modal when shift-clicking', async () => {
    const wrapper = mount(DialogDeleteMod, {
      props: { mod },
      attachTo: document.body,
    })

    await wrapper
      .find('button[aria-label="Delete mod"]')
      .trigger('click', { shiftKey: true })
    await flushPromises()

    expect(deleteModMock).toHaveBeenCalledWith({ path: { mod_id: '1' } })
    expect(wrapper.emitted('deleted')).toBeTruthy()
    expect(document.body.textContent).not.toContain('Are you sure')

    wrapper.unmount()
  })
})
