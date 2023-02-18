import service from '@/utils/request'

export const createProject = (data) => {
  return service({
    url: '/project/createProject',
    method: 'post',
    data
  })
}

export const deleteProject = (data) => {
  return service({
    url: '/project/deleteProject',
    method: 'delete',
    data
  })
}

export const deleteProjectByIds = (data) => {
  return service({
    url: '/project/deleteProjectByIds',
    method: 'delete',
    data
  })
}

export const updateProject = (data) => {
  return service({
    url: '/project/updateProject',
    method: 'put',
    data
  })
}

export const findProject = (params) => {
  return service({
    url: '/project/findProject',
    method: 'get',
    params
  })
}

export const getProjectList = (params) => {
  return service({
    url: '/project/getProjectList',
    method: 'get',
    params
  })
}

export const bindProject = (data) => {
  return service({
    url: '/project/bindProject',
    method: 'post',
    data
  })
}

export const unbindProject = (data) => {
  return service({
    url: '/project/unbindProject',
    method: 'post',
    data
  })
}