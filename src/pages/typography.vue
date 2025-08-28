<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

// Import CSS Bootstrap + DataTables agar rapi
import 'bootstrap/dist/css/bootstrap.min.css'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'

// DataTables
import DataTable from 'datatables.net-vue3'
import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'

// Aktifkan DataTables core + buttons
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

// === STATE ===
const users = ref([])
const loading = ref(true)
const error = ref(null)

const router = useRouter()

// === AXIOS INSTANCE (supaya selalu kirim token) ===
const api = axios.create({
  baseURL: 'http://localhost:8080/api',
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// (Opsional) Interceptor response global untuk 401
// api.interceptors.response.use(
//   r => r,
//   err => {
//     if (err.response?.status === 401) {
//       // Kalau mau redirect otomatis ke halaman login:
//       // router.push('/login')
//     }
//     return Promise.reject(err)
//   }
// )

// === AMBIL DATA DARI API /users ===
const fetchUsers = async () => {
  loading.value = true
  error.value = null

  try {
    const response = await api.get('users')
    console.log('Response API /users:', response.data)

    // cek apakah responsenya array langsung atau ada wrapper "data"
    if (Array.isArray(response.data)) {
      users.value = response.data
    } else if (Array.isArray(response.data.data)) {
      users.value = response.data.data
    } else {
      users.value = []
    }
  } catch (err) {
    // khusus 401 tampilkan pesan yang diminta
    if (err.response?.status === 401) {
      error.value = 'Silahkan login terlebih dahulu'
      // Jika ingin auto-redirect ke login, uncomment:
      // router.push('/login')
    } else {
      error.value = err.response?.data?.message || err.message || 'Gagal ambil data'
    }
  } finally {
    loading.value = false
  }
}

// === HAPUS USER ===
const deleteUser = async (id) => {
  if (!confirm('Yakin ingin hapus user ini?')) return

  try {
    await api.delete(`/users/${id}`)
    // refresh data setelah sukses hapus
    await fetchUsers()
  } catch (err) {
    if (err.response?.status === 401) {
      error.value = 'Silahkan login terlebih dahulu'
      // router.push('/login') // optional
    } else {
      alert('Gagal hapus user: ' + (err.response?.data?.message || err.message))
    }
  }
}

onMounted(fetchUsers)
</script>

<template>
  <VRow>
    <VCol cols="12">
      <VCard class="shadow-sm rounded-3">
        <!-- Header Card -->
        <VCardTitle class="d-flex align-center justify-space-between py-3 px-4 border-bottom">
          <span class="fw-bold fs-5">Data User</span>
          <div>
            <button class="btn btn-primary" @click="router.push('/users/add')">+ Tambah User</button>
          </div>
        </VCardTitle>

        <!-- Konten Card -->
        <VCardText class="p-4">
          <div v-if="loading" class="text-center py-4">
            <div class="spinner-border text-primary" role="status">
              <span class="visually-hidden">Loading...</span>
            </div>
            <p class="mt-2">Sedang memuat data...</p>
          </div>

          <div v-else-if="error" class="alert alert-danger text-center">
            {{ error }}
            <!-- contoh tombol cepat ke login jika error 401 -->
            <div v-if="error && error.toLowerCase().includes('login')" class="mt-2">
              <button class="btn btn-sm btn-outline-light" @click="router.push('/login')">Login Sekarang</button>
            </div>
          </div>

          <!-- DataTable -->
          <div v-else>
            <DataTable
              class="table table-striped table-hover table-bordered align-middle text-center w-100"
              :data="users"
              :columns="[
                { title: 'ID', data: 'id' },
                { title: 'Username', data: 'username' },
                { title: 'Email', data: 'email' },
                {
                  title: 'Aksi',
                  data: null,
                  render: (data, type, row) => {
                    return `
                      <button class='btn btn-warning btn-sm me-1 edit-btn'>Edit</button>
                      <button class='btn btn-danger btn-sm delete-btn'>Hapus</button>
                    `
                  }
                }
              ]"
              :options="{
                dom: 'Bfrtip',
                buttons: ['copy', 'csv', 'excel', 'pdf', 'print'],
                responsive: true,
                pageLength: 10,
                createdRow: (row, data) => {
                  // tambahkan event listener untuk tombol edit/hapus
                  const editBtn = row.querySelector('.edit-btn')
                  const delBtn = row.querySelector('.delete-btn')

                  if (editBtn) {
                    editBtn.addEventListener('click', () => {
                      router.push(`/users/${data.id}/edit`)
                    })
                  }
                  if (delBtn) {
                    delBtn.addEventListener('click', () => {
                      deleteUser(data.id)
                    })
                  }
                },
                language: {
                  search: 'Cari:',
                  lengthMenu: 'Tampilkan _MENU_ data',
                  info: 'Menampilkan _START_ - _END_ dari _TOTAL_ data',
                  paginate: { next: '›', previous: '‹' }
                }
              }"
            />
          </div>
        </VCardText>
      </VCard>
    </VCol>
  </VRow>
</template>

<style scoped>
/* Tambahan styling agar lebih rapi */
table.dataTable th {
  background-color: #f8f9fa;
  font-weight: 600;
  text-align: center;
}

table.dataTable td {
  vertical-align: middle;
  text-align: center;
}

table.dataTable tbody tr:hover {
  background-color: #f1f1f1;
  transition: background-color 0.2s;
}
</style>
