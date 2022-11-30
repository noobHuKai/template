import { createApp } from "vue";
import App from "./App.vue";
import { setupRouter } from "./router";
import { setupStore } from "./store";

function setupApp() {
    const app = createApp(App);

    // store plugin: pinia
    setupStore(app);

    // vue router
    setupRouter(app);

    // mount app
    app.mount("#app");
}

setupApp();
