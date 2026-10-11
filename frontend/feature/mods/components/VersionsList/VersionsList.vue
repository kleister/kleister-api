<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { Button } from '@/components/ui/button'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Eye } from '@lucide/vue'
import { type Version, listVersions } from '../../../../client'
import DialogDeleteVersion from '../DialogDeleteVersion'
import DialogEditVersion from '../DialogEditVersion'
import DialogNewVersion from '../DialogNewVersion'

const props = defineProps<{ modId: string }>()

const router = useRouter()

const versions = ref<Version[]>([])
const total = ref(0)
const isLoading = ref(false)

async function loadVersions() {
  isLoading.value = true

  const { data, error } = await listVersions({
    path: { mod_id: props.modId },
  })

  isLoading.value = false

  if (error) {
    throw error
  }

  versions.value = data.versions
  total.value = data.total
}

onMounted(loadVersions)
watch(() => props.modId, loadVersions)
</script>

<template>
  <div class="flex flex-col gap-4">
    <div class="flex justify-end">
      <DialogNewVersion :mod-id="modId" @created="loadVersions" />
    </div>

    <div class="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Public</TableHead>
            <TableHead class="text-end">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="versions.length">
            <TableRow v-for="version in versions" :key="version.id">
              <TableCell>{{ version.name }}</TableCell>
              <TableCell>{{ version.public ? 'Yes' : 'No' }}</TableCell>
              <TableCell class="text-end">
                <div class="flex items-center justify-end gap-1">
                  <Button
                    variant="ghost"
                    size="icon"
                    aria-label="View version details"
                    @click="
                      router.push({
                        name: 'mod-version-detail',
                        params: { mod_id: modId, version_id: version.id },
                      })
                    "
                  >
                    <Eye class="h-4 w-4" />
                  </Button>
                  <DialogEditVersion
                    :mod-id="modId"
                    :version="version"
                    @updated="loadVersions"
                  />
                  <DialogDeleteVersion
                    :mod-id="modId"
                    :version="version"
                    @deleted="loadVersions"
                  />
                </div>
              </TableCell>
            </TableRow>
          </template>

          <TableRow v-else>
            <TableCell :colspan="3" class="h-16 text-center">
              {{ isLoading ? 'Loading...' : 'No results' }}
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>

    <div class="text-sm text-muted-foreground">
      {{ total }} version(s) total
    </div>
  </div>
</template>
