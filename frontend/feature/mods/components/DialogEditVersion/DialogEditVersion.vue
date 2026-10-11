<script setup lang="ts">
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { z } from 'zod'
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Checkbox } from '@/components/ui/checkbox'
import { Button } from '@/components/ui/button'
import { useForm } from 'vee-validate'
import { ref, watch } from 'vue'
import { type Version, updateVersion } from '../../../../client'
import { toast } from '@/components/ui/toast'
import { Pencil } from 'lucide-vue-next'

const props = defineProps<{ modId: string; version: Version }>()
const emit = defineEmits<{ (e: 'updated'): void }>()

const formSchema = z.object({
  name: z.string().min(3).max(255).optional(),
  public: z.boolean().optional(),
})

const { handleSubmit, isSubmitting, isValidating, setFieldError, setValues } =
  useForm({
    validationSchema: formSchema,
  })

const isOpen = ref(false)
const uploadDataUrl = ref<string | undefined>(undefined)

function closeModal() {
  isOpen.value = false
  uploadDataUrl.value = undefined
}

function onFileChange(event: Event) {
  const file = (event.target as HTMLInputElement).files?.[0]

  if (!file) {
    uploadDataUrl.value = undefined
    return
  }

  const reader = new FileReader()
  reader.onload = () => {
    uploadDataUrl.value = reader.result as string
  }
  reader.readAsDataURL(file)
}

const onSubmit = handleSubmit(async (values) => {
  try {
    const { error } = await updateVersion({
      path: { mod_id: props.modId, version_id: props.version.id! },
      body: {
        ...values,
        upload: uploadDataUrl.value,
      },
    })

    if (error?.status === 422) {
      for (const e of error.errors!) {
        setFieldError(
          // @ts-expect-error would be cumbersome to type the field to the keys
          e.field,
          e.message
        )
      }

      return
    }

    if (error) {
      console.error(error)

      toast({
        title: 'An error occurred',
        description:
          'We have encountered an unexpected issue while updating the version. Try again later.',
        variant: 'destructive',
      })

      return
    }

    emit('updated')
    closeModal()

    toast({
      title: 'Success',
      description: 'Version successfully updated.',
    })
  } catch (error) {
    console.error(error)

    toast({
      title: 'An error occurred',
      description:
        'We have encountered an unexpected issue while updating the version. Try again later.',
      variant: 'destructive',
    })
  }
})

watch(isOpen, (value) => {
  if (!value) {
    return
  }

  setValues({
    name: props.version.name,
    public: props.version.public,
  })
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="ghost" size="icon" aria-label="Edit version">
        <Pencil class="h-4 w-4" />
      </Button>
    </DialogTrigger>
    <DialogContent class="max-w-lg">
      <DialogHeader>
        <DialogTitle>Edit version</DialogTitle>
        <DialogDescription class="sr-only">
          Update version details. Click on Save when you're done.
        </DialogDescription>
      </DialogHeader>

      <form class="grid gap-4" @submit="onSubmit">
        <FormField v-slot="{ componentField }" name="name">
          <FormItem>
            <FormLabel>Name</FormLabel>
            <FormControl>
              <Input type="text" autocomplete="off" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <div class="grid gap-1.5">
          <label class="text-sm font-medium" for="version-upload"
            >Replace file</label
          >
          <Input id="version-upload" type="file" @change="onFileChange" />
        </div>

        <FormField v-slot="{ value, handleChange }" name="public">
          <FormItem class="flex flex-row items-center gap-2">
            <FormControl>
              <Checkbox :model-value="value" @update:model-value="handleChange" />
            </FormControl>
            <FormLabel>Public</FormLabel>
          </FormItem>
        </FormField>

        <DialogFooter>
          <Button
            variant="secondary"
            :disabled="isSubmitting || isValidating"
            @click="closeModal"
            >Cancel</Button
          >
          <Button type="submit" :disabled="isSubmitting || isValidating"
            >Save</Button
          >
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>
