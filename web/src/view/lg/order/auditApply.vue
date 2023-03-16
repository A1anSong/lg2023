<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="申请编号">
          <el-input v-model="searchInfo.applyNo" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="标段编号">
          <el-input v-model="searchInfo.projectNo" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="标段名称">
          <el-input v-model="searchInfo.projectName" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="申请企业">
          <el-input v-model="searchInfo.insureName" placeholder="搜索条件" clearable />
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
          <el-input v-model="searchInfo.elogNo" placeholder="搜索条件" clearable />
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
        <el-form-item v-auth="btnAuth.all" label="工号">
          <el-input v-model="searchInfo.authCode" placeholder="搜索条件" clearable />
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
        <el-table-column v-auth="btnAuth.operation" align="center" label="操作" min-width="200" fixed="left">
          <template #default="scope">
            <el-tag
              v-if="scope.row.revoke"
              type="warning"
              effect="dark"
              size="large"
            >已撤单
            </el-tag>
            <el-button
              v-if="scope.row.revoke === null && scope.row.apply.auditStatus===1 && scope.row.project"
              type="success"
              icon="select"
              @click="approveApplyFunc(scope.row)"
            >通过
            </el-button>
            <el-button
              v-if="scope.row.revoke === null &&scope.row.apply.auditStatus===1 && scope.row.project"
              type="danger"
              icon="closeBold"
              @click="rejectApplyFunc(scope.row)"
            >拒绝
            </el-button>
            <el-tag
              v-if="scope.row.revoke === null && scope.row.apply.auditStatus===2"
              type="success"
              effect="dark"
              size="large"
            >已审批通过
            </el-tag>
            <el-tag
              v-if="scope.row.revoke === null && scope.row.apply.auditStatus===3"
              type="danger"
              effect="dark"
              size="large"
            >已审批拒绝
            </el-tag>
            <el-tag
              v-if="scope.row.revoke === null && scope.row.project === null && scope.row.apply.auditStatus===1"
              type="info"
              effect="dark"
              size="large"
            >待绑定项目才后可审核
            </el-tag>
          </template>
        </el-table-column>
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
        <el-table-column align="center" label="业务员" prop="employeeNo" min-width="120px" />
        <el-table-column v-if="authCheck !== 0" align="center" label="查看" :min-width="authCheck * 100" fixed="right">
          <template #default="scope">
            <el-button v-auth="btnAuth.detail" type="info" icon="list" @click="openDetailDialog(scope.row)">详情</el-button>
            <el-button
              v-auth="btnAuth.attach"
              type="primary"
              icon="paperclip"
              :disabled="scope.row.apply.attachInfo == null || scope.row.apply.attachInfo === '' || scope.row.apply.attachInfo === '[]'"
              @click="openAttachDialog(scope.row.apply.attachInfo)"
            >附件
            </el-button>
            <el-button
              v-auth="btnAuth.elog"
              type="success"
              icon="printer"
              :disabled="scope.row.letter == null"
              @click="downloadLetterFile(scope.row)"
            >保函
            </el-button>
            <el-button
              v-auth="btnAuth.encrypt"
              type="warning"
              icon="box"
              :disabled="scope.row.letter == null"
              @click="downloadLetterEncryptFile(scope.row)"
            >密文
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
      <el-descriptions style="margin: 10px;" title="订单信息" size="small" :column="1" border>
        <el-descriptions-item label="申请编号">{{ orderDetailData.apply.applyNo }}</el-descriptions-item>
        <el-descriptions-item label="产品类型">{{
          productType(orderDetailData.apply.productType)
        }}
        </el-descriptions-item>
        <el-descriptions-item label="标段名称">{{ orderDetailData.apply.projectName }}</el-descriptions-item>
        <el-descriptions-item label="标段编号">{{ orderDetailData.apply.projectNo }}</el-descriptions-item>
        <el-descriptions-item label="担保金额">{{ amount(orderDetailData.apply.tenderDeposit) }}</el-descriptions-item>
        <el-descriptions-item label="保证金缴纳开始时间">{{
          orderDetailData.apply.depositStartDate
        }}
        </el-descriptions-item>
        <el-descriptions-item label="保证金缴纳截止时间">{{
          orderDetailData.apply.depositEndDate
        }}
        </el-descriptions-item>
        <el-descriptions-item label="开标时间">{{ orderDetailData.apply.openBeginDate }}</el-descriptions-item>
        <el-descriptions-item label="保函格式编号">{{ orderDetailData.apply.elogTemplateNo }}</el-descriptions-item>
        <el-descriptions-item label="保函格式名称">{{ orderDetailData.apply.elogTemplateName }}</el-descriptions-item>
        <el-descriptions-item label="所属市">{{
          orderDetailData.project != null && orderDetailData.project.projectCity !== undefined ? orderDetailData.project.projectCity : ''
        }}
        </el-descriptions-item>
        <el-descriptions-item label="所属县">{{
          orderDetailData.project != null && orderDetailData.projectCounty !== undefined ? orderDetailData.project.projectCounty : ''
        }}
        </el-descriptions-item>
        <el-descriptions-item label="受益人名称">{{ orderDetailData.apply.insuredName }}</el-descriptions-item>
        <el-descriptions-item label="受益人社会信用代码">{{
          orderDetailData.apply.insuredCreditCode
        }}
        </el-descriptions-item>
        <el-descriptions-item label="受益人地址">
          {{ orderDetailData.apply.insuredAddress != null ? orderDetailData.apply.insuredAddress : '' }}
        </el-descriptions-item>
        <el-descriptions-item label="投保方名称">{{ orderDetailData.apply.insureName }}</el-descriptions-item>
        <el-descriptions-item label="投保方社会信用代码">{{
          orderDetailData.apply.insureCreditCode
        }}
        </el-descriptions-item>
        <el-descriptions-item label="投保方法人姓名">{{ orderDetailData.apply.insureLegalName }}</el-descriptions-item>
        <el-descriptions-item label="投保方法人身份证号">{{
          orderDetailData.apply.insureLegalIdCard
        }}
        </el-descriptions-item>
        <el-descriptions-item label="投保方地址">{{ orderDetailData.insureAddress }}</el-descriptions-item>
        <el-descriptions-item label="经办人姓名">{{ orderDetailData.applicantName }}</el-descriptions-item>
        <el-descriptions-item label="经办人身份证号">{{ orderDetailData.apply.applicantIdCard }}</el-descriptions-item>
        <el-descriptions-item label="经办人联系电话">{{ orderDetailData.apply.applicantTel }}</el-descriptions-item>
      </el-descriptions>
      <el-descriptions
        v-if="orderDetailData.pay"
        style="margin: 10px;"
        title="支付信息"
        size="small"
        :column="1"
        border
      >
        <el-descriptions-item label="支付编号" width="50px">{{ orderDetailData.pay.payNo }}</el-descriptions-item>
        <el-descriptions-item label="支付支付流水号">{{ orderDetailData.pay.payTransNo }}</el-descriptions-item>
      </el-descriptions>
      <el-descriptions
        v-if="orderStatus(orderDetailData)==='已开'"
        style="margin: 10px;"
        title="保函信息"
        size="small"
        :column="1"
        border
      >
        <el-descriptions-item label="保函编号" width="50px">{{ orderDetailData.letter.elogNo }}</el-descriptions-item>
        <el-descriptions-item label="担保期限（天）">{{ orderDetailData.letter.insureDay }}</el-descriptions-item>
        <el-descriptions-item label="验真码">{{ orderDetailData.letter.validateCode }}</el-descriptions-item>
      </el-descriptions>
      <el-descriptions
        v-if="orderStatus(orderDetailData)==='延期'"
        style="margin: 10px;"
        title="延期信息"
        size="small"
        :column="1"
        border
      >
        <el-descriptions-item label="延期" width="50px">延期信息展示（请将需要显示的信息提交管理员进行更新）
        </el-descriptions-item>
      </el-descriptions>
      <el-descriptions
        v-if="orderStatus(orderDetailData)==='退函'"
        style="margin: 10px;"
        title="退函信息"
        size="small"
        :column="1"
        border
      >
        <el-descriptions-item label="退函" width="50px">退函信息展示（请将需要显示的信息提交管理员进行更新）
        </el-descriptions-item>
      </el-descriptions>
      <el-descriptions
        v-if="orderStatus(orderDetailData)==='理赔'"
        style="margin: 10px;"
        title="理赔信息"
        size="small"
        :column="1"
        border
      >
        <el-descriptions-item label="理赔" width="50px">理赔信息展示（请将需要显示的信息提交管理员进行更新）
        </el-descriptions-item>
      </el-descriptions>
      <el-descriptions
        v-if="orderStatus(orderDetailData)==='销函'"
        style="margin: 10px;"
        title="销函信息"
        size="small"
        :column="1"
        border
      >
        <el-descriptions-item label="销函" width="50px">销函信息展示（请将需要显示的信息提交管理员进行更新）
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
    <el-dialog v-model="dialogAttachVisible" title="附件" :close-on-click-modal="false">
      <el-descriptions style="margin: 10px;" size="small" :column="1" border>
        <el-descriptions-item
          v-for="attachInfo in JSON.parse(attachInfoData)"
          :key="attachInfo.attachName"
          :label="attachType(attachInfo.attachType)"
        >
          <el-link type="primary" :href="attachInfo.attachUrl">{{ attachInfo.attachName }}</el-link>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getOrderList,
  approveApply,
  rejectApply
} from '@/api/lg/order'

