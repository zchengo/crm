<template>
    <div class="container">
        <div class="header">
            <img src="../assets/logo.svg" style="width: 60px;height: 60px;filter: drop-shadow(2px 2px 6px #79bbff);" />
            <span class="title">ZO<b style="color: #1283FF;">C</b>RM</span>
        </div>
        <div class="content">
            <div class="main">
                <transition name="fade">
                    <router-view v-slot="{ Component }">
                        <component :is="Component" />
                    </router-view>
                </transition>
            </div>
            <div class="footer">
                <div class="links">
                    <a href="https://github.com/zchengo/crm">开源地址</a>
                    <a href="https://github.com/zchengo/crm/blob/main/LICENSE">许可证</a>
                    <a href="https://mail.qq.com/">企鹅邮箱</a>
                </div>
                <div class="copyright">Copyright © 2022 www.zocrm.cloud</div>
            </div>
        </div>
    </div>
</template>

<script>
export default {

}
</script>

<style scoped>
.container {
    width: 100%;
    height: 100vh;
    background: #F0F2F5;
}

.header {
    width: 100%;
    height: 30%;
    display: flex;
    align-items: center;
    justify-content: center;
}

.content {
    width: 100%;
}

.main {
    width: 368px;
    margin: 0 auto;
}

.footer {
    width: 100%;
    bottom: 0;
    padding: 0 16px;
    margin: 80px 0 24px;
    text-align: center;
    color: rgba(0, 0, 0, .45);
    font-size: 14px;
}

.links {
    margin-bottom: 8px;
}

.links a {
    color: rgba(0, 0, 0, .45);
    padding: 5px 20px;
}

.links a:hover {
    color: #5b5b5c;
}

.copyright {
    padding: 10px;
}

.title {
    font-size: 42px;
    color: rgba(31, 31, 31, 0.85);
    font-family: Avenir, Helvetica Neue, Arial, Helvetica, sans-serif;
    font-weight: 600;
    position: relative;
    top: 2px;
    padding: 20px;
    letter-spacing: 1px;
}
</style>
