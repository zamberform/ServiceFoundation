import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/tag/list',
    method: 'post',
    params
  })
}

export function addTag(data) {
  return request({
    url: '/tag/add',
    method: 'post',
    data
  })
}

export function updateTag(data) {
  return request({
    url: '/tag',
    method: 'post',
    data
  })
}

export function deleteTag(tagId) {
  return request({
    url: '/tag/' + tagId,
    method: 'delete'
  })
}

