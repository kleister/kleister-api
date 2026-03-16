import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import NotFound from './NotFound.vue'

describe('NotFound', () => {
  it('renders the 404 status code', () => {
    const wrapper = mount(NotFound)
    expect(wrapper.text()).toContain('404')
  })

  it('renders the "Page not found" message', () => {
    const wrapper = mount(NotFound)
    expect(wrapper.text()).toContain('Page not found')
  })

  it('renders an h1 element', () => {
    const wrapper = mount(NotFound)
    expect(wrapper.find('h1').exists()).toBe(true)
  })
})
