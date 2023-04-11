<template>
  <el-card shadow="hover">
    <template #header>
      <div class="card-header">
        <span style="font-size: 16px">剩余额度</span>
      </div>
    </template>
    <div style="height: 272px; width: auto; text-align: center">
      <el-progress type="circle" :width="256" :color="balanceColors" :percentage="balancePercentage">
        <template #default="{ percentage }">
          <div style="font-size: 20px;">
            <span class="percentage-value">{{ percentage.toFixed(2) }}% 剩余</span>
            <br>
            <span class="percentage-label">总额度 {{ insuranceTotal.toFixed(4) }} 万元</span>
            <br>
            <span class="percentage-label">剩余额度 {{ insuranceBalance.toFixed(4) }} 万元</span>
          </div>
        </template>
      </el-progress>
    </div>
  </el-card>
</template>

<script>
export default {
  name: 'InsuranceBalance'
}
</script>

<script setup>
import { computed, ref } from 'vue'
import { getInsuranceBalance } from '@/api/lg/order'

const insuranceBalance = ref(0)
const insuranceTotal = ref(0)
const balancePercentage = computed(() => {
  if (insuranceTotal.value === 0) return 0
  return (insuranceBalance.value / insuranceTotal.value) * 100
})

const balanceColors = [
  { color: '#F56C6C', percentage: 20 },
  { color: '#E6A23C', percentage: 50 },
  { color: '#67C23A', percentage: 80 },
]

const getData = async() => {
  const res = await getInsuranceBalance()
  if (res.code === 0) {
    insuranceBalance.value = res.data.insuranceBalance.insuranceBalance
    insuranceTotal.value = res.data.insuranceBalance.insuranceTotal
  }
}

getData()
</script>

<style scoped>
.el-card {
    margin-bottom: 20px;
}

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 34px;
}
</style>
