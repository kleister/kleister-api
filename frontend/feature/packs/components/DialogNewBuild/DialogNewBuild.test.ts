import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const createBuildMock = vi.fn()

vi.mock('../../../../client', () => ({
  createBuild: (...args: unknown[]) => createBuildMock(...args),
}))

import DialogNewBuild from './DialogNewBuild.vue'

async function submitForm() {
  document
    .querySelector('form')!
    .dispatchEvent(new Event('submit', { cancelable: true, bubbles: true }))

  await new Promise((resolve) => setTimeout(resolve, 300))
  await flushPromises()
}

describe('DialogNewBuild', () => {
  beforeEach(() => {
    createBuildMock.mockReset()
    createBuildMock.mockResolvedValue({ data: { id: '1' }, error: undefined })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('creates a build scoped to the given pack', async () => {
    const wrapper = mount(DialogNewBuild, {
      props: { packId: 'pack-1' },
      attachTo: document.body,
    })

    await wrapper.find('button').trigger('click')
    await flushPromises()

    await document
      .querySelector<HTMLInputElement>('input[name="name"]')!
      .dispatchEvent(new Event('input'))
    document.querySelector<HTMLInputElement>(
      'input[name="name"]'
    )!.value = '1.0.0'
    document
      .querySelector<HTMLInputElement>('input[name="name"]')!
      .dispatchEvent(new Event('input'))

    await submitForm()

    expect(createBuildMock).toHaveBeenCalledWith(
      expect.objectContaining({ path: { pack_id: 'pack-1' } })
    )
    expect(wrapper.emitted('created')).toBeTruthy()

    wrapper.unmount()
  })
})
