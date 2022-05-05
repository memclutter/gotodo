import {createApp} from 'vue'
import {createPinia} from 'pinia'
import VueDraggable from 'vuedraggable'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import 'element-plus/es/components/message/style/css'

import App from './App.vue'
import router from './router'

const app = createApp(App)
const pinia = createPinia()

pinia.use(piniaPluginPersistedstate)

app.use(pinia)
app.use(router)

app.component('vue-draggable', VueDraggable)

app.mount('#app')
