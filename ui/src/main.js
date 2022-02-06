import Vue from 'vue';
import App from './App.vue';
import Main from './Main.vue';

Vue.config.productionTip = false;

new Vue({
  render: (h) => h(App),
}).$mount("#app");