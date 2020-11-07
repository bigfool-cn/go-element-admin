import request from '@/utils/request'

export function listRole(params) {
  return request({
    url: '/roles',
    method: 'get',
    params
  })
}

export function getRole(params) {
  return request({
    url: '/role/get',
    method: 'get',
    params
  })
}

export function delRole(data) {
  return request({
    url: '/roles',
    method: 'delete',
    data
  })
}

export function addRole(data) {
  return request({
    url: '/role',
    method: 'post',
    data
  })
}

export function updateRole(role_id, data) {
  return request({
    url: '/role/' + role_id,
    method: 'put',
    data
  })
}
