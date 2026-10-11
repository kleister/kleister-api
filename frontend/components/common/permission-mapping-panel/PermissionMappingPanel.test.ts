import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, describe, expect, it, vi } from 'vitest'
import PermissionMappingPanel from './PermissionMappingPanel.vue'

function makeProps() {
  return {
    subjectLabel: 'user',
    fetchRows: vi.fn().mockResolvedValue({
      rows: [{ subjectId: '1', label: 'alice', perm: 'user' }],
      total: 1,
    }),
    attachRow: vi.fn().mockResolvedValue({}),
    permitRow: vi.fn().mockResolvedValue({}),
    dropRow: vi.fn().mockResolvedValue({}),
  }
}

describe('PermissionMappingPanel', () => {
  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('loads and renders rows on mount', async () => {
    const props = makeProps()
    const wrapper = mount(PermissionMappingPanel, { props })
    await flushPromises()

    expect(props.fetchRows).toHaveBeenCalled()
    expect(wrapper.text()).toContain('alice')
    expect(wrapper.text()).toContain('1 total')
  })

  it('attaches a new subject and reloads', async () => {
    const props = makeProps()
    const wrapper = mount(PermissionMappingPanel, {
      props,
      attachTo: document.body,
    })
    await flushPromises()
    props.fetchRows.mockClear()

    await wrapper.find('input').setValue('bob')
    await wrapper.find('form').trigger('submit')
    await flushPromises()

    expect(props.attachRow).toHaveBeenCalledWith('bob', 'user')
    expect(props.fetchRows).toHaveBeenCalled()

    wrapper.unmount()
  })

  it('removes a row', async () => {
    const props = makeProps()
    const wrapper = mount(PermissionMappingPanel, {
      props,
      attachTo: document.body,
    })
    await flushPromises()
    props.fetchRows.mockClear()

    await wrapper
      .find('button[aria-label="Remove alice"]')
      .trigger('click')
    await flushPromises()

    expect(props.dropRow).toHaveBeenCalledWith('1')
    expect(props.fetchRows).toHaveBeenCalled()

    wrapper.unmount()
  })

  it('shows a loading state before rows arrive', async () => {
    const props = makeProps()
    props.fetchRows.mockReturnValue(new Promise(() => {}))

    const wrapper = mount(PermissionMappingPanel, { props })
    await flushPromises()

    expect(wrapper.text()).toContain('Loading...')
  })
})
