<script setup lang="ts">
import UsersList from '../components/UsersList'
import UsersProvider from '../providers'
</script>

<template>
  <UsersProvider>
    <UsersList />
  </UsersProvider>
</template>
