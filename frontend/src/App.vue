<script setup>
import Navbar from "@/views/layout/Navbar.vue";
import '@/index.css'
</script>

<template>
  <Navbar v-if="$route.meta.showNav" />

  <router-view></router-view>

</template>

<style scoped>

</style>
