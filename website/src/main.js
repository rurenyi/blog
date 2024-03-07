import Vue from 'vue'
import App from './App.vue'

import VueRouter from 'vue-router';
import router from './router'

Vue.config.productionTip = false
let config;
if (process.env.VUE_APP_ENV === 'production') {
  config = require('./config/prod').default;
} else {
  config = require('./config/dev').default;
}
Vue.prototype.$config = config;

Vue.use(VueRouter)

new Vue({
  render: h => h(App),
  router:router
}).$mount('#app')
