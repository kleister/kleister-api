import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const updateGroupMock = vi.fn()

vi.mock('../../../../client', () => ({
  updateGroup: (...args: unknown[]) => updateGroupMock(...args),
}))

import DialogEditGroup from './DialogEditGroup.vue'

const group = {
  id: '1',
  name: 'Admins',
  slug: 'admins',
  created_at: '2026-01-01T00:00:00Z',
  updated_at: '2026-01-01T00:00:00Z',
}

async function openDialog() {
  const wrapper = mount(DialogEditGroup, {
    props: { group },
    attachTo: document.body,
  })

  await wrapper.find('button[aria-label="Edit group"]').trigger('click')
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

describe('DialogEditGroup', () => {
  beforeEach(() => {
    updateGroupMock.mockReset()
    updateGroupMock.mockResolvedValue({ data: { ...group }, error: undefined })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('does not render the form until opened', () => {
    const wrapper = mount(DialogEditGroup, {
      props: { group },
      attachTo: document.body,
    })

    expect(document.querySelector('form')).toBeNull()
    wrapper.unmount()
  })

  it('prefills the form with the given group when opened', async () => {
    const wrapper = await openDialog()

    const nameInput =
      document.querySelector<HTMLInputElement>('input[name="name"]')
    const slugInput =
      document.querySelector<HTMLInputElement>('input[name="slug"]')

    expect(nameInput?.value).toBe('Admins')
    expect(slugInput?.value).toBe('admins')

    wrapper.unmount()
  })

  it('submits the update', async () => {
    const wrapper = await openDialog()

    await submitForm()

    expect(updateGroupMock).toHaveBeenCalledWith({
      path: { group_id: '1' },
      body: { name: 'Admins', slug: 'admins' },
    })
    expect(wrapper.emitted('updated')).toBeTruthy()

    wrapper.unmount()
  })
})
