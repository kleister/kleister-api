<script setup lang="ts">
import { ref } from 'vue'
import { provideModsContext } from '.'
import { type Mod, type ListModsData, listMods } from '../../../client'

const mods = ref<Mod[]>([])
const total = ref(0)
const isLoading = ref(false)

let lastQuery: ListModsData['query'] = {}

async function loadMods(query: ListModsData['query'] = lastQuery) {
  lastQuery = query
  isLoading.value = true

  const { data, error } = await listMods({ query })

  if (error) {
    isLoading.value = false
    throw error
  }

  mods.value = data.mods
  total.value = data.total
  isLoading.value = false
}

provideModsContext({
  mods,
  total,
  isLoading,
  loadMods,
})
</script>

<template>
  <slot />
</template>
