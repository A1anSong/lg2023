import service from '@/utils/request'
import { ElMessage } from 'element-plus'

export const createProject = (data) => {
  return service({
    url: '/project/createProject',
    method: 'post',
    data
  })
}

export const updateProject = (data) => {
  return service({
    url: '/project/updateProject',
    method: 'put',
    data
  })
}

export const enableProjectByIds = (data) => {
  return service({
    url: '/project/enableProjectByIds',
    method: 'put',
    data
  })
}

export const autoMaticProjectByIds = (data) => {
  return service({
    url: '/project/autoMaticProjectByIds',
    method: 'put',
    data
  })
}

export const findProject = (params) => {
  return service({
    url: '/project/findProject',
    method: 'get',
    params
  })
}

export const getProjectList = (params) => {
  return service({
    url: '/project/getProjectList',
    method: 'get',
    params
  })
}

export const bindProject = (data) => {
  return service({
    url: '/project/bindProject',
    method: 'post',
    data
  })
}

export const unbindProject = (data) => {
  return service({
    url: '/project/unbindProject',
    method: 'post',
    data
  })
}

export const autoMaticProject = (data) => {
  return service({
    url: '/project/autoMaticProject',
    method: 'post',
    data
  })
}

export const unAutoMaticProject = (data) => {
  return service({
    url: '/project/unAutoMaticProject',
    method: 'post',
    data
  })
}

export const downloadTemplate = (file) => {
  return service({
    url: '/project/downloadTemplate',
    method: 'get',
    params: {
      fileName: file
    },
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, file)
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
    const event = new MouseEvent('click')
    a.dispatchEvent(event)
  }
}
