<script lang="ts" setup>
import type { HTMLAttributes } from 'vue'
import { cn } from '@/lib/utils'
import { type RouterLinkProps } from 'vue-router'

const props = withDefaults(
  defineProps<RouterLinkProps & { class?: HTMLAttributes['class'] }>(),
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
