import type { App } from "vue";

import {
    createRouter,
    createWebHashHistory,
    createWebHistory,
} from "vue-router";

const { VITE_HASH_ROUTE = "N", VITE_BASE_URL } = import.meta.env;

export const Routes = [
    { path: "/", component: () => import("@/views/home/index.vue") },
];

export const router = createRouter({
    history:
        VITE_HASH_ROUTE === "Y"
            ? createWebHashHistory(VITE_BASE_URL)
            : createWebHistory(VITE_BASE_URL),
    routes: Routes,
});

/** setup vue router. - [安装vue路由] */
export async function setupRouter(app: App) {
    app.use(router);
    await router.isReady();
}
