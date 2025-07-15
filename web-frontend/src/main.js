import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus from 'element-plus'
import router from '@/router'
import 'element-plus/dist/index.css'
import { showError } from '@/utils/errorHandler'

const app = createApp(App)
app.use(ElementPlus)
app.use(router)
app.config.globalProperties.$showError = showError
app.mount('#app')
