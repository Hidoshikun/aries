import layoutHeaderAside from '@/layout/header-aside'

// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

/**
 * 在主框架内显示
 */
const frameIn = [
  {
    path: '/',
    redirect: { name: 'index' },
    component: layoutHeaderAside,
    children: [
      // 首页
      {
        path: 'index',
        name: 'index',
        meta: {
          title: '首页',
          auth: true
        },
        component: _import('aries/system/index')
      },
      {
        path: 'post/category',
        name: 'postCategory',
        meta: {
          title: '文章分类',
          auth: true
        },
        component: _import('aries/category/postCategory')
      },
      {
        path: 'tag',
        name: 'tag',
        meta: {
          title: '标签',
          auth: true
        },
        component: _import('aries/tag/tag')
      },
      {
        path: 'post',
        name: 'post',
        meta: {
          title: '文章',
          auth: true
        },
        component: _import('aries/post/post')
      },
      {
        path: 'comment',
        name: 'comment',
        meta: {
          title: '评论',
          auth: true
        },
        component: _import('aries/comment/comment')
      },
      {
        path: 'page',
        name: 'page',
        meta: {
          title: '页面',
          auth: true
        },
        component: _import('aries/page/page')
      },
      {
        path: 'page/journal',
        name: 'journal',
        meta: {
          title: '日志',
          auth: true
        },
        component: _import('aries/page/journal')
      },
      {
        path: 'page/gallery',
        name: 'gallery',
        meta: {
          title: '图库',
          auth: true
        },
        component: _import('aries/page/gallery')
      },
      {
        path: 'user',
        name: 'user',
        meta: {
          title: '用户信息',
          auth: true
        },
        component: _import('aries/user/user')
      },
      {
        path: 'setting',
        name: 'setting',
        meta: {
          title: '设置',
          auth: true
        },
        component: _import('aries/system/setting')
      },
      {
        path: 'link',
        name: 'links',
        meta: {
          title: '友链',
          auth: true
        },
        component: _import('aries/link/links')
      },
      {
        path: 'link/category',
        name: 'linkCategory',
        meta: {
          title: '友链分类',
          auth: true
        },
        component: _import('aries/category/linkCategory')
      },
      {
        path: 'nav',
        name: 'navs',
        meta: {
          title: '菜单',
          auth: true
        },
        component: _import('aries/nav/navs')
      },
      {
        path: 'attachment',
        name: 'attachment',
        meta: {
          title: '附件',
          auth: true
        },
        component: _import('aries/system/attachment')
      },
      {
        path: 'theme',
        name: 'theme',
        meta: {
          title: '主题',
          auth: true
        },
        component: _import('aries/theme/theme')
      },
      {
        path: 'doc',
        name: 'doc',
        meta: {
          title: 'API 文档',
          auth: true
        },
        component: _import('aries/system/doc')
      },
      // 系统 前端日志
      {
        path: 'log',
        name: 'log',
        meta: {
          title: '前端日志',
          auth: true
        },
        component: _import('system/log')
      },
      // 刷新页面 必须保留
      {
        path: 'refresh',
        name: 'refresh',
        hidden: true,
        component: _import('system/function/refresh')
      },
      // 页面重定向 必须保留
      {
        path: 'redirect/:route*',
        name: 'redirect',
        hidden: true,
        component: _import('system/function/redirect')
      }
    ]
  }
]

/**
 * 在主框架之外显示
 */
const frameOut = [
  // 登录
  {
    path: '/login',
    name: 'login',
    meta: {
      title: '登录',
      auth: false // 表明无需登录验证
    },
    component: _import('aries/auth/login')
  }
]

/**
 * 错误页面
 */
const errorPage = [
  {
    path: '*',
    name: '404',
    component: _import('system/error/404')
  }
]

// 导出需要显示菜单的
export const frameInRoutes = frameIn

// 重新组织后导出
export default [
  ...frameIn,
  ...frameOut,
  ...errorPage
]
