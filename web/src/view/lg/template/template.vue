<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="success" icon="plus" @click="openDialog">新增</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        height="800"
        @selection-change="handleSelectionChange"
      >
        <el-table-column align="left" label="模板名称" prop="templateName" min-width="120px" />
        <el-table-column align="center" label="操作" width="200" fixed="right">
          <template #default="scope">
            <el-button type="primary" icon="download" @click="downloadTemplateFile(scope.row.templateFile)">下载</el-button>
            <el-button type="danger" icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增模板" :close-on-click-modal="false">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="模板名称:" prop="templateName">
          <el-input v-model="formData.templateName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件:" prop="templateName">
          <el-upload
            v-model:file-list="fileList"
            :action="`${path}/file/upload`"
            :before-upload="checkFile"
            :headers="{ 'x-token': userStore.token }"
            :on-error="uploadError"
            :on-success="uploadSuccess"
            drag
          >
            <el-icon class="el-icon--upload">
              <upload-filled />
            </el-icon>
            <div class="el-upload__text">
              拖动文件到此或 <em>点击上传</em>
            </div>
            <template #tip>
              <div>
                必须是测试通过的单个 docx 格式文件(大小不超过16m)
              </div>
            </template>
          </el-upload>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Template'
}
</script>

<script setup>
import {
  createTemplate,
  updateTemplate,
  getTemplateList
} from '@/api/lg/template'

import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
import { downloadFile } from '@/api/lg/file'
import { UploadFilled } from '@element-plus/icons-vue'

const path = ref(import.meta.env.VITE_BASE_API)

const userStore = useUserStore()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  templateName: '',
  fileName: '',
})

// 验证规则
const rule = reactive({})

const elFormRef = ref()
const fileList = ref([])

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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
  const table = await getTemplateList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteTemplateFunc(row)
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  fileList.value = []
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    templateName: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createTemplate(formData.value)
        break
      case 'update':
        res = await updateTemplate(formData.value)
        break
      default:
        res = await createTemplate(formData.value)
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

const checkFile = (file) => {
  const isDOCX = file.type === 'application/vnd.openxmlformats-officedocument.wordprocessingml.document'
  const isLt16M = file.size / 1024 / 1024 < 16
  if (!isDOCX) {
    ElMessage.error('上传文件只能是 docx 格式!')
  }
  if (!isLt16M) {
    ElMessage.error('文件大小不能超过 16MB，请重新上传')
  }
  return isDOCX && isLt16M
}

const uploadSuccess = (res, uploadFile) => {
  const { data } = res
  if (data.fileName) {
    uploadFile.name = data.fileName
    formData.value.fileName = data.fileName
    fileList.value = [uploadFile]
    ElMessage({
      type: 'success',
      message: '上传成功'
    })
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
}

const downloadTemplateFile = (file) => {
  downloadFile(file)
}
</script>

<style>
</style>
