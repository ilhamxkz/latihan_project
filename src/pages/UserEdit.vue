<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter, useRoute } from 'vue-router'

const API_URL = 'http://localhost:8080/api/users'
const router = useRouter()
const route = useRoute()

const formUser = ref({ id: null, username: '', email: '' })

// Ambil data user berdasarkan ID
const fetchUser = async () => {
  try {
    const res = await axios.get(`${API_URL}/${route.params.id}`)
    formUser.value = res.data.data
  } catch (err) {
    alert('Gagal ambil data user: ' + (err.response?.data?.message || err.message))
    router.push('/users')
  }
}

const updateUser = async () => {
  try {
    await axios.put(`${API_URL}/${formUser.value.id}`, formUser.value)
    alert('User berhasil diupdate!')
    router.push('/users')
  } catch (err) {
    alert('Gagal update user: ' + (err.response?.data?.message || err.message))
  }
}

onMounted(fetchUser)
</script>

<template>
  <div class="container mt-4">
    <h3>Edit User</h3>
    <form @submit.prevent="updateUser" class="mt-3">
      <div class="mb-3">
        <label class="form-label">Username</label>
        <input v-model="formUser.username" type="text" class="form-control" required />
      </div>
      <div class="mb-3">
        <label class="form-label">Email</label>
        <input v-model="formUser.email" type="email" class="form-control" required />
      </div>
      <button type="submit" class="btn btn-primary">Update</button>
      <button type="button" class="btn btn-secondary ms-2" @click="router.push('/users')">Batal</button>
    </form>
  </div>
</template>
