import type { ListPacksData, Pack } from '../../../client'
import { createContext } from 'radix-vue'
import type { Ref } from 'vue'

export const [usePacks, providePacksContext] = createContext<{
  packs: Ref<Pack[]>
  total: Ref<number>
  isLoading: Ref<boolean>
  loadPacks: (query?: ListPacksData['query']) => Promise<void>
}>('Packs')
