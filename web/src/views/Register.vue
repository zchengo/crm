<template>
    <a-form :model="formData" layout="vertical" @finish="onRegister">
        <a-form-item name="email" :rules="[{ required: true, message: '请输入邮箱!' }]">
            <a-input v-model:value="formData.email" size="large" placeholder="邮箱">
            </a-input>
        </a-form-item>
        <a-form-item name="code" :rules="[{ required: true, message: '请输入验证码!' }]">
            <a-input v-model:value="formData.code" size="large" style="width: 55%;" placeholder="验证码" />
            <a-button @click="onGetCode" size="large" style="width: 40%;float: right;" :loading="loading" :disabled="disabled">
                {{ buttonText }}</a-button>
        </a-form-item>
        <a-form-item name="password1" :rules="[{ required: true, message: '请输入密码!' }]">
            <a-input-password v-model:value="formData.password1" size="large" placeholder="密码" />
        </a-form-item>
        <a-form-item name="password2" :rules="[{ required: true, message: '请输入密码!' }]">
            <a-input-password v-model:value="formData.password2" size="large" placeholder="确认密码" />
        </a-form-item>
        <a-form-item>
            <a-button type="primary" html-type="submit" size="large" style="width: 50%;">注册</a-button>
            <a-button type="link" style="width: 50%;" @click="onLogin">使用已有账户登录</a-button>
        </a-form-item>
    </a-form>
</template>

<script>
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { userRegister, getVerifyCode } from '../api/user';
import { message } from 'ant-design-vue';

export default {
    setup() {
        const router = useRouter()

        // 用户注册
        const formData = reactive({
            email: '',
            code: '',
            password1: '',
            password2: '',
        });

        const loading = ref(false)
        const disabled = ref(false)
        const buttonText = ref('获取验证码')

        const onRegister = () => {
            if (formData.password1 != formData.password2) {
                message.info('密码不一致');
                return
            }
            let param = {
                email: formData.email,
                code: formData.code,
                password: formData.password2,
            }
            userRegister(param).then((res) => {
                if (res.data.code == 0) {
                    message.success('注册成功');
                    onLogin()
                }
                if (res.data.code == 10001) {
                    message.warn('该用户已经存在');
                }
                if (res.data.code == 10005) {
                    message.error('验证码错误');
                }
            })
        };

        // 获取验证码
        const onGetCode = () => {
            if (formData.email == '') {
                message.warn('邮箱不能为空')
                return
            }
            loading.value = true
            let param = {
                email: formData.email
            }
            getVerifyCode(param).then((res) => {
                if (res.data.code == 0) {
                    loading.value = false
                    disabled.value = true
                    buttonText.value = '验证码已发送'
                }
                if (res.data.code == 10004) {
                    loading.value = false
                    message.error('验证码发送失败')
                }
            })
        }

        // 跳转到登录页面
        const onLogin = () => {
            router.push("/login")
        }

        return {
            formData,
            loading,
            disabled,
            buttonText,
            onRegister,
            onLogin,
            onGetCode,
        };
    },
}
</script>

<style scoped>

</style>