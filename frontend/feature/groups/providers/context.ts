import type { Group, ListGroupsData } from '../../../client'
import { createContext } from 'radix-vue'
import type { Ref } from 'vue'

export const [useGroups, provideGroupsContext] = createContext<{
  groups: Ref<Group[]>
  total: Ref<number>
  isLoading: Ref<boolean>
  loadGroups: (query?: ListGroupsData['query']) => Promise<void>
}>('Groups')
