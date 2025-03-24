import type { Group } from '../../../client'
import { createContext } from 'radix-vue'
import type { Ref } from 'vue'

export const [useGroups, provideGroupsContext] = createContext<{
  groups: Ref<Group[]>
  isLoading: Ref<boolean>
  loadGroups: () => Promise<void>
  addGroup: (group: Group) => void
}>('Groups')
