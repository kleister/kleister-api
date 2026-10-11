<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { type User, showUser } from '../../../client'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'

const route = useRoute()
const router = useRouter()

const user = ref<User | null>(null)
const isLoading = ref(false)

function formatDate(value?: string) {
  if (!value) {
    return ''
  }

  return new Intl.DateTimeFormat('en-US').format(new Date(value))
}

async function loadUser() {
  isLoading.value = true

  const { data, error } = await showUser({
    path: { user_id: String(route.params.user_id) },
  })

  isLoading.value = false

  if (error) {
    throw error
  }

  user.value = data
}

onMounted(loadUser)
</script>

<template>
  <div class="flex flex-col gap-4">
    <Button
      variant="outline"
      class="w-fit"
      @click="router.push({ name: 'users' })"
    >
      Back to users
    </Button>

    <Card v-if="user">
      <CardHeader>
        <CardTitle>{{ user.username }}</CardTitle>
      </CardHeader>
      <CardContent class="grid gap-2">
        <div><span class="font-medium">Email:</span> {{ user.email }}</div>
        <div>
          <span class="font-medium">Fullname:</span> {{ user.fullname }}
        </div>
        <div>
          <span class="font-medium">Admin:</span>
          {{ user.admin ? 'Yes' : 'No' }}
        </div>
        <div>
          <span class="font-medium">Active:</span>
          {{ user.active ? 'Yes' : 'No' }}
        </div>
        <div>
          <span class="font-medium">Created:</span>
          {{ formatDate(user.created_at) }}
        </div>
        <div>
          <span class="font-medium">Updated:</span>
          {{ formatDate(user.updated_at) }}
        </div>
      </CardContent>
    </Card>

    <p v-else-if="isLoading">Loading...</p>
  </div>
</template>
