<script setup lang="ts">
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from '@/components/ui/breadcrumb'
import { Separator } from '@/components/ui/separator'
import { NavigationSidebar } from '@/components/common/navigation-sidebar'
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from '@/components/ui/sidebar'
import { RouterView } from 'vue-router'
import { useAuthStore } from './feature/auth/store/auth'
import { Toaster } from './components/ui/toast'
import { storeToRefs } from 'pinia'
import { useBreadcrumbs } from './composables/useBreadcrumbs'

const authStore = useAuthStore()
const { isAuthed } = storeToRefs(authStore)
const { breadcrumbs } = useBreadcrumbs()
</script>

<template>
  <SidebarProvider>
    <NavigationSidebar v-if="isAuthed" />

    <SidebarInset>
      <header
        v-if="isAuthed"
        class="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-12"
      >
        <div class="flex items-center gap-2 px-4">
          <SidebarTrigger class="-ml-1" />
          <Separator orientation="vertical" class="mr-2 h-4" />
          <Breadcrumb>
            <BreadcrumbList>
              <template
                v-for="(breadcrumb, index) in breadcrumbs"
                :key="breadcrumb.url"
              >
                <BreadcrumbItem>
                  <BreadcrumbPage
                    v-if="index === breadcrumbs.length - 1"
                    :href="breadcrumb.url"
                  >
                    {{ breadcrumb.name }}
                  </BreadcrumbPage>

                  <BreadcrumbLink v-else :to="breadcrumb.url">
                    {{ breadcrumb.name }}
                  </BreadcrumbLink>
                </BreadcrumbItem>

                <BreadcrumbSeparator v-if="index !== breadcrumbs.length - 1" />
              </template>
            </BreadcrumbList>
          </Breadcrumb>
        </div>
      </header>

      <main class="px-4 pt-2 pb-4">
        <RouterView />
      </main>
    </SidebarInset>
  </SidebarProvider>

  <Toaster />
</template>
