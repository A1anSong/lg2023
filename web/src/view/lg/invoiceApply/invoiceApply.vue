<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
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
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <!--        <el-table-column type="selection" width="55" />-->
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="申请编号" prop="applyNo" width="120" />
        <el-table-column align="left" label="总金额" prop="invoiceTotalAmount" width="120" />
        <el-table-column align="left" label="发票类型" prop="invoiceType" width="120" />
        <el-table-column align="left" label="发票抬头类型" prop="invoiceTileType" width="120" />
        <el-table-column align="left" label="发票抬头" prop="invoiceTile" width="120" />
        <el-table-column align="left" label="税号" prop="taxNo" width="120" />
        <el-table-column align="left" label="开户银行" prop="bankName" width="120" />
        <el-table-column align="left" label="银行账号" prop="bankNo" width="120" />
        <el-table-column align="left" label="企业地址" prop="companyAddress" width="120" />
        <el-table-column align="left" label="企业电话" prop="companyTel" width="120" />
        <el-table-column align="left" label="开票备注" prop="remarks" width="120" />
        <el-table-column align="left" label="操作" min-width="200" fixed="right">
          <template #default="scope">
            <el-button
              v-if="scope.row.auditStatus!==2 && scope.row.auditStatus!==3"
              type="success"
              icon="select"
              size="small"
              @click="approveInvoiceApplyFunc(scope.row)"
            >通过
            </el-button>
            <el-button
              v-if="scope.row.auditStatus!==2 && scope.row.auditStatus!==3"
              type="danger"
              icon="closeBold"
              size="small"
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
        <el-form-item label="申请编号:" prop="applyNo">
          <el-input v-model="formData.applyNo" :clearable="true" placeholder="请输入" />
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
  name: 'TestInvoiceApply'
}
</script>

<script setup>
import {
  createInvoiceApply,
  deleteInvoiceApply,
  deleteInvoiceApplyByIds,
  updateInvoiceApply,
  findInvoiceApply,
  getInvoiceApplyList,
  approveInvoiceApply,
  rejectInvoiceApply
} from '@/api/lgjx/testInvoiceApply'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { approveRefund, rejectRefund } from '@/api/lgjx/testOrder'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  applyNo: '',
})

// 验证规则
const rule = reactive({
})

const elFormRef = ref()

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
    deleteInvoiceApplyFunc(row)
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
  const res = await deleteInvoiceApplyByIds({ ids })
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
const updateInvoiceApplyFunc = async(row) => {
  const res = await findInvoiceApply({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reinvoiceApply
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteInvoiceApplyFunc = async(row) => {
  const res = await deleteInvoiceApply({ ID: row.ID })
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
    applyNo: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createInvoiceApply(formData.value)
           break
         case 'update':
           res = await updateInvoiceApply(formData.value)
           break
         default:
           res = await createInvoiceApply(formData.value)
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
      getTableData()
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
      getTableData()
    }
  })
}
</script>

<style>
</style>
