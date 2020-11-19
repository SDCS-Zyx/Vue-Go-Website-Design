import Vue from 'vue'
import Router from 'vue-router'
import Login from '../components/login.vue'
import Register from '../components/register.vue'
import Home from '../components/home.vue'

import Main from '../components/index/main.vue'
import Profile from '../components/index/profile.vue'
import Classinfo from '../components/index/classinfo.vue'
import Checkin from '../components/index/checkin.vue'

Vue.use(Router)

export default new Router({
 mode: 'history',
 routes: [
 {
   path: '/Login',
   name: 'Login',
   component: Login
 },
 {
   path: '/',
   component: Login
 },
 {
   path: '/Register',
   name: 'Register',
   component: Register
 },
 {
   path: '/Home',
   name: 'Home',
   component: Home,
   children: [
   {
    path: '/main',
    name: 'main',
    component: Main,
    hidden: true
  },
  {
    path: '/profile',
    name: 'profile',
    component: Profile,
    hidden: true
  },
  {
    path: '/classinfo',
    name: 'classinfo',
    component: Classinfo,
    hidden: true
  },
  {
    path: '/checkin',
    name: 'checkin',
    component: Checkin,
    hidden: true
  }
  ]
}
]
})