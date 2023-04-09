<template>
  <el-card shadow="hover">
    <template #header>
      <div class="card-header">
        <span>开单占比（本周）</span>
      </div>
    </template>
    <div ref="echart" style="height: 291px; width: auto;" />
  </el-card>
</template>

<script>
export default {
  name: 'EmployeeStatistic'
}
</script>
<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
import { getEmployeeStatisticData } from '@/api/lg/order'
const chart = shallowRef(null)
const echart = ref(null)
const dataSet = ref([])
const initChart = () => {
  chart.value = echarts.init(echart.value, '', { locale: 'ZH' })
  getData()
}
// 绘制图表
const getData = async() => {
  chart.value.showLoading()
  const table = await getEmployeeStatisticData()
  if (table.code === 0) {
    dataSet.value.push(['来源', '数量'])
    table.data.employeeStatisticData.employeeData.forEach((item) => {
      dataSet.value.push([item.name, item.count])
    })
    chart.value.setOption({
      tooltip: {},
      series: [{ type: 'pie' }],
      dataset: {
        source: dataSet.value
      },
    })
    chart.value.hideLoading()
  }
}

onMounted(async() => {
  await nextTick()
  initChart()
  window.addEventListener('resize', () => {
    chart.value.resize()
  })
})

onUnmounted(() => {
  if (!chart.value) {
    return
  }
  chart.value.dispose()
  chart.value = null
})
</script>

<style scoped>

</style>
