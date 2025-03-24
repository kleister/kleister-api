import { computed } from 'vue'
import { useRoute } from 'vue-router'

export function useBreadcrumbs() {
  const route = useRoute()

  const breadcrumbs = computed(() => {
    const res = []

    res.push({
      name: route.name,
    })

    return res
  })

  return { breadcrumbs }
}
