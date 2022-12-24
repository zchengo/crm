import { createApp } from 'vue'
import App from './App.vue'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.less';
import router from './router/index';
import { createPinia } from 'pinia'

const pinia = createPinia()
const app = createApp(App)

app.use(Antd).use(router).use(pinia).mount('#app')
