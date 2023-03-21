export const invoiceForm = (invoiceForm) => {
  switch (invoiceForm) {
    case 'A1':
      return '纸质发票'
    case 'B1':
      return '电子发票'
    default:
      return invoiceForm
  }
}
