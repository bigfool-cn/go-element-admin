import request from '@/utils/request'

export function listPath(params) {
  return request({
    url: '/paths',
    method: 'get',
    params
  })
}

export function getPath(path_id) {
  return request({
    url: '/path/' + path_id,
    method: 'get'
  })
}

export function delPath(data) {
  return request({
    url: '/paths',
    method: 'delete',
    data
  })
}

export function addPath(data) {
  return request({
    url: '/path',
    method: 'post',
    data
  })
}

export function updatePath(path_id, data) {
  return request({
    url: '/path/' + path_id,
    method: 'put',
    data
  })
}
