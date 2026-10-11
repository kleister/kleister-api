<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Button } from '@/components/ui/button'
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

export interface RelationAttachRow {
  id: string
  label: string
}

const props = defineProps<{
  subjectLabel: string
  fetchRows: () => Promise<{ rows: RelationAttachRow[]; total: number }>
  attachRow: (id: string) => Promise<{ error?: unknown } | void>
  dropRow: (id: string) => Promise<{ error?: unknown } | void>
}>()

const rows = ref<RelationAttachRow[]>([])
const total = ref(0)
const isLoading = ref(false)

const newId = ref('')
const isAttaching = ref(false)

async function load() {
  isLoading.value = true

  try {
    const result = await props.fetchRows()
    rows.value = result.rows
    total.value = result.total
  } finally {
    isLoading.value = false
  }
}

function reportError(action: string) {
  toast({
    title: 'An error occurred',
    description: `We have encountered an unexpected issue while ${action}. Try again later.`,
    variant: 'destructive',
  })
}

async function onAttach() {
  if (!newId.value) {
    return
  }

  isAttaching.value = true

  try {
    const result = await props.attachRow(newId.value)

    if (result?.error) {
      reportError(`attaching the ${props.subjectLabel}`)
      return
    }

    newId.value = ''
    await load()
  } catch {
    reportError(`attaching the ${props.subjectLabel}`)
  } finally {
    isAttaching.value = false
  }
}

async function onDrop(row: RelationAttachRow) {
  try {
    const result = await props.dropRow(row.id)

    if (result?.error) {
      reportError(`removing the ${props.subjectLabel}`)
      return
    }

    await load()
  } catch {
    reportError(`removing the ${props.subjectLabel}`)
  }
}

onMounted(load)
</script>

<template>
  <div class="flex flex-col gap-4">
    <form class="flex flex-wrap items-end gap-2" @submit.prevent="onAttach">
      <div class="grid gap-1.5">
        <label class="text-sm font-medium" :for="`${subjectLabel}-id`">
          {{ subjectLabel }}
        </label>
        <Input
          :id="`${subjectLabel}-id`"
          v-model="newId"
          type="text"
          autocomplete="off"
          :placeholder="`${subjectLabel} id or slug`"
        />
      </div>

      <Button type="submit" :disabled="isAttaching || !newId"> Attach </Button>
    </form>

    <div class="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead class="text-end">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="rows.length">
            <TableRow v-for="row in rows" :key="row.id">
              <TableCell>{{ row.label }}</TableCell>
              <TableCell class="text-end">
                <Button
                  variant="ghost"
                  size="sm"
                  :aria-label="`Remove ${row.label}`"
                  @click="onDrop(row)"
                >
                  Remove
                </Button>
              </TableCell>
            </TableRow>
          </template>

          <TableRow v-else>
            <TableCell :colspan="2" class="h-16 text-center">
              {{ isLoading ? 'Loading...' : 'No results' }}
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>

    <div class="text-sm text-muted-foreground">{{ total }} total</div>
  </div>
</template>
