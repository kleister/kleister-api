import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

const updateBuildMock = vi.fn()

vi.mock('../../../../client', () => ({
  updateBuild: (...args: unknown[]) => updateBuildMock(...args),
}))

import DialogEditBuild from './DialogEditBuild.vue'

const build = {
  id: 'build-1',
  name: '1.0.0',
  java: '17',
  memory: '4G',
  latest: true,
  recommended: false,
  public: true,
}

async function submitForm() {
  document
    .querySelector('form')!
    .dispatchEvent(new Event('submit', { cancelable: true, bubbles: true }))

  await new Promise((resolve) => setTimeout(resolve, 300))
  await flushPromises()
}

describe('DialogEditBuild', () => {
  beforeEach(() => {
    updateBuildMock.mockReset()
    updateBuildMock.mockResolvedValue({ data: { ...build }, error: undefined })
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('prefills the form and submits the update scoped to the pack and build', async () => {
    const wrapper = mount(DialogEditBuild, {
      props: { packId: 'pack-1', build },
      attachTo: document.body,
    })

    await wrapper.find('button[aria-label="Edit build"]').trigger('click')
    await flushPromises()

    const nameInput =
      document.querySelector<HTMLInputElement>('input[name="name"]')
    expect(nameInput?.value).toBe('1.0.0')

    await submitForm()

    expect(updateBuildMock).toHaveBeenCalledWith({
      path: { pack_id: 'pack-1', build_id: 'build-1' },
      body: {
        name: '1.0.0',
        java: '17',
        memory: '4G',
        latest: true,
        recommended: false,
        public: true,
      },
    })
    expect(wrapper.emitted('updated')).toBeTruthy()

    wrapper.unmount()
  })
})
