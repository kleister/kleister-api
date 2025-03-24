<script setup lang="ts">
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarRail,
} from '@/components/ui/sidebar'
import { BadgeCheck, ChevronsUpDown, LogOut } from 'lucide-vue-next'
import { Avatar, AvatarFallback } from '@/components/ui/avatar'
import { useAuthStore } from '@/feature/auth/store/auth'
import { storeToRefs } from 'pinia'
import { RouterLink } from 'vue-router'
import { useNavigationLinks } from '@/composables/useNavigationLinks'

const authStore = useAuthStore()
const { user } = storeToRefs(authStore)

const { links } = useNavigationLinks()
</script>

<template>
  <Sidebar collapsible="icon">
    <!-- <SidebarHeader>
      <SidebarMenu>
        <SidebarMenuItem>
        </SidebarMenuItem>
      </SidebarMenu>
    </SidebarHeader> -->
    <SidebarContent>
      <SidebarGroup>
        <SidebarGroupLabel>General</SidebarGroupLabel>
        <SidebarMenu>
          <SidebarMenuItem v-for="item in links.general" :key="item.name">
            <SidebarMenuButton as-child>
              <RouterLink :to="item.url">
                <component :is="item.icon" />
                <span>{{ item.name }}</span>
              </RouterLink>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarGroup>
      <SidebarGroup v-if="user.isAdmin">
        <SidebarGroupLabel>Admin</SidebarGroupLabel>
        <SidebarMenu>
          <SidebarMenuItem v-for="item in links.admin" :key="item.name">
            <SidebarMenuButton as-child>
              <RouterLink :to="item.url">
                <component :is="item.icon" />
                <span>{{ item.name }}</span>
              </RouterLink>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarGroup>
    </SidebarContent>
    <SidebarFooter>
      <SidebarMenu>
        <SidebarMenuItem>
          <DropdownMenu>
            <DropdownMenuTrigger as-child>
              <SidebarMenuButton
                size="lg"
                class="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground"
              >
                <Avatar class="h-8 w-8 rounded-lg">
                  <AvatarFallback class="rounded-lg">{{
                    user.profile
                  }}</AvatarFallback>
                </Avatar>
                <div class="grid flex-1 text-left text-sm leading-tight">
                  <span class="truncate font-semibold">{{
                    user.displayName
                  }}</span>
                  <span class="truncate text-xs">{{ user.email }}</span>
                </div>
                <ChevronsUpDown class="ml-auto size-4" />
              </SidebarMenuButton>
            </DropdownMenuTrigger>
            <DropdownMenuContent
              class="w-[--radix-dropdown-menu-trigger-width] min-w-56 rounded-lg"
              side="bottom"
              align="end"
              :side-offset="4"
            >
              <DropdownMenuLabel class="p-0 font-normal">
                <div
                  class="flex items-center gap-2 px-1 py-1.5 text-left text-sm"
                >
                  <Avatar class="h-8 w-8 rounded-lg">
                    <AvatarFallback class="rounded-lg">{{
                      user.profile
                    }}</AvatarFallback>
                  </Avatar>
                  <div class="grid flex-1 text-left text-sm leading-tight">
                    <span class="truncate font-semibold">{{
                      user.displayName
                    }}</span>
                    <span class="truncate text-xs">{{ user.email }}</span>
                  </div>
                </div>
              </DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuGroup>
                <DropdownMenuItem>
                  <BadgeCheck />
                  Account
                </DropdownMenuItem>
              </DropdownMenuGroup>
              <DropdownMenuSeparator />
              <DropdownMenuItem @select="authStore.signOutUser">
                <LogOut />
                Signout
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </SidebarMenuItem>
      </SidebarMenu>
    </SidebarFooter>
    <SidebarRail />
  </Sidebar>
</template>
