<template>
  <div class="page">
    <div class="gva-card">
      <el-row :gutter="24">
        <el-col :span="8">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <span style="font-size: 16px">今日数据</span>
              </div>
            </template>
            <el-space wrap>
              <div class="statistic-card">
                <el-statistic :value="orderStatisticData.todayOrderCount">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      订单数量
                      <el-tooltip
                        effect="dark"
                        :content="'昨天成交 '+orderStatisticData.yesterdayOrderCount+' 单'"
                        placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <Warning />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>较昨天</span>
                    <span :class="compareValue(orderStatisticData.todayOrderCount, orderStatisticData.yesterdayOrderCount)>=0?'red':'green'">
                      {{ Math.abs(compareValue(orderStatisticData.todayOrderCount, orderStatisticData.yesterdayOrderCount))*100 }}%
                      <el-icon>
                        <CaretTop v-if="compareValue(orderStatisticData.todayOrderCount, orderStatisticData.yesterdayOrderCount)>=0" />
                        <CaretBottom v-else />
                      </el-icon>
                    </span>
                  </div>
                </div>
              </div>
              <div class="statistic-card">
                <el-statistic :value="orderStatisticData.todayGuaranteeAmount/10000" :precision="4" suffix="万元">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      担保金额
                      <el-tooltip
                        effect="dark"
                        :content="'昨天担保 '+(orderStatisticData.yesterdayGuaranteeAmount/10000).toFixed(4).replace(/(\d)(?=(\d{3})+\.)/g, '$1,')+' 万元'"
                        placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <Warning />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>较昨天</span>
                    <span :class="compareValue(orderStatisticData.todayGuaranteeAmount, orderStatisticData.yesterdayGuaranteeAmount)>=0?'red':'green'">
                      {{ Math.abs(compareValue(orderStatisticData.todayGuaranteeAmount, orderStatisticData.yesterdayGuaranteeAmount))*100 }}%
                      <el-icon>
                        <CaretTop v-if="compareValue(orderStatisticData.todayGuaranteeAmount, orderStatisticData.yesterdayGuaranteeAmount)>=0" />
                        <CaretBottom v-else />
                      </el-icon>
                    </span>
                  </div>
                </div>
              </div>
              <div v-auth="btnAuth.all" class="statistic-card">
                <el-statistic :value="(orderStatisticData.todayElogAmount/1).toFixed(2)/1" :precision="2" suffix="元">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      保费金额
                      <el-tooltip
                        effect="dark"
                        :content="'昨天保费 '+(orderStatisticData.yesterdayElogAmount/1).toFixed(2).replace(/(\d)(?=(\d{3})+\.)/g, '$1,')+' 元'"
                        placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <Warning />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>较昨天</span>
                    <span :class="compareValue(orderStatisticData.todayElogAmount, orderStatisticData.yesterdayElogAmount)>=0?'red':'green'">
                      {{ Math.abs(compareValue(orderStatisticData.todayElogAmount, orderStatisticData.yesterdayElogAmount))*100 }}%
                      <el-icon>
                        <CaretTop v-if="compareValue(orderStatisticData.todayElogAmount, orderStatisticData.yesterdayElogAmount)>=0" />
                        <CaretBottom v-else />
                      </el-icon>
                    </span>
                  </div>
                </div>
              </div>
            </el-space>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <span style="font-size: 16px">本周数据</span>
              </div>
            </template>
            <el-space wrap>
              <div class="statistic-card">
                <el-statistic :value="orderStatisticData.weekOrderCount">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      订单数量
                      <el-tooltip
                        effect="dark"
                        :content="'上周成交 '+orderStatisticData.lastWeekOrderCount+' 单'"
                        placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <Warning />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>较上周</span>
                    <span :class="compareValue(orderStatisticData.weekOrderCount, orderStatisticData.lastWeekOrderCount)>=0?'red':'green'">
                      {{ Math.abs(compareValue(orderStatisticData.weekOrderCount, orderStatisticData.lastWeekOrderCount))*100 }}%
                      <el-icon>
                        <CaretTop v-if="compareValue(orderStatisticData.weekOrderCount, orderStatisticData.lastWeekOrderCount)>=0" />
                        <CaretBottom v-else />
                      </el-icon>
                    </span>
                  </div>
                </div>
              </div>
              <div class="statistic-card">
                <el-statistic :value="orderStatisticData.weekGuaranteeAmount/10000" :precision="4" suffix="万元">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      担保金额
                      <el-tooltip
                        effect="dark"
                        :content="'上周担保 '+(orderStatisticData.lastWeekGuaranteeAmount/10000).toFixed(4).replace(/(\d)(?=(\d{3})+\.)/g, '$1,')+' 万元'"
                        placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <Warning />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>较上周</span>
                    <span :class="compareValue(orderStatisticData.weekGuaranteeAmount, orderStatisticData.lastWeekGuaranteeAmount)>=0?'red':'green'">
                      {{ Math.abs(compareValue(orderStatisticData.weekGuaranteeAmount, orderStatisticData.lastWeekGuaranteeAmount))*100 }}%
                      <el-icon>
                        <CaretTop v-if="compareValue(orderStatisticData.weekGuaranteeAmount, orderStatisticData.lastWeekGuaranteeAmount)>=0" />
                        <CaretBottom v-else />
                      </el-icon>
                    </span>
                  </div>
                </div>
              </div>
              <div v-auth="btnAuth.all" class="statistic-card">
                <el-statistic :value="(orderStatisticData.weekElogAmount/1).toFixed(2)/1" :precision="2" suffix="元">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      保费金额
                      <el-tooltip
                        effect="dark"
                        :content="'上周保费 '+(orderStatisticData.lastWeekElogAmount/1).toFixed(2).replace(/(\d)(?=(\d{3})+\.)/g, '$1,')+' 元'"
                        placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <Warning />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>较上周</span>
                    <span :class="compareValue(orderStatisticData.weekElogAmount, orderStatisticData.lastWeekElogAmount)>=0?'red':'green'">
                      {{ Math.abs(compareValue(orderStatisticData.weekElogAmount, orderStatisticData.lastWeekElogAmount))*100 }}%
                      <el-icon>
                        <CaretTop v-if="compareValue(orderStatisticData.weekElogAmount, orderStatisticData.lastWeekElogAmount)>=0" />
                        <CaretBottom v-else />
                      </el-icon>
                    </span>
                  </div>
                </div>
              </div>
            </el-space>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <span style="font-size: 16px">本月数据</span>
              </div>
            </template>
            <el-space wrap>
              <div class="statistic-card">
                <el-statistic :value="orderStatisticData.monthOrderCount">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      订单数量
                      <el-tooltip
                        effect="dark"
                        :content="'上月成交 '+orderStatisticData.lastMonthOrderCount+' 单'"
                        placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <Warning />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>较上月</span>
                    <span :class="compareValue(orderStatisticData.monthOrderCount, orderStatisticData.lastMonthOrderCount)>=0?'red':'green'">
                      {{ Math.abs(compareValue(orderStatisticData.monthOrderCount, orderStatisticData.lastMonthOrderCount))*100 }}%
                      <el-icon>
                        <CaretTop v-if="compareValue(orderStatisticData.monthOrderCount, orderStatisticData.lastMonthOrderCount)>=0" />
                        <CaretBottom v-else />
                      </el-icon>
                    </span>
                  </div>
                </div>
              </div>
              <div class="statistic-card">
                <el-statistic :value="orderStatisticData.monthGuaranteeAmount/10000" :precision="4" suffix="万元">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      担保金额
                      <el-tooltip
                        effect="dark"
                        :content="'上周担保 '+(orderStatisticData.lastMonthGuaranteeAmount/10000).toFixed(4).replace(/(\d)(?=(\d{3})+\.)/g, '$1,')+' 万元'"
                        placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <Warning />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>较上周</span>
                    <span :class="compareValue(orderStatisticData.monthGuaranteeAmount, orderStatisticData.lastMonthGuaranteeAmount)>=0?'red':'green'">
                      {{ Math.abs(compareValue(orderStatisticData.monthGuaranteeAmount, orderStatisticData.lastMonthGuaranteeAmount))*100 }}%
                      <el-icon>
                        <CaretTop v-if="compareValue(orderStatisticData.monthGuaranteeAmount, orderStatisticData.lastMonthGuaranteeAmount)>=0" />
                        <CaretBottom v-else />
                      </el-icon>
                    </span>
                  </div>
                </div>
              </div>
              <div v-auth="btnAuth.all" class="statistic-card">
                <el-statistic :value="(orderStatisticData.monthElogAmount/1).toFixed(2)/1" :precision="2" suffix="元">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      保费金额
                      <el-tooltip
                        effect="dark"
                        :content="'上周保费 '+(orderStatisticData.lastMonthElogAmount/1).toFixed(2).replace(/(\d)(?=(\d{3})+\.)/g, '$1,')+' 元'"
                        placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <Warning />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>较上周</span>
                    <span :class="compareValue(orderStatisticData.monthElogAmount, orderStatisticData.lastMonthElogAmount)>=0?'red':'green'">
                      {{ Math.abs(compareValue(orderStatisticData.monthElogAmount, orderStatisticData.lastMonthElogAmount))*100 }}%
                      <el-icon>
                        <CaretTop v-if="compareValue(orderStatisticData.monthElogAmount, orderStatisticData.lastMonthElogAmount)>=0" />
                        <CaretBottom v-else />
                      </el-icon>
                    </span>
                  </div>
                </div>
              </div>
            </el-space>
          </el-card>
        </el-col>
      </el-row>
      <el-row :gutter="24">
        <el-col :span="8">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <span>总计</span>
              </div>
            </template>
            <div class="statistic-card">
              <el-statistic :value="orderStatisticData.totalOrderCount">
                <template #title>
                  <div style="display: inline-flex; align-items: center">
                    订单数量
                  </div>
                </template>
              </el-statistic>
            </div>
            <div class="statistic-card">
              <el-statistic :value="orderStatisticData.totalGuaranteeAmount/10000" :precision="4" suffix="万元">
                <template #title>
                  <div style="display: inline-flex; align-items: center">
                    担保金额
                  </div>
                </template>
              </el-statistic>
            </div>
            <div v-auth="btnAuth.all" class="statistic-card">
              <el-statistic :value="(orderStatisticData.totalElogAmount/1).toFixed(2)/1" :precision="2" suffix="元">
                <template #title>
                  <div style="display: inline-flex; align-items: center">
                    保费金额
                  </div>
                </template>
              </el-statistic>
            </div>
          </el-card>
        </el-col>
        <el-col v-auth="btnAuth.all" :span="8">
          <EmployeePieChart />
        </el-col>
        <el-col :span="8">
          <GeoDistribution />
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dash'
}
</script>

<script setup>
import { ref } from 'vue'
import { Warning, CaretTop, CaretBottom } from '@element-plus/icons-vue'
import EmployeePieChart from '@/view/lg/dashboard/dashboardCharts/employeeStatistic.vue'
import GeoDistribution from '@/view/lg/dashboard/dashboardCharts/geoDistribution.vue'

import { useUserStore } from '@/pinia/modules/user'
import { useBtnAuth } from '@/utils/btnAuth'
import { getOrderStatisticData, getOrderTrendData } from '@/api/lg/order'
const userStore = useUserStore()
const userInfo = userStore.userInfo
const btnAuth = useBtnAuth()

const searchInfo = ref({})
const orderStatisticData = ref({})
const orderTrendData = ref([])

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

const getTrendData = async() => {
  const trend = await getOrderTrendData({ ...searchInfo.value })
  if (trend.code === 0) {
    orderTrendData.value = trend.data.orderTrendData
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
getTrendData()
</script>

<style scoped>
.gva-card {
    box-sizing: border-box;
    background-color: #fff;
    border-radius: 2px;
    height: auto;
    padding: 26px 30px;
    overflow: hidden;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
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
