<script setup>
import VerticalNavLink from '@layouts/components/VerticalNavLink.vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()

const handleLogout = async () => {
  try {
    const token = localStorage.getItem('token') || sessionStorage.getItem('token')

    await axios.post(
      'http://localhost:8080/api/logout',
      {},
      { headers: { Authorization: `Bearer ${token}` } }
    )

    // Hapus token di localStorage/sessionStorage
    localStorage.removeItem('token')
    sessionStorage.removeItem('token')

    alert('Logout berhasil!')
    router.push('/login')
  } catch (error) {
    console.error('Logout gagal:', error)
    alert('Logout gagal, silakan coba lagi')
  }
}
</script>
<template>
    <VerticalNavLink
      :item="{
        title: 'Logout',
        icon: 'ri-logout-box-line',
        to: '/',
      }"
      @click="handleLogout"
    />
</template>
