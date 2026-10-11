<script setup lang="ts">
import {
  columnFilteringFeature,
  columnOrderingFeature,
  globalFilteringFeature,
  rowPaginationFeature,
  rowSortingFeature,
  createColumnHelper,
  FlexRender,
  tableFeatures,
  useTable,
  type ColumnOrderState,
  type HeaderContext,
  type PaginationState,
  type SortingState,
} from '@tanstack/vue-table'
import { refDebounced } from '@vueuse/core'
import { ArrowDown, ArrowUp, ArrowUpDown, Eye } from 'lucide-vue-next'
import { h, ref, unref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
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
import { valueUpdater } from '@/lib/utils'
import type { Mod, ListModsData } from '../../../../client'
import { useMods } from '../../providers'
import DialogDeleteMod from '../DialogDeleteMod'
import DialogEditMod from '../DialogEditMod'
import DialogNewMod from '../DialogNewMod'

const { mods, total, isLoading, loadMods } = useMods()

const features = tableFeatures({
  columnFilteringFeature,
  globalFilteringFeature,
  columnOrderingFeature,
  rowSortingFeature,
  rowPaginationFeature,
})

const sortColumnMap: Record<string, string> = {
  name: 'name',
  slug: 'slug',
  author: 'author',
  public: 'public',
  created_at: 'created',
  updated_at: 'updated',
}

const columnLabels: Record<string, string> = {
  name: 'Name',
  slug: 'Slug',
  author: 'Author',
  public: 'Public',
  created_at: 'Created',
  updated_at: 'Updated',
}

const pageSizes = [10, 20, 30, 40, 50, 100]
const defaultPageSize = 50

const route = useRoute()
const router = useRouter()

function parseInitialSorting(): SortingState {
  const sort = route.query.sort
  const order = route.query.order

  if (typeof sort === 'string' && sort in sortColumnMap) {
    return [{ id: sort, desc: order === 'desc' }]
  }

  return [{ id: 'name', desc: false }]
}

function parseInitialPagination(): PaginationState {
  const page = Number(route.query.page)
  const limit = Number(route.query.limit)

  return {
    pageIndex: Number.isInteger(page) && page > 0 ? page - 1 : 0,
    pageSize: pageSizes.includes(limit) ? limit : defaultPageSize,
  }
}

function parseInitialGlobalFilter(): string {
  return typeof route.query.search === 'string' ? route.query.search : ''
}

const sorting = ref<SortingState>(parseInitialSorting())
const pagination = ref<PaginationState>(parseInitialPagination())
const globalFilter = ref(parseInitialGlobalFilter())
const columnOrder = ref<ColumnOrderState>([])

const debouncedGlobalFilter = refDebounced(globalFilter, 300)

function sortableHeader(label: string, align: 'start' | 'end' = 'start') {
  return ({ column }: HeaderContext<typeof features, Mod>) => {
    const sorted = column.getIsSorted()
    const icon =
      sorted === 'asc' ? ArrowUp : sorted === 'desc' ? ArrowDown : ArrowUpDown

    return h(
      Button,
      {
        variant: 'ghost',
        size: 'sm',
        class: align === 'end' ? '-mr-3 ml-auto flex' : '-ml-3',
        onClick: column.getToggleSortingHandler(),
      },
      () => [label, h(icon, { class: 'ml-2 h-4 w-4' })]
    )
  }
}

const columnHelper = createColumnHelper<typeof features, Mod>()

const columns = columnHelper.columns([
  columnHelper.display({
    id: 'rowNumber',
    header: '#',
    cell: ({ row }) =>
      pagination.value.pageIndex * pagination.value.pageSize + row.index + 1,
  }),
  columnHelper.accessor('name', {
    header: sortableHeader(columnLabels.name),
    cell: (info) => info.getValue(),
  }),
  columnHelper.accessor('slug', {
    header: sortableHeader(columnLabels.slug),
    cell: (info) => info.getValue(),
  }),
  columnHelper.accessor('author', {
    header: sortableHeader(columnLabels.author),
    cell: (info) => info.getValue(),
  }),
  columnHelper.accessor('public', {
    header: sortableHeader(columnLabels.public),
    cell: (info) => (info.getValue() ? 'Yes' : 'No'),
  }),
  columnHelper.accessor('created_at', {
    header: sortableHeader(columnLabels.created_at, 'end'),
    cell: (info) => {
      const formattedDate = new Intl.DateTimeFormat('en-US').format(
        new Date(info.getValue())
      )

      return h('div', { class: 'text-end' }, formattedDate)
    },
  }),
  columnHelper.accessor('updated_at', {
    header: sortableHeader(columnLabels.updated_at, 'end'),
    cell: (info) => {
      const formattedDate = new Intl.DateTimeFormat('en-US').format(
        new Date(info.getValue())
      )

      return h('div', { class: 'text-end' }, formattedDate)
    },
  }),
  columnHelper.display({
    id: 'actions',
    header: () => h('div', { class: 'text-end' }, 'Actions'),
    cell: ({ row }) =>
      h('div', { class: 'flex items-center justify-end gap-1' }, [
        h(
          Button,
          {
            variant: 'ghost',
            size: 'icon',
            'aria-label': 'View mod details',
            onClick: () =>
              router.push({
                name: 'mod-detail',
                params: { mod_id: row.original.id },
              }),
          },
          () => h(Eye, { class: 'h-4 w-4' })
        ),
        h(DialogEditMod, {
          mod: row.original,
          onUpdated: () => applyQuery(),
        }),
        h(DialogDeleteMod, {
          mod: row.original,
          onDeleted: () => applyQuery(),
        }),
      ]),
  }),
])

const table = useTable({
  features,
  get data() {
    return unref(mods)
  },
  columns,
  enableMultiSort: false,
  manualSorting: true,
  manualFiltering: true,
  manualPagination: true,
  get rowCount() {
    return total.value
  },
  onSortingChange: (updater) => valueUpdater(updater, sorting),
  onGlobalFilterChange: (updater) => valueUpdater(updater, globalFilter),
  onPaginationChange: (updater) => valueUpdater(updater, pagination),
  onColumnOrderChange: (updater) => valueUpdater(updater, columnOrder),
  state: {
    get sorting() {
      return sorting.value
    },
    get globalFilter() {
      return globalFilter.value
    },
    get pagination() {
      return pagination.value
    },
    get columnOrder() {
      return columnOrder.value
    },
  },
})

function buildQuery(): ListModsData['query'] {
  const sort = sorting.value[0]

  return {
    search: debouncedGlobalFilter.value || undefined,
    sort: sort ? sortColumnMap[sort.id] : undefined,
    order: sort ? (sort.desc ? 'desc' : 'asc') : undefined,
    limit: pagination.value.pageSize,
    offset: pagination.value.pageIndex * pagination.value.pageSize,
  }
}

function applyQuery() {
  const query = buildQuery()
  const sort = sorting.value[0]

  router.replace({
    query: {
      ...route.query,
      search: query.search,
      sort: sort?.id,
      order: query.order,
      page: pagination.value.pageIndex + 1,
      limit: query.limit,
    },
  })

  loadMods(query)
}

watch([sorting, debouncedGlobalFilter], () => {
  if (pagination.value.pageIndex !== 0) {
    table.setPageIndex(0)
    return
  }

  applyQuery()
})

watch(pagination, applyQuery, { deep: true, immediate: true })

const nonDraggableColumnIds = ['rowNumber', 'actions']

function isDraggableHeader(header: {
  isPlaceholder: boolean
  column: { id: string }
}) {
  return (
    !header.isPlaceholder && !nonDraggableColumnIds.includes(header.column.id)
  )
}

const draggedColumnId = ref<string | null>(null)

function onColumnDragStart(columnId: string) {
  draggedColumnId.value = columnId
}

function onColumnDrop(targetColumnId: string) {
  const draggedId = draggedColumnId.value
  draggedColumnId.value = null

  if (!draggedId || draggedId === targetColumnId) {
    return
  }

  const order = table.getAllLeafColumns().map((column) => column.id)
  const from = order.indexOf(draggedId)
  const to = order.indexOf(targetColumnId)

  if (from === -1 || to === -1) {
    return
  }

  order.splice(from, 1)
  order.splice(to, 0, draggedId)

  table.setColumnOrder(order)
}
</script>

<template>
  <div class="w-full">
    <div class="flex flex-wrap items-center justify-between gap-2 py-4">
      <Input v-model="globalFilter" class="max-w-sm" placeholder="Search..." />

      <DialogNewMod />
    </div>

    <div class="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow
            v-for="headerGroup in table.getHeaderGroups()"
            :key="headerGroup.id"
          >
            <TableHead
              v-for="header in headerGroup.headers"
              :key="header.id"
              :draggable="isDraggableHeader(header)"
              :class="{ 'cursor-move': isDraggableHeader(header) }"
              @dragstart="onColumnDragStart(header.column.id)"
              @dragover.prevent
              @drop="onColumnDrop(header.column.id)"
            >
              <FlexRender v-if="!header.isPlaceholder" :header="header" />
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="table.getRowModel().rows?.length">
            <TableRow v-for="row in table.getRowModel().rows" :key="row.id">
              <TableCell v-for="cell in row.getAllCells()" :key="cell.id">
                <FlexRender :cell="cell" />
              </TableCell>
            </TableRow>
          </template>

          <TableRow v-else>
            <TableCell :colspan="columns.length" class="h-24 text-center">
              {{ isLoading ? 'Loading...' : 'No results' }}
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>

    <div class="flex flex-wrap items-center justify-between gap-4 py-4">
      <div class="text-sm text-muted-foreground">{{ total }} mod(s) total</div>

      <div class="flex flex-wrap items-center gap-6">
        <div class="flex items-center gap-2">
          <span class="text-sm text-muted-foreground">Rows per page</span>
          <Select
            :model-value="String(pagination.pageSize)"
            @update:model-value="(value) => table.setPageSize(Number(value))"
          >
            <SelectTrigger class="h-8 w-[70px]">
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectItem
                v-for="pageSize in pageSizes"
                :key="pageSize"
                :value="String(pageSize)"
              >
                {{ pageSize }}
              </SelectItem>
            </SelectContent>
          </Select>
        </div>

        <div class="text-sm text-muted-foreground">
          Page {{ pagination.pageIndex + 1 }} of
          {{ Math.max(table.getPageCount(), 1) }}
        </div>

        <div class="flex items-center gap-2">
          <Button
            variant="outline"
            size="sm"
            :disabled="!table.getCanPreviousPage()"
            @click="table.firstPage()"
          >
            First
          </Button>
          <Button
            variant="outline"
            size="sm"
            :disabled="!table.getCanPreviousPage()"
            @click="table.previousPage()"
          >
            Previous
          </Button>
          <Button
            variant="outline"
            size="sm"
            :disabled="!table.getCanNextPage()"
            @click="table.nextPage()"
          >
            Next
          </Button>
          <Button
            variant="outline"
            size="sm"
            :disabled="!table.getCanNextPage()"
            @click="table.lastPage()"
          >
            Last
          </Button>
        </div>
      </div>
    </div>
  </div>
</template>
