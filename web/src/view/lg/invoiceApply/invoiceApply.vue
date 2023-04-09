<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="发票抬头">
          <el-input v-model.trim="searchInfo.invoiceTile" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="申请时间">
          <el-date-picker
            v-model="searchInfo.applyTime"
            type="daterange"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
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
        <el-table-column align="center" label="状态" min-width="150" fixed="left">
          <template #default="scope">
            <el-tag
              v-if="scope.row.auditStatus!==2 && scope.row.auditStatus!==3"
              type="warning"
              effect="dark"
              size="large"
            >未到审批时间
            </el-tag>
            <el-tag
              v-if="scope.row.auditStatus===2"
              type="success"
              effect="dark"
              size="large"
            >已审批通过
            </el-tag>
            <el-tag
              v-if="scope.row.auditStatus===3"
              type="danger"
              effect="dark"
              size="large"
            >已审批拒绝
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" label="申请时间" width="100px">
          <template #default="scope">{{ date(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="center" label="申请编号" prop="applyNo" min-width="120px" />
        <el-table-column align="center" label="总金额" min-width="80px">
          <template #default="scope">{{ amount(scope.row.invoiceTotalAmount) }}</template>
        </el-table-column>
        <el-table-column align="center" label="发票类型" min-width="120px">
          <template #default="scope">{{ invoiceType(scope.row.invoiceType) }}</template>
        </el-table-column>
        <el-table-column align="center" label="抬头类型" min-width="80px">
          <template #default="scope">{{ invoiceTileType(scope.row.invoiceTileType) }}</template>
        </el-table-column>
        <el-table-column align="center" label="发票抬头" prop="invoiceTile" min-width="300px" />
        <el-table-column align="center" label="税号" prop="taxNo" min-width="180px" />
        <el-table-column align="center" label="开户银行" prop="bankName" min-width="300px" />
        <el-table-column align="center" label="银行账号" prop="bankNo" min-width="200px" />
        <el-table-column align="center" label="企业地址" prop="companyAddress" min-width="300px" />
        <el-table-column align="center" label="企业电话" prop="companyTel" min-width="120px" />
        <el-table-column align="center" label="开票备注" prop="remarks" min-width="280px" />
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
  name: 'InvoiceApply'
}
</script>

<script setup>
import {
  getInvoiceApplyList
} from '@/api/lg/invoiceApply'

import { ref } from 'vue'

import { date } from '@/utils/lg/date'
import { amount } from '@/utils/lg/amount'
import { invoiceType } from '@/utils/lg/invoiceType'
import { invoiceTileType } from '@/utils/lg/invoiceTileType'

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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
  const table = await getInvoiceApplyList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()
</script>

<style>
</style>
