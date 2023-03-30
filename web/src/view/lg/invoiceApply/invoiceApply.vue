<template>
  <div>
    <warning-bar title="注意：从请求开票到查询结果之间请间隔1-2分钟，如超过5分钟仍开票失败，请联系管理员" />
    <warning-bar title="注意：当前开票模式为一对一（一票一函）模式，请确认您已知悉再进行操作，如有疑问，请咨询管理员！！！" />
    <warning-bar title="注意：目前未做审批检查，请审批通过前一定点击“详情”查看每个订单是否都已经开票成功（审批拒绝可以忽略）！！！" />
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="发票抬头">
          <el-input v-model="searchInfo.invoiceTile" placeholder="搜索条件" clearable />
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
        <el-table-column align="center" label="操作" min-width="200" fixed="left">
          <template #default="scope">
            <el-button
              v-if="scope.row.auditStatus!==2 && scope.row.auditStatus!==3"
              type="success"
              icon="select"
              @click="approveInvoiceApplyFunc(scope.row)"
            >通过
            </el-button>
            <el-button
              v-if="scope.row.auditStatus!==2 && scope.row.auditStatus!==3"
              type="danger"
              icon="closeBold"
              @click="rejectInvoiceApplyFunc(scope.row)"
            >拒绝
            </el-button>
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
        <el-table-column align="center" label="总金额" min-width="120px">
          <template #default="scope">{{ amount(scope.row.invoiceTotalAmount) }}</template>
        </el-table-column>
        <el-table-column align="center" label="发票类型" min-width="120px">
          <template #default="scope">{{ invoiceType(scope.row.invoiceType) }}</template>
        </el-table-column>
        <el-table-column align="center" label="发票抬头类型" min-width="120px">
          <template #default="scope">{{ invoiceTileType(scope.row.invoiceTileType) }}</template>
        </el-table-column>
        <el-table-column align="center" label="发票抬头" prop="invoiceTile" min-width="280px" />
        <el-table-column align="center" label="税号" prop="taxNo" min-width="160px" />
        <el-table-column align="center" label="开户银行" prop="bankName" min-width="160px" />
        <el-table-column align="center" label="银行账号" prop="bankNo" min-width="160px" />
        <el-table-column align="center" label="企业地址" prop="companyAddress" min-width="280px" />
        <el-table-column align="center" label="企业电话" prop="companyTel" min-width="120px" />
        <el-table-column align="center" label="开票备注" prop="remarks" min-width="280px" />
        <el-table-column align="center" label="查看" min-width="100" fixed="right">
          <template #default="scope">
            <el-button
              type="info"
              icon="list"
              @click="openDetailDialog(scope.row)"
            >详情
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
    <el-dialog v-model="dialogDetailVisible" title="详情" :close-on-click-modal="false">
      <el-descriptions style="margin: 10px;" size="small" :column="1" border>
        <el-descriptions-item label="申请编号">{{ invoiceApplyData.applyNo }}</el-descriptions-item>
        <el-descriptions-item label="开票总金额">{{ invoiceApplyData.invoiceTotalAmount }}</el-descriptions-item>
        <el-descriptions-item label="发票类型">{{ invoiceApplyData.invoiceType === 'A1' ? '增值税普通发票' : '' }}</el-descriptions-item>
        <el-descriptions-item label="发票抬头类型">{{ invoiceApplyData.invoiceTileType === 'A1' ? '个人或事业单位' : invoiceApplyData.invoiceTileType === 'B1' ? '企业' : '' }}</el-descriptions-item>
        <el-descriptions-item label="发票抬头">{{ invoiceApplyData.invoiceTile }}</el-descriptions-item>
        <el-descriptions-item label="税号">{{ invoiceApplyData.taxNo }}</el-descriptions-item>
        <el-descriptions-item label="开户银行">{{ invoiceApplyData.bankName }}</el-descriptions-item>
        <el-descriptions-item label="银行账号">{{ invoiceApplyData.bankNo }}</el-descriptions-item>
        <el-descriptions-item label="企业地址">{{ invoiceApplyData.companyAddress }}</el-descriptions-item>
        <el-descriptions-item label="企业电话">{{ invoiceApplyData.companyTel }}</el-descriptions-item>
        <el-descriptions-item label="开票备注">{{ invoiceApplyData.remarks }}</el-descriptions-item>
        <el-descriptions-item label="订单列表">
          <el-table
            style="width: 100%; --el-table-border-color: none"
            :data="orderListData"
            row-key="orderNo"
            border
            size="small"
            table-layout="fixed"
            empty-text="无数据"
          >
            <el-table-column align="center" label="订单编号" prop="orderNo" width="120px" />
            <el-table-column align="center" label="支付金额" min-width="120px">
              <template #default="scope">{{ amount(scope.row.pay.payAmount) }}</template>
            </el-table-column>
            <el-table-column align="center" label="开票状态" min-width="80px">
              <template #default="scope">
                <el-tag
                  v-if="scope.row.invoice == null"
                  type="danger"
                  effect="dark"
                  size="large"
                >未开票
                </el-tag>
                <el-tag
                  v-if="scope.row.invoice != null"
                  type="success"
                  effect="dark"
                  size="large"
                >已开票
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column align="center" label="操作" width="200" fixed="right">
              <template #default="scope">
                <el-button
                  v-if="scope.row.invoice == null"
                  type="success"
                  icon="list"
                  @click="requestInvoiceFunc(scope.row)"
                >请求开票
                </el-button>
                <el-button
                  v-if="scope.row.invoice != null && (scope.row.invoice.invoiceDownloadUrl == null || scope.row.invoice.invoiceDownloadUrl == '')"
                  type="primary"
                  icon="list"
                  @click="queryInvoiceFunc(scope.row)"
                >查询结果
                </el-button>
                <el-button
                  v-if="scope.row.invoice != null && scope.row.invoice.invoiceDownloadUrl != null && scope.row.invoice.invoiceDownloadUrl != ''"
                  type="info"
                  icon="list"
                  @click="downloadInvoiceFunc(scope.row)"
                >查看发票
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getInvoiceApplyList,
  approveInvoiceApply,
  rejectInvoiceApply
} from '@/api/lg/invoiceApply'

