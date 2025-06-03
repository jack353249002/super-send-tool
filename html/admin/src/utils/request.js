import axios from 'axios'
import store from '@/store'
import storage from 'store'
import notification from 'ant-design-vue/es/notification'
import { VueAxios } from './axios'
import { ACCESS_TOKEN } from '@/store/mutation-types'

// 创建 axios 实例
const request = axios.create({
  // API 请求的默认前缀
  baseURL: process.env.VUE_APP_API_BASE_URL,
  timeout: 6000 // 请求超时时间
})
// 创建 axios 实例
export const createRequestCon = function (superSendInfo) {
  var requestSuperSend = axios.create({
    // API 请求的默认前缀
    baseURL: process.env.VUE_APP_API_BASE_URL,
    timeout: 6000 // 请求超时时间
  })
  // request interceptor
  requestSuperSend.interceptors.request.use(config => {
    const token = superSendInfo.token
    const superSendID = superSendInfo.id
    // 如果 token 存在
    // 让每个请求携带自定义 token 请根据实际情况自行修改
    if (token) {
      config.headers['token'] = token
    }
    if (superSendID > 0) {
      config.headers['super_send_id'] = superSendID
    }
    return config
  }, errorHandler)
  requestSuperSend.interceptors.response.use((response) => {
    return response.data
  }, errorHandler)
  return { request: requestSuperSend, superSendInfo: superSendInfo }
}
export const createEtcdBridgeRequestCon = function (etcdBridgeInfo) {
  var requestEtcdBridge = axios.create({
    // API 请求的默认前缀
    baseURL: process.env.VUE_APP_API_BASE_URL,
    timeout: 6000 // 请求超时时间
  })
  // request interceptor
  requestEtcdBridge.interceptors.request.use(config => {
    const token = etcdBridgeInfo.token
    const etcdBridgeID = etcdBridgeInfo.id
    // 如果 token 存在
    // 让每个请求携带自定义 token 请根据实际情况自行修改
    if (token) {
      config.headers['token'] = token
    }
    if (etcdBridgeID > 0) {
      config.headers['etcd_bridge_id'] = etcdBridgeID
    }
    return config
  }, errorHandler)
  requestEtcdBridge().interceptors.response.use((response) => {
    return response.data
  }, errorHandler)
  return { request: requestEtcdBridge(), etcdBridgeInfo: etcdBridgeInfo }
}
// 异常拦截处理器
const errorHandler = (error) => {
  if (error.response) {
    const data = error.response.data
    // 从 localstorage 获取 token
    const token = storage.get(ACCESS_TOKEN)
    if (error.response.status === 403) {
      notification.error({
        message: 'Forbidden',
        description: data.message
      })
    }
    if (error.response.status === 401 && !(data.result && data.result.isLogin)) {
      notification.error({
        message: 'Unauthorized',
        description: 'Authorization verification failed'
      })
      if (token) {
        store.dispatch('Logout').then(() => {
          setTimeout(() => {
            window.location.reload()
          }, 1500)
        })
      }
    }
  }
  return Promise.reject(error)
}

// request interceptor
request.interceptors.request.use(config => {
  const token = storage.get(ACCESS_TOKEN)
  // 如果 token 存在
  // 让每个请求携带自定义 token 请根据实际情况自行修改
  if (token) {
    config.headers[ACCESS_TOKEN] = token
  }
  return config
}, errorHandler)

// response interceptor
request.interceptors.response.use((response) => {
  return response.data
}, errorHandler)

const installer = {
  vm: {},
  install (Vue) {
    Vue.use(VueAxios, request)
  }
}

export function getAction (url, parameter) {
  return axios({
    url: url,
    method: 'get',
    params: parameter
  })
}
export default request

export {
  installer as VueAxios,
  request as axios
}
