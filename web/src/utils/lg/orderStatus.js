export const orderStatus = (order) => {
  if (order.revoke != null) {
    return '已撤'
  } else if (order.logout != null) {
    return '销函'
  } else if (order.claim != null && order.claim.auditStatus === 2) {
    return '理赔'
  } else if (order.refund != null && order.refund.auditStatus === 2) {
    return '退函'
  } else if (order.delay != null && order.delay.auditStatus === 2) {
    return '延期'
  } else if (order.letter != null) {
    return '已开'
  } else {
    return '未开'
  }
}

export const orderStatusType = (order) => {
  if (order.revoke != null) {
    return 'warning'
  } else if (order.logout != null) {
    return 'info'
  } else if (order.claim != null && order.claim.auditStatus === 2) {
    return 'danger'
  } else if (order.refund != null && order.refund.auditStatus === 2) {
    return 'info'
  } else if (order.delay != null && order.delay.auditStatus === 2) {
    return ''
  } else if (order.letter != null) {
    return 'success'
  } else {
    return 'warning'
  }
}
