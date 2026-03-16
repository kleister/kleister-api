import { describe, expect, it } from 'vitest'
import { useNavigationLinks } from './useNavigationLinks'

describe('useNavigationLinks', () => {
  it('returns the expected general and admin navigation entries', () => {
    const { links } = useNavigationLinks()

    expect(links.value.general.map((entry) => entry.name)).toEqual(['Dashboard'])
    expect(links.value.general[0]?.url).toBe('/')

    expect(links.value.admin.map((entry) => entry.name)).toEqual([
      'Users',
      'Groups',
    ])
    expect(links.value.admin.map((entry) => entry.url)).toEqual([
      '/admin/users',
      '/admin/groups',
    ])
  })
})
