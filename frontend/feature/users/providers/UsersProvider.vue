<script setup lang="ts">
import { ref } from 'vue'
import { provideUsersContext } from '.'
import { type ListUsersData, type User, listUsers } from '../../../client'

const users = ref<User[]>([])
const total = ref(0)
const isLoading = ref(false)

let lastQuery: ListUsersData['query'] = {}

async function loadUsers(query: ListUsersData['query'] = lastQuery) {
  lastQuery = query
  isLoading.value = true

  const { data, error } = await listUsers({ query })

  if (error) {
    isLoading.value = false
    throw error
  }

  users.value = data.users
  total.value = data.total
  isLoading.value = false
}

provideUsersContext({
  users,
  total,
  isLoading,
  loadUsers,
})
</script>

<template>
  <slot />
</template>
