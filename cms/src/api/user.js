import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function getInfo(data) {
  return request({
    url: '/admin/info',
    method: 'post',
    data
  })
}

export function logout(token) {
  return request({
    url: '/admin/logout',
    method: 'post'
  })
}

export function getList(token) {
  return request({
    url: '/user/list',
    method: 'post'
  })
}

export function addUser(data) {
  return request({
    url: '/user/add',
    method: 'post',
    data
  })
}

export function delUser(userId) {
  return request({
    url: '/user/' + userId,
    method: 'delete'
  })
}

export function updateUserDesc(userId, data) {
  return request({
    url: '/user/desc/' + userId,
    method: 'post',
    data
  })
}
