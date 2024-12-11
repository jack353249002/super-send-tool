// eslint-disable-next-line
import { UserLayout, BasicLayout, BlankLayout } from '@/layouts'

const RouteView = {
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
            meta: { title: '用户列表', keepAlive: true, permission: [] }
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
            meta: { title: '用户列表', keepAlive: true, permission: [] }
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
