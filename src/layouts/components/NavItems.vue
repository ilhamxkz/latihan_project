<script setup>
import VerticalNavSectionTitle from '@/@layouts/components/VerticalNavSectionTitle.vue'
import VerticalNavGroup from '@layouts/components/VerticalNavGroup.vue'
import VerticalNavLink from '@layouts/components/VerticalNavLink.vue'
</script>

<template>
  <!-- 👉 Dashboards -->
  <VerticalNavGroup
    :item="{
      title: 'Dashboards',
      badgeContent: '5',
      badgeClass: 'bg-error',
      icon: 'ri-home-smile-line',
    }"
  >
    <VerticalNavLink
      :item="{
        title: 'User',
        to: '/dashboard',
      }"
    />
  </VerticalNavGroup>


  <!-- 👉 Apps & Pages -->
  <VerticalNavSectionTitle
    :item="{
      heading: 'Apps & Pages',
    }"
  />

  <VerticalNavLink
    :item="{
      title: 'Role',
      icon: 'ri-user-settings-line',
      to: '/account-settings',
    }"
  />

  <VerticalNavLink
    :item="{
      title: 'Akses',
      icon: 'ri-user-settings-line',
      to: '/typography',
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
