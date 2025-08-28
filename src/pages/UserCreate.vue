<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const API_URL = 'http://localhost:8080/api/users'
const router = useRouter()

// pastikan field sesuai dengan API backend (ganti id_role -> role_id kalau perlu)
const formUser = ref({ 
  name: '', 
  username: '', 
  password: '', 
  email: '', 
  role_id: ''   // biasanya pakai role_id
})

const createUser = async () => {
  try {
    await axios.post(API_URL, formUser.value)
    alert('User berhasil ditambahkan!')
    router.push('/dashboard') // balik ke list user biar kelihatan datanya
  } catch (err) {
    console.error(err)
    alert('Gagal tambah user: ' + (err.response?.data?.message || err.message))
  }
}
</script>

<template>
  <div class="container mt-4">
    <h3>Tambah User</h3>
    <form @submit.prevent="createUser" class="mt-3">
      <div class="mb-3">
        <label class="form-label">Name</label>
        <input v-model="formUser.name" type="text" class="form-control" required />
      </div>
      <div class="mb-3">
        <label class="form-label">Username</label>
        <input v-model="formUser.username" type="text" class="form-control" required />
      </div>
      <div class="mb-3">
        <label class="form-label">Password</label>
        <input v-model="formUser.password" type="password" class="form-control" required />
      </div>
      <div class="mb-3">
        <label class="form-label">Email</label>
        <input v-model="formUser.email" type="email" class="form-control" required />
      </div>
      <div class="mb-3">
        <label class="form-label">Role ID</label>
        <input v-model.number="formUser.role_id" type="number" class="form-control" required />
      </div>
      <button type="submit" class="btn btn-success">Simpan</button>
      <button type="button" class="btn btn-secondary ms-2" @click="router.push('/dashboard')">Batal</button>
    </form>
  </div>
</template>
