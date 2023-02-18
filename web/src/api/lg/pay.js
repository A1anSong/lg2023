import service from '@/utils/request'

export const createPay = (data) => {
  return service({
    url: '/pay/createPay',
    method: 'post',
    data
  })
}

export const deletePay = (data) => {
  return service({
    url: '/pay/deletePay',
    method: 'delete',
    data
  })
}

export const deletePayByIds = (data) => {
  return service({
    url: '/pay/deletePayByIds',
    method: 'delete',
    data
  })
}

export const updatePay = (data) => {
  return service({
    url: '/pay/updatePay',
    method: 'put',
    data
  })
}

export const findPay = (params) => {
  return service({
    url: '/pay/findPay',
    method: 'get',
    params
  })
}

export const getPayList = (params) => {
  return service({
    url: '/pay/getPayList',
    method: 'get',
    params
  })
}
