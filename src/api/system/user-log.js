import request from '@/utils/request'

export function listLog(params) {
  return request({
    url: '/user/logs',
    method: 'get',
    params
  })
}

export function delLog(data) {
  return request({
    url: '/user/logs',
    method: 'DELETE',
    data
  })
}
