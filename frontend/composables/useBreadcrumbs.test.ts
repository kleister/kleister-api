import { describe, expect, it, vi } from 'vitest'

vi.mock('vue-router', () => ({
  useRoute: () => ({
    name: 'users',
  }),
}))

import { useBreadcrumbs } from './useBreadcrumbs'

describe('useBreadcrumbs', () => {
  it('builds a breadcrumb entry from the current route name', () => {
    const { breadcrumbs } = useBreadcrumbs()

    expect(breadcrumbs.value).toEqual([
      {
        name: 'users',
      },
    ])
  })
})
