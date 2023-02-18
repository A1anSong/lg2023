export const auditStatus = (auditStatus) => {
  switch (auditStatus) {
    case 1:
      return '待审'
    case 2:
      return '通过'
    case 3:
      return '拒绝'
    default:
      return auditStatus
  }
}

export const auditType = (auditType) => {
  switch (auditType) {
    case 1:
      return 'warning'
    case 2:
      return 'success'
    case 3:
      return 'danger'
    default:
      return ''
  }
}