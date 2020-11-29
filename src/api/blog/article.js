import request from '@/utils/request'

export function listArticle(params) {
  return request({
    url: '/blog/articles',
    method: 'get',
    params
  })
}

export function getArticle(article_id) {
  return request({
    url: '/blog/article/' + article_id,
    method: 'get'
  })
}

export function delArticle(data) {
  return request({
    url: '/blog/articles',
    method: 'delete',
    data
  })
}

export function addArticle(data) {
  return request({
    url: '/blog/article',
    method: 'post',
    data
  })
}

export function updateArticle(article_id, data) {
  return request({
    url: '/blog/article/' + article_id,
    method: 'put',
    data
  })
}
