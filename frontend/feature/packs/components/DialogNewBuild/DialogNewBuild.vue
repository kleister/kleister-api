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
import { createBuild } from '../../../../client'
import { toast } from '@/components/ui/toast'

const props = defineProps<{ packId: string }>()
const emit = defineEmits<{ (e: 'created'): void }>()

const formSchema = z.object({
  name: z.string().min(3).max(255).optional(),
  java: z.string().max(255).optional(),
  memory: z.string().max(255).optional(),
  latest: z.boolean().optional(),
  recommended: z.boolean().optional(),
  public: z.boolean().optional(),
})

const { handleSubmit, isSubmitting, isValidating, setFieldError } = useForm({
  validationSchema: formSchema,
  initialValues: { latest: false, recommended: false, public: false },
})

const isOpen = ref(false)

function closeModal() {
  isOpen.value = false
}

const onSubmit = handleSubmit(async (values) => {
  try {
    const { error } = await createBuild({
      path: { pack_id: props.packId },
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
          'We have encountered an unexpected issue while creating the build. Try again later.',
        variant: 'destructive',
      })

      return
    }

    emit('created')
    closeModal()

    toast({
      title: 'Success',
      description: 'Build successfully created.',
    })
  } catch (error) {
    console.error(error)

    toast({
      title: 'An error occurred',
      description:
        'We have encountered an unexpected issue while creating the build. Try again later.',
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
        <DialogTitle>New build</DialogTitle>
        <DialogDescription class="sr-only">
          Enter build details. Click on Create when you're done.
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

        <FormField v-slot="{ componentField }" name="java">
          <FormItem>
            <FormLabel>Java</FormLabel>
            <FormControl>
              <Input type="text" autocomplete="off" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="memory">
          <FormItem>
            <FormLabel>Memory</FormLabel>
            <FormControl>
              <Input type="text" autocomplete="off" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ value, handleChange }" name="latest">
          <FormItem class="flex flex-row items-center gap-2">
            <FormControl>
              <Checkbox :model-value="value" @update:model-value="handleChange" />
            </FormControl>
            <FormLabel>Latest</FormLabel>
          </FormItem>
        </FormField>

        <FormField v-slot="{ value, handleChange }" name="recommended">
          <FormItem class="flex flex-row items-center gap-2">
            <FormControl>
              <Checkbox :model-value="value" @update:model-value="handleChange" />
            </FormControl>
            <FormLabel>Recommended</FormLabel>
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
            >Create</Button
          >
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>
