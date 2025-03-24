import { Group, LayoutDashboard, Users } from 'lucide-vue-next'
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
      ],
      admin: [
        {
          name: 'Users',
          url: `/admin/users`,
          icon: Users,
        },
        {
          name: 'Groups',
          url: `/admin/groups`,
          icon: Group,
        },
      ],
    }
  })

  return { links }
}
