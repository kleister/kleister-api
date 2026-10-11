import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const updatePackMock = vi.fn()

vi.mock('../../../../client', () => ({
  updatePack: (...args: unknown[]) => updatePackMock(...args),
}))

import DialogEditPack from './DialogEditPack.vue'

const pack = {
  id: '1',
  name: 'Vanilla Plus',
  slug: 'vanilla-plus',
  website: 'https://example.com',
  public: false,
  created_at: '2026-01-01T00:00:00Z',
  updated_at: '2026-01-01T00:00:00Z',
}

async function openDialog() {
  const wrapper = mount(DialogEditPack, {
    props: { pack },
    attachTo: document.body,
  })

  await wrapper.find('button[aria-label="Edit pack"]').trigger('click')
  await flushPromises()

  return wrapper
}

async function submitForm() {
  document
    .querySelector('form')!
    .dispatchEvent(new Event('submit', { cancelable: true, bubbles: true }))

  await new Promise((resolve) => setTimeout(resolve, 300))
  await flushPromises()
}

describe('DialogEditPack', () => {
  beforeEach(() => {
    updatePackMock.mockReset()
    updatePackMock.mockResolvedValue({ data: { ...pack }, error: undefined })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('does not render the form until opened', () => {
    const wrapper = mount(DialogEditPack, {
      props: { pack },
      attachTo: document.body,
    })

    expect(document.querySelector('form')).toBeNull()
    wrapper.unmount()
  })

  it('prefills the form with the given pack when opened', async () => {
    const wrapper = await openDialog()

    const nameInput =
      document.querySelector<HTMLInputElement>('input[name="name"]')
    const slugInput =
      document.querySelector<HTMLInputElement>('input[name="slug"]')

    expect(nameInput?.value).toBe('Vanilla Plus')
    expect(slugInput?.value).toBe('vanilla-plus')

    wrapper.unmount()
  })

  it('submits the update', async () => {
    const wrapper = await openDialog()

    await submitForm()

    expect(updatePackMock).toHaveBeenCalledWith({
      path: { pack_id: '1' },
      body: {
        name: 'Vanilla Plus',
        slug: 'vanilla-plus',
        website: 'https://example.com',
        public: false,
      },
    })
    expect(wrapper.emitted('updated')).toBeTruthy()

    wrapper.unmount()
  })
})
