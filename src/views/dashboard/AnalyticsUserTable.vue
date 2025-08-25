<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'


import avatar1 from '@images/avatars/avatar-1.png'
import avatar2 from '@images/avatars/avatar-2.png'
import avatar3 from '@images/avatars/avatar-3.png'
import avatar4 from '@images/avatars/avatar-4.png'

const headers = [
  { title: 'User', key: 'username' },
  { title: 'Email', key: 'email' },
  { title: 'Role', key: 'role' },
  { title: 'Status', key: 'status' },
]

const userData = ref([])


// Fungsi ambil data API
// AnalyticsUserTable.vue (ubah fetchUsers)
const fetchUsers = async () => {
  const token = sessionStorage.getItem('token') || localStorage.getItem('token')
  console.log('DEBUG token from storage:', token)

  if (!token) {
    console.warn('Tidak ada token di storage — silakan login dulu.')
    return
  }

  try {
    const res = await axios.get('http://localhost:8080/api/users', {
      headers: { Authorization: `Bearer ${token}` }
    })

    console.log('DEBUG res.data:', res.data)
    // cari array
    let usersArray = []
    if (Array.isArray(res.data)) usersArray = res.data
    else if (res.data && Array.isArray(res.data.data)) usersArray = res.data.data
    else if (res.data && Array.isArray(res.data.users)) usersArray = res.data.users
    else if (res.data && Array.isArray(res.data.result)) usersArray = res.data.result
    else usersArray = [res.data] // fallback

    userData.value = usersArray.map((u, i) => ({
      id: u.ID || u.id || i,
      fullName: u.username || u.name || 'No Name',
      username: u.username || u.email || `user${i+1}`,
      email: u.email || '',
      role: u.role || 'subscriber',
      status: u.status || 'active',
      avatar: [avatar1, avatar2, avatar3, avatar4][i % 4],
    }))
  } catch (err) {
    console.error('Gagal ambil data users:', err)
    if (err.response?.status === 401) {
      console.warn('Server balikan 401 — token mungkin tidak valid/tidak cocok dengan DB.')
    }
  }
}



onMounted(() => {
  const token = sessionStorage.getItem('token') || localStorage.getItem('token')
  if (!token) {
    console.log('DEBUG: belum login — fetchUsers dibatalkan')
    return
  }
  fetchUsers()
})


const resolveUserRoleVariant = role => {
  const roleLowerCase = role.toLowerCase()
  if (roleLowerCase === 'subscriber')
    return { color: 'success', icon: 'ri-user-line' }
  if (roleLowerCase === 'author')
    return { color: 'error', icon: 'ri-computer-line' }
  if (roleLowerCase === 'maintainer')
    return { color: 'info', icon: 'ri-pie-chart-line' }
  if (roleLowerCase === 'editor')
    return { color: 'warning', icon: 'ri-edit-box-line' }
  if (roleLowerCase === 'admin')
    return { color: 'primary', icon: 'ri-vip-crown-line' }
  return { color: 'success', icon: 'ri-user-line' }
}

const resolveUserStatusVariant = stat => {
  const statLowerCase = stat.toLowerCase()
  if (statLowerCase === 'pending') return 'warning'
  if (statLowerCase === 'active') return 'success'
  if (statLowerCase === 'inactive') return 'secondary'
  return 'primary'
}
</script>

<template>
  <VCard>
    <VDataTable
      :headers="headers"
      :items="userData"
      item-value="id"
      class="text-no-wrap"
    >
      <!-- User -->
      <template #item.username="{ item }">
        <div class="d-flex align-center gap-x-4">
          <VAvatar
            size="34"
            :variant="!item.avatar ? 'tonal' : undefined"
            :color="!item.avatar ? resolveUserRoleVariant(item.role).color : undefined"
          >
            <VImg v-if="item.avatar" :src="item.avatar" />
          </VAvatar>

          <div class="d-flex flex-column">
            <h6 class="text-h6 font-weight-medium user-list-name">
              {{ item.fullName }}
            </h6>
            <span class="text-sm text-medium-emphasis">@{{ item.username }}</span>
          </div>
        </div>
      </template>

      <!-- Role -->
      <template #item.role="{ item }">
        <div class="d-flex gap-4">
          <VIcon
            :icon="resolveUserRoleVariant(item.role).icon"
            :color="resolveUserRoleVariant(item.role).color"
            size="22"
          />
          <div class="text-capitalize text-high-emphasis">
            {{ item.role }}
          </div>
        </div>
      </template>

      <!-- Status -->
      <template #item.status="{ item }">
        <VChip
          :color="resolveUserStatusVariant(item.status)"
          size="small"
          class="text-capitalize"
        >
          {{ item.status }}
        </VChip>
      </template>

      <template #bottom />
    </VDataTable>
  </VCard>
</template>
