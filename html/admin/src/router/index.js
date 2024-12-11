import Vue from 'vue'
import Router from 'vue-router'
import { constantRouterMap, asyncRouterMap } from '@/config/router.config'

Vue.use(Router)

/* export default new Router({
  mode: 'history',
  routes: constantRouterMap.concat(asyncRouterMap)
}) */
const createRouter = () =>
  new Router({
    mode: 'history',
    routes: constantRouterMap
  })

const router = createRouter()
/* const router = new Router({
  mode: 'history',
  routes: constantRouterMap.concat(asyncRouterMap)
}) */
// 定义一个resetRouter 方法，在退出登录后或token过期后 需要重新登录时，调用即可
export function resetRouter () {
  // const newRouter = createRouter()
  // const newRouter = createRouter()
  const newRouter = new Router({
    mode: 'history',
    routes: asyncRouterMap
  })
  router.matcher = newRouter.matcher
}
export default router
