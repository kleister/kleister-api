import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const updateVersionMock = vi.fn()

vi.mock('../../../../client', () => ({
  updateVersion: (...args: unknown[]) => updateVersionMock(...args),
}))

import DialogEditVersion from './DialogEditVersion.vue'

const version = { id: 'version-1', name: '1.0.0', public: true }

async function submitForm() {
  document
    .querySelector('form')!
    .dispatchEvent(new Event('submit', { cancelable: true, bubbles: true }))

  await new Promise((resolve) => setTimeout(resolve, 300))
  await flushPromises()
}

describe('DialogEditVersion', () => {
  beforeEach(() => {
    updateVersionMock.mockReset()
    updateVersionMock.mockResolvedValue({
      data: { ...version },
      error: undefined,
    })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('prefills the form and submits the update scoped to the mod and version', async () => {
    const wrapper = mount(DialogEditVersion, {
      props: { modId: 'mod-1', version },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Edit version"]').trigger('click')
    await flushPromises()

    const nameInput =
      document.querySelector<HTMLInputElement>('input[name="name"]')
    expect(nameInput?.value).toBe('1.0.0')

    await submitForm()

    expect(updateVersionMock).toHaveBeenCalledWith({
      path: { mod_id: 'mod-1', version_id: 'version-1' },
      body: { name: '1.0.0', public: true, upload: undefined },
    })
    expect(wrapper.emitted('updated')).toBeTruthy()

    wrapper.unmount()
  })
})
