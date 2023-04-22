import service from '@/utils/request'

export const getInvoiceApplyList = (params) => {
  return service({
    url: '/invoiceApply/getInvoiceApplyList',
    method: 'get',
    params
  })
}