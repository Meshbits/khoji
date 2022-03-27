import Vue from 'vue';
import App from './App.vue';
import Main from './Main.vue';
import BlockView from './components/BlockView.vue';
import TransactionView from './components/TransactionView.vue';
import IdentityView from './components/IdentityView.vue';
import AddressView from './components/AddressView.vue';
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
  }, {
    path: "/block/:height",
    component: BlockView,
  }, {
    path: "/transaction/:hash",
    component: TransactionView,
  }, {
    path: "/identity/:name",
    component: IdentityView,
  }, {
    path: "/address/:address",
    component: AddressView,
  },
];

const router = new VueRouter({
  routes,
  mode: "history",
});

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");