import service from '@/utils/request'

export const createRefund = (data) => {
  return service({
    url: '/refund/createRefund',
    method: 'post',
    data
  })
}

export const deleteRefund = (data) => {
  return service({
    url: '/refund/deleteRefund',
    method: 'delete',
    data
  })
}

export const deleteRefundByIds = (data) => {
  return service({
    url: '/refund/deleteRefundByIds',
    method: 'delete',
    data
  })
}

export const updateRefund = (data) => {
  return service({
    url: '/refund/updateRefund',
    method: 'put',
    data
  })
}

export const findRefund = (params) => {
  return service({
    url: '/refund/findRefund',
    method: 'get',
    params
  })
}

export const getRefundList = (params) => {
  return service({
    url: '/refund/getRefundList',
    method: 'get',
    params
  })
}
