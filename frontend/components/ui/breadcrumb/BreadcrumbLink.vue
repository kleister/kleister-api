<script lang="ts" setup>
import type { HTMLAttributes } from 'vue'
import { cn } from '@/lib/utils'
import type { RouteLocationRaw } from 'vue-router'

const props = withDefaults(
  defineProps<{ to: RouteLocationRaw; class?: HTMLAttributes['class'] }>(),
  {
    class: '',
  }
)
</script>

<template>
  <RouterLink
    :to="props.to"
    :class="cn('transition-colors hover:text-foreground', props.class)"
  >
    <slot />
  </RouterLink>
</template>
