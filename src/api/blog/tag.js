import request from '@/utils/request'

export function listTag(params) {
  return request({
    url: '/blog/tags',
    method: 'get',
    params
  })
}

export function getTag(params) {
  return request({
    url: '/blog/tag',
    method: 'get',
    params
  })
}

export function delTag(data) {
  return request({
    url: '/blog/tags',
    method: 'delete',
    data
  })
}

export function addTag(data) {
  return request({
    url: '/blog/tag',
    method: 'post',
    data
  })
}

export function updateTag(tag_id, data) {
  return request({
    url: '/blog/tag/' + tag_id,
    method: 'put',
    data
  })
}
