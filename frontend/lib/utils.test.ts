import { describe, expect, it } from 'vitest'
import { ref } from 'vue'
import { cn, valueUpdater } from './utils'

describe('cn', () => {
  it('returns a single class unchanged', () => {
    expect(cn('foo')).toBe('foo')
  })

  it('merges multiple classes', () => {
    expect(cn('px-2', 'py-2')).toBe('px-2 py-2')
  })

  it('deduplicates conflicting Tailwind classes, keeping the last', () => {
    expect(cn('px-2', 'px-4')).toBe('px-4')
  })

  it('ignores falsy values', () => {
    expect(cn('foo', false, undefined, null, 'bar')).toBe('foo bar')
  })
})

describe('valueUpdater', () => {
  it('sets ref to a direct value', () => {
    const r = ref(0)
    valueUpdater(42, r)
    expect(r.value).toBe(42)
  })

  it('applies an updater function to the current ref value', () => {
    const r = ref(10)
    valueUpdater((prev: number) => prev + 5, r)
    expect(r.value).toBe(15)
  })
})
