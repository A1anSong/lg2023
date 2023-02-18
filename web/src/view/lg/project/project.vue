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
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="small"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除
            </el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="项目名称" prop="projectName" width="120" />
        <el-table-column align="left" label="项目编号" prop="projectNo" width="120" />
        <el-table-column align="left" label="项目金额" prop="projectAmount" width="120" />
        <el-table-column align="left" label="招标人名称" prop="tendereeName" width="120" />
        <el-table-column align="left" label="招标人地址" prop="tendereeAddress" width="120" />
        <el-table-column align="left" label="招标人电话" prop="tendereeTel" width="120" />
        <el-table-column align="left" label="招标代理电话" prop="agentTel" width="120" />
        <el-table-column align="left" label="担保金额" prop="tenderDeposit" width="120" />
        <el-table-column align="left" label="项目开标时间" prop="projectOpenTime" width="120" />
        <el-table-column align="left" label="项目发布日期" prop="projectPublishTime" width="120" />
        <el-table-column align="left" label="所属市" prop="projectCity" width="120" />
        <el-table-column align="left" label="所属县" prop="projectCounty" width="120" />
        <el-table-column align="left" label="保函有效期" prop="projectDay" width="120" />
        <el-table-column align="left" label="投标截止时间" prop="tenderEndDate" width="120" />
        <el-table-column align="left" label="项目类型" prop="projectType" width="120" />
        <el-table-column align="left" label="项目类别" prop="projectCategory" width="120" />
        <el-table-column align="left" label="是否启用" width="180">
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
        <el-table-column align="left" label="项目详情" min-width="160" fixed="right">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              size="small"
              class="table-button"
              @click="updateProjectFunc(scope.row)"
            >变更
            </el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
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
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="项目名称:" prop="projectName">
          <el-input v-model="formData.projectName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目编号:" prop="projectNo">
          <el-input v-model="formData.projectNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目金额:" prop="projectAmount">
          <el-input-number v-model="formData.projectAmount" style="width:100%" :precision="2" :clearable="true" placeholder="请输入" :controls="false" />
        </el-form-item>
        <el-form-item label="招标人名称:" prop="tendereeName">
          <el-input v-model="formData.tendereeName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="招标人地址:" prop="tendereeAddress">
          <el-input v-model="formData.tendereeAddress" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="招标人电话:" prop="agentTel">
          <el-input v-model="formData.tendereeTel" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="招标代理电话:" prop="agentTel">
          <el-input v-model="formData.agentTel" :clearable="true" placeholder="请输入" />
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
          <el-input v-model="formData.projectCity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属县:" prop="projectCounty">
          <el-input v-model="formData.projectCounty" :clearable="true" placeholder="请输入" />
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
          <el-input v-model="formData.projectType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目类别:" prop="projectType">
          <el-input v-model="formData.projectCategory" :clearable="true" placeholder="请输入" />
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

<script>
export default {
  name: 'TestProject'
}
</script>

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
} from '@/api/lgjx/testProject'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { getTemplateList } from '@/api/lgjx/testTemplate'

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
    getTableData()
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
      getTableData()
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
      getTableData()
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
      getTableData()
    } else {
      row.isEnable = true
    }
  }
}
</script>

<style>
.el-input-number .el-input__inner{
  text-align: left;
}
</style>
