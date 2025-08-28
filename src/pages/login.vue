<script setup>
import { ref, computed } from 'vue'
import { useTheme } from 'vuetify'
import { useRouter } from 'vue-router'
import AuthProvider from '@/views/pages/authentication/AuthProvider.vue'
import logo from '@images/dst.png'
import authV1MaskDark from '@images/pages/auth-v1-mask-dark.png'
import authV1MaskLight from '@images/pages/auth-v1-mask-light.png'
import authV1Tree2 from '@images/pages/auth-v1-tree-2.png'
import authV1Tree from '@images/pages/auth-v1-tree.png'
import axios from 'axios'

const router = useRouter()
const form = ref({
  username: '',
  password: '',
  remember: false,
})

const loading = ref(false)
const vuetifyTheme = useTheme()
const isPasswordVisible = ref(false)

// state untuk notif
const notif = ref({
  show: false,
  type: '',   // 'success' atau 'danger'
  message: ''
})

const authThemeMask = computed(() => {
  return vuetifyTheme.global.name.value === 'light' ? authV1MaskLight : authV1MaskDark
})

const API_URL = 'http://localhost:8080/api/login'

const login = async () => {
  if (!form.value.username || !form.value.password) {
    notif.value = { show: true, type: 'danger', message: 'Username dan password harus diisi.' }
    return
  }

  loading.value = true
  try {
    const response = await axios.post(API_URL, {
      username: form.value.username,
      password: form.value.password,
    })

    const token = response.data.token

    if (form.value.remember) {
      localStorage.setItem('token', token)
      sessionStorage.removeItem('token')
    } else {
      sessionStorage.setItem('token', token)
      localStorage.removeItem('token')
    }

    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`

    notif.value = { show: true, type: 'success', message: 'Login berhasil! Mengarahkan ke dashboard...' }

    setTimeout(() => {
      router.push('/dashboard')
    }, 1500)

  } catch (error) {
    const msg = error.response?.data?.error || error.message || 'Terjadi kesalahan'
    notif.value = { show: true, type: 'danger', message: 'Login gagal: ' + msg }
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <!-- Wrapper full height supaya card center -->
  <div class="auth-wrapper d-flex align-center justify-center pa-4" style="min-height: 100vh;">
    <VCard class="auth-card pa-4 pt-7" max-width="448">

      <!-- Logo -->
      <VCardItem class="justify-center">
        <RouterLink to="/" class="d-flex align-center gap-3">
          <img :src="logo" alt="Logo" style="width:300px; height:auto;" />
        </RouterLink>
      </VCardItem>

      <!-- Title -->
      <VCardText class="pt-2 text-center">
        <h4 class="text-h4 mb-1">Selamat datang di halaman login!</h4>
        <p class="mb-0">Silahkan login untuk masuk ke aplikasi</p>
      </VCardText>

      <VCardText>
        <VForm @submit.prevent="login">
          <VRow>
            <VCol cols="12">
              <VTextField v-model="form.username" label="Username" />
            </VCol>

            <VCol cols="12">
              <VTextField
                v-model="form.password"
                label="Password"
                placeholder="············"
                :type="isPasswordVisible ? 'text' : 'password'"
                autocomplete="password"
                :append-inner-icon="isPasswordVisible ? 'ri-eye-off-line' : 'ri-eye-line'"
                @click:append-inner="isPasswordVisible = !isPasswordVisible"
              />

              <!-- 🔔 Notif pindah ke bawah password -->
              <div v-if="notif.show" class="mt-2">
                <div :class="['alert', `alert-${notif.type}`, 'alert-dismissible', 'fade', 'show']" role="alert">
                  {{ notif.message }}
                  <button type="button" class="btn-close" @click="notif.show = false"></button>
                </div>
              </div>

              <div class="d-flex align-center justify-space-between flex-wrap my-6">
                <VCheckbox v-model="form.remember" label="Remember me" />
                <a class="text-primary" href="javascript:void(0)">Forgot Password?</a>
              </div>

              <VBtn :loading="loading" :disabled="loading" block type="submit">
                Login
              </VBtn>
            </VCol>

            <VCol cols="12" class="text-center text-base">
              <span>New on our platform?</span>
              <RouterLink class="text-primary ms-2" to="/register">Create an account</RouterLink>
            </VCol>

            <VCol cols="12" class="d-flex align-center">
              <VDivider />
              <span class="mx-4">or</span>
              <VDivider />
            </VCol>

            <VCol cols="12" class="text-center">
              <AuthProvider />
            </VCol>
          </VRow>
        </VForm>
      </VCardText>
    </VCard>

    <!-- Dekorasi -->
    
  </div>
</template>
