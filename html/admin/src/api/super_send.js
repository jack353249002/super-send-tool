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
  DelMessage: '/super_send/delMessage',
  GetSuperSendListCon: '/super_send/getSendInfoList',
  GetSmtpServerList: '/super_send/getSmtpServerList',
  SetSuperSendListCon: '/super_send/setSend',
  SetSuperSendInfoCon: '/super_send/setSendInfo',
  SetSendInfoActionCon: '/super_send/sendInfoAction',
  GetSendListCon: '/super_send/getSendList',
  GetSmtpServer: '/super_send/getSmtpServer',
  SetSmtpServer: '/super_send/setSmtpServer',
  DelSmtpServer: '/super_send/delSmtpServer',
  ReloadSmtpServer: '/super_send/reloadServer',
  GetConfList: '/super_send/getConfList',
  SetConf: '/super_send/setConf',
  DelConf: '/super_send/delConf',
  ReloadConf: '/super_send/reloadConf',
  GetReceiveList: '/super_send/getReceiveList',
  AddReceive: '/super_send/addReceive',
  SetReceive: '/super_send/setReceive',
  GetReceiveMessagesList: '/super_send/getReceiveMessagesList',
  GetReceiveMessagesInfo: '/super_send/getReceiveMessageInfo',
  ReceiveAction: '/super_send/receiveAction',
  GetImapServer: '/super_send/getImapServer',
  SetImapServer: '/super_send/setImapServer',
  DelImapServer: '/super_send/delImapServer',
  GetImapAllServer: '/super_send/getImapServerAllList',
  GetServerUsers: '/super_send/getUsers',
  GetServerUserInfo: '/super_send/getUserInfo',
  SetServerPassword: '/super_send/setPassword',
  SetServerUserPassword: '/super_send/setUserPassword',
  DelServerUser: '/super_send/delUser',
  CreateServerUser: '/super_send/createUser',
  GetServerRoles: '/super_send/getRolesList',
  AddServerRolesFroUser: '/super_send/addRoleForUser',
  DelServerUserRole: '/super_send/delUserRole',
  GetRolesList: '/super_send/getRolesList',
  AddRole: '/super_send/addRole',
  DelRole: '/super_send/delRole',
  GetRolesPermissions: '/super_send/getRolesPermissions',
  DelRolesPermissions: '/super_send/delRolesPermissions'
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
  parameter.refresh_token_interval = Number(parameter.refresh_token_interval)
  return request({
    url: superSendApi.AddSuperSend,
    method: 'post',
    data: parameter
  })
}
export function updateSuperSend (parameter) {
  parameter.is_ssl = Number(parameter.is_ssl)
  parameter.refresh_token_interval = Number(parameter.refresh_token_interval)
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
export function delMessage (con, id) {
  return con.request({
    url: superSendApi.DelMessage,
    method: 'post',
    data: { 'id': id }
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

export function setConfCon (con, parameter) {
  parameter.id = Number(parameter.id)

  return con.request({
    url: superSendApi.SetConf,
    method: 'post',
    data: parameter
  })
}

export function getConfListCon (con, parameter) {
  return con.request({
    url: superSendApi.GetConfList,
    method: 'post',
    data: parameter
  })
}
export function delConfCon (con, id) {
  return con.request({
    url: superSendApi.DelConf,
    method: 'post',
    data: { 'id': id }
  })
}
export function reloadConfCon (con) {
  return con.request({
    url: superSendApi.ReloadConf,
    method: 'post',
    data: { }
  })
}
export function getReceiveList (con, parameter) {
  parameter.pageNo = Number(parameter.pageNo)
  parameter.pageSize = Number(parameter.pageSize)
  if (parameter.status === undefined || parameter.status === null || parameter.status === '') {
    parameter.status = '-1'
  }
  parameter.status = Number(parameter.status)
  return con.request({
    url: superSendApi.GetReceiveList,
    method: 'post',
    data: parameter
  })
}
export function addReceive (con, parameter) {
  parameter.end_type = Number(parameter.end_type)
  return con.request({
    url: superSendApi.AddReceive,
    method: 'post',
    data: parameter
  })
}
export function getReceiveMessagesList (con, parameter) {
  return con.request({
    url: superSendApi.GetReceiveMessagesList,
    method: 'post',
    data: parameter
  })
}
export function getReceiveMessagesInfo (con, parameter) {
  return con.request({
    url: superSendApi.GetReceiveMessagesInfo,
    method: 'post',
    data: parameter
  })
}
export function getImapServerCon (con, parameter) {
  return con.request({
    url: superSendApi.GetImapServer,
    method: 'post',
    data: parameter
  })
}
export function setImapServerCon (con, parameter) {
  parameter.id = Number(parameter.id)
  parameter.port = Number(parameter.port)
  parameter.max_client = Number(parameter.max_client)
  return con.request({
    url: superSendApi.SetImapServer,
    method: 'post',
    data: parameter
  })
}
export function delImapServerCon (con, id) {
  return con.request({
    url: superSendApi.DelImapServer,
    method: 'post',
    data: { 'id': Number(id) }
  })
}
export function getImapAllServerCon (con, parameter) {
  return con.request({
    url: superSendApi.GetImapAllServer,
    method: 'post',
    data: parameter
  })
}
export function receiveActionCon (con, parameter) {
  return con.request({
    url: superSendApi.ReceiveAction,
    method: 'post',
    data: parameter
  })
}
export function getServerUsersCon (con, parameter) {
  return con.request({
    url: superSendApi.GetServerUsers,
    method: 'post',
    data: parameter
  })
}
export function getServerUserInfoCon (con, parameter) {
  return con.request({
    url: superSendApi.GetServerUserInfo,
    method: 'post',
    data: parameter
  })
}

export function setServerPasswordCon (con, parameter) {
  return con.request({
    url: superSendApi.SetServerPassword,
    method: 'post',
    data: parameter
  })
}

export function setServerUserPasswordCon (con, parameter) {
  return con.request({
    url: superSendApi.SetServerUserPassword,
    method: 'post',
    data: parameter
  })
}

export function delServerUserCon (con, parameter) {
  return con.request({
    url: superSendApi.DelServerUser,
    method: 'post',
    data: parameter
  })
}
export function createServerUserCon (con, parameter) {
  return con.request({
    url: superSendApi.CreateServerUser,
    method: 'post',
    data: parameter
  })
}
export function getServerRolesCon (con, parameter) {
  return con.request({
    url: superSendApi.GetServerRoles,
    method: 'post',
    data: parameter
  })
}
export function delServerUserRoleCon (con, parameter) {
  return con.request({
    url: superSendApi.DelServerUserRole,
    method: 'post',
    data: parameter
  })
}
export function addServerRolesFroUserCon (con, parameter) {
  return con.request({
    url: superSendApi.AddServerRolesFroUser,
    method: 'post',
    data: parameter
  })
}
export function getRolesListCon (con, parameter) {
  return con.request({
    url: superSendApi.GetRolesList,
    method: 'post',
    data: parameter
  })
}
export function addRoleCon (con, parameter) {
  return con.request({
    url: superSendApi.AddRole,
    method: 'post',
    data: parameter
  })
}
export function delRoleCon (con, parameter) {
  return con.request({
    url: superSendApi.DelRole,
    method: 'post',
    data: parameter
  })
}
export function getRolesPermissions (con, parameter) {
  return con.request({
    url: superSendApi.GetRolesPermissions,
    method: 'post',
    data: parameter
  })
}
export function delRolesPermissions (con, parameter) {
  return con.request({
    url: superSendApi.DelRolesPermissions,
    method: 'post',
    data: parameter
  })
}
export function setReceiveCon (con, parameter) {
  return con.request({
    url: superSendApi.SetReceive,
    method: 'post',
    data: parameter
  })
}
