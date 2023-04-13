<template>
  <div class="page">
    <div class="gva-card">
      <el-row :gutter="24">
        <el-col :xs="24" :md="12" :xl="8">
          <DataStatistic />
        </el-col>
        <el-col v-auth="btnAuth.all" :xs="24" :md="12" :xl="8">
          <InsuranceBalance />
        </el-col>
        <el-col :xs="24" :md="12" :xl="8">
          <GeoDistribution />
        </el-col>
        <el-col v-auth="btnAuth.all" :xs="24" :md="12" :xl="8">
          <EmployeePieChart />
        </el-col>
      </el-row>
      <el-row :gutter="24">
        <el-col :span="24">
          <OrderTrendData />
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
import DataStatistic from '@/view/lg/dashboard/dashboardTable/dataStatistic.vue'
import InsuranceBalance from '@/view/lg/dashboard/dashboardTable/insuranceBalance.vue'
import EmployeePieChart from '@/view/lg/dashboard/dashboardCharts/employeeStatistic.vue'
import GeoDistribution from '@/view/lg/dashboard/dashboardCharts/geoDistribution.vue'
import OrderTrendData from '@/view/lg/dashboard/dashboardCharts/orderTrendData.vue'

import { useBtnAuth } from '@/utils/btnAuth'
const btnAuth = useBtnAuth()
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
</style>
