<script setup lang="ts">
import {
  columnFilteringFeature,
  columnVisibilityFeature,
  createFilteredRowModel,
  FlexRender,
  rowExpandingFeature,
  rowSelectionFeature,
  tableFeatures,
  useTable,
  type ColumnDef,
  type ColumnFiltersState,
} from '@tanstack/vue-table'
import { useGroups } from '../../providers'
import type { Group } from '../../../../client'
import { Input } from '@/components/ui/input'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import DialogNewGroups from '../DialogNewGroup'
import { computed, h, ref, unref } from 'vue'
import { valueUpdater } from '@/lib/utils'

const { groups, loadGroups } = useGroups()

const features = tableFeatures({
  columnFilteringFeature,
  columnVisibilityFeature,
  rowExpandingFeature,
  rowSelectionFeature,
  filteredRowModel: createFilteredRowModel(),
})

const columns: ColumnDef<typeof features, Group>[] = [
  {
    accessorKey: 'name',
    header: 'Name',
    cell: ({ row }) => row.getValue('name'),
  },
  {
    accessorKey: 'slug',
    header: 'Slug',
    cell: ({ row }) => row.getValue('slug'),
  },
  {
    accessorKey: 'created_at',
    header: () => h('div', { class: 'text-end' }, 'Created'),
    cell: ({ row }) => {
      const formattedDate = new Intl.DateTimeFormat('en-US').format(
        new Date(row.getValue('created_at'))
      )

      return h('div', { class: 'text-end' }, formattedDate)
    },
  },
  {
    accessorKey: 'updated_at',
    header: () => h('div', { class: 'text-end' }, 'Updated'),
    cell: ({ row }) => {
      const formattedDate = new Intl.DateTimeFormat('en-US').format(
        new Date(row.getValue('updated_at'))
      )

      return h('div', { class: 'text-end' }, formattedDate)
    },
  },
]

const columnFilters = ref<ColumnFiltersState>([])

const table = useTable({
  features,
  // Using data directly without a getter was not reactive...
  // Might be due to the data coming from context
  // Not really worth the effort to investigate further since the getter works fine
  get data() {
    return unref(groups)
  },
  columns,
  onColumnFiltersChange: (updaterOrValue) =>
    valueUpdater(updaterOrValue, columnFilters),
  state: {
    get columnFilters() {
      return columnFilters.value
    },
  },
})

const filter = computed({
  get() {
    return table.getColumn('name')?.getFilterValue() as string
  },
  set(value: string) {
    table.getColumn('name')?.setFilterValue(value)
  },
})

loadGroups()
</script>

<template>
  <div class="w-full">
    <div class="flex gap-2 items-center justify-between py-4">
      <Input v-model="filter" class="max-w-sm" placeholder="Filter" />
      <DialogNewGroups />
    </div>
    <div class="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow
            v-for="headerGroup in table.getHeaderGroups()"
            :key="headerGroup.id"
          >
            <TableHead v-for="header in headerGroup.headers" :key="header.id">
              <FlexRender v-if="!header.isPlaceholder" :header="header" />
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="table.getRowModel().rows?.length">
            <template v-for="row in table.getRowModel().rows" :key="row.id">
              <TableRow :data-state="row.getIsSelected() && 'selected'">
                <TableCell v-for="cell in row.getVisibleCells()" :key="cell.id">
                  <FlexRender :cell="cell" />
                </TableCell>
              </TableRow>
              <TableRow v-if="row.getIsExpanded()">
                <TableCell :colspan="row.getAllCells().length">
                  {{ JSON.stringify(row.original) }}
                </TableCell>
              </TableRow>
            </template>
          </template>

          <TableRow v-else>
            <TableCell :colspan="columns.length" class="h-24 text-center">
              No results
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>
  </div>
</template>
