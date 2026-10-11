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
import { ref } from 'vue'
import { createVersion } from '../../../../client'
import { toast } from '@/components/ui/toast'

const props = defineProps<{ modId: string }>()
const emit = defineEmits<{ (e: 'created'): void }>()

const formSchema = z.object({
  name: z.string().min(3).max(255).optional(),
  public: z.boolean().optional(),
})

const { handleSubmit, isSubmitting, isValidating, setFieldError } = useForm({
  validationSchema: formSchema,
  initialValues: { public: false },
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
    const { error } = await createVersion({
      path: { mod_id: props.modId },
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
          'We have encountered an unexpected issue while creating the version. Try again later.',
        variant: 'destructive',
      })

      return
    }

    emit('created')
    closeModal()

    toast({
      title: 'Success',
      description: 'Version successfully created.',
    })
  } catch (error) {
    console.error(error)

    toast({
      title: 'An error occurred',
      description:
        'We have encountered an unexpected issue while creating the version. Try again later.',
      variant: 'destructive',
    })
  }
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="outline">New</Button>
    </DialogTrigger>
    <DialogContent class="max-w-lg">
      <DialogHeader>
        <DialogTitle>New version</DialogTitle>
        <DialogDescription class="sr-only">
          Enter version details. Click on Create when you're done.
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
          <label class="text-sm font-medium" for="version-upload">File</label>
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
            >Create</Button
          >
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>
