import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/comment/list',
    method: 'get',
    params
  })
}
