<script setup>
import { ref, onMounted, reactive } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

// CSS
import 'bootstrap/dist/css/bootstrap.min.css'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'

// DataTables
import DataTable from 'datatables.net-vue3'
import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

const router = useRouter()

// state
const users = ref([])
const loading = ref(true)
const error = ref(null)

// axios instance
const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})
// tambahan: interceptor response untuk logging & mudah debugging
api.interceptors.response.use(res => {
  console.log('[API RES]', res.config.method?.toUpperCase(), res.config.url, res.status, res.data)
  return res
}, err => {
  console.warn('[API ERR]', err.config?.method?.toUpperCase(), err.config?.url, err.response?.status, err.response?.data, err.message)
  return Promise.reject(err)
})

// helper error
const handleError = (err) => {
  console.error('handleError ->', err)
  const status = err?.response?.status
  if (status === 401) {
    error.value = 'Silahkan login terlebih dahulu'
  } else if (err?.response?.data) {
    // banyak backend mengirim detail error di res.data
    // tampilkan pesan ringkas
    error.value = err.response.data.message || JSON.stringify(err.response.data) || err.message
  } else {
    error.value = err?.message || 'Gagal ambil data'
  }
}

// Columns sebagai variable — mencegah DataTables meminta properti yang tidak ada
const userColumns = [
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
]

// fetch
const normalizeUser = (u) => ({
  id: u.id ?? u.ID ?? u.user_id ?? (u.user && u.user.id) ?? null,
  username: u.username ?? u.Username ?? u.name ?? u.nama ?? '',
  email: u.email ?? u.Email ?? u.email_address ?? ''
})

const fetchUsers = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/users')
    // console.log supaya kamu bisa lihat bentuk response
    console.log('raw /users response:', res.data)

    let arr = []
    if (Array.isArray(res.data)) arr = res.data
    else if (Array.isArray(res.data.users)) arr = res.data.users
    else if (Array.isArray(res.data.data)) arr = res.data.data
    else if (Array.isArray(res.data.result)) arr = res.data.result
    else arr = []

    // normalisasi tiap item -> pastikan ada property id, username, email
    users.value = arr.map(normalizeUser)
  } catch (err) {
    handleError(err)
  } finally {
    loading.value = false
  }
}

/* ----------------------------
   Modal simple untuk Create / Edit
   ---------------------------- */
const showModal = ref(false)
const modalMode = ref('create') // 'create' | 'edit'
const form = reactive({
  id: null,
  username: '',
  email: '',
  password: '' // optional for update
})

const openCreate = () => {
  modalMode.value = 'create'
  form.id = null
  form.username = ''
  form.email = ''
  form.password = ''
  showModal.value = true
}

const openEdit = (row) => {
  modalMode.value = 'edit'
  form.id = row.id
  form.username = row.username || ''
  form.email = row.email || ''
  form.password = ''
  showModal.value = true
}

const submit = async () => {
  error.value = null
  try {
    if (!form.username || !form.email || (modalMode.value === 'create' && !form.password)) {
      alert('Harap isi username, email, dan password (untuk create).')
      return
    }

    if (modalMode.value === 'create') {
      console.log('[CREATE] payload:', { username: form.username, email: form.email, password: form.password })
      const res = await api.post('/users', {
        username: form.username,
        email: form.email,
        password: form.password
      })
      console.log('[CREATE] server response', res.data)
    } else {
      const payload = { username: form.username, email: form.email }
      if (form.password) payload.password = form.password
      console.log('[UPDATE] id', form.id, 'payload:', payload)
      const res = await api.put(`/users/${form.id}`, payload)
      console.log('[UPDATE] server response', res.data)
    }
    showModal.value = false
    await fetchUsers()
  } catch (err) {
    // perlihatkan detail error agar mudah debugging
    console.error('submit err ->', err)
    if (err?.response) {
      alert('Error: ' + (err.response.data?.message || JSON.stringify(err.response.data)))
    } else {
      alert('Error: ' + err.message)
    }
    handleError(err)
  }
}

const deleteUser = async (id) => {
  if (!confirm('Yakin ingin hapus user ini?')) return
  try {
    console.log('[DELETE] id', id)
    // pastikan tidak mengirim body pada delete — axios.delete(url) default tanpa body
    const res = await api.delete(`/users/${id}`)
    console.log('[DELETE] server response', res.data)
    await fetchUsers()
  } catch (err) {
    console.error('delete err ->', err)
    if (err?.response) alert('Gagal hapus: ' + (err.response.data?.message || JSON.stringify(err.response.data)))
    else alert('Gagal hapus: ' + err.message)
    handleError(err)
  }

  try {
  // ... request create/update/delete
} catch (err) {
  console.error('AXIOS ERROR', err)
  if (err.response) {
    // info berguna: status, headers, body
    console.log('response.status =', err.response.status)
    console.log('response.data =', err.response.data)
    console.log('response.headers =', err.response.headers)
    alert('Server error: ' + (err.response.data?.message || JSON.stringify(err.response.data)))
  } else if (err.request) {
    console.log('no response received, request =', err.request)
    alert('No response dari server')
  } else {
    console.log('error message =', err.message)
    alert('Error: ' + err.message)
  }
}

}

onMounted(fetchUsers)
</script>

<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between mb-3">
      <h3>Data User</h3>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah User</button>
        <button class="btn btn-secondary" @click="fetchUsers">Refresh</button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-4">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Sedang memuat data...</p>
    </div>

    <div v-else-if="error" class="alert alert-danger text-center">
      {{ error }}
      <div v-if="error && error.toLowerCase().includes('login')" class="mt-2">
        <button class="btn btn-sm btn-outline-light" @click="router.push('/login')">Login Sekarang</button>
      </div>
    </div>

    <div v-else>
      <DataTable
        class="table table-striped table-hover table-bordered align-middle text-center w-100"
        :data="users"
        :columns="userColumns"
        :options="{
          dom: 'lBfrtip',
          buttons: ['copy', 'csv', 'excel', 'pdf', 'print'],
          pageLength: 10,
          lengthMenu: [[10, 25, 50], [10, 25, 50]],
          responsive: true,
          createdRow: function (row, data) {
            const editBtn = row.querySelector('.edit-btn')
            const delBtn = row.querySelector('.delete-btn')
            if (editBtn) {
              editBtn.addEventListener('click', () => { openEdit(data) })
            }
            if (delBtn) {
              delBtn.addEventListener('click', () => { deleteUser(data.id) })
            }
          }
        }"
      />
    </div>

    <!-- SIMPLE MODAL -->
    <div v-if="showModal" class="modal-backdrop" style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1050;">
      <div class="card p-3" style="width:420px;">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h5 class="m-0">{{ modalMode === 'create' ? 'Tambah User' : 'Edit User' }}</h5>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal = false">×</button>
        </div>

        <div class="mb-2">
          <label class="form-label">Username</label>
          <input class="form-control" v-model="form.username" />
        </div>
        <div class="mb-2">
          <label class="form-label">Email</label>
          <input class="form-control" v-model="form.email" />
        </div>
        <div class="mb-3">
          <label class="form-label">Password <small v-if="modalMode === 'edit'">(kosongkan jika tidak ganti)</small></label>
          <input type="password" class="form-control" v-model="form.password" />
        </div>

        <div class="d-flex justify-content-end">
          <button class="btn btn-secondary me-2" @click="showModal = false">Batal</button>
          <button class="btn btn-primary" @click="submit">{{ modalMode === 'create' ? 'Simpan' : 'Update' }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
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
/* modal-backdrop slight dim */
.modal-backdrop { background: rgba(0,0,0,0.35); }
</style>
