import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const deleteGroupMock = vi.fn()

vi.mock('../../../../client', () => ({
  deleteGroup: (...args: unknown[]) => deleteGroupMock(...args),
}))

import DialogDeleteGroup from './DialogDeleteGroup.vue'

const group = {
  id: '1',
  name: 'Admins',
  slug: 'admins',
  created_at: '2026-01-01T00:00:00Z',
  updated_at: '2026-01-01T00:00:00Z',
}

describe('DialogDeleteGroup', () => {
  beforeEach(() => {
    deleteGroupMock.mockReset()
    deleteGroupMock.mockResolvedValue({
      data: { message: 'ok' },
      error: undefined,
    })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('opens a confirmation modal on a plain click and does not delete yet', async () => {
    const wrapper = mount(DialogDeleteGroup, {
      props: { group },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete group"]').trigger('click')
    await flushPromises()

    expect(deleteGroupMock).not.toHaveBeenCalled()
    expect(document.body.textContent).toContain('Delete group')
    expect(document.body.textContent).toContain('Admins')

    wrapper.unmount()
  })

  it('deletes the group when confirming in the modal', async () => {
    const wrapper = mount(DialogDeleteGroup, {
      props: { group },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete group"]').trigger('click')
    await flushPromises()

    const deleteButton = Array.from(document.querySelectorAll('button')).find(
      (button) => button.textContent === 'Delete'
    )
    expect(deleteButton).toBeTruthy()

    deleteButton!.dispatchEvent(new MouseEvent('click', { bubbles: true }))
    await flushPromises()

    expect(deleteGroupMock).toHaveBeenCalledWith({ path: { group_id: '1' } })
    expect(wrapper.emitted('deleted')).toBeTruthy()

    wrapper.unmount()
  })

  it('cancels without deleting', async () => {
    const wrapper = mount(DialogDeleteGroup, {
      props: { group },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete group"]').trigger('click')
    await flushPromises()

    const cancelButton = Array.from(document.querySelectorAll('button')).find(
      (button) => button.textContent === 'Cancel'
    )

    cancelButton!.dispatchEvent(new MouseEvent('click', { bubbles: true }))
    await flushPromises()

    expect(deleteGroupMock).not.toHaveBeenCalled()

    wrapper.unmount()
  })

  it('deletes immediately without the modal when shift-clicking', async () => {
    const wrapper = mount(DialogDeleteGroup, {
      props: { group },
      attachTo: document.body,
    })

    await wrapper
      .find('button[aria-label="Delete group"]')
      .trigger('click', { shiftKey: true })
    await flushPromises()

    expect(deleteGroupMock).toHaveBeenCalledWith({ path: { group_id: '1' } })
    expect(wrapper.emitted('deleted')).toBeTruthy()
    expect(document.body.textContent).not.toContain('Are you sure')

    wrapper.unmount()
  })
})
