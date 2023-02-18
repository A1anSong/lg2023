import service from '@/utils/request'

export const createInvoiceApply = (data) => {
  return service({
    url: '/invoiceApply/createInvoiceApply',
    method: 'post',
    data
  })
}

export const deleteInvoiceApply = (data) => {
  return service({
    url: '/invoiceApply/deleteInvoiceApply',
    method: 'delete',
    data
  })
}

export const deleteInvoiceApplyByIds = (data) => {
  return service({
    url: '/invoiceApply/deleteInvoiceApplyByIds',
    method: 'delete',
    data
  })
}

export const updateInvoiceApply = (data) => {
  return service({
    url: '/invoiceApply/updateInvoiceApply',
    method: 'put',
    data
  })
}

export const findInvoiceApply = (params) => {
  return service({
    url: '/invoiceApply/findInvoiceApply',
    method: 'get',
    params
  })
}

export const getInvoiceApplyList = (params) => {
  return service({
    url: '/invoiceApply/getInvoiceApplyList',
    method: 'get',
    params
  })
}

export const approveInvoiceApply = (data) => {
  return service({
    url: '/invoiceApply/approveInvoiceApply',
    method: 'put',
    data
  })
}

export const rejectInvoiceApply = (data) => {
  return service({
    url: '/invoiceApply/rejectInvoiceApply',
    method: 'put',
    data
  })
}