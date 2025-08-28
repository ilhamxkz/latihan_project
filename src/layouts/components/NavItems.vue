<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

import VerticalNavSectionTitle from '@/@layouts/components/VerticalNavSectionTitle.vue'
import VerticalNavLink from '@layouts/components/VerticalNavLink.vue'

const menus = ref([])
const loading = ref(true)
const error = ref(null)

const iconMap = {
  dashboard: 'ri-dashboard-line',
  users: 'ri-user-line',
  user: 'ri-user-line',
  roles: 'ri-id-card-line',
  akses: 'ri-accessibility-line',
  login: 'ri-login-box-line'
}

function mapMenuItem(m) {
  const rawName = (m.nama_menu || '').toString()
  const title = rawName
    ? rawName.charAt(0).toUpperCase() + rawName.slice(1)
    : 'Untitled'
  const key = rawName.toLowerCase()
  const icon = iconMap[key] || 'ri-menu-line'
  return {
    id: m.id,
    title,
    to: m.routes || '#',
    icon
  }
}

onMounted(async () => {
  loading.value = true
  error.value = null

  try {
    const token = localStorage.getItem('token') || localStorage.getItem('authToken') || null

    const res = await axios.get('http://localhost:8080/api/menus', {
      headers: token ? { Authorization: `Bearer ${token}` } : {}
    })

    const dataMenus = res && res.data && Array.isArray(res.data.menus) ? res.data.menus : []
    menus.value = dataMenus.map(mapMenuItem)

  } catch (err) {
    console.error('Failed to load menus', err)
    if (err.response && err.response.status === 401) {
      error.value = 'Untuk mengakses silahkan login terlebih dahulu'
      menus.value = []
    } else {
      error.value = err.response?.data?.message || err.message || 'Gagal mengambil menu'
    }
  } finally {
    loading.value = false
  }
})
</script>

<template>

  <VerticalNavSectionTitle :item="{ heading: 'Apps & Pages' }" />

  <div v-if="loading" class="message loading">Memuat menu...</div>

  <div v-else-if="error" class="message error">Error: {{ error }}</div>

  <template v-else>
    <VerticalNavLink
      :item="{
        title: 'User',
        to: '/dashboard',
      }"
    />

    <VerticalNavLink
      v-for="menu in menus"
      :key="menu.id"
      :item="{ title: menu.title, icon: menu.icon, to: menu.to }"
    />
  </template>

</template>

<style scoped>
/* Pesan umum (loading / error) */
.message {
  padding-inline-start: 1rem;
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
  font-size: 0.875rem; /* setara text-sm */
  line-height: 1.25rem;
  margin: 0 0 0.25rem 0;
}

/* Tampilan saat memuat */
.loading {
  color: #6b7280; /* abu/ muted, kira-kira Tailwind gray-500 */
}

/* Tampilan error */
.error {
  color: #dc2626; /* merah, mirip Tailwind red-600 */
  font-weight: 500;
}

/* Optional: spacing pada list menu agar konsisten */
:deep(.vertical-nav-list) {
  margin-top: 0.5rem;
}

/* Jika VerticalNavLink internal menggunakan class tertentu dan perlu override,
   tambahkan selector di sini, misalnya:
   :deep(.vertical-nav-link) { padding: 0.5rem 0; }
*/
</style>
