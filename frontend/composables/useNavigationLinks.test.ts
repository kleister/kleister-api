import { describe, expect, it } from 'vitest'
import { useNavigationLinks } from './useNavigationLinks'

describe('useNavigationLinks', () => {
  it('returns the expected general and admin navigation entries', () => {
    const { links } = useNavigationLinks()

    expect(links.value.general.map((entry) => entry.name)).toEqual([
      'Dashboard',
      'Packs',
      'Mods',
    ])
    expect(links.value.general.map((entry) => entry.url)).toEqual([
      '/',
      '/packs',
      '/mods',
    ])

    expect(links.value.admin.map((entry) => entry.name)).toEqual([
      'Users',
      'Groups',
    ])
    expect(links.value.admin.map((entry) => entry.url)).toEqual([
      '/users',
      '/groups',
    ])
  })
})
