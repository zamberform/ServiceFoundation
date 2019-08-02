import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/comment/list',
    method: 'post',
    params
  })
}

export function delComment(commentId) {
  return request({
    url: '/comment/' + commentId,
    method: 'delete'
  })
}

export function publishComment(commentId) {
  return request({
    url: '/comment/publish',
    method: 'post',
    commentId
  })
}

export function hideComment(commentId) {
  return request({
    url: '/comment/cancel',
    method: 'post',
    commentId
  })
}
