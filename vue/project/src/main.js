// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import Qs from 'qs'

import 'default-passive-events'
 
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
 
// 注册element-ui
Vue.use(ElementUI)
 
Vue.config.productionTip = false

Vue.prototype.axios = axios
Vue.prototype.qs = Qs

axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
 
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})