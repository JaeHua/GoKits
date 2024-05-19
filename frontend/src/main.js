import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from '@/router'
import store from './store'
import App from './App.vue'

const app = createApp(App)
app.config.devtools = true
// app.config.globalProperties.$axios=axios;
app.use(ElementPlus)
app.use(router)
app.use(store)

app.mount('#app')