import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

import { date } from '@/utils/lg/date'
import { amount } from '@/utils/lg/amount'
import { invoiceType } from '@/utils/lg/invoiceType'
import { invoiceTileType } from '@/utils/lg/invoiceTileType'
import { findOrderByNos, queryInvoice, requestInvoice } from '@/api/lg/order'
import WarningBar from '@/components/warningBar/warningBar.vue'

const invoiceApplyData = ref({})
const orderListData = ref({})

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

const orderNos = ref([])

const getOrdersData = async() => {
  const table = await findOrderByNos({ orderNos: orderNos.value })
  if (table.code === 0) {
    orderListData.value = table.data.orders
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

const dialogDetailVisible = ref(false)

const openDetailDialog = async(invoiceApply) => {
  invoiceApplyData.value = invoiceApply
  orderNos.value = []
  JSON.parse(invoiceApplyData.value.orderList).forEach((order) => {
    orderNos.value.push(order.orderNo)
  })
  await getOrdersData()
  dialogDetailVisible.value = true
}

const approveInvoiceApplyFunc = async(apply) => {
  ElMessageBox.confirm('确定要通过吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await approveInvoiceApply(apply)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '提交成功'
      })
      await getTableData()
    }
  })
}

const rejectInvoiceApplyFunc = async(apply) => {
  ElMessageBox.confirm('确定要拒绝吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await rejectInvoiceApply(apply)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '提交成功'
      })
      await getTableData()
    }
  })
}

const requestInvoiceFunc = async(order) => {
  const res = await requestInvoice({ order: order, invoiceApply: invoiceApplyData.value })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '提交成功'
    })
    await getOrdersData()
  }
}

const queryInvoiceFunc = async(order) => {
  const res = await queryInvoice({ order: order })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '提交成功'
    })
    await getOrdersData()
  }
}

const downloadInvoiceFunc = async(order) => {
  if (order.invoice != null && order.invoice.invoiceDownloadUrl != null) {
    window.open(order.invoice.invoiceDownloadUrl)
  }
}
</script>

<style>
</style>
