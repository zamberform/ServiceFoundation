import request from '@/utils/request'

export function getPackageList(data) {
  return request({
    url: '/package/list',
    method: 'post',
    data
  })
}

