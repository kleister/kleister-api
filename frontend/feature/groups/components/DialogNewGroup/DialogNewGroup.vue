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
import { toTypedSchema } from '@vee-validate/zod'
import { z } from 'zod'
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { useForm } from 'vee-validate'
import { ref } from 'vue'
import { createGroup } from '../../../../client'
import { toast } from '@/components/ui/toast'
import { useGroups } from '../../providers'

const { addGroup } = useGroups()

const formSchema = toTypedSchema(
  z.object({
    name: z.string().min(3).max(255),
    slug: z.string().max(255).optional(),
  })
)

const { handleSubmit, isSubmitting, isValidating, setFieldError } = useForm({
  validationSchema: formSchema,
  initialValues: {},
})

const isOpen = ref(false)

function closeModal() {
  isOpen.value = false
}

const onSubmit = handleSubmit(async (values) => {
  try {
    const { data, error } = await createGroup({
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
          'We have encountered an unexpected issue while creating the project. Try again later.',
        variant: 'destructive',
      })

      return
    }

    addGroup(data)
    closeModal()

    toast({
      title: 'Success',
      description: 'Groups successfully created.',
    })
  } catch (error) {
    console.error(error)

    toast({
      title: 'An error occurred',
      description:
        'We have encountered an unexpected issue while creating the project. Try again later.',
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
        <DialogTitle>New group</DialogTitle>
        <DialogDescription class="sr-only">
          Enter group details. Click on Create when you're done.
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
