import service from '@/utils/request'

export const createLogout = (data) => {
  return service({
    url: '/logout/createLogout',
    method: 'post',
    data
  })
}

export const deleteLogout = (data) => {
  return service({
    url: '/logout/deleteLogout',
    method: 'delete',
    data
  })
}

export const deleteLogoutByIds = (data) => {
  return service({
    url: '/logout/deleteLogoutByIds',
    method: 'delete',
    data
  })
}

export const updateLogout = (data) => {
  return service({
    url: '/logout/updateLogout',
    method: 'put',
    data
  })
}

export const findLogout = (params) => {
  return service({
    url: '/logout/findLogout',
    method: 'get',
    params
  })
}

export const getLogoutList = (params) => {
  return service({
    url: '/logout/getLogoutList',
    method: 'get',
    params
  })
}
