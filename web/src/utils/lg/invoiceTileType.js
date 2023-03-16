export const invoiceTileType = (invoiceTileType) => {
  switch (invoiceTileType) {
    case 'A1':
      return '个人或事业单位'
    case 'B1':
      return '企业'
    default:
      return invoiceTileType
  }
}
