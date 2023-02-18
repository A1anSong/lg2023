<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="申请编号">
          <el-input v-model="searchInfo.applyNo" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="标段名称">
          <el-input v-model="searchInfo.projectName" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="申请企业">
          <el-input v-model="searchInfo.insureName" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="保函格式">
          <el-select v-model="searchInfo.elogTemplateId" clearable>
            <el-option v-for="template in templateData" :key="template.ID" :label="template.templateName" :value="template.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="保函编号">
          <el-input v-model="searchInfo.elogNo" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="订单状态">
          <el-select v-model="searchInfo.orderStatus" clearable>
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
          <el-select v-model="searchInfo.auditStatus" clearable>
            <el-option label="待审" value="1" />
            <el-option label="通过" value="2" />
            <el-option label="拒绝" value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="开标时间">
          <el-date-picker v-model="searchInfo.openBeginDate" type="daterange" start-placeholder="开始时间" end-placeholder="结束时间" />
        </el-form-item>
        <el-form-item label="申请时间">
          <el-date-picker v-model="searchInfo.applyCreatedAt" type="daterange" start-placeholder="开始时间" end-placeholder="结束时间" />
        </el-form-item>
        <el-form-item label="开函时间">
          <el-date-picker v-model="searchInfo.letterCreatedAt" type="daterange" start-placeholder="开始时间" end-placeholder="结束时间" />
        </el-form-item>
        <el-form-item label="担保期限" clearable>
          <el-input v-model.number="searchInfo.insureDay" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="success" icon="document" @click="exportExcel">导出excel</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" :data="tableData" row-key="ID" border size="small" table-layout="fixed" scrollbar-always-on @selection-change="handleSelectionChange">
        <el-table-column align="center" label="订单编号" prop="orderNo" width="120px" />
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
        <el-table-column align="center" label="申请时间" width="100px">
          <template #default="scope">{{ date(scope.row.apply.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="center" label="审核时间" width="100px">
          <template #default="scope">{{ date(scope.row.apply.auditDate) }}</template>
        </el-table-column>
        <el-table-column align="center" label="审核状态" min-width="80px">
          <template #default="scope">
            <el-tag :type="scope.row.revoke!=null?'info':auditType(scope.row.apply.auditStatus)" effect="dark" round>{{ scope.row.revoke != null ? '已撤' : auditStatus(scope.row.apply.auditStatus) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" label="付款时间" width="100px">
          <template #default="scope">{{ scope.row.pay?date(scope.row.pay.payTime):'' }}</template>
        </el-table-column>
        <el-table-column align="center" label="付款金额" min-width="120px">
          <template #default="scope">{{ scope.row.pay != null ? amount(scope.row.pay.payAmount) : '' }}</template>
        </el-table-column>
        <el-table-column align="center" label="付款状态" min-width="80px">
          <template #default="scope">
            <el-tag :type="scope.row.pay != null ? 'success' : 'info'" effect="dark" round>{{ scope.row.pay != null ? "已付" : "未付" }}</el-tag>
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
        <el-table-column align="center" label="业务员" prop="employeeNo" min-width="120px" />
        <el-table-column align="center" label="查看" min-width="360" fixed="right">
          <template #default="scope">
            <el-button type="info" icon="list" size="small" @click="openDetailDialog(scope.row)">详情</el-button>
            <el-button type="primary" icon="paperclip" size="small" :disabled="scope.row.apply.attachInfo == null || scope.row.apply.attachInfo === '' || scope.row.apply.attachInfo === '[]'" @click="openAttachDialog(scope.row.apply.attachInfo)">附件</el-button>
            <el-button type="success" icon="printer" size="small" :disabled="scope.row.letter == null" @click="downloadLetterFile(scope.row)">保函</el-button>
            <el-button type="warning" icon="box" size="small" :disabled="scope.row.letter == null" @click="downloadLetterEncryptFile(scope.row)">密文</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-card shadow="always" style="float: left">
        累计成功(未理赔，未退函)担保金额为
        <span style="color: red">{{ statisticData.totalGuaranteeAmount.toFixed(2).replace(/(\d)(?=(\d{3})+\.)/g, '$1,') }}</span>
        元，收取保函费用为
        <span style="color: red">{{ statisticData.totalElogAmount.toFixed(2).replace(/(\d)(?=(\d{3})+\.)/g, '$1,') }}</span>
        元
      </el-card>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange" @size-change="handleSizeChange"/>
      </div>
    </div>
    <el-dialog v-model="dialogDetailVisble" title="详情">
      <el-descriptions style="margin: 10px;" title="订单信息" size="small" :column="1" border>
        <el-descriptions-item label="申请编号">{{ orderDetailData.apply.applyNo }}</el-descriptions-item>
        <el-descriptions-item label="产品类型">{{ productType(orderDetailData.apply.productType) }}</el-descriptions-item>
        <el-descriptions-item label="标段名称">{{ orderDetailData.apply.projectName }}</el-descriptions-item>
        <el-descriptions-item label="标段编号">{{ orderDetailData.apply.projectNo }}</el-descriptions-item>
        <el-descriptions-item label="担保金额">{{ amount(orderDetailData.apply.tenderDeposit) }}</el-descriptions-item>
        <el-descriptions-item label="保证金缴纳开始时间">{{ orderDetailData.apply.depositStartDate }}</el-descriptions-item>
        <el-descriptions-item label="保证金缴纳截止时间">{{ orderDetailData.apply.depositEndDate }}</el-descriptions-item>
        <el-descriptions-item label="开标时间">{{ orderDetailData.apply.openBeginDate }}</el-descriptions-item>
        <el-descriptions-item label="保函格式编号">{{ orderDetailData.apply.elogTemplateNo }}</el-descriptions-item>
        <el-descriptions-item label="保函格式名称">{{ orderDetailData.apply.elogTemplateName }}</el-descriptions-item>
        <el-descriptions-item label="所属市">{{ orderDetailData.project != null && orderDetailData.project.projectCity !== undefined ? orderDetailData.project.projectCity : '' }}</el-descriptions-item>
        <el-descriptions-item label="所属县">{{ orderDetailData.project != null && orderDetailData.projectCounty !== undefined ? orderDetailData.project.projectCounty : '' }}</el-descriptions-item>
        <el-descriptions-item label="受益人名称">{{ orderDetailData.apply.insuredName }}</el-descriptions-item>
        <el-descriptions-item label="受益人社会信用代码">{{ orderDetailData.apply.insuredCreditCode }}</el-descriptions-item>
        <el-descriptions-item label="受益人地址">{{ orderDetailData.apply.insuredAddress != null ? orderDetailData.apply.insuredAddress : '' }}</el-descriptions-item>
        <el-descriptions-item label="投保方名称">{{ orderDetailData.apply.insureName }}</el-descriptions-item>
        <el-descriptions-item label="投保方社会信用代码">{{ orderDetailData.apply.insureCreditCode }}</el-descriptions-item>
        <el-descriptions-item label="投保方法人姓名">{{ orderDetailData.apply.insureLegalName }}</el-descriptions-item>
        <el-descriptions-item label="投保方法人身份证号">{{ orderDetailData.apply.insureLegalIdCard }}</el-descriptions-item>
        <el-descriptions-item label="投保方地址">{{ orderDetailData.insureAddress }}</el-descriptions-item>
        <el-descriptions-item label="经办人姓名">{{ orderDetailData.applicantName }}</el-descriptions-item>
        <el-descriptions-item label="经办人身份证号">{{ orderDetailData.apply.applicantIdCard }}</el-descriptions-item>
        <el-descriptions-item label="经办人联系电话">{{ orderDetailData.apply.applicantTel }}</el-descriptions-item>
      </el-descriptions>
      <el-descriptions v-if="orderDetailData.pay" style="margin: 10px;" title="支付信息" size="small" :column="1" border>
        <el-descriptions-item label="支付编号" width="50px">{{ orderDetailData.pay.payNo }}</el-descriptions-item>
        <el-descriptions-item label="支付支付流水号">{{ orderDetailData.pay.payTransNo }}</el-descriptions-item>
      </el-descriptions>
      <el-descriptions v-if="orderStatus(orderDetailData)==='已开'" style="margin: 10px;" title="保函信息" size="small" :column="1" border>
        <el-descriptions-item label="保函编号" width="50px">{{ orderDetailData.letter.elogNo }}</el-descriptions-item>
        <el-descriptions-item label="担保期限（天）">{{ orderDetailData.letter.insureDay }}</el-descriptions-item>
        <el-descriptions-item label="验真码">{{ orderDetailData.letter.validateCode }}</el-descriptions-item>
      </el-descriptions>
      <el-descriptions v-if="orderStatus(orderDetailData)==='延期'" style="margin: 10px;" title="延期信息" size="small" :column="1" border>
        <el-descriptions-item label="延期" width="50px">延期信息展示（请将需要显示的信息提交管理员进行更新）</el-descriptions-item>
      </el-descriptions>
      <el-descriptions v-if="orderStatus(orderDetailData)==='退函'" style="margin: 10px;" title="退函信息" size="small" :column="1" border>
        <el-descriptions-item label="退函" width="50px">退函信息展示（请将需要显示的信息提交管理员进行更新）</el-descriptions-item>
      </el-descriptions>
      <el-descriptions v-if="orderStatus(orderDetailData)==='理赔'" style="margin: 10px;" title="理赔信息" size="small" :column="1" border>
        <el-descriptions-item label="理赔" width="50px">理赔信息展示（请将需要显示的信息提交管理员进行更新）</el-descriptions-item>
      </el-descriptions>
      <el-descriptions v-if="orderStatus(orderDetailData)==='销函'" style="margin: 10px;" title="销函信息" size="small" :column="1" border>
        <el-descriptions-item label="销函" width="50px">销函信息展示（请将需要显示的信息提交管理员进行更新）</el-descriptions-item>
      </el-descriptions>
    </el-dialog>
    <el-dialog v-model="dialogAttachVisble" title="附件">
      <el-descriptions style="margin: 10px;" size="small" :column="1" border>
        <el-descriptions-item v-for="attachInfo in JSON.parse(attachInfoData)" :key="attachInfo.attachName" :label="attachType(attachInfo.attachType)">
          <el-link type="primary" :href="attachInfo.attachUrl">{{ attachInfo.attachName }}</el-link>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getOrderList,
  getOrderStatisticData, downloadExcel
} from '@/api/lgjx/testOrder'

import { ref } from 'vue'

import { date } from '@/utils/jxlg/date'
import { auditStatus, auditType } from '@/utils/jxlg/auditStatus'
import { productType } from '@/utils/jxlg/productType'
import { amount } from '@/utils/jxlg/amount'
import { attachType } from '@/utils/jxlg/attachType'
import { orderStatus, orderStatusType } from '@/utils/jxlg/orderStatus'
import { getTemplateList } from '@/api/lgjx/testTemplate'
import { downloadFile } from '@/api/lgjx/testFile'

const orderDetailData = ref({})
const attachInfoData = ref({})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const templateData = ref([])
const statisticData = ref({
  totalGuaranteeAmount: 0.0,
  totalElogAmount: 0.0,
})

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
  const statistic = await getOrderStatisticData()
  if (statistic.code === 0) {
    statisticData.value = statistic.data.orderStatisticData
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

const dialogDetailVisble = ref(false)
const dialogAttachVisble = ref(false)

const openDetailDialog = (order) => {
  orderDetailData.value = order
  dialogDetailVisble.value = true
}

const openAttachDialog = (attachInfo) => {
  attachInfoData.value = attachInfo
  dialogAttachVisble.value = true
}

// 导出excel
const exportExcel = async() => {
  await downloadExcel({ page: 1, pageSize: 9999, ...searchInfo.value })
}

const downloadLetterFile = (order) => {
  if (order.delay.elogFile !== null) {
    downloadFile(order.delay.elogFile)
  } else {
    downloadFile(order.letter.elogFile)
  }
}

const downloadLetterEncryptFile = (order) => {
  if (order.delay.elogEncryptFile !== null) {
    downloadFile(order.delay.elogEncryptFile)
  } else {
    downloadFile(order.letter.elogEncryptFile)
  }
}
</script>

<style>
</style>
