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
import { type Build, listBuilds } from '../../../../client'
import DialogDeleteBuild from '../DialogDeleteBuild'
import DialogEditBuild from '../DialogEditBuild'
import DialogNewBuild from '../DialogNewBuild'

const props = defineProps<{ packId: string }>()

const router = useRouter()

const builds = ref<Build[]>([])
const total = ref(0)
const isLoading = ref(false)

async function loadBuilds() {
  isLoading.value = true

  const { data, error } = await listBuilds({
    path: { pack_id: props.packId },
  })

  isLoading.value = false

  if (error) {
    throw error
  }

  builds.value = data.builds
  total.value = data.total
}

onMounted(loadBuilds)
watch(() => props.packId, loadBuilds)
</script>

<template>
  <div class="flex flex-col gap-4">
    <div class="flex justify-end">
      <DialogNewBuild :pack-id="packId" @created="loadBuilds" />
    </div>

    <div class="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Latest</TableHead>
            <TableHead>Recommended</TableHead>
            <TableHead>Public</TableHead>
            <TableHead class="text-end">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="builds.length">
            <TableRow v-for="build in builds" :key="build.id">
              <TableCell>{{ build.name }}</TableCell>
              <TableCell>{{ build.latest ? 'Yes' : 'No' }}</TableCell>
              <TableCell>{{ build.recommended ? 'Yes' : 'No' }}</TableCell>
              <TableCell>{{ build.public ? 'Yes' : 'No' }}</TableCell>
              <TableCell class="text-end">
                <div class="flex items-center justify-end gap-1">
                  <Button
                    variant="ghost"
                    size="icon"
                    aria-label="View build details"
                    @click="
                      router.push({
                        name: 'pack-build-detail',
                        params: { pack_id: packId, build_id: build.id },
                      })
                    "
                  >
                    <Eye class="h-4 w-4" />
                  </Button>
                  <DialogEditBuild
                    :pack-id="packId"
                    :build="build"
                    @updated="loadBuilds"
                  />
                  <DialogDeleteBuild
                    :pack-id="packId"
                    :build="build"
                    @deleted="loadBuilds"
                  />
                </div>
              </TableCell>
            </TableRow>
          </template>

          <TableRow v-else>
            <TableCell :colspan="5" class="h-16 text-center">
              {{ isLoading ? 'Loading...' : 'No results' }}
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>

    <div class="text-sm text-muted-foreground">{{ total }} build(s) total</div>
  </div>
</template>
