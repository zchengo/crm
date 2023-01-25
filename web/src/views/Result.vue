<template>
    <div style="margin-top: 50px;">
        <a-result status="403" sub-title="您需要订阅专业版或高级版才能使用该功能哦！">
            <template #extra>
                <a-button type="primary" @click="onBuy">立即订阅</a-button>
            </template>
        </a-result>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useStore } from '../store/index';

const router = useRouter()
const store = useStore()

const onBuy = () => {
    store.selectedKeys = 'subscribe'
    router.push('/subscribe')
}
</script>

<style>

</style>