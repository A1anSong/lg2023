<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="项目名称:" prop="projectName">
          <el-input v-model="formData.projectName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目编号:" prop="projectNo">
          <el-input v-model="formData.projectNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目金额:" prop="projectAmount">
          <el-input-number v-model="formData.projectAmount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="招标人名称:" prop="tendereeName">
          <el-input v-model="formData.tendereeName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="招标人地址:" prop="tendereeAddress">
          <el-input v-model="formData.tendereeAddress" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="招标人电话:" prop="tendereeTel">
          <el-input v-model="formData.tendereeTel" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="担保金额:" prop="tenderDeposit">
          <el-input-number v-model="formData.tenderDeposit" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="项目开标时间:" prop="projectOpenTime">
          <el-input v-model="formData.projectOpenTime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目发布日期:" prop="projectPublishTime">
          <el-input v-model="formData.projectPublishTime" :clearable="true" placeholder="请输入" />
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
          <el-input v-model="formData.tenderEndDate" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目类型:" prop="projectType">
          <el-input v-model="formData.projectType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Project'
}
</script>

<script setup>
import {
  createProject,
  updateProject,
  findProject
} from '@/api/lgjx/project'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            projectName: '',
            projectNo: '',
            projectAmount: 0,
            tendereeName: '',
            tendereeAddress: '',
            tendereeTel: '',
            tenderDeposit: 0,
            projectOpenTime: '',
            projectPublishTime: '',
            projectCity: '',
            projectCounty: '',
            projectDay: 0,
            tenderEndDate: '',
            projectType: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProject({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reproject
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
