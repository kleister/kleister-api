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
import { ref, watch } from 'vue'
import { createUser } from '../../../../client'
import { toast } from '@/components/ui/toast'
import { Checkbox } from '@/components/ui/checkbox'
import { useUsers } from '../../providers'

const { addUser } = useUsers()

const formSchema = toTypedSchema(
  z.object({
    username: z.string().min(3).max(255),
    password: z.string().min(3).max(255),
    email: z.string().email(),
    fullname: z.string().min(1).max(255),
    admin: z.boolean().optional(),
    active: z.boolean().optional(),
  })
)

const {
  handleSubmit,
  isSubmitting,
  isValidating,
  setFieldValue,
  setFieldError,
} = useForm({
  validationSchema: formSchema,
  initialValues: {
    active: true,
  },
})

const isOpen = ref(false)

function closeModal() {
  isOpen.value = false
}

const onSubmit = handleSubmit(async (values) => {
  try {
    const { data, error } = await createUser({
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

    addUser(data)
    closeModal()

    toast({
      title: 'Success',
      description: 'Users successfully created.',
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

watch(isOpen, (value) => {
  if (!value) {
    return
  }

  setFieldValue('active', true)
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="outline">New</Button>
    </DialogTrigger>
    <DialogContent class="max-w-lg">
      <DialogHeader>
        <DialogTitle>New user</DialogTitle>
        <DialogDescription class="sr-only">
          Enter user details. Click on Create when you're done.
        </DialogDescription>
      </DialogHeader>

      <form class="grid gap-4" @submit="onSubmit">
        <FormField v-slot="{ componentField }" name="username">
          <FormItem>
            <FormLabel>Username</FormLabel>
            <FormControl>
              <Input type="text" autocomplete="off" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="password">
          <FormItem>
            <FormLabel>Password</FormLabel>
            <FormControl>
              <Input
                type="password"
                autocomplete="off"
                v-bind="componentField"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="email">
          <FormItem>
            <FormLabel>Email</FormLabel>
            <FormControl>
              <Input type="text" autocomplete="off" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="fullname">
          <FormItem>
            <FormLabel>Fullname</FormLabel>
            <FormControl>
              <Input type="text" autocomplete="off" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="admin">
          <FormItem>
            <FormLabel>Admin</FormLabel>
            <FormControl>
              <Checkbox v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="active">
          <FormItem>
            <FormLabel>Active</FormLabel>
            <FormControl>
              <Checkbox v-bind="componentField" />
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
