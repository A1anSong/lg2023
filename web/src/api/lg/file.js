import service from '@/utils/request'
import { ElMessage } from 'element-plus'

export const downloadFile = (file) => {
  return service({
    url: '/file/download',
    method: 'get',
    params: file,
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, file.fileName)
  })
}

const handleFileError = (res, fileName) => {
  if (typeof (res.data) !== 'undefined') {
    if (res.data.type === 'application/json') {
      const reader = new FileReader()
      reader.onload = function() {
        const message = JSON.parse(reader.result).msg
        ElMessage({
          showClose: true,
          message: message,
          type: 'error'
        })
      }
      reader.readAsText(new Blob([res.data]))
    }
  } else {
    const downloadUrl = window.URL.createObjectURL(new Blob([res]))
    const a = document.createElement('a')
    a.style.display = 'none'
    a.href = downloadUrl
    a.download = fileName
    a.click()
  }
}
