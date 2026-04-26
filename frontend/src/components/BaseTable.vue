<template>
  <div class="table-wrapper">
    <div v-if="loading">
      <table class="table">
        <tbody>
          <tr v-for="n in 5" :key="n">
            <td v-for="c in columns" :key="c.key">
              <div class="skeleton"></div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-else-if="!rows.length" class="empty">
      No data available
    </div>

    <table v-else class="table">
      <thead>
        <tr>
          <th
            v-for="col in columns"
            :key="col.key"
            @click="handleSort(col)"
            class="sortable"
          >
            {{ col.label }}
            <span v-if="sortKey === col.key">
              {{ sortOrder === 'asc' ? '↑' : '↓' }}
            </span>
          </th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="row in rows" :key="row.id">
          <td v-for="col in columns" :key="col.key">
            <slot
              :name="col.key"
              :value="row[col.key]"
              :row="row"
            >
              {{ row[col.key] }}
            </slot>
          </td>
        </tr>
      </tbody>
    </table>

    <div class="pagination">
      <button
        class="page-btn"
        @click="emit('prev')"
        :disabled="isFirstPage"
      >
        ← Prev
      </button>

      <div class="page-info">
        <span class="page-current">{{ page }}</span>
        <span class="page-separator">/</span>
        <span class="page-total">{{ totalPages }}</span>
      </div>

      <button
        class="page-btn primary"
        @click="emit('next')"
        :disabled="isLastPage"
      >
        Next →
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue"

const props = defineProps({
  columns: Array,
  rows: Array,
  loading: Boolean,
  page: Number,
  totalPages: {
    type: Number,
    default: 1
  }
})

const emit = defineEmits(["sort", "next", "prev"])

const sortKey = ref("")
const sortOrder = ref("asc")

const isFirstPage = computed(() => props.page <= 1)
const isLastPage = computed(() => props.page >= props.totalPages)

const handleSort = (col) => {
  if (!col.sortable) return

  if (sortKey.value === col.key) {
    sortOrder.value = sortOrder.value === "asc" ? "desc" : "asc"
  } else {
    sortKey.value = col.key
    sortOrder.value = "asc"
  }

  emit("sort", {
    key: sortKey.value,
    order: sortOrder.value
  })
}
</script>