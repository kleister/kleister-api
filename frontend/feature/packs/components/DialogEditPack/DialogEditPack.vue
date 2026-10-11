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
import { type Pack, updatePack } from '../../../../client'
import { toast } from '@/components/ui/toast'
import { Pencil } from '@lucide/vue'

const props = defineProps<{ pack: Pack }>()
const emit = defineEmits<{ (e: 'updated'): void }>()

const formSchema = z.object({
  name: z.string().min(3).max(255),
  slug: z.string().max(255).optional(),
  website: z.string().max(255).optional(),
  public: z.boolean().optional(),
})

const { handleSubmit, isSubmitting, isValidating, setFieldError, setValues } =
  useForm({
    validationSchema: formSchema,
  })

const isOpen = ref(false)

function closeModal() {
  isOpen.value = false
}

const onSubmit = handleSubmit(async (values) => {
  try {
    const { error } = await updatePack({
      path: { pack_id: props.pack.id! },
      body: {
        ...values,
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
          'We have encountered an unexpected issue while updating the pack. Try again later.',
        variant: 'destructive',
      })

      return
    }

    emit('updated')
    closeModal()

    toast({
      title: 'Success',
      description: 'Pack successfully updated.',
    })
  } catch (error) {
    console.error(error)

    toast({
      title: 'An error occurred',
      description:
        'We have encountered an unexpected issue while updating the pack. Try again later.',
      variant: 'destructive',
    })
  }
})

watch(isOpen, (value) => {
  if (!value) {
    return
  }

  setValues({
    name: props.pack.name,
    slug: props.pack.slug,
    website: props.pack.website,
    public: props.pack.public,
  })
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="ghost" size="icon" aria-label="Edit pack">
        <Pencil class="h-4 w-4" />
      </Button>
    </DialogTrigger>
    <DialogContent class="max-w-lg">
      <DialogHeader>
        <DialogTitle>Edit pack</DialogTitle>
        <DialogDescription class="sr-only">
          Update pack details. Click on Save when you're done.
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

        <FormField v-slot="{ componentField }" name="slug">
          <FormItem>
            <FormLabel>Slug</FormLabel>
            <FormControl>
              <Input type="text" autocomplete="off" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="website">
          <FormItem>
            <FormLabel>Website</FormLabel>
            <FormControl>
              <Input type="text" autocomplete="off" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

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
