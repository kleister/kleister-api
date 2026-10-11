import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const deleteUserMock = vi.fn()

vi.mock('../../../../client', () => ({
  deleteUser: (...args: unknown[]) => deleteUserMock(...args),
}))

import DialogDeleteUser from './DialogDeleteUser.vue'

const user = {
  id: '1',
  username: 'john',
  email: 'john@example.com',
  fullname: 'John Doe',
  admin: false,
  active: true,
  created_at: '2026-01-01T00:00:00Z',
  updated_at: '2026-01-01T00:00:00Z',
}

describe('DialogDeleteUser', () => {
  beforeEach(() => {
    deleteUserMock.mockReset()
    deleteUserMock.mockResolvedValue({
      data: { message: 'ok' },
      error: undefined,
    })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('opens a confirmation modal on a plain click and does not delete yet', async () => {
    const wrapper = mount(DialogDeleteUser, {
      props: { user },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete user"]').trigger('click')
    await flushPromises()

    expect(deleteUserMock).not.toHaveBeenCalled()
    expect(document.body.textContent).toContain('Delete user')
    expect(document.body.textContent).toContain('john')

    wrapper.unmount()
  })

  it('deletes the user when confirming in the modal', async () => {
    const wrapper = mount(DialogDeleteUser, {
      props: { user },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete user"]').trigger('click')
    await flushPromises()

    const deleteButton = Array.from(document.querySelectorAll('button')).find(
      (button) => button.textContent === 'Delete'
    )
    expect(deleteButton).toBeTruthy()

    deleteButton!.dispatchEvent(new MouseEvent('click', { bubbles: true }))
    await flushPromises()

    expect(deleteUserMock).toHaveBeenCalledWith({ path: { user_id: '1' } })
    expect(wrapper.emitted('deleted')).toBeTruthy()

    wrapper.unmount()
  })

  it('cancels without deleting', async () => {
    const wrapper = mount(DialogDeleteUser, {
      props: { user },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Delete user"]').trigger('click')
    await flushPromises()

    const cancelButton = Array.from(document.querySelectorAll('button')).find(
      (button) => button.textContent === 'Cancel'
    )

    cancelButton!.dispatchEvent(new MouseEvent('click', { bubbles: true }))
    await flushPromises()

    expect(deleteUserMock).not.toHaveBeenCalled()

    wrapper.unmount()
  })

  it('deletes immediately without the modal when shift-clicking', async () => {
    const wrapper = mount(DialogDeleteUser, {
      props: { user },
      attachTo: document.body,
    })

    await wrapper
      .find('button[aria-label="Delete user"]')
      .trigger('click', { shiftKey: true })
    await flushPromises()

    expect(deleteUserMock).toHaveBeenCalledWith({ path: { user_id: '1' } })
    expect(wrapper.emitted('deleted')).toBeTruthy()
    expect(document.body.textContent).not.toContain('Are you sure')

    wrapper.unmount()
  })
})
