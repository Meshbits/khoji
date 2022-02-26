import Vue from 'vue';
import App from './App.vue';
import Main from './Main.vue';
import BlockView from './components/BlockView.vue';
import { BootstrapVue } from 'bootstrap-vue';
import VueRouter from 'vue-router';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

Vue.config.productionTip = false;

[VueRouter, BootstrapVue].forEach((x) => Vue.use(x));

const routes = [
  {
    path: "",
    component: Main,
  }
];

const router = new VueRouter({
  routes,
  mode: "history",
});

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");