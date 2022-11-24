import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { NaiveUiResolver } from "unplugin-vue-components/resolvers";
import { resolve } from "path";
import DefineOptions from "unplugin-vue-define-options/vite";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        DefineOptions(),
        AutoImport({
            imports: [
                "vue",
                {
                    "naive-ui": [
                        "useDialog",
                        "useMessage",
                        "useNotification",
                        "useLoadingBar",
                    ],
                },
            ],
        }),
        Components({
            resolvers: [NaiveUiResolver()],
        }),
    ],
    server: {
        proxy: {
            "/api": {
                target: "http://127.0.0.1:8080/",
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, ""),
            },
        },
    },
    resolve: {
        alias: [
            {
                find: "@",
                replacement: resolve(__dirname, "./src"),
            },
            {
                find: "~",
                replacement: resolve(process.cwd()),
            },
        ],
    },
    build: {
        minify: "terser",
        terserOptions: {
            compress: {
                // 生产环境时移除console
                drop_console: true,
                drop_debugger: true,
            },
        },
        chunkSizeWarningLimit: 1500,
        rollupOptions: {
            output: {
                chunkFileNames: "static/js/[name]-[hash].js",
                entryFileNames: "static/js/[name]-[hash].js",
                assetFileNames: "static/[ext]/[name]-[hash].[ext]",
            },
        },
    },
});
