<script setup>
import DefaultTheme from 'vitepress/theme'
const { Layout } = DefaultTheme

</script>

<template>
  <Layout>
    <template #aside-ads-before>
        <div class="contaniner">
            <img src="../../public/sponsor.png"/>
            <p style="font-size: 10px;margin: 20px 0 5px;">微信扫一扫赞赏码，赞赏支持作者！</p>
        </div>
    </template>
  </Layout>
</template>

<style scoped>
.contaniner{
    margin-top: 24px;
    display: flex;
    align-items: center;
    flex-direction: column;
    border-radius: 12px;
    min-height: 256px;
    padding: 15px;
    text-align: center;
    line-height: 18px;
    font-size: 12px;
    font-weight: 500;
    background-color: var(--vp-c-bg-soft);
}

.contaniner img {
  border-radius: 10px;
}
</style>