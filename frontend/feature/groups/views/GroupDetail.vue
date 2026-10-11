<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { type Group, showGroup } from '../../../client'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'

const route = useRoute()
const router = useRouter()

const group = ref<Group | null>(null)
const isLoading = ref(false)

function formatDate(value?: string) {
  if (!value) {
    return ''
  }

  return new Intl.DateTimeFormat('en-US').format(new Date(value))
}

async function loadGroup() {
  isLoading.value = true

  const { data, error } = await showGroup({
    path: { group_id: String(route.params.group_id) },
  })

  isLoading.value = false

  if (error) {
    throw error
  }

  group.value = data
}

onMounted(loadGroup)
</script>

<template>
  <div class="flex flex-col gap-4">
    <Button
      variant="outline"
      class="w-fit"
      @click="router.push({ name: 'groups' })"
    >
      Back to groups
    </Button>

    <Card v-if="group">
      <CardHeader>
        <CardTitle>{{ group.name }}</CardTitle>
      </CardHeader>
      <CardContent class="grid gap-2">
        <div><span class="font-medium">Slug:</span> {{ group.slug }}</div>
        <div>
          <span class="font-medium">Created:</span>
          {{ formatDate(group.created_at) }}
        </div>
        <div>
          <span class="font-medium">Updated:</span>
          {{ formatDate(group.updated_at) }}
        </div>
      </CardContent>
    </Card>

    <p v-else-if="isLoading">Loading...</p>
  </div>
</template>
