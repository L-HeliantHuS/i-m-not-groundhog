import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('./views/About.vue')
    },
    {
      path: '/pushvideo',
      name: 'pushVideo',
      component: () => import('./views/PostVideo.vue')
    },
    {
      path: '/video/:videoID',
      name: 'showVideo',
      component: () => import('./views/ShowVideo.vue')
    },
    {
      path: '/live/:uid',
      name: 'livehome',
      component: () => import('./views/LiveHome.vue')
    },
    {
      path: '/user/home',
      name: 'userhome',
      component: () => import('./views/user/Room.vue')
    },
    {
      path: '/user/register',
      name: 'register',
      component: () => import('./views/user/Register.vue')
    },
    {
      path: '/user/login',
      name: 'login',
      component: () => import('./views/user/Login.vue')
    },
  ]
})
