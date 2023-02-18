import service from '@/utils/request'

export const createApply = (data) => {
  return service({
    url: '/apply/createApply',
    method: 'post',
    data
  })
}

export const deleteApply = (data) => {
  return service({
    url: '/apply/deleteApply',
    method: 'delete',
    data
  })
}

export const deleteApplyByIds = (data) => {
  return service({
    url: '/apply/deleteApplyByIds',
    method: 'delete',
    data
  })
}

export const updateApply = (data) => {
  return service({
    url: '/apply/updateApply',
    method: 'put',
    data
  })
}

export const findApply = (params) => {
  return service({
    url: '/apply/findApply',
    method: 'get',
    params
  })
}

export const getApplyList = (params) => {
  return service({
    url: '/apply/getApplyList',
    method: 'get',
    params
  })
}
