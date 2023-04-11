
export const lastPeriod = (period) => {
  switch (period) {
    case 'day':
      return '昨天'
    case 'week':
      return '上周'
    case 'month':
      return '上个月'
    default:
      return ''
  }
}
