<script setup lang="ts">
import { ref } from 'vue'
import { provideGroupsContext } from '.'
import { type Group, listGroups } from '../../../client'

const groups = ref<Group[]>([])
const isLoading = ref(false)

async function loadGroups() {
  isLoading.value = true

  const { data, error } = await listGroups()

  if (error) {
    isLoading.value = false
    throw error
  }

  groups.value = data.groups
  isLoading.value = false
}

function addGroup(group: Group) {
  groups.value = [group, ...groups.value]
}

provideGroupsContext({
  groups,
  isLoading,
  loadGroups,
  addGroup,
})
</script>

<template>
  <slot />
</template>
