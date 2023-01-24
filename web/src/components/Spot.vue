<script setup>
import { ref } from 'vue';
const spotType = ref('spot-' + props.type)
const props = defineProps({
    type: String,
    title: String
})
</script>

<template>
    <div class="spot-container">
        <span class="spot-base" :class="spotType"></span>
        <span>{{ title }}</span>
    </div>
</template>

<style scoped>
.spot-container {
    display: flex;
    align-items: center;
    justify-content: flex-start;
}

.spot-base {
    width: 12px;
    height: 12px;
    margin-right: 3px;
    border-radius: 8px;
    display: inline-block;
}

.spot-success {
    background-color: #52c41a;
}

.spot-info {
    background-color: #1890ff;
}

.spot-warning {
    background-color: #faad14;
}

.spot-danger {
    background-color: #ff4d4f;
}
</style>