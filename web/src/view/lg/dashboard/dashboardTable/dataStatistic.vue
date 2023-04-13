<template>
  <el-card v-if="statisticData" shadow="hover">
    <template #header>
      <div class="card-header">
        <span style="font-size: 16px">统计数据</span>
        <el-select v-model="periodType" style="width: 100px">
          <el-option
            v-for="type in periodCollection"
            :key="type.value"
            :label="type.label"
            :value="type.value"
          />
        </el-select>
      </div>
    </template>
    <el-space wrap style="height: 280px">
      <div class="statistic-card">
        <el-statistic :value="statisticData.orderCount">
          <template #title>
            <div style="display: inline-flex; align-items: center">
              订单数量
              <el-tooltip
                v-if="periodType !== 'total'"
                effect="dark"
                :content="lastPeriod(periodType)+'成交 '+comparedStatisticData.orderCount+' 单'"
                placement="top"
              >
                <el-icon style="margin-left: 4px" :size="12">
                  <Warning />
                </el-icon>
              </el-tooltip>
            </div>
          </template>
        </el-statistic>
        <div v-if="periodType !== 'total'" class="statistic-footer">
          <div class="footer-item">
            <span>较{{ lastPeriod(periodType) }}</span>
            <span :class="compareValue(statisticData.orderCount, comparedStatisticData.orderCount)>=0?'red':'green'">
              {{ (Math.abs(compareValue(statisticData.orderCount, comparedStatisticData.orderCount))*100).toFixed(2) }}%
              <el-icon>
                <CaretTop v-if="compareValue(statisticData.orderCount, comparedStatisticData.orderCount)>=0" />
                <CaretBottom v-else />
              </el-icon>
            </span>
          </div>
        </div>
      </div>
      <div class="statistic-card">
        <el-statistic :value="statisticData.guaranteeAmount/10000" :precision="4" suffix="万元">
          <template #title>
            <div style="display: inline-flex; align-items: center">
              担保金额
              <el-tooltip
                v-if="periodType !== 'total'"
                effect="dark"
                :content="lastPeriod(periodType)+'担保 '+(comparedStatisticData.guaranteeAmount/10000).toFixed(4).replace(/(\d)(?=(\d{3})+\.)/g, '$1,')+' 万元'"
                placement="top"
              >
                <el-icon style="margin-left: 4px" :size="12">
                  <Warning />
                </el-icon>
              </el-tooltip>
            </div>
          </template>
        </el-statistic>
        <div v-if="periodType !== 'total'" class="statistic-footer">
          <div class="footer-item">
            <span>较{{ lastPeriod(periodType) }}</span>
            <span :class="compareValue(statisticData.guaranteeAmount, comparedStatisticData.guaranteeAmount)>=0?'red':'green'">
              {{ (Math.abs(compareValue(statisticData.guaranteeAmount, comparedStatisticData.guaranteeAmount))*100).toFixed(2) }}%
              <el-icon>
                <CaretTop v-if="compareValue(statisticData.guaranteeAmount, comparedStatisticData.guaranteeAmount)>=0" />
                <CaretBottom v-else />
              </el-icon>
            </span>
          </div>
        </div>
      </div>
      <div v-auth="btnAuth.all" class="statistic-card">
        <el-statistic :value="(statisticData.elogAmount/1).toFixed(2)/1" :precision="2" suffix="元">
          <template #title>
            <div style="display: inline-flex; align-items: center">
              保费金额
              <el-tooltip
                v-if="periodType !== 'total'"
                effect="dark"
                :content="lastPeriod(periodType)+'保费 '+(comparedStatisticData.elogAmount/1).toFixed(2).replace(/(\d)(?=(\d{3})+\.)/g, '$1,')+' 元'"
                placement="top"
              >
                <el-icon style="margin-left: 4px" :size="12">
                  <Warning />
                </el-icon>
              </el-tooltip>
            </div>
          </template>
        </el-statistic>
        <div v-if="periodType !== 'total'" class="statistic-footer">
          <div class="footer-item">
            <span>较{{ lastPeriod(periodType) }}</span>
            <span :class="compareValue(statisticData.elogAmount, comparedStatisticData.elogAmount)>=0?'red':'green'">
              {{ (Math.abs(compareValue(statisticData.elogAmount, comparedStatisticData.elogAmount))*100).toFixed(2) }}%
              <el-icon>
                <CaretTop v-if="compareValue(statisticData.elogAmount, comparedStatisticData.elogAmount)>=0" />
                <CaretBottom v-else />
              </el-icon>
            </span>
          </div>
        </div>
      </div>
    </el-space>
  </el-card>
</template>

<script>

export default {
  name: 'DataStatistic'
}
</script>

<script setup>
import { CaretBottom, CaretTop, Warning } from '@element-plus/icons-vue'

import { computed, ref } from 'vue'
import { getOrderStatisticData } from '@/api/lg/order'
import { lastPeriod } from '@/utils/lg/lastPeriod'

import { useUserStore } from '@/pinia/modules/user'
import { useBtnAuth } from '@/utils/btnAuth'
const userStore = useUserStore()
const userInfo = userStore.userInfo
const btnAuth = useBtnAuth()

const searchInfo = ref({})
const orderStatisticData = ref({})
const statisticData = computed(() => {
  return orderStatisticData.value[periodType.value]
})
const comparedStatisticData = computed(() => {
  return orderStatisticData.value['last' + periodType.value]
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

const getStatisticData = async() => {
  const statistic = await getOrderStatisticData({ ...searchInfo.value })
  if (statistic.code === 0) {
    orderStatisticData.value = statistic.data.orderStatisticData
  }
}

const compareValue = (current, last) => {
  if (last === 0) {
    return (0).toFixed(4)
  } else {
    return (current / last - 1).toFixed(4)
  }
}

checkAuthorityAll()
getStatisticData()
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

.el-statistic {
    --el-statistic-content-font-size: 28px;
}

.statistic-card{
    height: 100%;
    padding: 20px;
    border-radius: 4px;
    background-color: var(--el-bg-color-overlay);
}

.statistic-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-wrap: wrap;
    font-size: 12px;
    color: var(--el-text-color-regular);
    margin-top: 16px;
}

.statistic-footer .footer-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.statistic-footer .footer-item span:last-child {
    display: inline-flex;
    align-items: center;
    margin-left: 4px;
}

.green {
    color: var(--el-color-success);
}
.red {
    color: var(--el-color-error);
}
</style>
