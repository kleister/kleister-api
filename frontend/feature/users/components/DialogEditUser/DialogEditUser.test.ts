import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const updateUserMock = vi.fn()

vi.mock('../../../../client', () => ({
  updateUser: (...args: unknown[]) => updateUserMock(...args),
}))

import DialogEditUser from './DialogEditUser.vue'

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

async function openDialog() {
  const wrapper = mount(DialogEditUser, {
    props: { user },
    attachTo: document.body,
  })

  await wrapper.find('button[aria-label="Edit user"]').trigger('click')
  await flushPromises()

  return wrapper
}

// vee-validate's async schema validation resolves over several macrotasks,
// so flushPromises() alone isn't enough to observe the submit outcome.
async function submitForm() {
  document
    .querySelector('form')!
    .dispatchEvent(new Event('submit', { cancelable: true, bubbles: true }))

  await new Promise((resolve) => setTimeout(resolve, 300))
  await flushPromises()
}

describe('DialogEditUser', () => {
  beforeEach(() => {
    updateUserMock.mockReset()
    updateUserMock.mockResolvedValue({ data: { ...user }, error: undefined })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('does not render the form until opened', () => {
    const wrapper = mount(DialogEditUser, {
      props: { user },
      attachTo: document.body,
    })

    expect(document.querySelector('form')).toBeNull()
    wrapper.unmount()
  })

  it('prefills the form with the given user when opened', async () => {
    const wrapper = await openDialog()

    const usernameInput = document.querySelector<HTMLInputElement>(
      'input[name="username"]'
    )
    const emailInput = document.querySelector<HTMLInputElement>(
      'input[name="email"]'
    )

    expect(usernameInput?.value).toBe('john')
    expect(emailInput?.value).toBe('john@example.com')

    wrapper.unmount()
  })

  it('submits the update without a password when left blank', async () => {
    const wrapper = await openDialog()

    await submitForm()

    expect(updateUserMock).toHaveBeenCalledWith({
      path: { user_id: '1' },
      body: {
        username: 'john',
        email: 'john@example.com',
        fullname: 'John Doe',
        admin: false,
        active: true,
      },
    })
    expect(wrapper.emitted('updated')).toBeTruthy()

    wrapper.unmount()
  })

  it('includes the password when one is provided', async () => {
    const wrapper = await openDialog()

    const passwordInput = document.querySelector<HTMLInputElement>(
      'input[name="password"]'
    )!
    passwordInput.value = 'supersecret'
    passwordInput.dispatchEvent(new Event('input'))
    await flushPromises()

    await submitForm()

    expect(updateUserMock).toHaveBeenCalledWith({
      path: { user_id: '1' },
      body: {
        username: 'john',
        email: 'john@example.com',
        fullname: 'John Doe',
        admin: false,
        active: true,
        password: 'supersecret',
      },
    })

    wrapper.unmount()
  })
})
