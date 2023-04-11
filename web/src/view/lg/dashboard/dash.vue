<template>
  <div class="page">
    <div class="gva-card">
      <el-row :gutter="24">
        <el-col :xs="24" :md="12" :xl="8">
          <DataStatistic />
        </el-col>
        <el-col :xs="24" :md="12" :xl="8">
          <InsuranceBalance />
        </el-col>
        <el-col :xs="24" :md="12" :xl="8">
          <GeoDistribution />
        </el-col>
        <el-col v-auth="btnAuth.all" :xs="24" :md="12" :xl="8">
          <EmployeePieChart />
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
import DataStatistic from '@/view/lg/dashboard/dashboardTable/dataStatistic.vue'
import InsuranceBalance from '@/view/lg/dashboard/dashboardTable/insuranceBalance.vue'
import EmployeePieChart from '@/view/lg/dashboard/dashboardCharts/employeeStatistic.vue'
import GeoDistribution from '@/view/lg/dashboard/dashboardCharts/geoDistribution.vue'

import { useUserStore } from '@/pinia/modules/user'
import { useBtnAuth } from '@/utils/btnAuth'
import { getOrderTrendData } from '@/api/lg/order'
const userStore = useUserStore()
const userInfo = userStore.userInfo
const btnAuth = useBtnAuth()

const searchInfo = ref({})
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

const getTrendData = async() => {
  const trend = await getOrderTrendData({ ...searchInfo.value })
  if (trend.code === 0) {
    orderTrendData.value = trend.data.orderTrendData
  }
}

checkAuthorityAll()
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
