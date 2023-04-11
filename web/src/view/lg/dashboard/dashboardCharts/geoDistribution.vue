<template>
  <el-card shadow="hover">
    <template #header>
      <div class="card-header">
        <span style="font-size: 16px">分布图</span>
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
  name: 'GeoDistribution'
}
</script>

<script setup>
import * as echarts from 'echarts'
import { computed, nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
import { getGEODistributionData } from '@/api/lg/order'
import jxData from '@/assets/jiangxi.json'
const chart = shallowRef(null)
const echart = ref(null)

import { useUserStore } from '@/pinia/modules/user'
import { useBtnAuth } from '@/utils/btnAuth'
const userStore = useUserStore()
const userInfo = userStore.userInfo
const btnAuth = useBtnAuth()

const searchInfo = ref({})

const dataTitle = ref(['城市单量', '值'])
const dataAll = ref({
  day: [],
  week: [],
  month: [],
  total: []
})
const dataSet = computed(() => {
  return [dataTitle.value, ...dataAll.value[periodType.value]]
})
const dataSetMin = computed(() => {
  return Math.min.apply(null, dataAll.value[periodType.value].map((item) => {
    return item[1]
  }))
})
const dataSetMax = computed(() => {
  return Math.max.apply(null, dataAll.value[periodType.value].map((item) => {
    return item[1]
  }))
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
  echarts.registerMap('jiangxi', jxData)
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
  const table = await getGEODistributionData({ ...searchInfo.value })
  if (table.code === 0) {
    table.data.geoDistributionData.day.forEach((item) => {
      dataAll.value.day.push([item.city, item.count])
    })
    table.data.geoDistributionData.week.forEach((item) => {
      dataAll.value.week.push([item.city, item.count])
    })
    table.data.geoDistributionData.month.forEach((item) => {
      dataAll.value.month.push([item.city, item.count])
    })
    table.data.geoDistributionData.total.forEach((item) => {
      dataAll.value.total.push([item.city, item.count])
    })
    chart.value.setOption({
      tooltip: {},
      visualMap: {
        min: dataSetMin.value,
        max: dataSetMax.value,
        text: ['多', '少'],
        calculable: true,
        inRange: {
          color: ['white', 'yellow', 'red']
        }
      },
      series: [{
        type: 'map',
        map: 'jiangxi',
      }],
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
    visualMap: {
      min: dataSetMin.value,
      max: dataSetMax.value,
    },
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
