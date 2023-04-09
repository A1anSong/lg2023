<template>
  <div>
    <el-row v-loading="loading" :gutter="24" justify="center">
      <el-col :sm="12" :lg="6">
        <el-result
          :icon="queryStatus.status"
          :title="queryStatus.message"
          :sub-title="queryStatus.subMessage"
        />
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'ElogValidate'
}
</script>
<script setup>
import { ref } from 'vue'
import { elogValidate } from '@/api/lg/order'

const loading = ref(true)
const queryStatus = ref({
  status: 'info',
  message: '正在查询',
  subMessage: '请耐心等待'
})

const elogNo = ref('')
const validateCode = ref('')

const params = new URLSearchParams(location.search)

elogNo.value = params.get('elogNo')
validateCode.value = params.get('validateCode')

const validate = async() => {
  if (elogNo.value === null || validateCode.value === null) {
    queryStatus.value = {
      status: 'warning',
      message: '无法鉴真',
      subMessage: '请从保函上扫描二维码进行鉴真'
    }
    loading.value = false
    return
  }
  const res = await elogValidate({ elogNo: elogNo.value, validateCode: validateCode.value })
  if (res.code === 0) {
    queryStatus.value = {
      status: res.data.elogValidateMessage.status,
      message: res.data.elogValidateMessage.message,
      subMessage: res.data.elogValidateMessage.subMessage
    }
    loading.value = false
  } else {
    queryStatus.value = {
      status: 'error',
      message: '服务器异常',
      subMessage: '请联系管理员或稍后再试'
    }
    loading.value = false
  }
  // const res = await fetch('/api/elog/validate', {
  //   method: 'POST',
  //   headers: {
  //     'Content-Type': 'application/json'
  //   },
  //   body: JSON.stringify({
  //     elogNo: elogNo.value,
  //     validateCode: validateCode.value
  //   })
  // })
  // const data = await res.json()
  // if (data.status === 'success') {
  //   queryStatus.value = {
  //     status: 'success',
  //     message: '验证成功',
  //     subMessage: '请耐心等待'
  //   }
  // } else {
  //   queryStatus.value = {
  //     status: 'error',
  //     message: '验证失败',
  //     subMessage: '请耐心等待'
  //   }
  // }
  // loading.value = false
}

validate()
</script>
<style scoped>
</style>
