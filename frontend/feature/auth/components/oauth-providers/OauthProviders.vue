<script setup lang="ts">
import { onBeforeMount, ref } from 'vue'
import { listProviders, type Provider } from '../../../../client'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { library } from '@fortawesome/fontawesome-svg-core'
import {
  faGithub,
  faGoogle,
  faMicrosoft,
  faGit,
  faGitlab,
  faOpenid,
} from '@fortawesome/free-brands-svg-icons'
import { Button } from '@/components/ui/button'

library.add(faGit, faGithub, faGoogle, faMicrosoft, faGitlab, faOpenid)

const oauthProviders = ref<Provider[]>([])

async function getProviders() {
  try {
    const { data, error } = await listProviders()

    if (error) {
      // TODO: handle API error
      console.error(error)

      return
    }

    oauthProviders.value = data!.providers
  } catch (error) {
    // TODO: handle generic error
    console.error(error)
  }
}

function getIconName(i: string): string {
  const [_prefix, icon] = i.split(' ')

  return icon.replace('fa-', '')
}

function getProviderUrl(driver: string): string {
  return '/api/v1/auth/' + driver + '/request'
}

onBeforeMount(getProviders)
</script>

<template>
  <div class="grid grid-cols-5 gap-2">
    <Button
      v-for="provider in oauthProviders"
      :key="provider.driver"
      as="a"
      variant="outline"
      :title="provider.display"
      :href="getProviderUrl(provider.driver)"
    >
      <FontAwesomeIcon
        v-if="provider.icon"
        :icon="{
          prefix: 'fab',
          iconName: getIconName(provider.icon),
        }"
      />
    </Button>
  </div>
</template>
