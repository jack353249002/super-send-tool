// eslint-disable-next-line
import { UserLayout, BasicLayout, BlankLayout, PageView } from '@/layouts'

export const RouteView = {
  name: 'RouteView',
  render: h => h('router-view')
}
export const asyncRouterMap = [
  {
    path: '/',
    name: 'index',
    component: BasicLayout,
    meta: { title: 'menu.home' },
    redirect: '/super_send',
    // 你需要动态引入的页面组件
    super_send_users: () => import('@/views/super_send/users'),
    super_send_message: () => import('@/views/super_send_conn/message'),
    children: [
    ]
  },
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/',
    name: 'index',
    component: BasicLayout,
    meta: { title: 'menu.home' },
    redirect: '/super_send',
    children: [
      // list
      {
        path: '/super_send',
        name: 'super_send',
        component: RouteView,
        redirect: '/super_send/users',
        meta: { title: 'super-send工具', icon: 'table', permission: [] },
        children: [
          {
            path: '/super_send/users',
            name: 'super_send_users',
            component: () => import('@/views/super_send/users'),
            meta: { title: '节点列表', keepAlive: true, permission: [] }
          }
        ]
      }
    ]
  },
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]
