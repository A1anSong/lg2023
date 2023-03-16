<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="标段编号">
          <el-input v-model="searchInfo.projectNo" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="标段名称">
          <el-input v-model="searchInfo.projectName" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="开标时间">
          <el-date-picker
            v-model="searchInfo.openTime"
            type="daterange"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>
        <el-form-item label="所属市">
          <el-input v-model="searchInfo.projectCity" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="所属县">
          <el-input v-model="searchInfo.projectCounty" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="保函模板">
          <el-select v-model="searchInfo.templateID" clearable placeholder="请输入">
            <el-option
              v-for="template in templateData"
              :key="template.ID"
              :label="template.templateName"
              :value="template.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="上架状态">
          <el-select v-model="searchInfo.isEnable" clearable placeholder="请输入">
            <el-option
              label="已上架"
              value="true"
            />
            <el-option
              label="未上架"
              value="false"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="success" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              type="danger"
              icon="delete"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除
            </el-button>
          </template>
        </el-popover>
      </div>
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
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" min-width="55" />
        <el-table-column align="center" label="标段名称" prop="projectName" min-width="300px" />
        <el-table-column align="center" label="标段编号" prop="projectNo" min-width="160px" />
        <el-table-column align="center" label="标段金额" min-width="120px">
          <template #default="scope">{{ amount(scope.row.projectAmount) }}</template>
        </el-table-column>
        <el-table-column align="center" label="招标人名称" prop="tendereeName" min-width="280px" />
        <el-table-column align="center" label="招标人地址" prop="tendereeAddress" min-width="300px" />
        <el-table-column align="center" label="招标人电话" prop="tendereeTel" min-width="120px" />
        <el-table-column align="center" label="招标代理电话" prop="agentTel" min-width="120px" />
        <el-table-column align="center" label="担保金额" min-width="120px">
          <template #default="scope">{{ amount(scope.row.tenderDeposit) }}</template>
        </el-table-column>
        <el-table-column align="center" label="开标时间" width="100px">
          <template #default="scope">{{ date(scope.row.projectOpenTime) }}</template>
        </el-table-column>
        <el-table-column align="center" label="发布日期" width="99px">
          <template #default="scope">{{ date(scope.row.projectPublishTime) }}</template>
        </el-table-column>
        <el-table-column align="center" label="所属市" prop="projectCity" min-width="120px" />
        <el-table-column align="center" label="所属县" prop="projectCounty" min-width="120px" />
        <el-table-column align="center" label="保函有效期" min-width="100px">
          <template #default="scope">{{ scope.row.projectDay + '天' }}</template>
        </el-table-column>
        <el-table-column align="center" label="截止时间" width="100px">
          <template #default="scope">{{ date(scope.row.tenderEndDate) }}</template>
        </el-table-column>
        <el-table-column align="center" label="项目类型" prop="projectType" min-width="120px" />
        <el-table-column align="center" label="项目类别" prop="projectCategory" min-width="120px" />
        <el-table-column align="center" label="上架状态" min-width="120px">
          <template #default="scope">
            <el-switch
              v-model="scope.row.isEnable"
              inline-prompt
              active-text="是"
              inactive-text="否"
              style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
              @change="changeProjectEnable($event, scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" min-width="200" fixed="right">
          <template #default="scope">
            <el-button
              type="warning"
              icon="edit"
              @click="updateProjectFunc(scope.row)"
            >变更
            </el-button>
            <el-button
              type="danger"
              icon="delete"
              @click="deleteRow(scope.row)"
            >删除
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="项目操作" :close-on-click-modal="false">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="项目名称:" prop="projectName">
          <el-input v-model.trim="formData.projectName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目编号:" prop="projectNo">
          <el-input v-model.trim="formData.projectNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目金额:" prop="projectAmount">
          <el-input-number v-model="formData.projectAmount" style="width:100%" :precision="2" :clearable="true" placeholder="请输入" :controls="false" />
        </el-form-item>
        <el-form-item label="招标人名称:" prop="tendereeName">
          <el-input v-model.trim="formData.tendereeName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="招标人地址:" prop="tendereeAddress">
          <el-input v-model.trim="formData.tendereeAddress" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="招标人电话:" prop="agentTel">
          <el-input v-model.trim="formData.tendereeTel" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="招标代理电话:" prop="agentTel">
          <el-input v-model.trim="formData.agentTel" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="担保金额:" prop="tenderDeposit">
          <el-input-number v-model="formData.tenderDeposit" style="width:100%" :precision="2" :clearable="true" placeholder="请输入" :controls="false" />
        </el-form-item>
        <el-form-item label="项目开标时间:" prop="projectOpenTime">
          <el-date-picker
            v-model="formData.projectOpenTime"
            type="datetime"
            placeholder="请输入"
            value-format="YYYY-MM-DD HH:mm:ss"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="项目发布日期:" prop="projectPublishTime">
          <el-date-picker
            v-model="formData.projectPublishTime"
            type="datetime"
            placeholder="请输入"
            value-format="YYYY-MM-DD HH:mm:ss"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="所属市:" prop="projectCity">
          <el-input v-model.trim="formData.projectCity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属县:" prop="projectCounty">
          <el-input v-model.trim="formData.projectCounty" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="保函有效期:" prop="projectDay">
          <el-input v-model.number="formData.projectDay" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投标截止时间:" prop="tenderEndDate">
          <el-date-picker
            v-model="formData.tenderEndDate"
            type="datetime"
            placeholder="请输入"
            value-format="YYYY-MM-DD HH:mm:ss"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="项目类型:" prop="projectType">
          <el-input v-model.trim="formData.projectType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目类别:" prop="projectType">
          <el-input v-model.trim="formData.projectCategory" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目模板:" prop="templateID">
          <el-select v-model="formData.templateID" clearable placeholder="请输入" style="width: 100%">
            <el-option
              v-for="template in templateData"
              :key="template.ID"
              :label="template.templateName"
              :value="template.ID"
            />
          </el-select>
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

