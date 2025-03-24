<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import OauthProviders from '../components/oauth-providers'
import { loginAuth } from '../../../client'
import { useAuthStore } from '../store/auth'
import { useRouter } from 'vue-router'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import { useForm } from 'vee-validate'
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { computed, unref } from 'vue'
import { toast } from '@/components/ui/toast'

const { signInUser } = useAuthStore()
const router = useRouter()

const formSchema = toTypedSchema(
  z.object({
    username: z.string().min(1).max(255),
    password: z.string().min(1).max(255),
  })
)

const { handleSubmit, isSubmitting, isValidating, errors, values } = useForm({
  validationSchema: formSchema,
})

const isSubmitDisabled = computed(() => {
  return (
    unref(isSubmitting) ||
    unref(isValidating) ||
    unref(errors).password ||
    unref(errors).username ||
    !values.username ||
    !values.password
  )
})

const onSubmit = handleSubmit(async (values, actions) => {
  try {
    const { data, error } = await loginAuth({
      body: values as Required<typeof values>,
    })

    if (error?.status === 401) {
      actions.setErrors({
        username: 'Invalid username or password',
      })
      return
    }

    if (error) {
      console.error(error)
      toast({
        title: 'An error occurred',
        description:
          'We have encountered an unexpected issue while signing you in. Try again later.',
        variant: 'destructive',
      })

      return
    }

    signInUser(data.token)
    await router.replace('/')
  } catch (error) {
    console.error(error)
    toast({
      title: 'An error occurred',
      description:
        'We have encountered an unexpected issue while signing you in. Try again later.',
      variant: 'destructive',
    })
  }
})
</script>

<template>
  <div class="min-h-dvh flex items-center justify-center">
    <Card class="mx-auto w-full max-w-sm">
      <CardHeader>
        <CardTitle class="text-2xl text-center">Welcome</CardTitle>
      </CardHeader>
      <CardContent>
        <form @submit="onSubmit">
          <div class="grid gap-4">
            <FormField v-slot="{ componentField }" name="username">
              <FormItem>
                <FormLabel>Username</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    autocomplete="username"
                    v-bind="componentField"
                  />
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
                    autocomplete="current-password"
                    v-bind="componentField"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            </FormField>
            <Button type="submit" class="w-full" :disabled="isSubmitDisabled"
              >Signin</Button
            >

            <div class="text-center text-sm text-gray-500 my-4">Or</div>
            <OauthProviders />
          </div>
        </form>
      </CardContent>
    </Card>
  </div>
</template>
