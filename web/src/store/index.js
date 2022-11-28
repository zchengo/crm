import { createPinia, defineStore } from 'pinia'
import { ref } from 'vue'

const pinia = createPinia();

export const useSpinStore = defineStore('spin', () => {
  const spinning = ref(true)
  return { spinning }
})

export const spinStore = useSpinStore(pinia)