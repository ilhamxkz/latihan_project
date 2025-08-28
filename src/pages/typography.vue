<script setup>
import { ref, reactive, onMounted } from 'vue'
import axios from 'axios'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'
import DataTable from 'datatables.net-vue3'
import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

const aksesList = ref([])
const roles = ref([])
const menus = ref([])
const loading = ref(true)
const error = ref(null)
const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

const fetchAll = async () => {
  loading.value = true
  error.value = null
  try {
    const [rAkses, rRoles, rMenus] = await Promise.all([
      api.get('/akses'),
      api.get('/roles'),
      api.get('/menus')
    ])
    aksesList.value = Array.isArray(rAkses.data.akses) ? rAkses.data.akses : (Array.isArray(rAkses.data) ? rAkses.data : [])
    roles.value = Array.isArray(rRoles.data.roles) ? rRoles.data.roles : (Array.isArray(rRoles.data)? rRoles.data : [])
    menus.value = Array.isArray(rMenus.data.menus) ? rMenus.data.menus : (Array.isArray(rMenus.data)? rMenus.data : [])
  } catch (err) {
    if (err?.response?.status === 401) error.value = 'Silahkan login terlebih dahulu'
    else error.value = err?.response?.data?.message || err.message
  } finally { loading.value = false }
}

/* modal */
const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({ id: null, role_id: null, menu_id: null })

const openCreate = () => { modalMode.value='create'; form.id=null; form.role_id=null; form.menu_id=null; showModal.value=true }
const openEdit = (row) => { modalMode.value='edit'; form.id=row.id; form.role_id=row.role_id; form.menu_id=row.menu_id; showModal.value=true }

const submit = async () => {
  try {
    if (modalMode.value === 'create') {
      await api.post('/akses', { role_id: form.role_id, menu_id: form.menu_id })
    } else {
      await api.put(`/akses/${form.id}`, { role_id: form.role_id, menu_id: form.menu_id })
    }
    showModal.value=false
    await fetchAll()
  } catch (err) {
    if (err?.response?.status === 401) error.value='Silahkan login terlebih dahulu'
    else alert(err?.response?.data?.message || err.message)
  }
}

const deleteAkses = async (id) => {
  if (!confirm('Hapus akses ini?')) return
  try {
    await api.delete(`/akses/${id}`)
    await fetchAll()
  } catch (err) {
    if (err?.response?.status === 401) error.value='Silahkan login terlebih dahulu'
    else alert(err?.response?.data?.message || err.message)
  }
}

onMounted(fetchAll)
</script>

<template>
  <div class="container mt-3">
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h4>Data Akses</h4>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Akses</button>
        <button class="btn btn-secondary" @click="fetchAll">Refresh</button>
      </div>
    </div>

    <div v-if="loading">Loading...</div>
    <div v-else-if="error" class="alert alert-danger">{{ error }}</div>
    <div v-else>
      <DataTable
        class="table table-hover table-bordered align-middle text-center w-100"
        :data="aksesList"
        :columns="[
          { title: 'ID', data: 'id' },
          { title: 'Role', data: row => (row.role?.nama_role || row.Role?.nama_role || row.role_id) },
          { title: 'Menu', data: row => (row.menu?.nama_menu || row.Menu?.nama_menu || row.menu_id) },
          { title: 'Aksi', data: null, render: (d,t,row)=>`<button class='btn btn-warning btn-sm me-1 edit-btn'>Edit</button><button class='btn btn-danger btn-sm delete-btn'>Hapus</button>` }
        ]"
        :options="{
          dom: 'Bfrtip',
          buttons: ['copy','csv','excel','pdf','print'],
          createdRow: function(row,data) {
            const e = row.querySelector('.edit-btn'), del = row.querySelector('.delete-btn')
            if (e) e.addEventListener('click', ()=> openEdit(data))
            if (del) del.addEventListener('click', ()=> deleteAkses(data.id))
          },
          pageLength: 10
        }"
      />
    </div>

    <!-- modal -->
    <div v-if="showModal" class="modal-backdrop" style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1050;">
      <div class="card p-3" style="width:420px;">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h6 class="m-0">{{ modalMode==='create' ? 'Tambah Akses' : 'Edit Akses' }}</h6>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal=false">×</button>
        </div>
        <div class="mb-2">
          <label class="form-label">Role</label>
          <select class="form-select" v-model="form.role_id">
            <option :value="null">-- pilih role --</option>
            <option v-for="r in roles" :key="r.id" :value="r.id">{{ r.nama_role || r.NamaRole || r.name || r.username }}</option>
          </select>
        </div>
        <div class="mb-3">
          <label class="form-label">Menu</label>
          <select class="form-select" v-model="form.menu_id">
            <option :value="null">-- pilih menu --</option>
            <option v-for="m in menus" :key="m.id" :value="m.id">{{ m.nama_menu || m.NamaMenu || m.routes }}</option>
          </select>
        </div>
        <div class="d-flex justify-content-end">
          <button class="btn btn-secondary me-2" @click="showModal=false">Batal</button>
          <button class="btn btn-primary" @click="submit">{{ modalMode==='create' ? 'Simpan' : 'Update' }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
table.dataTable th { background:#f5f5f5; font-weight:600; text-align:center }
.modal-backdrop { background: rgba(0,0,0,0.35) }
</style>
