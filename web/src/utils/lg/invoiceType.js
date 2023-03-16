export const invoiceType = (invoiceType) => {
  switch (invoiceType) {
    case 'A1':
      return '增值税电子发票'
    default:
      return invoiceType
  }
}
