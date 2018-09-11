import request from '@/utils/request'

export function getUsers(params) {
  return request({
    url: '/user/list',
    method: 'get',
    params
  })
}
