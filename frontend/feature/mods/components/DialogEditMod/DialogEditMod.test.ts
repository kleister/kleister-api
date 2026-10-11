import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const updateModMock = vi.fn()

vi.mock('../../../../client', () => ({
  updateMod: (...args: unknown[]) => updateModMock(...args),
}))

import DialogEditMod from './DialogEditMod.vue'

const mod = {
  id: '1',
  name: 'JEI',
  slug: 'jei',
  author: 'mezz',
  website: 'https://example.com',
  public: false,
  created_at: '2026-01-01T00:00:00Z',
  updated_at: '2026-01-01T00:00:00Z',
}

async function openDialog() {
  const wrapper = mount(DialogEditMod, {
    props: { mod },
    attachTo: document.body,
  })

  await wrapper.find('button[aria-label="Edit mod"]').trigger('click')
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

describe('DialogEditMod', () => {
  beforeEach(() => {
    updateModMock.mockReset()
    updateModMock.mockResolvedValue({ data: { ...mod }, error: undefined })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('does not render the form until opened', () => {
    const wrapper = mount(DialogEditMod, {
      props: { mod },
      attachTo: document.body,
    })

    expect(document.querySelector('form')).toBeNull()
    wrapper.unmount()
  })

  it('prefills the form with the given mod when opened', async () => {
    const wrapper = await openDialog()

    const nameInput =
      document.querySelector<HTMLInputElement>('input[name="name"]')
    const slugInput =
      document.querySelector<HTMLInputElement>('input[name="slug"]')

    expect(nameInput?.value).toBe('JEI')
    expect(slugInput?.value).toBe('jei')

    wrapper.unmount()
  })

  it('submits the update', async () => {
    const wrapper = await openDialog()

    await submitForm()

    expect(updateModMock).toHaveBeenCalledWith({
      path: { mod_id: '1' },
      body: {
        name: 'JEI',
        slug: 'jei',
        author: 'mezz',
        website: 'https://example.com',
        public: false,
      },
    })
    expect(wrapper.emitted('updated')).toBeTruthy()

    wrapper.unmount()
  })
})
