import service from '@/utils/request'

export const createInvoice = (data) => {
  return service({
    url: '/invoice/createInvoice',
    method: 'post',
    data
  })
}

export const deleteInvoice = (data) => {
  return service({
    url: '/invoice/deleteInvoice',
    method: 'delete',
    data
  })
}

export const deleteInvoiceByIds = (data) => {
  return service({
    url: '/invoice/deleteInvoiceByIds',
    method: 'delete',
    data
  })
}

export const updateInvoice = (data) => {
  return service({
    url: '/invoice/updateInvoice',
    method: 'put',
    data
  })
}

export const findInvoice = (params) => {
  return service({
    url: '/invoice/findInvoice',
    method: 'get',
    params
  })
}

export const getInvoiceList = (params) => {
  return service({
    url: '/invoice/getInvoiceList',
    method: 'get',
    params
  })
}
