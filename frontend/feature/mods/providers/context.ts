import type { ListModsData, Mod } from '../../../client'
import { createContext } from 'radix-vue'
import type { Ref } from 'vue'

export const [useMods, provideModsContext] = createContext<{
  mods: Ref<Mod[]>
  total: Ref<number>
  isLoading: Ref<boolean>
  loadMods: (query?: ListModsData['query']) => Promise<void>
}>('Mods')
