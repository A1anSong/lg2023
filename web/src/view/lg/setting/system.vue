<template>
  <div class="system">
    <warning-bar title="注意：此处部分配置会影响到正式环境，请谨慎和管理员确认后修改" />
    <warning-bar title="修改完配置后建议重启服务生效" />
    <el-form ref="form" :model="config" label-width="240px">
      <!--  System start  -->
      <el-collapse v-model="activeNames">

        <el-collapse-item title="金融机构信息配置" name="1">
          <el-form-item label="单位名称">
            <el-input v-model="config.insurance['name']" />
          </el-form-item>
          <el-form-item label="信用代码">
            <el-input v-model="config.insurance['credit-code']" />
          </el-form-item>
          <el-form-item label="单位地址">
            <el-input v-model="config.insurance['address']" />
          </el-form-item>
          <el-form-item label="邮政编码">
            <el-input v-model="config.insurance['zip-code']" />
          </el-form-item>
          <el-form-item label="电话">
            <el-input v-model="config.insurance['tel']" />
          </el-form-item>
          <el-form-item label="开户银行">
            <el-input v-model="config.insurance['bankName']" />
          </el-form-item>
          <el-form-item label="银行账户">
            <el-input v-model="config.insurance['bankNo']" />
          </el-form-item>
          <el-form-item label="额度token">
            <el-input v-model="config.insurance['insurance-token']" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="开函配置" name="2">
          <el-form-item label="域名地址">
            <el-input v-model="config.insurance['host-domain']" />
          </el-form-item>
          <el-form-item label="保函编号前缀">
            <el-input v-model="config.insurance['elog-prefix']" />
          </el-form-item>
          <el-form-item label="临时文件路径">
            <el-input v-model="config.insurance['temp-dir']" />
          </el-form-item>
          <el-form-item label="LOGO文件路径">
            <el-input v-model="config.insurance['logo-file']" />
          </el-form-item>
          <el-form-item label="签章证书文件">
            <el-input v-model="config.insurance['key-file']" />
          </el-form-item>
          <el-form-item label="单位公章文件">
            <el-input v-model="config.insurance['stamp-file']" />
          </el-form-item>
          <el-form-item label="法人章文件">
            <el-input v-model="config.insurance['legal-file']" />
          </el-form-item>
          <el-form-item label="签章程序路径">
            <el-input v-model="config.insurance['sign-program']" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="系统接入配置" name="3">
          <el-form-item label="正式环境地址">
            <el-input v-model="config.insurance['api-domain']" />
          </el-form-item>
          <el-form-item label="交易中心API地址">
            <el-input v-model="config.insurance['jr-api-domain']" />
          </el-form-item>
          <el-form-item label="交易中心AppKey">
            <el-input v-model="config.insurance['app-key']" />
          </el-form-item>
          <el-form-item label="交易中心AppSecret">
            <el-input v-model="config.insurance['app-secret']" />
          </el-form-item>
          <el-form-item label="磁盘路径">
            <el-input v-model="config.insurance['disk-path']" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="保函费用配置" name="4">
          <el-form-item label="保函费率">
            <el-input-number v-model="config.insurance['elog-rate']" style="width:100%" :precision="4" :controls="false" />
          </el-form-item>
          <el-form-item label="保函最低收费">
            <el-input-number v-model="config.insurance['elog-min-amount']" style="width:100%" :precision="2" :controls="false" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="电子发票配置" name="5">
          <el-form-item label="请求地址">
            <el-input v-model="config.insurance['nn-request-url']" />
          </el-form-item>
          <el-form-item label="AppKey">
            <el-input v-model="config.insurance['nn-app-key']" />
          </el-form-item>
          <el-form-item label="AppSecret">
            <el-input v-model="config.insurance['nn-app-secret']" />
          </el-form-item>
          <el-form-item label="税号">
            <el-input v-model="config.insurance['nn-tax-no']" />
          </el-form-item>
          <el-form-item label="税率">
            <el-input-number v-model="config.insurance['nn-tax-rate']" style="width:100%" :precision="2" :controls="false" />
          </el-form-item>
          <el-form-item label="复核人">
            <el-input v-model="config.insurance['nn-checker']" />
          </el-form-item>
          <el-form-item label="收款人">
            <el-input v-model="config.insurance['nn-payee']" />
          </el-form-item>
          <el-form-item label="开票人">
            <el-input v-model="config.insurance['nn-clerk']" />
          </el-form-item>
          <el-form-item label="AccessToken">
            <el-input v-model="config.insurance['nn-access-token']" :disabled="true" />
          </el-form-item>
          <el-form-item label="商品税收分类编码">
            <el-input v-model="config.insurance['nn-goods-code']" />
          </el-form-item>
        </el-collapse-item>
      </el-collapse>
    </el-form>
    <div class="gva-btn-list">
      <el-button type="primary" size="small" @click="update">立即更新</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SysSetting'
}
</script>
<script setup>
import { getSystemConfig, setSystemConfig } from '@/api/system'
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import WarningBar from '@/components/warningBar/warningBar.vue'

const activeNames = reactive([])
const config = ref({
  system: {
    'iplimit-count': 0,
    'iplimit-time': 0
  },
  jwt: {},
  mysql: {},
  pgsql: {},
  excel: {},
  autocode: {},
  redis: {},
  qiniu: {},
  'tencent-cos': {},
  'aliyun-oss': {},
  'hua-wei-obs': {},
  captcha: {},
  zap: {},
  local: {},
  email: {},
  timer: {
    detail: {}
  },
  insurance: {},
})

const initForm = async() => {
  const res = await getSystemConfig()
  if (res.code === 0) {
    config.value = res.data.config
  }
}
initForm()
const update = async() => {
  const res = await setSystemConfig({ config: config.value })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '配置文件设置成功'
    })
    await initForm()
  }
}

</script>

<style lang="scss">
.system {
  background: #fff;
  padding:36px;
  border-radius: 2px;
  h2 {
    padding: 10px;
    margin: 10px 0;
    font-size: 16px;
    box-shadow: -4px 0px 0px 0px #e7e8e8;
  }
  ::v-deep(.el-input-number__increase){
    top:5px !important;
  }
  .gva-btn-list{
    margin-top:16px;
  }
}

.el-input-number .el-input__inner{
  text-align: left;
}
</style>