<script setup>
import {
  createProject,
  deleteProject,
  deleteProjectByIds,
  updateProject,
  findProject,
  getProjectList,
  bindProject,
  unbindProject
} from '@/api/lg/project'

import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

import { date } from '@/utils/lg/date'
import { amount } from '@/utils/lg/amount'
import { getTemplateList } from '@/api/lg/template'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  projectName: null,
  projectNo: null,
  projectAmount: null,
  tendereeName: null,
  tendereeAddress: null,
  tendereeTel: null,
  agentTel: null,
  tenderDeposit: null,
  projectOpenTime: null,
  projectPublishTime: null,
  projectCity: null,
  projectCounty: null,
  projectDay: null,
  tenderEndDate: null,
  projectType: null,
  projectCategory: null,
  templateID: null,
  isEnable: false,
})

// 验证规则
const rule = reactive({})

const elFormRef = ref()

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
  const table = await getProjectList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteProjectFunc(row)
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
  const res = await deleteProjectByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    await getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateProjectFunc = async(row) => {
  const res = await findProject({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reproject
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteProjectFunc = async(row) => {
  const res = await deleteProject({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    await getTableData()
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
    projectName: null,
    projectNo: null,
    projectAmount: null,
    tendereeName: null,
    tendereeAddress: null,
    tendereeTel: null,
    agentTel: null,
    tenderDeposit: null,
    projectOpenTime: null,
    projectPublishTime: null,
    projectCity: null,
    projectCounty: null,
    projectDay: null,
    tenderEndDate: null,
    projectType: null,
    projectCategory: null,
    templateID: null,
    isEnable: false,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createProject(formData.value)
        break
      case 'update':
        res = await updateProject(formData.value)
        break
      default:
        res = await createProject(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      await getTableData()
    }
  })
}

const changeProjectEnable = async(val, row) => {
  if (val === true) {
    const res = await bindProject(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '绑定成功'
      })
      await getTableData()
    } else {
      row.isEnable = false
    }
  } else {
    const res = await unbindProject(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '解绑成功'
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
