import {createApp} from 'vue'
import {createPinia} from 'pinia'
import VueDraggable from 'vuedraggable'

import 'element-plus/es/components/message/style/css'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.component('vue-draggable', VueDraggable)

app.mount('#app')
