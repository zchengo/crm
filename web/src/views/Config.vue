<template>
    <div :style="{ padding: '0 20px 12px 20px' }">
        <a-tabs v-model:activeKey="activeKey">
            <a-tab-pane key="1" tab="邮件服务配置">
                <a-card size="small" title="配置STMP服务" :headStyle="{ padding: '0 0' }" :bodyStyle="{ padding: '10px 0' }"
                    :bordered="false">
                    <template #extra><a-switch v-model:checked="mailConfig.status" :checkedValue="1" :unCheckedValue="2"
                            @change="onSwitch" /></template>
                    <a-alert message="您需要配置SMTP服务器才能使用邮件发送功能，目前仅支持企鹅邮箱或网易邮箱提供的免费SMTP服务器。" type="info" show-icon /><br />
                    <a-form ref="mailConfigRef" :model="mailConfig" layout="vertical" name="mailConfig" :rules="rules"
                        @finish="onSave">
                        <a-row :gutter="50">
                            <a-col :span="12">
                                <a-form-item label="STMP服务器" name="stmp">
                                    <a-input v-model:value="mailConfig.stmp" placeholder="请输入域名">
                                        <template #suffix>
                                            <a-tooltip title="如 smtp.qq.com">
                                                <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
                                            </a-tooltip>
                                        </template>
                                    </a-input>
                                </a-form-item>
                            </a-col>
                            <a-col :span="12">
                                <a-form-item label="端口号" name="port">
                                    <a-input-number v-model:value="mailConfig.port" placeholder="请输入端口号"
                                        :controls="false" style="width: 100%;">
                                        <template #suffix>
                                            <a-tooltip title="STMP服务对应的端口号">
                                                <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
                                            </a-tooltip>
                                        </template>
                                    </a-input-number>
                                </a-form-item>
                            </a-col>
                            <a-col :span="12">
                                <a-form-item label="授权码" name="authCode">
                                    <a-input v-model:value="mailConfig.authCode" placeholder="请输入授权码">
                                        <template #suffix>
                                            <a-tooltip title="从STMP服务器提供商获得的授权码">
                                                <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
                                            </a-tooltip>
                                        </template>
                                    </a-input>
                                </a-form-item>
                            </a-col>
                            <a-col :span="12">
                                <a-form-item label="邮箱账号" name="email">
                                    <a-input v-model:value="mailConfig.email" placeholder="请输入邮箱">
                                        <template #suffix>
                                            <a-tooltip title="开启STMP服务时所使用的邮箱账号">
                                                <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
                                            </a-tooltip>
                                        </template>
                                    </a-input>
                                </a-form-item>
                            </a-col>
                            <a-col :span="12">
                                <a-form-item label="服务有效性检测" name="usability">
                                    <a-select ref="select" v-model:value="usability" @focus="focus"
                                        @change="handleChange" placeholder="自动检测服务有效性" disabled>
                                        <a-select-option :value="1">
                                            <Spot type="info" title="正在检测，请勿离开页面！" />
                                        </a-select-option>
                                        <a-select-option :value="2">
                                            <Spot type="success" title="服务配置有效" />
                                        </a-select-option>
                                        <a-select-option :value="3">
                                            <Spot type="danger" title="配置无效，请检查STMP服务器、端口号、授权码和邮箱账号是否匹配。" />
                                        </a-select-option>
                                    </a-select>
                                </a-form-item>
                            </a-col>
                            <a-col :span="12">
                                <a-form-item label="服务状态" name="status">
                                    <a-select ref="select" v-model:value="mailConfig.status" placeholder="自动检测服务状态"
                                        disabled>
                                        <a-select-option :value="1">
                                            <Spot type="success" title="服务已开启" />
                                        </a-select-option>
                                        <a-select-option :value="2">
                                            <Spot type="danger" title="服务已关闭" />
                                        </a-select-option>
                                    </a-select>
                                </a-form-item>
                            </a-col>
                            <a-col :span="12">
                                <a-form-item><br />
                                    <a-button type="primary" html-type="submit">保存</a-button>
                                    <a-button type="primary" style="margin-left: 20px;" @click="reset" ghost>重置</a-button>
                                    <a-button style="margin-left: 20px;" @click="delConfig" danger>删除</a-button>
                                </a-form-item>
                            </a-col>
                        </a-row>
                    </a-form>
                </a-card>
            </a-tab-pane>
        </a-tabs>
    </div>
