<script setup>
import axios from 'axios'
import 'bootstrap/dist/css/bootstrap.min.css'
import DataTablesCore from 'datatables.net-bs5'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import Buttons from 'datatables.net-buttons-bs5'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'
import DataTable from 'datatables.net-vue3'
import { onMounted, reactive, ref } from 'vue'
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

const roles = ref([])
const loading = ref(true)
const error = ref(null)
const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

const fetchRoles = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/roles')
    if (Array.isArray(res.data)) roles.value = res.data
    else if (Array.isArray(res.data.roles)) roles.value = res.data.roles
    else if (Array.isArray(res.data.data)) roles.value = res.data.data
    else roles.value = []
  } catch (err) {
    if (err?.response?.status === 401) error.value = 'Silahkan login terlebih dahulu'
    else error.value = err?.response?.data?.message || err.message
  } finally { loading.value = false }
}

/* modal simple */
const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({ id: null, nama_role: '' })

const openCreate = () => { modalMode.value = 'create'; form.id = null; form.nama_role = ''; showModal.value = true }
const openEdit = (row) => { modalMode.value = 'edit'; form.id = row.id; form.nama_role = row.nama_role || ''; showModal.value = true }

const submit = async () => {
  try {
    if (modalMode.value === 'create') {
      await api.post('/roles', { nama_role: form.nama_role })
    } else {
      await api.put(`/roles/${form.id}`, { nama_role: form.nama_role })
    }
    showModal.value = false
    await fetchRoles()
  } catch (err) {
    if (err?.response?.status === 401) error.value = 'Silahkan login terlebih dahulu'
    else alert(err?.response?.data?.message || err.message)
  }
}

const deleteRole = async (id) => {
  if (!confirm('Hapus role ini?')) return
  try {
    await api.delete(`/roles/${id}`)
    await fetchRoles()
  } catch (err) {
    if (err?.response?.status === 401) error.value = 'Silahkan login terlebih dahulu'
    else alert(err?.response?.data?.message || err.message)
  }
}

onMounted(fetchRoles)
</script>

<template>
  <div class="container mt-3">
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h4>Data Roles</h4>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Role</button>
        <button class="btn btn-secondary" @click="fetchRoles">Refresh</button>
      </div>
    </div>

    <div v-if="loading">Loading...</div>
    <div v-else-if="error" class="alert alert-danger">{{ error }}</div>
    <div v-else>
      <DataTable
        class="table table-hover table-bordered align-middle text-center w-100"
        :data="roles"
        :columns="[
          { title: 'ID', data: 'id' },
          { title: 'Role Name', data: 'nama_role' },
          { title: 'Aksi', data: null, render: (d,t,row)=>`<button class='btn btn-warning btn-sm me-1 edit-btn'>Edit</button><button class='btn btn-danger btn-sm delete-btn'>Hapus</button>` }
        ]"
        :options="{
                  dom: 'lBfrtip',       // 'l' = length menu (entries per page) default DataTables
          buttons: ['copy','csv','excel','pdf','print'],
          responsive: true,
          pageLength: 5,
          lengthMenu: [[5,10,25,50,100],[5,10,25,50,100]],
          createdRow: function(row,data) {
            const e = row.querySelector('.edit-btn'), del = row.querySelector('.delete-btn')
            if (e) e.addEventListener('click', ()=> openEdit(data))
            if (del) del.addEventListener('click', ()=> deleteRole(data.id))
          },

        }"
      />
    </div>

    <div v-if="showModal" class="modal-backdrop" style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1050;">
      <div class="card p-3" style="width:360px;">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h6 class="m-0">{{ modalMode==='create' ? 'Tambah Role' : 'Edit Role' }}</h6>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal=false">×</button>
        </div>
        <div class="mb-2">
          <label class="form-label">Nama Role</label>
          <input class="form-control" v-model="form.nama_role" />
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
table.dataTable th { background:#f5f5f5; font-weight:600;text-align:center }
table.dataTable td { text-align:center }
.modal-backdrop { background: rgba(0,0,0,0.35) }
</style>
