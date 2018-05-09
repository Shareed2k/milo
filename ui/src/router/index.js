import Vue from 'vue'
import Router from 'vue-router'
import store from '../store/modules/User'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      name: 'HomePage',
      component: require('@/components/HomePage').default,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/regions',
      name: 'RegionPage',
      component: require('@/components/RegionPage').default,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/datacenters',
      name: 'DataCenterPage',
      component: require('@/components/DataCenterPage').default,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/login',
      name: 'LoginPage',
      component: require('@/components/LoginPage').default,
      meta: {
        requiresAuth: false
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!store.state.isAuthorized) {
      next('/login')
    } else {
      next()
    }
  } else {
    if (store.state.isAuthorized && to.matched.some(record => {
      switch (record.name) {
        case 'LoginPage':
          return true
      }

      return false
    })) {
      next('/')
    } else {
      next()
    }
  }
})

export default router
