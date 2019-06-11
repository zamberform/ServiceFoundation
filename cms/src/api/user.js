import request from '@/utils/request'

export function getUserList(data) {
  return request({
    url: '/user/list',
    method: 'post',
    data
  })
}

