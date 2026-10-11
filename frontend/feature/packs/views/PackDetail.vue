<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  type Pack,
  showPack,
  listPackUsers,
  attachPackToUser,
  permitPackUser,
  deletePackFromUser,
  listPackGroups,
  attachPackToGroup,
  permitPackGroup,
  deletePackFromGroup,
} from '../../../client'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { PermissionMappingPanel } from '@/components/common/permission-mapping-panel'
import BuildsList from '../components/BuildsList'

const route = useRoute()
const router = useRouter()

const pack = ref<Pack | null>(null)
const isLoading = ref(false)

function formatDate(value?: string) {
  if (!value) {
    return ''
  }

  return new Intl.DateTimeFormat('en-US').format(new Date(value))
}

async function loadPack() {
  isLoading.value = true

  const { data, error } = await showPack({
    path: { pack_id: String(route.params.pack_id) },
  })

  isLoading.value = false

  if (error) {
    throw error
  }

  pack.value = data
}

onMounted(loadPack)

async function fetchUsers() {
  const { data, error } = await listPackUsers({
    path: { pack_id: String(route.params.pack_id) },
  })

  if (error) {
    throw error
  }

  return {
    total: data.total,
    rows: data.users.map((entry) => ({
      subjectId: entry.user_id,
      label: entry.user?.username ?? entry.user_id,
      perm: entry.perm ?? 'user',
    })),
  }
}

function attachUser(subjectId: string, perm: string) {
  return attachPackToUser({
    path: { pack_id: String(route.params.pack_id) },
    body: { user: subjectId, perm },
  })
}

function permitUser(subjectId: string, perm: string) {
  return permitPackUser({
    path: { pack_id: String(route.params.pack_id) },
    body: { user: subjectId, perm },
  })
}

function dropUser(subjectId: string) {
  return deletePackFromUser({
    path: { pack_id: String(route.params.pack_id) },
    body: { user: subjectId },
  })
}

async function fetchGroups() {
  const { data, error } = await listPackGroups({
    path: { pack_id: String(route.params.pack_id) },
  })

  if (error) {
    throw error
  }

  return {
    total: data.total,
    rows: data.groups.map((entry) => ({
      subjectId: entry.group_id,
      label: entry.group?.name ?? entry.group_id,
      perm: entry.perm ?? 'user',
    })),
  }
}

function attachGroup(subjectId: string, perm: string) {
  return attachPackToGroup({
    path: { pack_id: String(route.params.pack_id) },
    body: { group: subjectId, perm },
  })
}

function permitGroup(subjectId: string, perm: string) {
  return permitPackGroup({
    path: { pack_id: String(route.params.pack_id) },
    body: { group: subjectId, perm },
  })
}

function dropGroup(subjectId: string) {
  return deletePackFromGroup({
    path: { pack_id: String(route.params.pack_id) },
    body: { group: subjectId },
  })
}
</script>

<template>
  <div class="flex flex-col gap-4">
    <Button
      variant="outline"
      class="w-fit"
      @click="router.push({ name: 'packs' })"
    >
      Back to packs
    </Button>

    <Card v-if="pack">
      <CardHeader>
        <CardTitle>{{ pack.name }}</CardTitle>
      </CardHeader>
      <CardContent>
        <Tabs default-value="overview">
          <TabsList>
            <TabsTrigger value="overview">Overview</TabsTrigger>
            <TabsTrigger value="builds">Builds</TabsTrigger>
            <TabsTrigger value="users">Users</TabsTrigger>
            <TabsTrigger value="groups">Groups</TabsTrigger>
          </TabsList>

          <TabsContent value="overview" class="grid gap-2">
            <div><span class="font-medium">Slug:</span> {{ pack.slug }}</div>
            <div>
              <span class="font-medium">Website:</span> {{ pack.website }}
            </div>
            <div>
              <span class="font-medium">Public:</span>
              {{ pack.public ? 'Yes' : 'No' }}
            </div>
            <div>
              <span class="font-medium">Created:</span>
              {{ formatDate(pack.created_at) }}
            </div>
            <div>
              <span class="font-medium">Updated:</span>
              {{ formatDate(pack.updated_at) }}
            </div>
          </TabsContent>

          <TabsContent value="builds">
            <BuildsList :pack-id="pack.id!" />
          </TabsContent>

          <TabsContent value="users">
            <PermissionMappingPanel
              subject-label="user"
              :fetch-rows="fetchUsers"
              :attach-row="attachUser"
              :permit-row="permitUser"
              :drop-row="dropUser"
            />
          </TabsContent>

          <TabsContent value="groups">
            <PermissionMappingPanel
              subject-label="group"
              :fetch-rows="fetchGroups"
              :attach-row="attachGroup"
              :permit-row="permitGroup"
              :drop-row="dropGroup"
            />
          </TabsContent>
        </Tabs>
      </CardContent>
    </Card>

    <p v-else-if="isLoading">Loading...</p>
  </div>
</template>
