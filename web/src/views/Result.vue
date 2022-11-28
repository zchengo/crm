<template>
    <div style="margin-top: 50px;">
        <a-result status="403" sub-title="您需要订阅个人版才能使用该功能哦！">
            <template #extra>
                <a-button type="primary" @click="onBuy">立即订阅</a-button>
            </template>
        </a-result>
    </div>
</template>

<script>
import { useRouter } from 'vue-router'

export default {
    setup() {

        const router = useRouter()
        
        const onBuy = () => {
            router.push('/subscribe')
        }
        return {
            onBuy
        }
    }
}
</script>

<style>

</style>