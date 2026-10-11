<script setup lang="ts">
import { ref } from 'vue'
import { provideGroupsContext } from '.'
import { type Group, type ListGroupsData, listGroups } from '../../../client'

const groups = ref<Group[]>([])
const total = ref(0)
const isLoading = ref(false)

let lastQuery: ListGroupsData['query'] = {}

async function loadGroups(query: ListGroupsData['query'] = lastQuery) {
  lastQuery = query
  isLoading.value = true

  const { data, error } = await listGroups({ query })

  if (error) {
    isLoading.value = false
    throw error
  }

  groups.value = data.groups
  total.value = data.total
  isLoading.value = false
}

provideGroupsContext({
  groups,
  total,
  isLoading,
  loadGroups,
})
</script>

<template>
  <slot />
</template>
