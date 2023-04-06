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
          <el-select v-model="searchInfo.templateID" clearable placeholder="选择条件">
            <el-option
              v-for="template in templateData"
              :key="template.ID"
              :label="template.templateName"
              :value="template.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="上架状态">
          <el-select v-model="searchInfo.isEnable" clearable placeholder="选择条件">
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
        <el-table-column align="center" label="标段名称" prop="projectName" min-width="300px" />
        <el-table-column align="center" label="标段编号" prop="projectNo" min-width="160px" />
        <el-table-column align="center" label="标段金额" min-width="120px">
          <template #default="scope">{{ scope.row.projectAmount == null ? "" : amount(scope.row.projectAmount) }}</template>
        </el-table-column>
        <el-table-column align="center" label="招标人名称" prop="tendereeName" min-width="280px" />
        <el-table-column align="center" label="招标人地址" prop="tendereeAddress" min-width="300px" />
        <el-table-column align="center" label="招标人电话" prop="tendereeTel" min-width="120px" />
        <el-table-column align="center" label="招标代理电话" prop="agentTel" min-width="120px" />
        <el-table-column align="center" label="担保金额" min-width="120px">
          <template #default="scope">{{ scope.row.tenderDeposit == null ? "" : amount(scope.row.tenderDeposit) }}</template>
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
        <el-table-column align="center" label="招标文件" prop="tendereeFile" min-width="120px" />
        <el-table-column align="center" label="上架状态" min-width="100">
          <template #default="scope">
            <el-switch
              v-model="scope.row.isEnable"
              inline-prompt
              active-text="是"
              inactive-text="否"
              style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
              disabled
            />
          </template>
        </el-table-column>
        <el-table-column align="center" label="自动审批" min-width="100">
          <template #default="scope">
            <el-switch
              v-model="scope.row.isAutoMatic"
              inline-prompt
              active-text="是"
              inactive-text="否"
              style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
              disabled
            />
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
  name: 'ProjectView'
}
</script>

<script setup>
import {
  getProjectList
} from '@/api/lg/project'

import { ref } from 'vue'

import { date } from '@/utils/lg/date'
import { amount } from '@/utils/lg/amount'
import { getTemplateList } from '@/api/lg/template'

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
</script>

<style>
</style>
