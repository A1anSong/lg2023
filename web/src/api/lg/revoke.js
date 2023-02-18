import service from '@/utils/request'

export const createRevoke = (data) => {
  return service({
    url: '/revoke/createRevoke',
    method: 'post',
    data
  })
}

export const deleteRevoke = (data) => {
  return service({
    url: '/revoke/deleteRevoke',
    method: 'delete',
    data
  })
}

export const deleteRevokeByIds = (data) => {
  return service({
    url: '/revoke/deleteRevokeByIds',
    method: 'delete',
    data
  })
}

export const updateRevoke = (data) => {
  return service({
    url: '/revoke/updateRevoke',
    method: 'put',
    data
  })
}

export const findRevoke = (params) => {
  return service({
    url: '/revoke/findRevoke',
    method: 'get',
    params
  })
}

export const getRevokeList = (params) => {
  return service({
    url: '/revoke/getRevokeList',
    method: 'get',
    params
  })
}
