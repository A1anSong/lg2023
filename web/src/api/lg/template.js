import service from '@/utils/request'

export const createTemplate = (data) => {
  return service({
    url: '/template/createTemplate',
    method: 'post',
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

export const getTemplateList = (params) => {
  return service({
    url: '/template/getTemplateList',
    method: 'get',
    params
  })
}
