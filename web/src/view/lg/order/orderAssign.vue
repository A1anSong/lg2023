<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item label="订单编号">
          <el-input v-model.trim="orderNo" placeholder="搜索条件" clearable />
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
          <el-button type="primary" icon="share" @click="onSubmit">认领</el-button>
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
import { assignOrder } from '@/api/lg/order'
import { ElMessage } from 'element-plus'

// =========== 表格控制部分 ===========
const orderNo = ref('')
const employeeID = ref('')
const employeeData = ref([])

// 搜索
const onSubmit = async() => {
  const res = await assignOrder({ orderNo: orderNo.value, employeeId: employeeID.value })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '认领成功'
    })
  }
}

// 获取模板列表
const getEmployeeData = async() => {
  const res = await getEmployeeListWithNo({ page: 1, pageSize: 999 })
  if (res.code === 0) {
    employeeData.value = res.data.list
  }
}

getEmployeeData()
</script>

<style>
</style>
