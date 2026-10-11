import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const createVersionMock = vi.fn()

vi.mock('../../../../client', () => ({
  createVersion: (...args: unknown[]) => createVersionMock(...args),
}))

import DialogNewVersion from './DialogNewVersion.vue'

async function submitForm() {
  document
    .querySelector('form')!
    .dispatchEvent(new Event('submit', { cancelable: true, bubbles: true }))

  await new Promise((resolve) => setTimeout(resolve, 300))
  await flushPromises()
}

describe('DialogNewVersion', () => {
  beforeEach(() => {
    createVersionMock.mockReset()
    createVersionMock.mockResolvedValue({
      data: { id: '1' },
      error: undefined,
    })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('creates a version scoped to the given mod', async () => {
    const wrapper = mount(DialogNewVersion, {
      props: { modId: 'mod-1' },
      attachTo: document.body,
    })

    await wrapper.find('button').trigger('click')
    await flushPromises()

    const nameInput =
      document.querySelector<HTMLInputElement>('input[name="name"]')!
    nameInput.value = '1.0.0'
    nameInput.dispatchEvent(new Event('input'))

    await submitForm()

    expect(createVersionMock).toHaveBeenCalledWith(
      expect.objectContaining({ path: { mod_id: 'mod-1' } })
    )
    expect(wrapper.emitted('created')).toBeTruthy()

    wrapper.unmount()
  })
})
