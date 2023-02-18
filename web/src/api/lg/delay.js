import service from '@/utils/request'

export const createDelay = (data) => {
  return service({
    url: '/delay/createDelay',
    method: 'post',
    data
  })
}

export const deleteDelay = (data) => {
  return service({
    url: '/delay/deleteDelay',
    method: 'delete',
    data
  })
}

export const deleteDelayByIds = (data) => {
  return service({
    url: '/delay/deleteDelayByIds',
    method: 'delete',
    data
  })
}

export const updateDelay = (data) => {
  return service({
    url: '/delay/updateDelay',
    method: 'put',
    data
  })
}

export const findDelay = (params) => {
  return service({
    url: '/delay/findDelay',
    method: 'get',
    params
  })
}

export const getDelayList = (params) => {
  return service({
    url: '/delay/getDelayList',
    method: 'get',
    params
  })
}
