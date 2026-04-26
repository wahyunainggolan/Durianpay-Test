<template>
  <div class="dashboard">
    <div class="dashboard-header">
      <h2 class="title">Dashboard</h2>
      <button class="logout-btn" @click="logout">
        Logout
      </button>
    </div>
    <SummaryCard :summary="summary" />
    <div class="toolbar">
      <select v-model="status" @change="onFilterChange" class="select">
        <option value="">All Status</option>
        <option value="completed">Completed</option>
        <option value="processing">Processing</option>
        <option value="failed">Failed</option>
      </select>
      <div class="actions">
        <button class="refresh-btn" @click="fetchPayments">
          Refresh
        </button>
      </div>
    </div>
    <div v-if="loading" class="state">
      Loading payments...
    </div>
    <div v-else-if="error" class="error">
      {{ error }}
    </div>

    <div v-else>
      <BaseTable
        :columns="columns"
        :rows="payments"
        :loading="loading"
        :page="page"
        @sort="handleSort"
        @next="nextPage"
        @prev="prevPage"
      >
        <template #amount="{ value }">
          {{ Formatter.currency(value) }}
        </template>

        <template #status="{ value }">
          <span :class="['badge', value]">
            {{ value }}
          </span>
        </template>
        <template #created_at="{ value }">
          {{ Formatter.date(value) }}
        </template>
      </BaseTable>
    </div>
  </div>
</template>

<script setup>
import "../styles/pages/dashboard.css"
import { ref, onMounted } from "vue"
import { getPayments } from "../services/paymentService"
import BaseTable from "../components/BaseTable.vue"
import SummaryCard from "../components/SummaryCard.vue"
import { Formatter } from "../utils/formatter"
import { useRouter } from "vue-router"

const payments = ref([])
const summary = ref({
  total: 0,
  completed: 0,
  processing: 0,
  failed: 0
})
const router = useRouter()
const status = ref("")
const loading = ref(false)
const error = ref(null)
const page = ref(1)
const sort = ref({
  key: "",
  order: ""
})

const columns = [
  { key: "id", label: "ID", sortable: true },
  { key: "merchant", label: "Merchant" },
  { key: "amount", label: "Amount" },
  { key: "status", label: "Status" },
  { key: "created_at", label: "Date", sortable: true }
]

onMounted(() => {
  fetchPayments()
})

const fetchPayments = async () => {
  try {
    loading.value = true
    error.value = null
    const res = await getPayments(buildQuery())
    console.log(res)
    
    const data = res.data.data
    payments.value = data.payments ?? []
    summary.value = data.summary ?? fallbackSummary()
  } catch (err) {
    error.value = "Failed to load payments"
  } finally {
    loading.value = false
  }
}

const logout = () => {
  localStorage.removeItem("token")
  localStorage.removeItem("role")
  router.push("/login")
}

const buildQuery = () => {
  const query = {
    page: page.value,
  }

  if (status.value) {
    query.status = status.value
  }

  if (sort.value.key) {
    query.sort = `${sort.value.order === "desc" ? "-" : ""}${sort.value.key}`
  }

  return query
}

const onFilterChange = () => {
  page.value = 1
  fetchPayments()
}

const handleSort = ({ key, order }) => {
  sort.value = { key, order }
  fetchPayments()
}

const nextPage = () => {
  page.value++
  fetchPayments()
}

const prevPage = () => {
  if (page.value > 1) {
    page.value--
    fetchPayments()
  }
}

const fallbackSummary = () => ({
  total: 0,
  completed: 0,
  processing: 0,
  failed: 0
})
</script>