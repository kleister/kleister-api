import { flushPromises, mount } from '@vue/test-utils'
import { afterEach, describe, expect, it, vi } from 'vitest'
import RelationAttachPanel from './RelationAttachPanel.vue'

function makeProps() {
  return {
    subjectLabel: 'version',
    fetchRows: vi.fn().mockResolvedValue({
      rows: [{ id: '1', label: '1.0.0' }],
      total: 1,
    }),
    attachRow: vi.fn().mockResolvedValue({}),
    dropRow: vi.fn().mockResolvedValue({}),
  }
}

describe('RelationAttachPanel', () => {
  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('loads and renders rows on mount', async () => {
    const props = makeProps()
    const wrapper = mount(RelationAttachPanel, { props })
    await flushPromises()

    expect(props.fetchRows).toHaveBeenCalled()
    expect(wrapper.text()).toContain('1.0.0')
    expect(wrapper.text()).toContain('1 total')
  })

  it('attaches a new relation and reloads', async () => {
    const props = makeProps()
    const wrapper = mount(RelationAttachPanel, {
      props,
      attachTo: document.body,
    })
    await flushPromises()
    props.fetchRows.mockClear()

    await wrapper.find('input').setValue('2.0.0')
    await wrapper.find('form').trigger('submit')
    await flushPromises()

    expect(props.attachRow).toHaveBeenCalledWith('2.0.0')
    expect(props.fetchRows).toHaveBeenCalled()

    wrapper.unmount()
  })

  it('removes a row', async () => {
    const props = makeProps()
    const wrapper = mount(RelationAttachPanel, {
      props,
      attachTo: document.body,
    })
    await flushPromises()
    props.fetchRows.mockClear()

    await wrapper.find('button[aria-label="Remove 1.0.0"]').trigger('click')
    await flushPromises()

    expect(props.dropRow).toHaveBeenCalledWith('1')
    expect(props.fetchRows).toHaveBeenCalled()

    wrapper.unmount()
  })
})
