import service from '@/utils/request'

export const createLetter = (data) => {
  return service({
    url: '/letter/createLetter',
    method: 'post',
    data
  })
}

export const deleteLetter = (data) => {
  return service({
    url: '/letter/deleteLetter',
    method: 'delete',
    data
  })
}

export const deleteLetterByIds = (data) => {
  return service({
    url: '/letter/deleteLetterByIds',
    method: 'delete',
    data
  })
}

export const updateLetter = (data) => {
  return service({
    url: '/letter/updateLetter',
    method: 'put',
    data
  })
}

export const findLetter = (params) => {
  return service({
    url: '/letter/findLetter',
    method: 'get',
    params
  })
}

export const getLetterList = (params) => {
  return service({
    url: '/letter/getLetterList',
    method: 'get',
    params
  })
}
