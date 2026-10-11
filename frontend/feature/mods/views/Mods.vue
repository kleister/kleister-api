<script setup lang="ts">
import ModsList from '../components/ModsList'
import ModsProvider from '../providers'
</script>

<template>
  <ModsProvider>
    <ModsList />
  </ModsProvider>
</template>
