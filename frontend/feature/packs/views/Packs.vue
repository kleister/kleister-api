<script setup lang="ts">
import PacksList from '../components/PacksList'
import PacksProvider from '../providers'
</script>

<template>
  <PacksProvider>
    <PacksList />
  </PacksProvider>
</template>
