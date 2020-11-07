import request from '@/utils/request'

export function listUser(params) {
  return request({
    url: '/users',
    method: 'get',
    params
  })
}

export function addUser(data) {
  return request({
    url: '/user',
    method: 'post',
    data
  })
}

export function delUser(data) {
  return request({
    url: '/users',
    method: 'DELETE',
    data
  })
}

export function updateUser(user_id, data) {
  return request({
    url: '/user/' + user_id,
    method: 'put',
    data
  })
}

export function updatePwd(user_id, data) {
  return request({
    url: '/user/pwd/' + user_id,
    method: 'post',
    data
  })
}
