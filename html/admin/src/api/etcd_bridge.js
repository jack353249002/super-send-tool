import request from '@/utils/request'
export const SUPERSENDCON = {}
export const etcdBirdgeApi = {
  GetEtcdBridgeList: '/etcd_bridge/getEtcdBridgeList'

}
export function getEtcdBridgeList (parameter) {
  parameter.is_ssl = Number(parameter.is_ssl)
  return request({
    url: etcdBirdgeApi.GetEtcdBridgeList,
    method: 'post',
    data: parameter
  })
}
