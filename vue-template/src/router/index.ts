import type { App } from "vue";
import { constantRoutes } from './routes';

import {
    createRouter,
    createWebHashHistory,
    createWebHistory,
} from "vue-router";

const { VITE_HASH_ROUTE = "N", VITE_BASE_URL } = import.meta.env;

export const router = createRouter({
    history:
        VITE_HASH_ROUTE === "Y"
            ? createWebHashHistory(VITE_BASE_URL)
            : createWebHistory(VITE_BASE_URL),
    routes: constantRoutes,
});

/** setup vue router. - [安装vue路由] */
export async function setupRouter(app: App) {
    app.use(router);
    await router.isReady();
}

export * from "./routes";
