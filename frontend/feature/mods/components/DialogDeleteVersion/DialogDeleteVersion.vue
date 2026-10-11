<script setup lang="ts">
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { ref } from 'vue'
import { type Version, deleteVersion } from '../../../../client'
import { toast } from '@/components/ui/toast'
import { Trash2 } from '@lucide/vue'

const props = defineProps<{ modId: string; version: Version }>()
const emit = defineEmits<{ (e: 'deleted'): void }>()

const isOpen = ref(false)
const isDeleting = ref(false)

async function performDelete() {
  isDeleting.value = true

  try {
    const { error } = await deleteVersion({
      path: { mod_id: props.modId, version_id: props.version.id! },
    })

    if (error) {
      console.error(error)

      toast({
        title: 'An error occurred',
        description:
          'We have encountered an unexpected issue while deleting the version. Try again later.',
        variant: 'destructive',
      })

      return
    }

    emit('deleted')

    toast({
      title: 'Success',
      description: 'Version successfully deleted.',
    })
  } catch (error) {
    console.error(error)

    toast({
      title: 'An error occurred',
      description:
        'We have encountered an unexpected issue while deleting the version. Try again later.',
      variant: 'destructive',
    })
  } finally {
    isDeleting.value = false
    isOpen.value = false
  }
}

function onTriggerClick(event: MouseEvent) {
  if (event.shiftKey) {
    void performDelete()
    return
  }

  isOpen.value = true
}
</script>

<template>
  <Button
    variant="ghost"
    size="icon"
    aria-label="Delete version"
    @click="onTriggerClick"
  >
    <Trash2 class="h-4 w-4" />
  </Button>

  <Dialog v-model:open="isOpen">
    <DialogContent class="max-w-sm">
      <DialogHeader>
        <DialogTitle>Delete version</DialogTitle>
        <DialogDescription>
          Are you sure you want to delete
          <strong>{{ version.name }}</strong
          >? This action cannot be undone. Tip: hold Shift while clicking
          delete to skip this confirmation.
        </DialogDescription>
      </DialogHeader>

      <DialogFooter>
        <Button
          variant="secondary"
          :disabled="isDeleting"
          @click="isOpen = false"
          >Cancel</Button
        >
        <Button
          variant="destructive"
          :disabled="isDeleting"
          @click="performDelete"
          >Delete</Button
        >
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
