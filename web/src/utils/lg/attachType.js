export const attachType = (attachType) => {
  switch (attachType) {
    case '001':
      return '招标公告'
    case '002':
      return '出让公告'
    case '003':
      return '更正公告'
    case '004':
      return '中标结果公告'
    case '005':
      return '出让结果公告'
    case '006':
      return '招标文件'
    case '007':
      return '企业营业执照'
    case '008':
      return '法人身份证背面'
    case '009':
      return '法人身份证正面'
    case '010':
      return '经办人身份证背面'
    case '011':
      return '经办人身份证正面'
    case '012':
      return '授信协议书'
    case '013':
      return '担保协议书'
    case '014':
      return '传输协议书'
    case '999':
      return '其他文件'
    default:
      return attachType
  }
}
