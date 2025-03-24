<script setup lang="ts">
import { ref } from 'vue'
import { provideUsersContext } from '.'
import { type User, listUsers } from '../../../client'

const users = ref<User[]>([])
const isLoading = ref(false)

async function loadUsers() {
  isLoading.value = true

  const { data, error } = await listUsers()

  if (error) {
    isLoading.value = false
    throw error
  }

  users.value = data.users
  isLoading.value = false
}

function addUser(user: User) {
  users.value = [user, ...users.value]
}

provideUsersContext({
  users,
  isLoading,
  loadUsers,
  addUser,
})
</script>

<template>
  <slot />
</template>
