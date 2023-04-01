<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="订单编号">
          <el-input v-model="searchInfo.orderNo" placeholder="搜索条件" clearable />
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
        <el-table-column align="center" label="发票编号" prop="invoice.invoiceNo" width="120px" />
        <el-table-column align="center" label="发票类型" width="120px">
          <template #default="scope">{{ invoiceType(scope.row.invoice.invoiceType) }}</template>
        </el-table-column>
        <el-table-column align="center" label="开票方式" width="120px">
          <template #default="scope">{{ invoiceForm(scope.row.invoice.invoiceForm) }}</template>
        </el-table-column>
        <el-table-column align="center" label="发票金额" min-width="120px">
          <template #default="scope">{{ amount(scope.row.invoice.invoiceAmount) }}</template>
        </el-table-column>
        <el-table-column align="center" label="发票抬头类型" width="120px">
          <template #default="scope">{{ invoiceTileType(scope.row.invoice.invoiceTileType) }}</template>
        </el-table-column>
        <el-table-column align="center" label="发票抬头" prop="invoice.InvoiceTile" width="120px" />
        <el-table-column align="center" label="税号" prop="invoice.TaxNo" width="120px" />
        <el-table-column align="center" label="开户银行" prop="invoice.BankName" width="120px" />
        <el-table-column align="center" label="银行账号" prop="invoice.BankNo" width="120px" />
        <el-table-column align="center" label="企业地址" prop="invoice.CompanyAddress" width="120px" />
        <el-table-column align="center" label="企业电话" prop="invoice.CompanyTel" width="120px" />
        <el-table-column align="center" label="发票备注" prop="invoice.Remarks" width="120px" />
        <el-table-column align="center" label="发票内容" prop="invoice.invoiceContent" width="120px" />
        <el-table-column align="center" label="开票时间" prop="invoice.invoiceTime" width="100px" />
        <el-table-column align="center" label="查看" :min-width="100" fixed="right">
          <template #default="scope">
            <el-tag
              v-if="scope.row.invoice.invoiceDownloadUrl == null"
              type="warning"
              effect="dark"
              size="large"
            >未更新发票
            </el-tag>
            <el-button
              v-if="scope.row.invoice.invoiceDownloadUrl != null"
              type="primary"
              icon="list"
              @click="downloadInvoiceFunc(scope.row)"
            >发票
            </el-button>
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
  name: 'InvoiceList'
}
</script>

<script setup>
import { getOrderList } from '@/api/lg/order'

// 全量引入格式化工具 请按需保留
import { ref } from 'vue'

import { amount } from '@/utils/lg/amount'
import { invoiceType } from '@/utils/lg/invoiceType'
import { invoiceTileType } from '@/utils/lg/invoiceTileType'
import { invoiceForm } from '@/utils/lg/invoiceForm'

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({ onlyInvoice: true })

// 重置
const onReset = () => {
  searchInfo.value = { onlyInvoice: true }
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

getTableData()

// ============== 表格控制部分结束 ===============
const downloadInvoiceFunc = async(order) => {
  if (order.invoice != null && order.invoice.invoiceDownloadUrl != null) {
    window.open(order.invoice.invoiceDownloadUrl)
  }
}
</script>

<style>
</style>
