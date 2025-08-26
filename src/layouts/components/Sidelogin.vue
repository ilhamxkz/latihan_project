<script setup>
import VerticalNavLink from '@layouts/components/VerticalNavLink.vue'
</script>
<template>
    <VerticalNavLink
    :item="{
      title: 'Dashbord',
      icon: 'ri-dashboard-line',
      to: '/dashboard',
    }"
  />
    <VerticalNavLink
    :item="{
      title: 'Login',
      icon: 'ri-login-box-line',
      to: '/login',
    }"
  />
</template>
