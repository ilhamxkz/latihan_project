<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const username = ref('')
const email = ref('')
const password = ref('')

const router = useRouter()
const API_URL = 'http://localhost:8000/api/users'

const saveUser = async () => {
  try {
    await axios.post(API_URL, {
      username: username.value,
      email: email.value,
      password: password.value,
    })
    router.push('/users')
  } catch (err) {
    alert('Gagal tambah user')
  }
}
</script>

<template>
  <VCard>
    <VCardTitle>Tambah User</VCardTitle>
    <VCardText>
      <VForm @submit.prevent="saveUser">
        <VTextField v-model="name" label="Name" required />
        <VTextField v-model="username" label="Username" required />
        <VTextField v-model="password" label="Password" type="password" required />
        <VTextField v-model="email" label="Email" required />
        <VBtn type="submit" color="primary">Simpan</VBtn>
        <VBtn @click="$router.push('/users')" color="secondary">Batal</VBtn>
      </VForm>
    </VCardText>
  </VCard>
</template>
