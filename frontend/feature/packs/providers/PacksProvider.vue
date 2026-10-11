<script setup lang="ts">
import { ref } from 'vue'
import { providePacksContext } from '.'
import { type Pack, type ListPacksData, listPacks } from '../../../client'

const packs = ref<Pack[]>([])
const total = ref(0)
const isLoading = ref(false)

let lastQuery: ListPacksData['query'] = {}

async function loadPacks(query: ListPacksData['query'] = lastQuery) {
  lastQuery = query
  isLoading.value = true

  const { data, error } = await listPacks({ query })

  if (error) {
    isLoading.value = false
    throw error
  }

  packs.value = data.packs
  total.value = data.total
  isLoading.value = false
}

providePacksContext({
  packs,
  total,
  isLoading,
  loadPacks,
})
</script>

<template>
  <slot />
</template>
