import type { ListUsersData, User } from '../../../client'
import { createContext } from 'radix-vue'
import type { Ref } from 'vue'

export const [useUsers, provideUsersContext] = createContext<{
  users: Ref<User[]>
  total: Ref<number>
  isLoading: Ref<boolean>
  loadUsers: (query?: ListUsersData['query']) => Promise<void>
}>('Users')
