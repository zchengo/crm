<template>
    <a-form :model="formData" name="normal_login" class="login-form" @finish="onLogin" @finishFailed="onLoginFailed">
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

<script>
import { reactive } from 'vue';
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';
import { useRouter } from 'vue-router'
import { userLogin } from '../api/user';
import { message } from 'ant-design-vue';

export default {
    components: {
        UserOutlined,
        LockOutlined,
    },
    setup() {
        const router = useRouter()

        // 用户登录
        const formData = reactive({
            email: '1655064994@qq.com',
            password: '1655064994',
            remember: true,
        });
        const onLogin = () => {
            let param = {
                email: formData.email,
                password: formData.password
            }
            userLogin(param).then((res) => {
                if (res.data.code == 0) {
                    localStorage.setItem('uid', res.data.data.uid)
                    localStorage.setItem('ver', res.data.data.ver)
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

        return {
            formData,
            onLogin,
            onLoginFailed,
            forgotPass,
            toRegister
        };
    }
};
</script>

<style scoped>
.site-form-item-icon {
    color: rgba(0, 0, 0, .45);
}
</style>