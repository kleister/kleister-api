import { Blocks, Group, LayoutDashboard, Package, Users } from '@lucide/vue'
import { computed } from 'vue'

export function useNavigationLinks() {
  const links = computed(() => {
    return {
      general: [
        {
          name: 'Dashboard',
          url: `/`,
          icon: LayoutDashboard,
        },
        {
          name: 'Packs',
          url: `/packs`,
          icon: Package,
        },
        {
          name: 'Mods',
          url: `/mods`,
          icon: Blocks,
        },
      ],
      admin: [
        {
          name: 'Users',
          url: `/users`,
          icon: Users,
        },
        {
          name: 'Groups',
          url: `/groups`,
          icon: Group,
        },
      ],
    }
  })

  return { links }
}
