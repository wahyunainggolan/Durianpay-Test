<template>
  <div class="summary-grid">
    <div class="card total">
      <p class="label">Total</p>
      <p class="value">{{ summary.total }}</p>
    </div>

    <div class="card success">
      <p class="label">Completed</p>
      <p class="value">{{ summary.completed }}</p>
    </div>

    <div class="card warning">
      <p class="label">Processing</p>
      <p class="value">{{ summary.processing }}</p>
    </div>

    <div class="card danger">
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
.summary-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}
.card {
  padding: 18px 20px;
  border-radius: 14px;
  background: #ffffff;
  border: 1px solid #eef0f3;
  box-shadow: 0 6px 18px rgba(16, 24, 40, 0.06);
  transition: all 0.25s ease;
  position: relative;
  overflow: hidden;
}

.card:hover {
  transform: translateY(-4px);
  box-shadow: 0 10px 28px rgba(16, 24, 40, 0.12);
}
.label {
  font-size: 13px;
  color: #6b7280;
  font-weight: 500;
  margin-bottom: 6px;
}
.value {
  font-size: 26px;
  font-weight: 700;
  color: #111827;
  letter-spacing: -0.5px;
}
.total {
  border-left: 4px solid #111827;
}

.success {
  border-left: 4px solid #22c55e;
}

.warning {
  border-left: 4px solid #f59e0b;
}

.danger {
  border-left: 4px solid #ef4444;
}

.success::after,
.warning::after,
.danger::after,
.total::after {
  content: "";
  position: absolute;
  top: -40px;
  right: -40px;
  width: 120px;
  height: 120px;
  opacity: 0.06;
  border-radius: 50%;
}

.success::after {
  background: #22c55e;
}

.warning::after {
  background: #f59e0b;
}

.danger::after {
  background: #ef4444;
}

.total::after {
  background: #111827;
}
</style>