<script setup>
import { ref, computed } from 'vue'
import { useTheme } from 'vuetify'
import { useRouter } from 'vue-router'
import AuthProvider from '@/views/pages/authentication/AuthProvider.vue'
import authV1MaskDark from '@images/pages/auth-v1-mask-dark.png'
import authV1MaskLight from '@images/pages/auth-v1-mask-light.png'
import authV1Tree2 from '@images/pages/auth-v1-tree-2.png'
import authV1Tree from '@images/pages/auth-v1-tree.png'
import axios from 'axios'
import logodst from '@/assets/images/logos/logodst.png'

const router = useRouter()
const form = ref({
  username: '',
  password: '',
  remember: false,
})

const loading = ref(false)
const vuetifyTheme = useTheme()
const isPasswordVisible = ref(false)

const authThemeMask = computed(() => {
  return vuetifyTheme.global.name.value === 'light' ? authV1MaskLight : authV1MaskDark
})

// Ubah URL ini ke endpoint login backend kamu
const API_URL = 'http://localhost:8080/api/login'

const login = async () => {
  if (!form.value.username || !form.value.password) {
    window.alert('Username dan password harus diisi.')
    return
  }

  loading.value = true
  try {
    const response = await axios.post(API_URL, {
      username: form.value.username,
      password: form.value.password,
    })

// setelah response dari axios.post login
const token = response.data.token

// simpan sesuai pilihan remember
if (form.value.remember) {
  localStorage.setItem('token', token)
    sessionStorage.setItem('token', token)

} else {
    localStorage.setItem('token', token)

  sessionStorage.setItem('token', token)
}

// set default header untuk semua request axios
axios.defaults.headers.common['Authorization'] = `Bearer ${token}`
    // tampilkan pesan sukses
    window.alert('Login berhasil! Mengarahkan ke dashboard...')

    // redirect ke dashboard
    router.push('/dashboard')
  } catch (error) {
    const msg = error.response?.data?.error || error.message || 'Terjadi kesalahan'
    window.alert('Login gagal: ' + msg)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="auth-wrapper d-flex align-center justify-center pa-4">
    <VCard class="auth-card pa-4 pt-7" max-width="448">
      <VCardItem class="justify-center">
        <img :src="logodst" alt="Logo" style="height: 40px;" />
      </VCardItem>

      <VCardText class="pt-2">
        <h4 class="text-h4 mb-1">selamat datang di halaman login! 👋🏻</h4>
        <p class="mb-0">silahkan login untuk masuk ke aplikasi</p>
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

              <div class="d-flex align-center justify-space-between flex-wrap my-6">
                <VCheckbox v-model="form.remember" label="Remember me" />
                <a class="text-primary" href="javascript:void(0)">Forgot Password?</a>
              </div>

              <!-- tombol disable saat loading -->
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

    <VImg class="auth-footer-start-tree d-none d-md-block" :src="authV1Tree" :width="250" />
    <VImg :src="authV1Tree2" class="auth-footer-end-tree d-none d-md-block" :width="350" />
    <VImg class="auth-footer-mask d-none d-md-block" :src="authThemeMask" />
  </div>
</template>

<style lang="scss">
@use "@core/scss/template/pages/page-auth";
</style>
