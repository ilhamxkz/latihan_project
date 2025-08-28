<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

import 'bootstrap/dist/css/bootstrap.min.css'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'

import DataTable from 'datatables.net-vue3'
import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'

DataTable.use(DataTablesCore)
DataTable.use(Buttons)

const users = ref([])
const loading = ref(true)
const error = ref(null)

const API_URL = 'http://localhost:8000/api/users'
const router = useRouter()

const fetchUsers = async () => {
  loading.value = true
  error.value = null
  try {
    const response = await axios.get(API_URL)
    users.value = Array.isArray(response.data.data) ? response.data.data : []
  } catch (err) {
    error.value = err.response?.data?.message || err.message || 'Gagal ambil data'
  } finally {
    loading.value = false
  }
}

const deleteUser = async (id) => {
  if (confirm('Yakin ingin hapus user ini?')) {
    try {
      await axios.delete(`${API_URL}/${id}`)
      await fetchUsers()
    } catch (err) {
      alert('Gagal hapus user: ' + (err.response?.data?.message || err.message))
    }
  }
}

onMounted(fetchUsers)
</script>

<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between mb-3">
      <h3>Data User</h3>
      <button class="btn btn-primary" @click="router.push('/users/add')">+ Tambah User</button>
    </div>

    <div v-if="loading" class="text-center py-4">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Sedang memuat data...</p>
    </div>

    <div v-else-if="error" class="alert alert-danger text-center">
      {{ error }}
    </div>

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
          dom: 'lBfrtip',
          buttons: ['copy', 'csv', 'excel', 'pdf', 'print'],
          pageLength: 10,
          lengthMenu: [[10, 25, 50], [10, 25, 50]],
          createdRow: (row, data) => {
            row.querySelector('.edit-btn').addEventListener('click', () => router.push(`/users/${data.id}/edit`))
            row.querySelector('.delete-btn').addEventListener('click', () => deleteUser(data.id))
          }
        }"
      />
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
</style>
