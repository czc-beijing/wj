// app.js

import Token from "./models/token";
import { createStoreBindings } from "mobx-miniprogram-bindings";
import { timStore } from "./store/tim";

App({
    async onLaunch() {
        storeBindings.destroyStoreBindings()
    }
})
