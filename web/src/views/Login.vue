<template>
    <a-form :model="formData" @finish="onLogin" @finishFailed="onLoginFailed" style="width: 65%;">
        <a-form-item>
            <div class="title">账号登录</div>
        </a-form-item>
        <a-form-item name="email" :rules="[{ required: true, message: '请输入邮箱!' }]">
            <a-input v-model:value="formData.email" size="large" placeholder="邮箱">
                <template #prefix>
                    <UserOutlined class="site-form-item-icon" />
                </template>
            </a-input>
        </a-form-item>
        <a-form-item name="password" :rules="[{ required: true, message: '请输入密码!' }]">
            <a-input-password v-model:value="formData.password" size="large" placeholder="密码">
                <template #prefix>
                    <LockOutlined class="site-form-item-icon" />
                </template>
            </a-input-password>
        </a-form-item>
        <a-form-item>
            <a-form-item name="remember" no-style>
                <a-checkbox v-model:checked="formData.remember" style="float: left;">记住我</a-checkbox>
            </a-form-item>
            <a class="login-form-forgot" style="float: right;" @click="forgotPass">忘记密码？</a>
        </a-form-item>
        <a-form-item>
            <a-button size="large" type="primary" html-type="submit" class="login-form-button" style="width: 100%">
                登录
            </a-button><br />
        </a-form-item>
        <a-form-item>
            还没有账号？<a @click="toRegister"> 立即注册</a>
        </a-form-item>
    </a-form>
</template>

<script setup>
import { reactive, onMounted } from 'vue';
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';
import { useRouter } from 'vue-router';
import { userLogin } from '../api/user';
import { message } from 'ant-design-vue';
import { initDatabase } from '../api/common';

// 初始化数据库（只会在生产环境中初始化）
onMounted(() => {
    if (formData.email == '1655064994@qq.com') {
        message.loading('正在初始化数据...', 2.5).then(() => {
            initDatabase().then((res) => {
                if (res.data.code == 0) {
                    message.success('初始化成功！', 2)
                }
                if (res.data.code == 10) {
                    message.error('初始化失败！')
                }
            })
        })
    }
})

const router = useRouter()

const formData = reactive({
    email: '1655064994@qq.com',
    password: '1655064994',
    remember: true,
})

// 用户登录
const onLogin = () => {
    let param = {
        email: formData.email,
        password: formData.password
    }
    userLogin(param).then((res) => {
        if (res.data.code == 0) {
            localStorage.setItem('uid', res.data.data.uid)
            localStorage.setItem('token', res.data.data.token)
            router.push("/home")
        }
        if (res.data.code == 10002) {
            message.error('用户不存在');
        }
        if (res.data.code == 10003) {
            message.error('用户名或密码错误');
        }
    })
};
const onLoginFailed = errorInfo => {
    console.log('Failed:', errorInfo);
};

// 忘记密码
const forgotPass = () => {
    router.push("/pass")
}

// 用户注册
const toRegister = () => {
    router.push("/register")
}
</script>

<style scoped>
.title {
    color: #303133;
    font-size: 30px;
    line-height: 65px;
    font-weight: 500;
}

.site-form-item-icon {
    color: rgba(0, 0, 0, .45);
}
</style>