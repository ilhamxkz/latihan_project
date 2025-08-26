<script setup>
import VerticalNavSectionTitle from '@/@layouts/components/VerticalNavSectionTitle.vue'
import VerticalNavLink from '@layouts/components/VerticalNavLink.vue'
</script>

<template>
  <!-- 👉 Apps & Pages -->
  <VerticalNavSectionTitle
    :item="{
      heading: 'Apps & Pages',
    }"
  />

  <VerticalNavLink
    :item="{
      title: 'Users table',
      icon: 'ri-user-line',
      to: '/account-settings', //belum diseting
    }"
  />

  <VerticalNavLink
    :item="{
      title: 'Roles table',
      icon: 'role-binding-outlined',
      to: '/register', //belum diseting
    }"
  />
  <VerticalNavLink
    :item="{
      title: 'Akses table',
      icon: 'ri-information-line',
      to: '/no-existence', //belum diseting
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
