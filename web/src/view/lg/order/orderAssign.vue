<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item label="订单编号">
          <el-input v-model="orderNo" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="选择人员">
          <el-select v-model="employeeID" placeholder="选择条件" clearable>
            <el-option
              v-for="employee in employeeData"
              :key="employee.ID"
              :label="employee.nickName + '：' + employee.employeeNo"
              :value="employee.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'OrderAssign'
}
</script>

<script setup>
import { ref } from 'vue'

import { getEmployeeListWithNo } from '@/api/user'

// =========== 表格控制部分 ===========
const orderNo = ref('')
const employeeID = ref('')
const employeeData = ref([])

// 搜索
const onSubmit = () => {
  console.log('searchInfo', searchInfo)
}

// 获取模板列表
const getEmployeeData = async() => {
  const employee = await getEmployeeListWithNo({ page: 1, pageSize: 999 })
  if (employee.code === 0) {
    employeeData.value = employee.data.list
  }
}

getEmployeeData()
</script>

<style>
</style>
