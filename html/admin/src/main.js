// with polyfills
import 'core-js/stable'
import 'regenerator-runtime/runtime'

import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store/'
import i18n from './locales'
import { VueAxios } from './utils/request'
import ProLayout, { PageHeaderWrapper } from '@ant-design-vue/pro-layout'
import themePluginConfig from '../config/themePluginConfig'
import SelectPage from 'v-selectpage'
// mock
// WARNING: `mockjs` NOT SUPPORT `IE` PLEASE DO NOT USE IN `production` ENV.
import './mock'

import bootstrap from './core/bootstrap'
import './core/lazy_use' // use lazy load components
import './utils/filter' // global filter
import './global.less'
import { getMessageList, getSmtpServerList, RequestConFactory, superSendApi } from '@/api/super_send'

Vue.config.productionTip = false
console.log('init')
Vue.prototype.getMessageList = getMessageList
Vue.prototype.getSmtpServerList = getSmtpServerList
// mount axios to `Vue.$http` and `this.$http`
Vue.use(VueAxios)
 Vue.use(SelectPage, {

  dataLoad: function (vue, url, params) {
    console.log('dataLoad', url)
    const con = RequestConFactory(params)
    if (url === superSendApi.GetMessageList) {
      return new Promise((resolve, reject) => {
        const requestParams = { 'keyWords': params.name, 'pageNo': params.pageNumber, 'pageSize': params.pageSize }
        vue.getMessageList(con, requestParams).then(res => {
          resolve(res)
        }, e => reject(e))
      })
    }
    if (url === superSendApi.GetSmtpServerList) {
      return new Promise((resolve, reject) => {
        const requestParams = {}
        vue.getSmtpServerList(con, requestParams).then(res => {
          resolve(res)
        }, e => reject(e))
      })
    }
  }

})
// use pro-layout components
Vue.component('pro-layout', ProLayout)
Vue.component('page-container', PageHeaderWrapper)
Vue.component('page-header-wrapper', PageHeaderWrapper)
console.log(Object.keys(Vue.options.components))
window.umi_plugin_ant_themeVar = themePluginConfig.theme

new Vue({
  router,
  store,
  i18n,
  // init localstorage, vuex, Logo message
  created: bootstrap,
  render: h => h(App)
}).$mount('#app')
