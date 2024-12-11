import request from '@/utils/request'

const superSendApi = {
  GetSuperSendList: '/super_send/getSuperSendList',
  AddSuperSend: '/super_send/addSuperSend',
  UpdateSuperSend: '/super_send/updateSuperSend',
  DeleteSuperSend: '/super_send/deleteSuperSend',
  GetOnlineSuperSend: '/super_send/getOnlineSuperSend',
  SetSuperSendOnline: '/super_send/setSuperSendOnline'
}
export function getOnlineSuperSend () {
  return request({
    url: superSendApi.GetOnlineSuperSend,
    method: 'post',
    data: {}
  })
}
export function getSuperSendList (parameter) {
  parameter.is_ssl = Number(parameter.is_ssl)
  return request({
    url: superSendApi.GetSuperSendList,
    method: 'post',
    data: parameter
  })
}
export function addSuperSend (parameter) {
  parameter.is_ssl = Number(parameter.is_ssl)
  return request({
    url: superSendApi.AddSuperSend,
    method: 'post',
    data: parameter
  })
}
export function updateSuperSend (parameter) {
  parameter.is_ssl = Number(parameter.is_ssl)
  return request({
    url: superSendApi.UpdateSuperSend,
    method: 'post',
    data: parameter
  })
}
export function deleteSuperSend (id) {
  return request({
    url: superSendApi.DeleteSuperSend,
    method: 'post',
    data: { 'id': id }
  })
}
export function setSuperSendOnline (id, online) {
  return request({
    url: superSendApi.SetSuperSendOnline,
    method: 'post',
    data: { 'id': id, 'online': online }
  })
}
