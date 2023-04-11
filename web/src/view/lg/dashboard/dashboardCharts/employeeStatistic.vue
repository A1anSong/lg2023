<template>
  <el-card shadow="hover">
    <template #header>
      <div class="card-header">
        <span style="font-size: 16px">订单占比</span>
        <el-select v-model="periodType" style="width: 100px" @change="refreshData">
          <el-option
            v-for="type in periodCollection"
            :key="type.value"
            :label="type.label"
            :value="type.value"
          />
        </el-select>
      </div>
    </template>
    <div ref="echart" style="height: 272px; width: auto;" />
  </el-card>
</template>

<script>
export default {
  name: 'EmployeeStatistic'
}
</script>
<script setup>
import * as echarts from 'echarts'
import { computed, nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
import { getEmployeeStatisticData } from '@/api/lg/order'
const chart = shallowRef(null)
const echart = ref(null)
const dataTitle = ref(['来源', '数量'])
const dataAll = ref({
  day: [],
  week: [],
  month: [],
  total: []
})
const dataSet = computed(() => {
  return [dataTitle.value, ...dataAll.value[periodType.value]]
})

const periodCollection = ref([
  {
    label: '今天',
    value: 'day'
  },
  {
    label: '这周',
    value: 'week'
  },
  {
    label: '这个月',
    value: 'month'
  },
  {
    label: '总计',
    value: 'total'
  }
])
const periodType = ref('day')

const initChart = () => {
  chart.value = echarts.init(echart.value, '', { locale: 'ZH' })
  getData()
}
// 绘制图表
const getData = async() => {
  chart.value.showLoading()
  const table = await getEmployeeStatisticData()
  if (table.code === 0) {
    table.data.employeeStatisticData.day.forEach((item) => {
      dataAll.value.day.push([item.name, item.count])
    })
    table.data.employeeStatisticData.week.forEach((item) => {
      dataAll.value.week.push([item.name, item.count])
    })
    table.data.employeeStatisticData.month.forEach((item) => {
      dataAll.value.month.push([item.name, item.count])
    })
    table.data.employeeStatisticData.total.forEach((item) => {
      dataAll.value.total.push([item.name, item.count])
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

const refreshData = async() => {
  chart.value.showLoading()
  chart.value.setOption({
    dataset: {
      source: dataSet.value
    },
  })
  chart.value.hideLoading()
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
.el-card {
    margin-bottom: 20px;
}

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
</style>