import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, computed } from 'vue'

import { date } from '@/utils/lg/date'
import { auditStatus, auditType } from '@/utils/lg/auditStatus'
import { productType } from '@/utils/lg/productType'
import { amount } from '@/utils/lg/amount'
import { attachType } from '@/utils/lg/attachType'
import { orderStatus, orderStatusType } from '@/utils/lg/orderStatus'
import { getTemplateList } from '@/api/lg/template'
import { downloadFile } from '@/api/lg/file'

import { useUserStore } from '@/pinia/modules/user'
import { useBtnAuth } from '@/utils/btnAuth'

const orderDetailData = ref({})
const attachInfoData = ref({})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const templateData = ref([])
const userStore = useUserStore()
const userInfo = userStore.userInfo
const btnAuth = useBtnAuth()
const authDetail = ref(false)
const authAttach = ref(false)
const authElog = ref(false)
const authEncrypt = ref(false)
const authCheck = computed(() => {
  return (authDetail.value === true ? 1 : 0) + (authAttach.value === true ? 1 : 0) + (authElog.value === true ? 1 : 0) + (authEncrypt.value === true ? 1 : 0)
})

// 重置
const onReset = () => {
  searchInfo.value = {}
  checkAuthorityAll()
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

const checkAuthorityDetail = () => {
  const authority = btnAuth.detail
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
    return
  }
  const waitUse = authority.toString().split(',')
  const flag = waitUse.some(item => item === userInfo.authorityId.toString())
  if (flag) {
    authDetail.value = true
  }
}

