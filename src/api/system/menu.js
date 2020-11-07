import request from '@/utils/request'

export function listMenu(params) {
  return request({
    url: '/menus',
    method: 'get',
    params
  })
}

export function getMenu(menu_id) {
  return request({
    url: '/menu/' + menu_id,
    method: 'get'
  })
}

export function delMenu(data) {
  return request({
    url: '/menus',
    method: 'delete',
    data
  })
}

export function addMenu(data) {
  return request({
    url: '/menu',
    method: 'post',
    data
  })
}

export function updateMenu(menu_id, data) {
  return request({
    url: '/menu/' + menu_id,
    method: 'put',
    data
  })
}
