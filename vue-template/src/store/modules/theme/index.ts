import { defineStore } from "pinia";
import { darkTheme } from "naive-ui";
import { localStg } from "@/utils";
import { initThemeSettings } from "./helpers";

type ThemeState = Theme.Setting;

export const useThemeStore = defineStore("theme-store", {
    state: (): ThemeState => initThemeSettings(),
    getters: {
        /** naive-ui暗黑主题 */
        naiveTheme(state) {
            return state.darkMode ? darkTheme : undefined;
        },
    },
    actions: {
        /** 重置theme状态 */
        resetThemeStore() {
            localStg.remove("themeSettings");
            this.$reset();
        },
        /** 缓存主题配置 */
        cacheThemeSettings() {
            const isProd = import.meta.env.PROD;
            if (isProd) {
                localStg.set("themeSettings", this.$state);
            }
        },
        /** 设置暗黑模式 */
        setDarkMode(darkMode: boolean) {
            this.darkMode = darkMode;
        },
        /** 设置自动跟随系统主题 */
        setFollowSystemTheme(visible: boolean) {
            this.followSystemTheme = visible;
        },
        /** 自动跟随系统主题 */
        setAutoFollowSystemMode(darkMode: boolean) {
            if (this.followSystemTheme) {
                this.darkMode = darkMode;
            }
        },
        /** 切换/关闭 暗黑模式 */
        toggleDarkMode() {
            this.darkMode = !this.darkMode;
        },
    },
});
