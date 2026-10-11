<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  type Version,
  showVersion,
  listVersionBuilds,
  attachVersionToBuild,
  deleteVersionFromBuild,
} from '../../../client'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { toast } from '@/components/ui/toast'

const route = useRoute()
const router = useRouter()

const modId = String(route.params.mod_id)
const versionId = String(route.params.version_id)

const version = ref<Version | null>(null)
const isLoading = ref(false)

const builds = ref<Array<{ packId: string; buildId: string; label: string }>>(
  []
)
const total = ref(0)

const newPackId = ref('')
const newBuildId = ref('')
const isAttaching = ref(false)

async function loadVersion() {
  isLoading.value = true

  const { data, error } = await showVersion({
    path: { mod_id: modId, version_id: versionId },
  })

  isLoading.value = false

  if (error) {
    throw error
  }

  version.value = data
}

async function loadBuilds() {
  const { data, error } = await listVersionBuilds({
    path: { mod_id: modId, version_id: versionId },
  })

  if (error) {
    throw error
  }

  total.value = data.total
  builds.value = data.builds.map((entry) => ({
    packId: entry.build?.pack?.id ?? '',
    buildId: entry.build_id,
    label: entry.build?.name ?? entry.build_id,
  }))
}

onMounted(() => {
  loadVersion()
  loadBuilds()
})

function reportError(action: string) {
  toast({
    title: 'An error occurred',
    description: `We have encountered an unexpected issue while ${action}. Try again later.`,
    variant: 'destructive',
  })
}

async function onAttach() {
  if (!newPackId.value || !newBuildId.value) {
    return
  }

  isAttaching.value = true

  try {
    const { error } = await attachVersionToBuild({
      path: { mod_id: modId, version_id: versionId },
      body: { pack: newPackId.value, build: newBuildId.value },
    })

    if (error) {
      reportError('attaching the build')
      return
    }

    newPackId.value = ''
    newBuildId.value = ''
    await loadBuilds()
  } finally {
    isAttaching.value = false
  }
}

async function onDrop(entry: { packId: string; buildId: string }) {
  const { error } = await deleteVersionFromBuild({
    path: { mod_id: modId, version_id: versionId },
    body: { pack: entry.packId, build: entry.buildId },
  })

  if (error) {
    reportError('removing the build')
    return
  }

  await loadBuilds()
}
</script>

<template>
  <div class="flex flex-col gap-4">
    <Button
      variant="outline"
      class="w-fit"
      @click="router.push({ name: 'mod-detail', params: { mod_id: modId } })"
    >
      Back to mod
    </Button>

    <Card v-if="version">
      <CardHeader>
        <CardTitle>{{ version.name }}</CardTitle>
      </CardHeader>
      <CardContent class="flex flex-col gap-4">
        <form class="flex flex-wrap items-end gap-2" @submit.prevent="onAttach">
          <div class="grid gap-1.5">
            <label class="text-sm font-medium" for="attach-pack">Pack</label>
            <Input
              id="attach-pack"
              v-model="newPackId"
              type="text"
              autocomplete="off"
              placeholder="pack id or slug"
            />
          </div>
          <div class="grid gap-1.5">
            <label class="text-sm font-medium" for="attach-build">Build</label>
            <Input
              id="attach-build"
              v-model="newBuildId"
              type="text"
              autocomplete="off"
              placeholder="build id or slug"
            />
          </div>
          <Button
            type="submit"
            :disabled="isAttaching || !newPackId || !newBuildId"
          >
            Attach
          </Button>
        </form>

        <div class="rounded-md border">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Build</TableHead>
                <TableHead class="text-end">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <template v-if="builds.length">
                <TableRow v-for="entry in builds" :key="entry.buildId">
                  <TableCell>{{ entry.label }}</TableCell>
                  <TableCell class="text-end">
                    <Button
                      variant="ghost"
                      size="sm"
                      :aria-label="`Remove ${entry.label}`"
                      @click="onDrop(entry)"
                    >
                      Remove
                    </Button>
                  </TableCell>
                </TableRow>
              </template>

              <TableRow v-else>
                <TableCell :colspan="2" class="h-16 text-center">
                  No results
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </div>

        <div class="text-sm text-muted-foreground">
          {{ total }} build(s) attached
        </div>
      </CardContent>
    </Card>

    <p v-else-if="isLoading">Loading...</p>
  </div>
</template>
