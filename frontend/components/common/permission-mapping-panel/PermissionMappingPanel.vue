<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { toast } from '@/components/ui/toast'

export interface PermissionMappingRow {
  subjectId: string
  label: string
  perm: string
}

const props = defineProps<{
  subjectLabel: string
  fetchRows: () => Promise<{ rows: PermissionMappingRow[]; total: number }>
  attachRow: (
    subjectId: string,
    perm: string
  ) => Promise<{ error?: unknown } | void>
  permitRow: (
    subjectId: string,
    perm: string
  ) => Promise<{ error?: unknown } | void>
  dropRow: (subjectId: string) => Promise<{ error?: unknown } | void>
}>()

const rows = ref<PermissionMappingRow[]>([])
const total = ref(0)
const isLoading = ref(false)

const newSubjectId = ref('')
const newPerm = ref('user')
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
  if (!newSubjectId.value) {
    return
  }

  isAttaching.value = true

  try {
    const result = await props.attachRow(newSubjectId.value, newPerm.value)

    if (result?.error) {
      reportError(`attaching the ${props.subjectLabel}`)
      return
    }

    newSubjectId.value = ''
    newPerm.value = 'user'
    await load()
  } catch {
    reportError(`attaching the ${props.subjectLabel}`)
  } finally {
    isAttaching.value = false
  }
}

async function onPermit(row: PermissionMappingRow, perm: string) {
  try {
    const result = await props.permitRow(row.subjectId, perm)

    if (result?.error) {
      reportError(`updating the ${props.subjectLabel}`)
      return
    }

    await load()
  } catch {
    reportError(`updating the ${props.subjectLabel}`)
  }
}

async function onDrop(row: PermissionMappingRow) {
  try {
    const result = await props.dropRow(row.subjectId)

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
    <form
      class="flex flex-wrap items-end gap-2"
      @submit.prevent="onAttach"
    >
      <div class="grid gap-1.5">
        <label class="text-sm font-medium" :for="`${subjectLabel}-id`">
          {{ subjectLabel }}
        </label>
        <Input
          :id="`${subjectLabel}-id`"
          v-model="newSubjectId"
          type="text"
          autocomplete="off"
          :placeholder="`${subjectLabel} id or slug`"
        />
      </div>

      <Select v-model="newPerm">
        <SelectTrigger class="w-[120px]">
          <SelectValue />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="user">User</SelectItem>
          <SelectItem value="admin">Admin</SelectItem>
        </SelectContent>
      </Select>

      <Button type="submit" :disabled="isAttaching || !newSubjectId">
        Attach
      </Button>
    </form>

    <div class="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Permission</TableHead>
            <TableHead class="text-end">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="rows.length">
            <TableRow v-for="row in rows" :key="row.subjectId">
              <TableCell>{{ row.label }}</TableCell>
              <TableCell>
                <Select
                  :model-value="row.perm"
                  @update:model-value="
                    (value) => onPermit(row, String(value))
                  "
                >
                  <SelectTrigger class="w-[120px]">
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="user">User</SelectItem>
                    <SelectItem value="admin">Admin</SelectItem>
                    <SelectItem value="owner" disabled>Owner</SelectItem>
                  </SelectContent>
                </Select>
              </TableCell>
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
            <TableCell :colspan="3" class="h-16 text-center">
              {{ isLoading ? 'Loading...' : 'No results' }}
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>

    <div class="text-sm text-muted-foreground">{{ total }} total</div>
  </div>
</template>
