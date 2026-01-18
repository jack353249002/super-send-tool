<template>
  <pro-layout
    :menus="mainMenuList"
    :collapsed="collapsed"
    :mediaQuery="query"
    :isMobile="isMobile"
    :handleMediaQuery="handleMediaQuery"
    :handleCollapse="handleCollapse"
    :i18nRender="i18nRender"
    v-bind="settings"
  >
    <!-- Ads begin
      广告代码 真实项目中请移除
      production remove this Ads
    -->
    <ads v-if="isProPreviewSite && !collapsed"/>
    <!-- Ads end -->

    <!-- 1.0.0+ 版本 pro-layout 提供 API，
          我们推荐使用这种方式进行 LOGO 和 title 自定义
    -->
    <template v-slot:menuHeaderRender>
      <div>
        <img src="@/assets/logo.svg" />
        <h1>{{ title }}</h1>
      </div>
    </template>
    <!-- 1.0.0+ 版本 pro-layout 提供 API,
          增加 Header 左侧内容区自定义
    -->
    <template v-slot:headerContentRender>
      <div>
<!--        <a-tooltip title="刷新页面">
          <a-icon type="reload" style="font-size: 18px;cursor: pointer;" @click="() => { $message.info('只是一个DEMO') }" />
        </a-tooltip>-->
      </div>
    </template>

    <setting-drawer v-if="isDev" :settings="settings" @change="handleSettingChange">
      <div style="margin: 12px 0;">
        This is SettingDrawer custom footer content.
      </div>
    </setting-drawer>
    <template v-slot:rightContentRender>
      <right-content :top-menu="settings.layout === 'topmenu'" :is-mobile="isMobile" :theme="settings.theme" />
    </template>
    <!-- custom footer / 自定义Footer -->
    <template v-slot:footerRender>
      <global-footer />
    </template>
    <router-view :key="$route.fullPath" :nowSendInfo="nowSendInfo"/>
  </pro-layout>
</template>

<script>
import { asyncRouterMap, constantRouterMap, RouteView } from '@/config/router.config.js'
import { SettingDrawer, updateTheme } from '@ant-design-vue/pro-layout'
import { i18nRender } from '@/locales'
import {
  CONTENT_WIDTH_TYPE,
  SET_SENDINFO,
  SIDEBAR_TYPE,
  TOGGLE_MOBILE_TYPE
} from '@/store/mutation-types'

import defaultSettings from '@/config/defaultSettings'
import RightContent from '@/components/GlobalHeader/RightContent'
import GlobalFooter from '@/components/GlobalFooter'
import Ads from '@/components/Other/CarbonAds'
import { getOnlineSuperSend, getSuperSendInfo } from '@/api/super_send'
import { mapState } from 'vuex'
import { resetRouter } from '@/router'

