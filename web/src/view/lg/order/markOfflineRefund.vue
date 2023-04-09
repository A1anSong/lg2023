<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="订单编号">
          <el-input v-model.trim="searchInfo.orderNo" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="标段编号">
          <el-input v-model.trim="searchInfo.projectNo" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="标段名称">
          <el-input v-model.trim="searchInfo.projectName" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="申请企业">
          <el-input v-model.trim="searchInfo.insureName" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="保函格式">
          <el-select v-model="searchInfo.elogTemplateId" placeholder="选择条件" clearable>
            <el-option
              v-for="template in templateData"
              :key="template.ID"
              :label="template.templateName"
              :value="template.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="保函编号">
          <el-input v-model.trim="searchInfo.elogNo" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="订单状态">
          <el-select v-model="searchInfo.orderStatus" placeholder="选择条件" clearable>
            <el-option value="已撤" />
            <el-option value="未开" />
            <el-option value="已开" />
            <el-option value="延期" />
            <el-option value="退函" />
            <el-option value="理赔" />
            <el-option value="销函" />
          </el-select>
        </el-form-item>
        <el-form-item label="审核状态">
          <el-select v-model="searchInfo.auditStatus" placeholder="选择条件" clearable>
            <el-option label="待审" value="1" />
            <el-option label="通过" value="2" />
            <el-option label="拒绝" value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="开标时间">
          <el-date-picker
            v-model="searchInfo.openBeginDate"
            type="daterange"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>
        <el-form-item label="申请时间">
          <el-date-picker
            v-model="searchInfo.applyCreatedAt"
            type="daterange"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>
        <el-form-item label="开函时间">
          <el-date-picker
            v-model="searchInfo.letterCreatedAt"
            type="daterange"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>
        <el-form-item label="担保期限" clearable>
          <el-input v-model.number="searchInfo.insureDay" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="工号">
          <el-input v-model.trim="searchInfo.authCode" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
        style="width: 100%"
        :data="tableData"
        row-key="ID"
        border
        size="small"
        table-layout="fixed"
        empty-text="无数据"
        scrollbar-always-on
        height="800"
      >
        <el-table-column align="center" label="订单编号" prop="orderNo" width="120px" />
        <el-table-column align="center" label="申请时间" width="100px">
          <template #default="scope">{{ date(scope.row.apply.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="center" label="产品类型" width="80px">
          <template #default="scope">{{ productType(scope.row.apply.productType) }}</template>
        </el-table-column>
        <el-table-column align="center" label="标段名称" prop="apply.projectName" min-width="300px" />
        <el-table-column align="center" label="开标时间" prop="apply.openBeginDate" width="100px" />
        <el-table-column align="center" label="标段编号" prop="apply.projectNo" min-width="160px" />
        <el-table-column align="center" label="受益方名称" prop="apply.insuredName" min-width="280px" />
        <el-table-column align="center" label="担保金额" min-width="120px">
          <template #default="scope">{{ amount(scope.row.apply.tenderDeposit) }}</template>
        </el-table-column>
        <el-table-column align="center" label="所属市" prop="project.projectCity" min-width="120px" />
        <el-table-column align="center" label="保函格式名称" prop="project.template.templateName" min-width="120px" />
        <el-table-column align="center" label="申请企业" prop="apply.insureName" min-width="280px" />
        <el-table-column align="center" label="审核时间" width="100px">
          <template #default="scope">{{ date(scope.row.apply.auditDate) }}</template>
        </el-table-column>
        <el-table-column align="center" label="审核状态" min-width="80px">
          <template #default="scope">
            <el-tag :type="scope.row.revoke!=null?'info':auditType(scope.row.apply.auditStatus)" effect="dark" round>
              {{ scope.row.revoke != null ? '已撤' : auditStatus(scope.row.apply.auditStatus) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" label="付款时间" width="100px">
          <template #default="scope">{{ scope.row.pay ? date(scope.row.pay.payTime) : '' }}</template>
        </el-table-column>
        <el-table-column align="center" label="付款金额" min-width="120px">
          <template #default="scope">{{ scope.row.pay != null ? amount(scope.row.pay.payAmount) : '' }}</template>
        </el-table-column>
        <el-table-column align="center" label="付款状态" min-width="80px">
          <template #default="scope">
            <el-tag :type="scope.row.pay != null ? 'success' : 'info'" effect="dark" round>
              {{ scope.row.pay != null ? "已付" : "未付" }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" label="开函时间" width="100px">
          <template #default="scope">{{ scope.row.letter ? date(scope.row.letter.CreatedAt) : '' }}</template>
        </el-table-column>
        <el-table-column align="center" label="订单状态" min-width="80px">
          <template #default="scope">
            <el-tag :type="orderStatusType(scope.row)" effect="dark" round>{{ orderStatus(scope.row) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" label="工号" prop="apply.applicantAuthCode" min-width="120px" />
        <el-table-column align="center" label="业务员" prop="employee.nickName" min-width="120px" />
        <el-table-column align="center" label="线下退款" min-width="100" fixed="right">
          <template #default="scope">
            <el-switch
              v-model="scope.row.isOfflineRefund"
              inline-prompt
              active-text="是"
              inactive-text="否"
              style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
              @change="changeOfflineRefund($event, scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MarkOfflineRefund'
}
</script>

<script setup>
import {
  getOrderList,
  markOfflineRefund,
  unmarkOfflineRefund
} from '@/api/lg/order'

import { ref } from 'vue'

import { date } from '@/utils/lg/date'
import { auditStatus, auditType } from '@/utils/lg/auditStatus'
import { productType } from '@/utils/lg/productType'
import { amount } from '@/utils/lg/amount'
import { orderStatus, orderStatusType } from '@/utils/lg/orderStatus'
import { getTemplateList } from '@/api/lg/template'
import { ElMessage } from 'element-plus'

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const templateData = ref([])

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 获取模板列表
const getTemplateData = async() => {
  const template = await getTemplateList({ page: 1, pageSize: 999 })
  if (template.code === 0) {
    templateData.value = template.data.list
  }
}

getTemplateData()
getTableData()

// ============== 表格控制部分结束 ===============

const changeOfflineRefund = async(val, row) => {
  if (val === true) {
    const res = await markOfflineRefund(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '标记成功'
      })
      await getTableData()
    } else {
      row.isEnable = false
    }
  } else {
    const res = await unmarkOfflineRefund(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '标记成功'
      })
      await getTableData()
    } else {
      row.isEnable = true
    }
  }
}
</script>

<style>
</style>
