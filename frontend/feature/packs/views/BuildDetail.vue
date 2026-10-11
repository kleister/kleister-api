<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  type Build,
  showBuild,
  listBuildVersions,
  attachBuildToVersion,
  deleteBuildFromVersion,
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

const packId = String(route.params.pack_id)
const buildId = String(route.params.build_id)

const build = ref<Build | null>(null)
const isLoading = ref(false)

const versions = ref<
  Array<{ modId: string; versionId: string; label: string }>
>([])
const total = ref(0)

const newModId = ref('')
const newVersionId = ref('')
const isAttaching = ref(false)

async function loadBuild() {
  isLoading.value = true

  const { data, error } = await showBuild({
    path: { pack_id: packId, build_id: buildId },
  })

  isLoading.value = false

  if (error) {
    throw error
  }

  build.value = data
}

async function loadVersions() {
  const { data, error } = await listBuildVersions({
    path: { pack_id: packId, build_id: buildId },
  })

  if (error) {
    throw error
  }

  total.value = data.total
  versions.value = data.versions.map((entry) => ({
    modId: entry.version?.mod?.id ?? '',
    versionId: entry.version_id,
    label: entry.version?.name ?? entry.version_id,
  }))
}

onMounted(() => {
  loadBuild()
  loadVersions()
})

function reportError(action: string) {
  toast({
    title: 'An error occurred',
    description: `We have encountered an unexpected issue while ${action}. Try again later.`,
    variant: 'destructive',
  })
}

async function onAttach() {
  if (!newModId.value || !newVersionId.value) {
    return
  }

  isAttaching.value = true

  try {
    const { error } = await attachBuildToVersion({
      path: { pack_id: packId, build_id: buildId },
      body: { mod: newModId.value, version: newVersionId.value },
    })

    if (error) {
      reportError('attaching the version')
      return
    }

    newModId.value = ''
    newVersionId.value = ''
    await loadVersions()
  } finally {
    isAttaching.value = false
  }
}

async function onDrop(entry: { modId: string; versionId: string }) {
  const { error } = await deleteBuildFromVersion({
    path: { pack_id: packId, build_id: buildId },
    body: { mod: entry.modId, version: entry.versionId },
  })

  if (error) {
    reportError('removing the version')
    return
  }

  await loadVersions()
}
</script>

<template>
  <div class="flex flex-col gap-4">
    <Button
      variant="outline"
      class="w-fit"
      @click="router.push({ name: 'pack-detail', params: { pack_id: packId } })"
    >
      Back to pack
    </Button>

    <Card v-if="build">
      <CardHeader>
        <CardTitle>{{ build.name }}</CardTitle>
      </CardHeader>
      <CardContent class="flex flex-col gap-4">
        <form class="flex flex-wrap items-end gap-2" @submit.prevent="onAttach">
          <div class="grid gap-1.5">
            <label class="text-sm font-medium" for="attach-mod">Mod</label>
            <Input
              id="attach-mod"
              v-model="newModId"
              type="text"
              autocomplete="off"
              placeholder="mod id or slug"
            />
          </div>
          <div class="grid gap-1.5">
            <label class="text-sm font-medium" for="attach-version"
              >Version</label
            >
            <Input
              id="attach-version"
              v-model="newVersionId"
              type="text"
              autocomplete="off"
              placeholder="version id or slug"
            />
          </div>
          <Button
            type="submit"
            :disabled="isAttaching || !newModId || !newVersionId"
          >
            Attach
          </Button>
        </form>

        <div class="rounded-md border">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Version</TableHead>
                <TableHead class="text-end">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <template v-if="versions.length">
                <TableRow v-for="entry in versions" :key="entry.versionId">
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
          {{ total }} version(s) attached
        </div>
      </CardContent>
    </Card>

    <p v-else-if="isLoading">Loading...</p>
  </div>
</template>
