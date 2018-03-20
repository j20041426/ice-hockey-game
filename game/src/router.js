import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: "/",
            redirect: "/home"
        },
        {
            path: "/home",
            component: () => import("@/views/Home"),
            meta: { order: 1 }
        },
        {
            path: "/game/:code",
            component: () => import("@/views/Game"),
            meta: { order: 2 }
        }
    ]
});
