import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/main',
    children: [{
      path: 'main',
      name: 'Main',
      component: () => import('@/views/main/index'),
      meta: { title: 'メイン', icon: 'dashboard' }
    }]
  },
  {
    path: '/user',
    component: Layout,
    name: 'User',
    meta: { title: 'ユーザー', icon: 'user' },
    children: [
      {
        path: 'user',
        name: 'List',
        component: () => import('@/views/user/list'),
        meta: { title: 'リスト', icon: 'table' }
      },
      {
        path: 'user',
        name: 'Editor',
        component: () => import('@/views/user/editor'),
        meta: { title: '編集', icon: 'edit' }
      }
    ]
  },
  {
    path: '/article',
    component: Layout,
    name: 'Articles',
    meta: { title: '文書', icon: 'documentation' },
    children: [
      {
        path: 'article',
        name: 'Articles',
        component: () => import('@/views/article/list'),
        meta: { title: 'リスト', icon: 'table' }
      },
      {
        path: 'article',
        name: 'Editor',
        component: () => import('@/views/article/editor'),
        meta: { title: '編集', icon: 'edit' }
      }
    ]
  },
  {
    path: '/tag',
    component: Layout,
    name: 'Tags',
    meta: { title: 'タグ', icon: 'list' },
    children: [
      {
        path: 'tag',
        name: 'Tags',
        component: () => import('@/views/tag/list'),
        meta: { title: 'リスト', icon: 'table' }
      },
      {
        path: 'tag',
        name: 'Editor',
        component: () => import('@/views/tag/editor'),
        meta: { title: '編集', icon: 'edit' }
      }
    ]
  },
  {
    path: '/comment',
    component: Layout,
    children: [
      {
        path: 'comment',
        name: 'Comments',
        component: () => import('@/views/comment/list'),
        meta: { title: 'コメント', icon: 'message' }
      }
    ]
  },
  {
    path: 'external-link',
    component: Layout,
    children: [
      {
        path: 'http://zamberform.github.io/blog/',
        meta: { title: 'ブログ', icon: 'link' }
      }
    ]
  },
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
