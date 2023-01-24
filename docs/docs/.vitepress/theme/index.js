import Theme from 'vitepress/theme'
import './index.css'
import adsense from './adsense.vue'

export default {
  ...Theme,
  Layout: adsense
}