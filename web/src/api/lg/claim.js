import service from '@/utils/request'

export const createClaim = (data) => {
  return service({
    url: '/claim/createClaim',
    method: 'post',
    data
  })
}

export const deleteClaim = (data) => {
  return service({
    url: '/claim/deleteClaim',
    method: 'delete',
    data
  })
}

export const deleteClaimByIds = (data) => {
  return service({
    url: '/claim/deleteClaimByIds',
    method: 'delete',
    data
  })
}

export const updateClaim = (data) => {
  return service({
    url: '/claim/updateClaim',
    method: 'put',
    data
  })
}

export const findClaim = (params) => {
  return service({
    url: '/claim/findClaim',
    method: 'get',
    params
  })
}

export const getClaimList = (params) => {
  return service({
    url: '/claim/getClaimList',
    method: 'get',
    params
  })
}
