import request, { createRequestCon } from '@/utils/request'
export const SUPERSENDCON = {}
export const superSendApi = {
  GetSuperSendList: '/super_send/getSuperSendList',
  AddSuperSend: '/super_send/addSuperSend',
  UpdateSuperSend: '/super_send/updateSuperSend',
  DeleteSuperSend: '/super_send/deleteSuperSend',
  GetOnlineSuperSend: '/super_send/getOnlineSuperSend',
  SetSuperSendOnline: '/super_send/setSuperSendOnline',
  GetMessageList: '/super_send/getMessageList',
  GetSuperSendInfo: '/super_send/getSuperSendInfo',
  Login: '/super_send/loginSuperSend',
  Register: '/super_send/registerSuperSend',
  SetMessage: '/super_send/setMessage',
  GetSuperSendListCon: '/super_send/getSendInfoList',
  GetSmtpServerList: '/super_send/getSmtpServerList',
  SetSuperSendListCon: '/super_send/setSend',
  SetSuperSendInfoCon: '/super_send/setSendInfo',
  SetSendInfoActionCon: '/super_send/sendInfoAction',
  GetSendListCon: '/super_send/getSendList',
  GetSmtpServer: '/super_send/getSmtpServer',
  SetSmtpServer: '/super_send/setSmtpServer',
  DelSmtpServer: '/super_send/delSmtpServer',
  ReloadSmtpServer: '/super_send/reloadServer'
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
export function login (id) {
  const con = createRequestCon({ 'token': '', 'id': id })
  return con.request({
    url: superSendApi.Login,
    method: 'post',
    data: {}
  })
}
export function register (token, id, parameter) {
  const con = createRequestCon({ 'token': token, 'id': id })
  return con.request({
    url: superSendApi.Register,
    method: 'post',
    data: parameter
  })
}
export function RequestConFactory (superSendInfo) {
  var con = SUPERSENDCON[superSendInfo.token]
  if (con !== undefined) {
    return con
  } else {
    // eslint-disable-next-line no-unused-expressions
    con = createRequestCon(superSendInfo)
    SUPERSENDCON[superSendInfo.token] = con
    return con
  }
}
export function getMessageList (con, parameter) {
  return con.request({
    url: superSendApi.GetMessageList,
    method: 'post',
    data: parameter
  })
}
export function getSuperSendInfo (id) {
  return request({
    url: superSendApi.GetSuperSendInfo,
    method: 'post',
    data: {
      'id': Number(id)
    }
  })
}
export function setMessage (con, parameter) {
  return con.request({
    url: superSendApi.SetMessage,
    method: 'post',
    data: parameter
  })
}
export function getSuperSendListCon (con, parameter) {
  if (parameter.status === undefined || parameter.status === null || parameter.status === '') {
    parameter.status = '-1'
  }
  parameter.status = Number(parameter.status)
  return con.request({
    url: superSendApi.GetSuperSendListCon,
    method: 'post',
    data: parameter
  })
}
export function setSuperSendListCon (con, parameter) {
  parameter.send_model = 0
  parameter.dispatch_model = 0
  parameter.message_id = Number(parameter.message_id)
  return con.request({
    url: superSendApi.SetSuperSendListCon,
    method: 'post',
    data: parameter
  })
}
export function getSmtpServerList (con, parameter) {
  return con.request({
    url: superSendApi.GetSmtpServerList,
    method: 'post',
    data: parameter
  })
}
export function setSuperSendInfoCon (con, parameter) {
  return con.request({
    url: superSendApi.SetSuperSendInfoCon,
    method: 'post',
    data: parameter
  })
}
export function sendInfoActionCon (con, parameter) {
  return con.request({
    url: superSendApi.SetSendInfoActionCon,
    method: 'post',
    data: parameter
  })
}
export function getSendListCon (con, parameter) {
  parameter.status = Number(parameter.status)
  parameter.send_id = Number(parameter.send_id)
  return con.request({
    url: superSendApi.GetSendListCon,
    method: 'post',
    data: parameter
  })
}
export function getSmtpServerCon (con, parameter) {
  return con.request({
    url: superSendApi.GetSmtpServer,
    method: 'post',
    data: parameter
  })
}
export function setSmtpServerCon (con, parameter) {
  parameter.port = Number(parameter.port)
  parameter.max_concurrency = Number(parameter.max_concurrency)

  return con.request({
    url: superSendApi.SetSmtpServer,
    method: 'post',
    data: parameter
  })
}
export function delSmtpServerCon (con, id) {
  return con.request({
    url: superSendApi.DelSmtpServer,
    method: 'post',
    data: { 'id': Number(id) }
  })
}
export function reloadSmtpServerCon (con, token) {
  return con.request({
    url: superSendApi.ReloadSmtpServer,
    method: 'post',
    data: { 'token': token }
  })
}
