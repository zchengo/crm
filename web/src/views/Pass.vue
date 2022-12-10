<template>
    <a-form ref="passFormRef" :model="formData" layout="vertical" @finish="onSubmit" :rules="rules">
        <a-form-item name="email">
            <a-input v-model:value="formData.email" size="large" placeholder="邮箱" />
        </a-form-item>
        <a-form-item name="code">
            <a-input v-model:value="formData.code" size="large" style="width: 55%;" placeholder="验证码" />
            <a-button @click="onGetCode" size="large" :disabled="disabled" :loading="loading"
                style="width: 40%;float: right;">
                {{ buttonText }}</a-button>
        </a-form-item>
        <a-form-item name="password1">
            <a-input-password v-model:value="formData.password1" size="large" placeholder="密码" />
        </a-form-item>
        <a-form-item name="password2">
            <a-input-password v-model:value="formData.password2" size="large" placeholder="确认密码" />
        </a-form-item>
        <a-form-item>
            <a-button type="primary" html-type="submit" size="large" style="width: 50%;">确定</a-button>
            <a-button type="link" style="width: 50%;" @click="onLogin">使用已有账户登录</a-button>
        </a-form-item>
    </a-form>
</template>

<script>
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { userForgotPass, getVerifyCode } from '../api/user'
import { message } from 'ant-design-vue';

export default {
    setup() {
        const router = useRouter()

        // 重置密码
        const formData = reactive({
            email: '',
            code: '',
            password1: '',
            password2: '',
        });

        // 表单校验
        const rules = {
            email: [{
                required: true,
                message: '请输入邮箱!'
            }, {
                pattern: /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/,
                message: '邮箱格式不正确',
                trigger: 'blur',
            }],
            code: [{ required: true, message: '请输入验证码!' }],
            password1: [{ required: true, message: '请输入密码!' }],
            password2: [{ required: true, message: '请输入密码!' }],
        };

        const loading = ref(false)
        const disabled = ref(false)
        const passFormRef = ref()
        const buttonText = ref('获取验证码')

        const onSubmit = () => {
            passFormRef.value.validateFields().then(() => {
                let param = {
                    email: formData.email,
                    code: formData.code,
                    password: formData.password2,
                }
                userForgotPass(param).then((res) => {
                    if (res.data.code == 0) {
                        message.success('密码已重置')
                        router.push("/login")
                    }
                    if (res.data.code == 10005) {
                        message.error('验证码错误');
                    }
                })
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
            rules,
            passFormRef,
            loading,
            disabled,
            buttonText,
            onSubmit,
            onGetCode,
            onLogin,
        };
    },
}
</script>

<style scoped>

</style>