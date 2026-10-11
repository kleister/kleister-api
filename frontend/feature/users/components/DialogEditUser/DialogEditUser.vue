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
import { Button } from '@/components/ui/button'
import { useForm } from 'vee-validate'
import { ref, watch } from 'vue'
import { type User, updateUser } from '../../../../client'
import { toast } from '@/components/ui/toast'
import { Checkbox } from '@/components/ui/checkbox'
import { Pencil } from 'lucide-vue-next'

const props = defineProps<{ user: User }>()
const emit = defineEmits<{ (e: 'updated'): void }>()

const formSchema = z.object({
  username: z.string().min(3).max(255),
  password: z.union([z.literal(''), z.string().min(3).max(255)]),
  email: z.string().email(),
  fullname: z.string().min(1).max(255),
  admin: z.boolean().optional(),
  active: z.boolean().optional(),
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
    const { password, ...rest } = values

    const { error } = await updateUser({
      path: { user_id: props.user.id! },
      body: {
        ...rest,
        ...(password ? { password } : {}),
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
          'We have encountered an unexpected issue while updating the user. Try again later.',
        variant: 'destructive',
      })

      return
    }

    emit('updated')
    closeModal()

    toast({
      title: 'Success',
      description: 'User successfully updated.',
    })
  } catch (error) {
    console.error(error)

    toast({
      title: 'An error occurred',
      description:
        'We have encountered an unexpected issue while updating the user. Try again later.',
      variant: 'destructive',
    })
  }
})

watch(isOpen, (value) => {
  if (!value) {
    return
  }

  setValues({
    username: props.user.username,
    password: '',
    email: props.user.email,
    fullname: props.user.fullname,
    admin: props.user.admin ?? false,
    active: props.user.active ?? false,
  })
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="ghost" size="icon" aria-label="Edit user">
        <Pencil class="h-4 w-4" />
      </Button>
    </DialogTrigger>
    <DialogContent class="max-w-lg">
      <DialogHeader>
        <DialogTitle>Edit user</DialogTitle>
        <DialogDescription class="sr-only">
          Update user details. Click on Save when you're done.
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
                placeholder="Leave blank to keep the current password"
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
            >Save</Button
          >
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>
