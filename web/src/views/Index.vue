<template>
    <div class="container">
        <a-row>
            <a-col :span="14">
                <div class="back"></div>
            </a-col>
            <a-col :span="10">
                <div class="form">
                    <transition name="fade">
                        <router-view v-slot="{ Component }">
                            <component :is="Component" />
                        </router-view>
                    </transition>
                </div>
            </a-col>
        </a-row>
    </div>
</template>

<script setup>

</script>

<style scoped>
.container {
    width: 100%;
    height: 100vh;
    background: #FFFFFF;
}

.back {
    width: 100%;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #f9f9fb;
    background-image: url('../assets/back.png');
    background-size: 75% 75%;
    background-position: center;
    background-repeat: no-repeat;
}

.form {
    width: 100%;
    height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
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