const checkAuthorityAttach = () => {
  const authority = btnAuth.attach
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
    return
  }
  const waitUse = authority.toString().split(',')
  const flag = waitUse.some(item => item === userInfo.authorityId.toString())
  if (flag) {
    authAttach.value = true
  }
}

const checkAuthorityElog = () => {
  const authority = btnAuth.elog
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
    return
  }
  const waitUse = authority.toString().split(',')
  const flag = waitUse.some(item => item === userInfo.authorityId.toString())
  if (flag) {
    authElog.value = true
  }
}

const checkAuthorityEncrypt = () => {
  const authority = btnAuth.encrypt
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
    return
  }
  const waitUse = authority.toString().split(',')
  const flag = waitUse.some(item => item === userInfo.authorityId.toString())
  if (flag) {
    authEncrypt.value = true
  }
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

checkAuthorityAll()
checkAuthorityDetail()
checkAuthorityAttach()
checkAuthorityElog()
checkAuthorityEncrypt()
getTemplateData()
getTableData()

// ============== 表格控制部分结束 ===============

const dialogDetailVisible = ref(false)
const dialogAttachVisible = ref(false)

const openDetailDialog = (order) => {
  orderDetailData.value = order
  dialogDetailVisible.value = true
}

const openAttachDialog = (attachInfo) => {
  attachInfoData.value = attachInfo
  dialogAttachVisible.value = true
}

const downloadLetterFile = (order) => {
  if (order.delay_id != null && order.delay.elogFile != null) {
    downloadFile(order.delay.elogFile)
  } else {
    downloadFile(order.letter.elogFile)
  }
}

const downloadLetterEncryptFile = (order) => {
  if (order.delay_id != null && order.delay.elogEncryptFile != null) {
    downloadFile(order.delay.elogEncryptFile)
  } else {
    downloadFile(order.letter.elogEncryptFile)
  }
}

const approveApplyFunc = async(apply) => {
  ElMessageBox.confirm('确定要通过吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await approveApply(apply)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '提交成功'
      })
      await getTableData()
    }
  })
}

const rejectApplyFunc = async(apply) => {
  ElMessageBox.confirm('确定要拒绝吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await rejectApply(apply)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '提交成功'
      })
      await getTableData()
    }
  })
}
</script>

<style>
</style>
