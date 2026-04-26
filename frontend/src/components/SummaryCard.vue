<template>
  <div class="summary-container">
    <div class="card total">
      <p class="label">Total</p>
      <p class="value">{{ summary.total }}</p>
    </div>

    <div class="card completed">
      <p class="label">Completed</p>
      <p class="value">{{ summary.completed }}</p>
    </div>

    <div class="card processing">
      <p class="label">Processing</p>
      <p class="value">{{ summary.processing }}</p>
    </div>

    <div class="card failed">
      <p class="label">Failed</p>
      <p class="value">{{ summary.failed }}</p>
    </div>
  </div>
</template>

<script setup>
defineProps({
  summary: {
    type: Object,
    required: true
  }
})
</script>

<style scoped>
.summary-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 16px;
  margin-bottom: 20px;
}

.card {
  padding: 16px;
  border-radius: 12px;
  background: #f5f7fa;
  box-shadow: 0 2px 6px rgba(0,0,0,0.05);
}

.label {
  font-size: 14px;
  color: #666;
}

.value {
  font-size: 22px;
  font-weight: bold;
  margin-top: 4px;
}

/* warna biar beda-beda */
.total {
  border-left: 5px solid #333;
}

.completed {
  border-left: 5px solid #28a745;
}

.processing {
  border-left: 5px solid #ffc107;
}

.failed {
  border-left: 5px solid #dc3545;
}
</style>