export default {
  name: 'BasicLayout',
  components: {
    SettingDrawer,
    RightContent,
    GlobalFooter,
    Ads
  },
  data () {
    return {
      // preview.pro.antdv.com only use.
      isProPreviewSite: process.env.VUE_APP_PREVIEW === 'true' && process.env.NODE_ENV !== 'development',
      // end
      isDev: process.env.NODE_ENV === 'development' || process.env.VUE_APP_PREVIEW === 'true',
      // base
      menus: [],
      // 侧栏收起状态
      collapsed: false,
      title: defaultSettings.title,
      mainMenuList: [],
      settings: {
        // 布局类型
        layout: defaultSettings.layout, // 'sidemenu', 'topmenu'
        // CONTENT_WIDTH_TYPE
        contentWidth: defaultSettings.layout === 'sidemenu' ? CONTENT_WIDTH_TYPE.Fluid : defaultSettings.contentWidth,
        // 主题 'dark' | 'light'
        theme: defaultSettings.navTheme,
        // 主色调
        primaryColor: defaultSettings.primaryColor,
        fixedHeader: defaultSettings.fixedHeader,
        fixSiderbar: defaultSettings.fixSiderbar,
        colorWeak: defaultSettings.colorWeak,

        hideHintAlert: false,
        hideCopyButton: false
      },
      // 媒体查询
      query: {},

      // 是否手机模式
      isMobile: false,
      nowSendInfo: {},
      sendID: 0
    }
  },
  computed: {
    ...mapState({
      updateMainMenu: state => state.app.updateMainMenuParent
    })
  },
  watch: {
    $route (to, from) {
      console.log('Route changed, new key:', to.meta.key)
      // 找到 '?' 的位置
      const questionMarkIndex = to.path.indexOf('?')

      // 检查是否存在 '?'
      if (questionMarkIndex !== -1) {
        // 提取从 '?' 开始到最后的字符串
        const queryString = to.path.substring(questionMarkIndex)
        console.log(queryString) // 输出类似 "?id=24"

        // 解析查询参数
        const urlParams = new URLSearchParams(queryString)
        const queryParams = Object.fromEntries(urlParams.entries())
        console.log(queryParams) // 输出类似 { id: "24" }

        this.sendID = queryParams.id

        if (this.sendID) {
          getSuperSendInfo(this.sendID).then(res => {
            if (res != null) {
              this.$store.commit(SET_SENDINFO, res.result)
              // this.nowSendInfo = res.result
              console.log('aas' + this.nowSendInfo.token)
            }
          })
        }
      } else {
        // 如果没有查询字符串，则将 sendID 设置为默认值或其他逻辑
        this.sendID = 0
        console.log('No query string found')
      }
    },
    updateMainMenu: {
      handler (val) {
        this.updateThisMainMenu()
      },
      deep: true,
      immediate: true
    }
  },
  created () {
    // 处理侧栏收起状态
    this.$watch('collapsed', () => {
      this.$store.commit(SIDEBAR_TYPE, this.collapsed)
    })
    this.$watch('isMobile', () => {
      this.$store.commit(TOGGLE_MOBILE_TYPE, this.isMobile)
    })
    console.log(this.mainMenuList)
  },
  mounted () {
    const userAgent = navigator.userAgent
    if (userAgent.indexOf('Edge') > -1) {
      this.$nextTick(() => {
        this.collapsed = !this.collapsed
        setTimeout(() => {
          this.collapsed = !this.collapsed
        }, 16)
      })
    }
    // first update color
    // TIPS: THEME COLOR HANDLER!! PLEASE CHECK THAT!!
    if (process.env.NODE_ENV !== 'production' || process.env.VUE_APP_PREVIEW === 'true') {
      updateTheme(this.settings.primaryColor)
    }
  },
  methods: {
    i18nRender,
    getThisOnlineSuperSend () {
      getOnlineSuperSend().then(
        res => {
          if (res == null) {
            return null
          }
          const lists = res.result
          const routes = asyncRouterMap.find((item) => item.path === '/')
          routes.children = []
          const staticRoutes = constantRouterMap.find((item) => item.path === '/')
          for (let i = 0; i < staticRoutes.children.length; i++) {
                routes.children.push(staticRoutes.children[i])
          }
          if (lists != null) {
            for (let i = 0; i < lists.length; i++) {
              const id = lists[i].id // 假设 lists[i]
              routes.children.push({
                path: `/super_send_conn${id}`,
                name: `super_send_conn_${id}`,
                component: RouteView,
                redirect: '/super_send_conn/message',
                meta: { title: lists[i].title, icon: 'table', permission: [] },
                children: [
                  {
                    path: `/super_send_conn${id}/message?id=${id}`,
                    name: `super_send_conn_${id}_message`,
                    component: () => import(`@/views/super_send_conn/message`),
                    meta: { title: '消息管理', forceReload: true, keepAlive: false, permission: [] }
                  },
                  {
                    path: `/super_send_conn${id}/send_info?id=${id}`,
                    name: `super_send_conn_${id}_send_info`,
                    component: () => import(`@/views/super_send_conn/send_info`),
                    meta: { title: '发送配置管理', keepAlive: false, permission: [] }
                  },
                  {
                    path: `/super_send_conn${id}/smtp?id=${id}`,
                    name: `super_send_conn_${id}_smtp`,
                    component: () => import(`@/views/super_send_conn/smtp`),
                    meta: { title: 'smtp服务器', keepAlive: false, permission: [] }
                  },
                  {
                    path: `/super_send_conn${id}/conf?id=${id}`,
                    name: `super_send_conn_${id}_conf`,
                    component: () => import(`@/views/super_send_conn/conf`),
                    meta: { title: '配置文件管理', keepAlive: false, permission: [] }
                  },
                  {
                    path: `/super_send_conn${id}/imap?id=${id}`,
                    name: `super_send_conn_${id}imap`,
                    component: () => import(`@/views/super_send_conn/imap`),
                    meta: { title: 'imap服务器', keepAlive: false, permission: [] }
                  },
                  {
                    path: `/super_send_conn${id}/receive?id=${id}`,
                    name: `super_send_conn_${id}_receive`,
                    component: () => import(`@/views/super_send_conn/receive`),
                    meta: { title: '接收配置管理', keepAlive: false, permission: [] }
                  },
                  {
                    path: `/super_send_conn${id}/server_user?id=${id}`,
                    name: `super_send_conn_${id}_server_user`,
                    component: () => import(`@/views/super_send_conn/server_user`),
                    meta: { title: '服务用户管理', keepAlive: false, permission: [] }
                  },
                  {
                    path: `/super_send_conn${id}/auth?id=${id}`,
                    name: `super_send_conn_${id}_auth`,
                    component: () => import(`@/views/super_send_conn/auth`),
                    meta: { title: '权限管理', keepAlive: false, permission: [] }
                  }
                ]
              })
              // RequestConFactory(lists[i])
              // this.$store.commit('SET_ROUTERS', routes)
            }
          }
          this.mainMenuList = []
          // const routes = this.mainMenu.find((item) => item.path === '/')
          this.mainMenuList = (routes && routes.children) || []
          this.$store.commit('SET_ROUTERS', routes)
          resetRouter()
          console.log(this.mainMenuList)
        }
      )
    },
    updateThisMainMenu () {
       this.getThisOnlineSuperSend()
    },
    handleMediaQuery (val) {
      this.query = val
      if (this.isMobile && !val['screen-xs']) {
        this.isMobile = false
        return
      }
      if (!this.isMobile && val['screen-xs']) {
        this.isMobile = true
        this.collapsed = false
        this.settings.contentWidth = CONTENT_WIDTH_TYPE.Fluid
        // this.settings.fixSiderbar = false
      }
    },
    handleCollapse (val) {
      this.collapsed = val
    },
    handleSettingChange ({ type, value }) {
      console.log('type', type, value)
      type && (this.settings[type] = value)
      switch (type) {
        case 'contentWidth':
          this.settings[type] = value
          break
        case 'layout':
          if (value === 'sidemenu') {
            this.settings.contentWidth = CONTENT_WIDTH_TYPE.Fluid
          } else {
            this.settings.fixSiderbar = false
            this.settings.contentWidth = CONTENT_WIDTH_TYPE.Fixed
          }
          break
      }
    }
  }
}
</script>

<style lang="less">
@import "./BasicLayout.less";
</style>
