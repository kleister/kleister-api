import type { User } from '../../../client'
import { createContext } from 'radix-vue'
import type { Ref } from 'vue'

export const [useUsers, provideUsersContext] = createContext<{
  users: Ref<User[]>
  isLoading: Ref<boolean>
  loadUsers: () => Promise<void>
  addUser: (user: User) => void
}>('Users')
