import Vue from "vue";
import App from "./App.vue";
import router from "./router";

import MuseUI from "muse-ui";
import "muse-ui/dist/muse-ui.css";
import "muse-ui/dist/theme-carbon.css";
Vue.use(MuseUI);

Vue.config.productionTip = false;

import axios from "axios";
Vue.prototype.$axios = axios;

new Vue({
    router,
    render: h => h(App)
}).$mount("#app");
