<template>
  <el-card shadow="hover">
    <template #header>
      <div class="card-header">
        <span style="font-size: 16px">订单趋势</span>
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
  name: 'OrderTrendData'
}
</script>
<script setup>
import * as echarts from 'echarts'
import { computed, nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
const chart = shallowRef(null)
const echart = ref(null)

import { useUserStore } from '@/pinia/modules/user'
import { useBtnAuth } from '@/utils/btnAuth'
import { getOrderTrendData } from '@/api/lg/order'
const userStore = useUserStore()
const userInfo = userStore.userInfo
const btnAuth = useBtnAuth()

const searchInfo = ref({})

const dataAll = ref({
  seven: [],
  thirty: []
})
const dataSet = computed(() => {
  return dataAll.value[periodType.value]
})

const periodCollection = ref([
  {
    label: '近7天',
    value: 'seven'
  },
  {
    label: '近30天',
    value: 'thirty'
  }
])
const periodType = ref('seven')
const weekArr = ['日', '一', '二', '三', '四', '五', '六']

const initChart = () => {
  chart.value = echarts.init(echart.value, '', { locale: 'ZH' })
  checkAuthorityAll()
  getData()
}

const checkAuthorityAll = () => {
  const authority = btnAuth.all
  let type = ''
  switch (Object.prototype.toString.call(authority)) {
    case '[object Array]':
      type = 'Array'
      break
    case '[object String]':
      type = 'String'
      break
    case '[object Number]':
      type = 'Number'
      break
    default:
      type = ''
      break
  }
  if (type === '') {
    searchInfo.value.employeeNo = userInfo.ID
    return
  }
  const waitUse = authority.toString().split(',')
  const flag = waitUse.some(item => item === userInfo.authorityId.toString())
  if (!flag) {
    searchInfo.value.employeeNo = userInfo.ID
  }
}

// 绘制图表
const getData = async() => {
  chart.value.showLoading()
  const table = await getOrderTrendData({ ...searchInfo.value })
  if (table.code === 0) {
    table.data.orderTrendData.seven.forEach((item) => {
      dataAll.value.seven.push([item.date, item.orderCount, item.guaranteeAmount])
    })
    table.data.orderTrendData.thirty.forEach((item) => {
      dataAll.value.thirty.push([item.date, item.orderCount, item.guaranteeAmount])
    })
    chart.value.setOption({
      tooltip: {
        trigger: 'axis',
        axisPointer: { type: 'cross' },
      },
      legend: {},
      xAxis: {
        type: 'time',
        axisLabel: {
          formatter: '{MMM}{d}日'
        },
        axisPointer: {
          label: {
            formatter: (params) => {
              // return echarts.format.formatTime('M月d日', params.value)
              //   return params.value.format('M月d日')
              const date = new Date(params.value)
              return `${date.getMonth() + 1}月${date.getDate()}日 星期${weekArr[date.getDay()]}`
            }
          }
        }
      },
      yAxis: [
        {
          name: '订单数量',
          type: 'value',
          position: 'left',
          splitLine: {
            show: false,
          },
        },
        {
          name: '担保金额',
          type: 'value',
          position: 'right',
          axisLabel: {
            formatter: (value) => {
              return (value / 10000).toFixed(0) + '万元'
            }
          },
        }
      ],
      dataset: {
        dimensions: ['日期', '订单数量', '担保金额'],
        source: dataSet.value
      },
      series: [
        {
          name: '订单数量',
          type: 'bar',
          yAxisIndex: 0
        },
        {
          name: '担保金额',
          type: 'line',
          smooth: true,
          yAxisIndex: 1,
          encode: {
            y: 2
          },
          tooltip: {
            valueFormatter: (value) => {
              return (value / 10000).toFixed(0) + '万元'
            }
          }
        }
      ],
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
.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
</style>
