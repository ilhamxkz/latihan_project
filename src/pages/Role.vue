<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

// DataTables + Bootstrap
import DataTable from 'datatables.net-vue3'
import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'

// Aktifkan DataTables core + buttons
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

const roles = ref([])
const loading = ref(true)
const error = ref(null)

const API_ROLES = 'http://localhost:8000/api/roles'

// Ambil data roles
const fetchRoles = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await axios.get(API_ROLES)

    // cek response backend
    if (Array.isArray(res.data)) {
      roles.value = res.data
    } else if (Array.isArray(res.data.data)) {
      roles.value = res.data.data
    } else {
      roles.value = []
    }
  } catch (err) {
    error.value = err.response?.data?.message || err.message || 'Gagal ambil data'
  } finally {
    loading.value = false
  }
}

onMounted(fetchRoles)
</script>

<template>
  <VRow>
    <VCol cols="12">
      <VCard>
        <!-- Header Card -->
        <VCardTitle class="d-flex align-center justify-space-between">
          <span class="font-weight-bold">Data Roles</span>
        </VCardTitle>

        <VCardText>
          <div v-if="loading">Loading...</div>
          <div v-else-if="error">Error: {{ error }}</div>

          <!-- DataTable -->
          <div v-else>
            <DataTable
              class="table table-hover table-bordered align-middle text-center w-100"
              :data="roles"
              :columns="[ 
                { title: 'ID', data: 'id' },
                { title: 'Role Name', data: 'nama_role' }
              ]"
              :options="{
                dom: 'Bfrtip',
                buttons: ['copy', 'csv', 'excel', 'pdf', 'print'],
                responsive: true,
                pageLength: 5,
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
/* Rapihin tabel */
table.dataTable {
  border-collapse: collapse !important;
  width: 100% !important;
}

table.dataTable th {
  background-color: #f5f5f5;
  font-weight: 600;
  text-align: center;
}

table.dataTable td {
  vertical-align: middle;
  text-align: center;
}

table.dataTable tbody tr:hover {
  background-color: #f9f9f9;
}
</style>
