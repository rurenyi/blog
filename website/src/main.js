import Vue from 'vue'
import App from './App.vue'

import VueRouter from 'vue-router';
import router from './router'
import axios from 'axios';

Vue.config.productionTip = false
let config;
if (process.env.VUE_APP_ENV === 'production') {
  config = require('./config/prod').default;
} else {
  config = require('./config/dev').default;
}
Vue.prototype.$config = config;

const BaseAxios = axios.create({
  "baseURL":config.serverUrl,
  withCredentials: true,
  timeout: 1000,
})
Vue.prototype.baseaxios = BaseAxios

BaseAxios.interceptors.response.use(
  (config) => {
    if (config.data.status === 403) {
      router.push("/login")
    }
    if (config.headers["auth_token"]) {
      localStorage.setItem("auth_token", config.headers["auth_token"])
    }
    return config
  },
)

BaseAxios.interceptors.request.use(
  (config) => {
    config.headers["auth_token"] = localStorage.getItem("auth_token")
    return config
  },
)

Vue.use(VueRouter)

new Vue({
  render: h => h(App),
  router:router
}).$mount('#app')
