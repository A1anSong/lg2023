import service from '@/utils/request'

export const createTemplate = (data) => {
  return service({
    url: '/template/createTemplate',
    method: 'post',
    data
  })
}

export const deleteTemplate = (data) => {
  return service({
    url: '/template/deleteTemplate',
    method: 'delete',
    data
  })
}

export const deleteTemplateByIds = (data) => {
  return service({
    url: '/template/deleteTemplateByIds',
    method: 'delete',
    data
  })
}

export const updateTemplate = (data) => {
  return service({
    url: '/template/updateTemplate',
    method: 'put',
    data
  })
}

export const findTemplate = (params) => {
  return service({
    url: '/template/findTemplate',
    method: 'get',
    params
  })
}

export const getTemplateList = (params) => {
  return service({
    url: '/template/getTemplateList',
    method: 'get',
    params
  })
}
