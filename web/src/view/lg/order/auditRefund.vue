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
          <el-select v-model="searchInfo.orderStatus" clearable>
            <el-option
              value="已撤"
            />
            <el-option
              value="未开"
            />
            <el-option
              value="已开"
            />
            <el-option
              value="延期"
            />
            <el-option
              value="退函"
            />
            <el-option
              value="理赔"
            />
            <el-option
              value="销函"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="审核状态">
          <el-select v-model="searchInfo.auditStatus" clearable>
            <el-option
              label="待审"
              value="1"
            />
            <el-option
              label="通过"
              value="2"
            />
            <el-option
              label="拒绝"
              value="3"
            />
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
        <el-form-item label="担保期限">
          <el-input v-model.number="searchInfo.insureDay" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <!--      <div class="gva-btn-list">-->
      <!--        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>-->
      <!--        <el-popover v-model:visible="deleteVisible" placement="top" width="160">-->
      <!--          <p>确定要删除吗？</p>-->
      <!--          <div style="text-align: right; margin-top: 8px;">-->
      <!--            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>-->
      <!--            <el-button size="small" type="primary" @click="onDelete">确定</el-button>-->
      <!--          </div>-->
      <!--          <template #reference>-->
      <!--            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>-->
      <!--          </template>-->
      <!--        </el-popover>-->
      <!--      </div>-->
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="light"
        :data="tableData"
        row-key="ID"
        border
        size="small"
        table-layout="fixed"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="expand" label="详情">
          <template #default="scope">
            <div>
              <el-row :gutter="20">
                <el-col :xs="24" :sm="12">
                  <el-descriptions
                    style="margin: 10px;"
                    title="订单信息"
                    size="small"
                    :column="2"
                    border
                  >
                    <el-descriptions-item label="申请编号">{{ scope.row.apply.applyNo }}</el-descriptions-item>
                    <el-descriptions-item label="产品类型">{{
                      productType(scope.row.apply.productType)
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="标段名称">{{ scope.row.apply.projectName }}</el-descriptions-item>
                    <el-descriptions-item label="标段编号">{{ scope.row.apply.projectNo }}</el-descriptions-item>
                    <el-descriptions-item label="担保金额">{{
                      amount(scope.row.apply.tenderDeposit)
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="保证金缴纳开始时间">{{
                      scope.row.apply.depositStartDate
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="保证金缴纳截止时间">{{
                      scope.row.apply.depositEndDate
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="开标时间">{{ scope.row.apply.openBeginDate }}</el-descriptions-item>
                    <el-descriptions-item label="保函格式编号">{{
                      scope.row.apply.elogTemplateNo
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="保函格式名称">{{
                      scope.row.apply.elogTemplateName
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="所属市">{{
                      scope.row.project != null && scope.row.project.city !== undefined ? scope.row.project.city : ''
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="所属县">{{
                      scope.row.project != null && scope.row.project.county !== undefined ? scope.row.project.county : ''
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="受益人名称">{{ scope.row.apply.insuredName }}</el-descriptions-item>
                    <el-descriptions-item label="受益人社会信用代码">{{
                      scope.row.apply.insuredCreditCode
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="受益人地址">
                      {{ scope.row.apply.insuredAddress != null ? scope.row.apply.insuredAddress : '' }}
                    </el-descriptions-item>
                    <el-descriptions-item label="投保方名称">{{ scope.row.apply.insureName }}</el-descriptions-item>
                    <el-descriptions-item label="投保方社会信用代码">{{
                      scope.row.apply.insureCreditCode
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="投保方法人姓名">{{
                      scope.row.apply.insureLegalName
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="投保方法人身份证号">{{
                      scope.row.apply.insureLegalIdCard
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="投保方地址">{{ scope.row.apply.insureAddress }}</el-descriptions-item>
                    <el-descriptions-item label="经办人姓名">{{ scope.row.apply.applicantName }}</el-descriptions-item>
                    <el-descriptions-item label="经办人身份证号">{{
                      scope.row.apply.applicantIdCard
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="经办人联系电话">{{
                      scope.row.apply.applicantTel
                    }}
                    </el-descriptions-item>
                  </el-descriptions>
                </el-col>
                <el-col :xs="24" :sm="12">
                  <el-descriptions
                    v-if="scope.row.apply.attachInfo"
                    style="margin: 10px;"
                    title="附件信息"
                    size="small"
                    :column="3"
                    border
                  >
                    <template v-for="attachInfo in JSON.parse(scope.row.apply.attachInfo)">
                      <el-descriptions-item :label="attachType(attachInfo.attachType)">
                        <el-link type="primary"><a
                          :href="attachInfo.attachUrl"
                          target="_blank"
                        >{{ attachInfo.attachName }}</a>
                        </el-link>
                      </el-descriptions-item>
                    </template>
                  </el-descriptions>
                  <el-descriptions
                    v-if="scope.row.pay"
                    style="margin: 10px;"
                    title="支付信息"
                    size="small"
                    :column="1"
                    border
                  >
                    <el-descriptions-item label="支付编号" width="50px">{{ scope.row.pay.payNo }}</el-descriptions-item>
                    <el-descriptions-item label="支付支付流水号">{{ scope.row.pay.payTransNo }}</el-descriptions-item>
                  </el-descriptions>
                  <el-descriptions
                    v-if="orderStatus(scope.row)==='已开'"
                    style="margin: 10px;"
                    title="保函信息"
                    size="small"
                    :column="1"
                    border
                  >
                    <el-descriptions-item label="保函编号" width="50px">{{
                      scope.row.letter.elogNo
                    }}
                    </el-descriptions-item>
                    <el-descriptions-item label="保函文件">{{ scope.row.letter.elogUrl }}</el-descriptions-item>
                    <el-descriptions-item label="担保期限（天）">{{ scope.row.letter.insureDay }}</el-descriptions-item>
                    <el-descriptions-item label="验真码">{{ scope.row.letter.validateCode }}</el-descriptions-item>
                  </el-descriptions>
                  <el-descriptions
                    v-if="orderStatus(scope.row)==='延期'"
                    style="margin: 10px;"
                    title="延期信息"
                    size="small"
                    :column="1"
                    border
                  >
                    <el-descriptions-item label="延期" width="50px">延期</el-descriptions-item>
                  </el-descriptions>
                  <el-descriptions
                    v-if="orderStatus(scope.row)==='退函'"
                    style="margin: 10px;"
                    title="退函信息"
                    size="small"
                    :column="1"
                    border
                  >
                    <el-descriptions-item label="退函" width="50px">退函</el-descriptions-item>
                  </el-descriptions>
                  <el-descriptions
                    v-if="orderStatus(scope.row)==='理赔'"
                    style="margin: 10px;"
                    title="理赔信息"
                    size="small"
                    :column="1"
                    border
                  >
                    <el-descriptions-item label="理赔" width="50px">理赔</el-descriptions-item>
                  </el-descriptions>
                  <el-descriptions
                    v-if="orderStatus(scope.row)==='销函'"
                    style="margin: 10px;"
                    title="销函信息"
                    size="small"
                    :column="1"
                    border
                  >
                    <el-descriptions-item label="销函" width="50px">销函</el-descriptions-item>
                  </el-descriptions>
                </el-col>
              </el-row>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="申请编号" prop="apply.applyNo" min-width="160px" />
        <el-table-column align="center" label="来源" min-width="80px">
          <template #default="scope">{{ productType(scope.row.apply.productType) }}</template>
        </el-table-column>
        <el-table-column align="center" label="标段信息">
          <el-table-column align="center" label="标段名称" prop="apply.projectName" min-width="300px" />
          <el-table-column align="center" label="开标时间" prop="apply.openBeginDate" min-width="160px" />
          <el-table-column align="center" label="标段编号" prop="apply.projectNo" min-width="160px" />
          <el-table-column align="center" label="受益方名称" prop="apply.insuredName" min-width="280px" />
          <el-table-column align="center" label="担保金额" min-width="120px">
            <template #default="scope">{{ amount(scope.row.apply.tenderDeposit) }}</template>
          </el-table-column>
          <el-table-column align="center" label="所属市" prop="project" min-width="120px" />
          <el-table-column align="center" label="保函格式名称" prop="apply.elogTemplateName" min-width="120px" />
        </el-table-column>
        <el-table-column align="center" label="申请人信息">
          <el-table-column align="center" label="申请企业" prop="apply.insureName" min-width="280px" />
          <el-table-column align="center" label="申请时间" min-width="160px">
            <template #default="scope">{{ date(scope.row.apply.CreatedAt) }}</template>
          </el-table-column>
        </el-table-column>
        <el-table-column align="center" label="订单信息">
          <el-table-column align="center" label="审核时间" prop="apply.auditDate" min-width="160px" />
          <el-table-column align="center" label="审核状态" min-width="80px">
            <template #default="scope">
              <el-tag
                :type="scope.row.revoke!=null?'info':auditType(scope.row.apply.auditStatus)"
                effect="dark"
                round
              >
                {{ scope.row.revoke != null ? '已撤' : auditStatus(scope.row.apply.auditStatus) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column align="center" label="付款时间" prop="pay.payTime" min-width="160px" />
          <el-table-column align="center" label="付款金额" min-width="120px">
            <template #default="scope">{{ scope.row.pay != null ? amount(scope.row.pay.payAmount) : '' }}</template>
          </el-table-column>
          <el-table-column align="center" label="付款状态" min-width="80px">
            <template #default="scope">
              <el-tag
                :type="scope.row.pay != null ? 'success' : 'info'"
                effect="dark"
                round
              >

                {{ scope.row.pay != null ? "已付" : "未付" }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column align="center" label="开函时间" min-width="160px">
            <template #default="scope">{{
              scope.row.letter !== null ? date(scope.row.letter.CreatedAt) : ''
            }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="订单状态" min-width="80px">
            <template #default="scope">
              <el-tag
                :type="orderStatusType(scope.row)"
                effect="dark"
                round
              >
                {{ orderStatus(scope.row) }}
              </el-tag>
            </template>
          </el-table-column>
        </el-table-column>
        <el-table-column align="center" label="推荐人">
          <el-table-column align="center" label="工号" prop="employeeNo" min-width="120px" />
          <el-table-column align="center" label="业务员" prop="employeeNo" min-width="120px" />
        </el-table-column>
        <el-table-column align="center" label="操作" min-width="200" fixed="right">
          <template #default="scope">
            <el-button
              v-if="scope.row.refund.auditStatus!==2&&scope.row.refund.auditStatus!==3"
              type="success"
              icon="select"
              size="small"
              @click="approveRefundFunc(scope.row)"
            >通过
            </el-button>
            <el-button
              v-if="scope.row.refund.auditStatus!==2&&scope.row.refund.auditStatus!==3"
              type="danger"
              icon="closeBold"
              size="small"
              @click="rejectRefundFunc(scope.row)"
            >拒绝
            </el-button>
            <el-tag
              v-if="scope.row.refund.auditStatus===2"
              type="success"
              effect="dark"
              size="large"
            >已审批通过
            </el-tag>
            <el-tag
              v-if="scope.row.refund.auditStatus===3"
              type="danger"
              effect="dark"
              size="large"
            >已审批拒绝
            </el-tag>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单编号:" prop="orderNo">
          <el-input v-model="formData.orderNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'TestRefundDelay'
}
</script>

<script setup>
import {
  createOrder,
  deleteOrder,
  deleteOrderByIds,
  updateOrder,
  findOrder,
  getOrderList,
  approveRefund,
  rejectRefund
} from '@/api/lgjx/testOrder'

import { updateApply } from '@/api/lgjx/testApply'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

import { date } from '@/utils/jxlg/date'
import { auditStatus, auditType } from '@/utils/jxlg/auditStatus'
import { productType } from '@/utils/jxlg/productType'
import { amount } from '@/utils/jxlg/amount'
import { attachType } from '@/utils/jxlg/attachType'
import { orderStatus, orderStatusType } from '@/utils/jxlg/orderStatus'
import { getTemplateList } from '@/api/lgjx/testTemplate'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  orderNo: '',
})

// 验证规则
const rule = reactive({})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({ auditRefund: true })
const templateData = ref([])

// 重置
const onReset = () => {
  searchInfo.value = { auditRefund: true }
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

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteOrderFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
  multipleSelection.value.map(item => {
    ids.push(item.ID)
  })
  const res = await deleteOrderByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateOrderFunc = async(row) => {
  const res = await findOrder({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reorder
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteOrderFunc = async(row) => {
  const res = await deleteOrder({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    orderNo: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createOrder(formData.value)
        break
      case 'update':
        res = await updateOrder(formData.value)
        break
      default:
        res = await createOrder(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

const approveRefundFunc = async(apply) => {
  ElMessageBox.confirm('确定要通过吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await approveRefund(apply)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '提交成功'
      })
      getTableData()
    }
  })
}

const rejectRefundFunc = async(apply) => {
  ElMessageBox.confirm('确定要拒绝吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await rejectRefund(apply)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '提交成功'
      })
      getTableData()
    }
  })
}
</script>

<style>
</style>
