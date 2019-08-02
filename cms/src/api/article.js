import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/article/list',
    method: 'post',
    params
  })
}

export function updateArticle(data) {
  return request({
    url: '/article',
    method: 'post',
    data
  })
}

export function addArticle(data) {
  return request({
    url: '/article/add',
    method: 'post',
    data
  })
}

export function findArticle(articleId) {
  return request({
    url: '/article/' + articleId,
    method: 'get'
  })
}

export function delArticle(articleId) {
  return request({
    url: '/article/' + articleId,
    method: 'delete'
  })
}