</template>

<script setup>
import { reactive, ref, onBeforeMount, createVNode } from 'vue';
import { InfoCircleOutlined, ExclamationCircleOutlined } from '@ant-design/icons-vue';
import Spot from '../components/Spot.vue';
import { saveMailConfig, deleteMailConfig, updateMailConfigStmpStatus, getMailConfig, checkMailConfig } from '../api/config';
import { message, Modal } from 'ant-design-vue';

// 初始化
onBeforeMount(() => {
    getMailConfigInfo()
})

// 邮件服务配置
const mailConfig = reactive({
    id: undefined,
    stmp: undefined,
    port: undefined,
    authCode: undefined,
    email: undefined,
    status: undefined
})
const mailConfigRef = ref()
const usability = ref(undefined)

// 表单校验
const rules = {
    stmp: [{
        required: true,
        pattern: /^(?=^.{3,255}$)(http(s)?:\/\/)?(www\.)?[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+(:\d+)*(\/\w+\.\w+)*$/,
        message: '域名格式不正确',
        trigger: 'blur',
    }],
    port: [{
        required: true,
        pattern: /^([0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{4}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$/,
        message: '端口号格式不正确',
        trigger: 'blur',
    }],
    authCode: [{ required: true, message: '产品编码格式不正确', trigger: 'blur' }],
    email: [{
        required: true,
        pattern: /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/,
        message: '邮箱格式不正确',
        trigger: 'blur',
    }]
};

// 保存邮件服务配置
const onSave = () => {
    saveMailConfig(mailConfig).then((res) => {
        if (res.data.code == 0) {
            setData(res.data.data)
            message.success("保存成功")
            checkConfig()
        } else {
            message.error("保存失败")
        }
    })
}

// 删除配置
const delConfig = () => {
    Modal.confirm({
        title: '您确定要删除邮件服务配置？',
        icon: createVNode(ExclamationCircleOutlined),
        centered: true,
        cancelText: '取消',
        okText: '确定',
        onOk() {
            deleteMailConfig({ id: mailConfig.id }).then((res) => {
                if (res.data.code == 0) {
                    message.success("删除成功")
                    reset()
                } else {
                    message.error("删除失败")
                }
            })
        },
        onCancel() {
            console.log('Cancel');
        },
    });
}

// 开启或关闭邮件服务
const onSwitch = () => {
    updateMailConfigStmpStatus({
        id: mailConfig.id,
        status: mailConfig.status
    }).then((res) => {
        if (res.data.code == 0) {
            if (mailConfig.status == 1) {
                message.success("服务已开启")
            } else {
                message.warning("服务已关闭")
            }
        }
        if (res.data.code == 50004) {
            mailConfig.status = undefined
            message.info("您还没有配置邮件服务！")
        }
    })
}

// 获取邮件服务配置
const getMailConfigInfo = () => {
    getMailConfig().then((res) => {
        if (res.data.code == 0) {
            setData(res.data.data)
            checkConfig()
        }
        if (res.data.code == 50004) {
            usability.value = undefined
            mailConfig.status = undefined
        }
    })
}

// 检查邮件服务配置
const checkConfig = () => {
    usability.value = 1
    checkMailConfig().then((res) => {
        if (res.data.code == 0) {
            usability.value = 2
        } else {
            usability.value = 3
        }
    })
}

// 重置
const reset = () => {
    mailConfigRef.value.resetFields()
    usability.value = undefined
}

// 数据赋值
const setData = data => {
    mailConfig.id = data.id
    mailConfig.stmp = data.stmp
    mailConfig.port = data.port
    if (data.port === 0) {
        mailConfig.port = undefined
    }
    mailConfig.authCode = data.authCode
    mailConfig.email = data.email
    mailConfig.status = data.status
    mailConfig.usability = data.usability
}
</script>