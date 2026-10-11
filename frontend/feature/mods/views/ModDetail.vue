<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  type Mod,
  showMod,
  listModUsers,
  attachModToUser,
  permitModUser,
  deleteModFromUser,
  listModGroups,
  attachModToGroup,
  permitModGroup,
  deleteModFromGroup,
} from '../../../client'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { PermissionMappingPanel } from '@/components/common/permission-mapping-panel'
import VersionsList from '../components/VersionsList'

const route = useRoute()
const router = useRouter()

const mod = ref<Mod | null>(null)
const isLoading = ref(false)

function formatDate(value?: string) {
  if (!value) {
    return ''
  }

  return new Intl.DateTimeFormat('en-US').format(new Date(value))
}

async function loadMod() {
  isLoading.value = true

  const { data, error } = await showMod({
    path: { mod_id: String(route.params.mod_id) },
  })

  isLoading.value = false

  if (error) {
    throw error
  }

  mod.value = data
}

onMounted(loadMod)

async function fetchUsers() {
  const { data, error } = await listModUsers({
    path: { mod_id: String(route.params.mod_id) },
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
  return attachModToUser({
    path: { mod_id: String(route.params.mod_id) },
    body: { user: subjectId, perm },
  })
}

function permitUser(subjectId: string, perm: string) {
  return permitModUser({
    path: { mod_id: String(route.params.mod_id) },
    body: { user: subjectId, perm },
  })
}

function dropUser(subjectId: string) {
  return deleteModFromUser({
    path: { mod_id: String(route.params.mod_id) },
    body: { user: subjectId },
  })
}

async function fetchGroups() {
  const { data, error } = await listModGroups({
    path: { mod_id: String(route.params.mod_id) },
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
  return attachModToGroup({
    path: { mod_id: String(route.params.mod_id) },
    body: { group: subjectId, perm },
  })
}

function permitGroup(subjectId: string, perm: string) {
  return permitModGroup({
    path: { mod_id: String(route.params.mod_id) },
    body: { group: subjectId, perm },
  })
}

function dropGroup(subjectId: string) {
  return deleteModFromGroup({
    path: { mod_id: String(route.params.mod_id) },
    body: { group: subjectId },
  })
}
</script>

<template>
  <div class="flex flex-col gap-4">
    <Button
      variant="outline"
      class="w-fit"
      @click="router.push({ name: 'mods' })"
    >
      Back to mods
    </Button>

    <Card v-if="mod">
      <CardHeader>
        <CardTitle>{{ mod.name }}</CardTitle>
      </CardHeader>
      <CardContent>
        <Tabs default-value="overview">
          <TabsList>
            <TabsTrigger value="overview">Overview</TabsTrigger>
            <TabsTrigger value="versions">Versions</TabsTrigger>
            <TabsTrigger value="users">Users</TabsTrigger>
            <TabsTrigger value="groups">Groups</TabsTrigger>
          </TabsList>

          <TabsContent value="overview" class="grid gap-2">
            <div><span class="font-medium">Slug:</span> {{ mod.slug }}</div>
            <div>
              <span class="font-medium">Author:</span> {{ mod.author }}
            </div>
            <div>
              <span class="font-medium">Website:</span> {{ mod.website }}
            </div>
            <div>
              <span class="font-medium">Public:</span>
              {{ mod.public ? 'Yes' : 'No' }}
            </div>
            <div>
              <span class="font-medium">Created:</span>
              {{ formatDate(mod.created_at) }}
            </div>
            <div>
              <span class="font-medium">Updated:</span>
              {{ formatDate(mod.updated_at) }}
            </div>
          </TabsContent>

          <TabsContent value="versions">
            <VersionsList :mod-id="mod.id!" />
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
