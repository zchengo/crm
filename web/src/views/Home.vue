<template>
    <a-layout has-sider>
        <a-layout-sider class="layout-sider" width="150">
            <div class="logo">
                <div><img src="../assets/logo.svg"
                        style="width: 32px;height: 32px;filter: drop-shadow(2px 2px 6px #79bbff);" /></div>
                <div v-if="collapsed == false" class="title"><b>Z</b>O<b style="color: #1283FF;">C</b>RM</div>
            </div>
            <a-menu style="border-right: none;width: 149px;" v-model:selectedKeys="selectedKeys" mode="inline">
                <a-menu-item :key="item.key" v-for="item in menuItem">
                    <router-link :to="item.to" @click="store.selectedKeys = item.key" replace>
                        <component :is="item.icon" />
                        <span>{{ item.name }}</span>
                    </router-link>
                </a-menu-item>
            </a-menu>
        </a-layout-sider>
        <a-layout :style="{ marginLeft: '150px', background: '#FAFAFA' }">
            <a-layout-header class="header">
                <div style="display: flex;align-items: center;justify-items: center;margin-right: 10px;">
                    <QuestionCircleFilled @click="toDocs"
                        style="color: #909399;font-size: 18px;margin: 2px 15px 0 0;cursor: pointer;" />
                    <a-popover placement="bottomRight" :overlayStyle="{ width: '320px' }" trigger="click">
                        <template #content>
                            <div style="max-height: 250px;overflow-y: scroll;">
                                <a-list item-layout="horizontal" :data-source="data.noticeList" size="small">
                                    <template #renderItem="{ item }">
                                        <a-list-item style="cursor: pointer;" @click="onReadNotice(item.id)">
                                            <div style="display: inline-flex;align-items: center;">
                                                <a-avatar shape="square" :size="20" v-if="item.status == 1">
                                                    <template #icon>
                                                        <BellFilled style="font-size: 18px;" />
                                                    </template>
                                                </a-avatar>
                                                <a-avatar shape="square"
                                                    style="color: #476FFF; background-color: #ccd6fa" :size="20" v-else>
                                                    <template #icon>
                                                        <BellFilled style="font-size: 18px;" />
                                                    </template>
                                                </a-avatar>
                                                <div v-if="item.status == 1" style="color: #717171;">
                                                    &nbsp;&nbsp;&nbsp;{{ item.content }}</div>
                                                <div v-else>&nbsp;&nbsp;&nbsp;{{ item.content }}</div>
                                            </div>
                                            <template #actions>
                                                <span v-if="item.status == 2" style="color: black;">{{
                                                    formatDate(item.created)
                                                }}</span>
                                                <span v-else>{{ formatDate(item.created) }}</span>
                                            </template>
                                        </a-list-item>
                                    </template>
                                </a-list>
                            </div>
                            <div style="margin-top: 10px;display: flex;align-items: center;justify-content: center;">
                                <a-button v-if="data.noticeList.length != 0" @click="onDeleteNotice" type="primary"
                                    style="width: 92%;" shape="round">清空全部 {{ data.noticeList.length }} 条通知</a-button>
                            </div>
                        </template>
                        <a-badge :count="data.noticeCount">
                            <BellFilled style="color: #909399; font-size: 20px;cursor: pointer;" @click="onNotice" />
                        </a-badge>
                    </a-popover>
                    <a-dropdown :trigger="['click']">
                        <a-avatar @click="onUserAvatar" class="avatar" :size="28">U</a-avatar>
                        <template #overlay>
                            <a-menu style="width: 120px;">
                                <a-menu-item @click="visible = true">
                                    <ExclamationCircleOutlined /> 注销账号
                                </a-menu-item>
                                <a-menu-item @click="onLogout">
                                    <LogoutOutlined /> 退出登录
                                </a-menu-item>
                            </a-menu>
                        </template>
                    </a-dropdown>
                </div>
                <!-- 注销账号弹出框 -->
                <a-modal v-model:visible="visible" title="注销账号" @ok="onConfirm" @cancel="onCancel" cancelText="取消"
                    okText="注销" width="400px" :centered="true">
                    <a-alert message="账号注销后，会清空账号相关的所有数据。" type="warning" show-icon /><br />
                    <a-form :model="user" layout="vertical" @finish="onSubmit" :rules="rules">
                        <a-form-item name="email">
                            <a-input v-model:value="user.email" placeholder="邮箱" disabled />
                        </a-form-item>
                        <a-form-item name="code">
                            <a-input v-model:value="user.code" style="width: 55%;" placeholder="验证码" />
                            <a-button @click="onGetCode" style="width: 40%;float: right;" :loading="loading"
                                :disabled="disabled">
                                {{ buttonText }}</a-button>
                        </a-form-item>
                    </a-form>
                </a-modal>
            </a-layout-header>
            <a-layout-content :style="{ margin: '10px', background: '#fff', overflow: 'initial', borderRadius: '5px' }">
                <transition name="fade">
                    <router-view v-slot="{ Component }">
                        <component :is="Component" />
                    </router-view>
                </transition>
            </a-layout-content>
        </a-layout>
    </a-layout>
</template>

<script setup>
import { reactive, ref, onBeforeMount } from 'vue';
import { useRouter } from 'vue-router'
import { useStore } from '../store/index';
import { message } from 'ant-design-vue';
import { getUserInfo, getVerifyCode, userDelete } from '../api/user';
import { updateNotice, getNoticeCount, getNoticeList, deleteNotice } from '../api/notice';
import { DashboardOutlined, SmileOutlined, MehOutlined, ShoppingOutlined, ProfileOutlined, CrownOutlined } from '@ant-design/icons-vue';
import { QuestionCircleFilled, BellFilled, ExclamationCircleOutlined, LogoutOutlined } from '@ant-design/icons-vue';
import moment from 'moment'

// 菜单选项
const menuItem = reactive([{
    key: "dashboard",
    to: "/dashboard",
    icon: DashboardOutlined,
    name: "仪表盘"
}, {
    key: "customer",
    to: "/customer",
    icon: SmileOutlined,
    name: "客户"
}, {
    key: "contract",
    to: "/contract",
    icon: MehOutlined,
    name: "合同"
}, {
    key: "product",
    to: "/product",
    icon: ShoppingOutlined,
    name: "产品"
}, {
    key: "config",
    to: "/config",
    icon: ProfileOutlined,
    name: "配置"
}, {
    key: "subscribe",
    to: "/subscribe",
    icon: CrownOutlined,
    name: "订阅"
}])

// 表单校验
const rules = {
    email: [{
        required: true,
        message: '请输入邮箱!',
        trigger: 'blur',
    }, {
        pattern: /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/,
        message: '邮箱格式不正确',
        trigger: 'blur',
    }],
    code: [{ required: true, message: '请输入验证码!' }],
};

const store = useStore();

const selectedKeys = ref([store.selectedKeys])
const collapsed = ref(false)

const data = reactive({
    noticeCount: 0,
    noticeList: []
})

store.$subscribe((mutation, state) => {
    selectedKeys.value = [state.selectedKeys]
})

const router = useRouter()

const user = reactive({
    name: undefined,
    email: undefined,
    verison: undefined,
    code: undefined,
    versionText: undefined
})

const visible = ref(false)
const visibleLogo = ref(false)
const loading = ref(false)
const disabled = ref(false)
const buttonText = ref('获取验证码')

// 初始化数据
onBeforeMount(() => {
    store.selectedKeys = 'dashboard'
    router.push('dashboard')
    noticeCount()
})

// 点击用户头像
const onUserAvatar = () => {
    getUserInfo().then((res) => {
        if (res.data.code == 0) {
            user.name = res.data.data.name
            user.email = res.data.data.email
            user.version = res.data.data.version
        }
    })
}

// 跳转到项目文档
const toDocs = () => {
    window.open("https://docs.zocrm.cloud")
}

// 点击获取验证码
const onGetCode = () => {
    if (user.email == '') {
        message.warn('邮箱不能为空')
        return
    }
    loading.value = true
    let param = {
        email: user.email
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

// 点击确认注销账号
const onConfirm = () => {
    let param = {
        email: user.email,
        code: user.code
    }
    userDelete(param).then((res) => {
        if (res.data.code == 0) {
            router.push('/')
            message.success('账号已注销')
        }
    })
}

// 日期格式化
const formatDate = (timeStamp) => {
    let now = (Date.parse(new Date())) / 1000
    if (now - timeStamp < 60) {
        return "刚刚"
    }
    if ((new Date().getDate()) == (new Date(timeStamp * 1000).getDate())) {
        return "今天 " + moment(timeStamp * 1000).format('HH:mm')
    }
    return moment(timeStamp * 1000).format('YYYY-MM-DD')
}

// 点击读取通知
const onReadNotice = (id) => {
    updateNotice({ id: id }).then((res) => {
        if (res.data.code == 0) {
            onNotice()
            noticeCount()
        }
    })
}

// 获取通知数量
const noticeCount = () => {
    getNoticeCount().then((res) => {
        if (res.data.code == 0) {
            data.noticeCount = res.data.data.count
        }
    })
}

// 获取通知列表
const onNotice = () => {
    getNoticeList().then((res) => {
        if (res.data.code == 0) {
            data.noticeList = res.data.data
        }
    })
}

// 删除通知
const onDeleteNotice = () => {
    let ids = []
    for (let index = 0; index < data.noticeList.length; index++) {
        ids[index] = data.noticeList[index].id
    }
    deleteNotice({ ids: ids }).then((res) => {
        if (res.data.code == 0) {
            data.noticeList = res.data.data
            onNotice()
        }
    })
}

// 点击退出账号
const onLogout = () => {
    localStorage.removeItem("uid")
    localStorage.removeItem("token")
    router.push('/')
}

// 点击取消按钮
const onCancel = () => {
    disabled.value = false
    modalFormRef.value.resetFields()
    visible.value = false
}
</script>

<style scoped>
.layout-sider {
    left: 0;
    top: 0;
    bottom: 0;
    overflow: auto;
    height: 100vh;
    position: fixed;
    background: #fff;
    border-right: 0.5px solid #F0F2F5;
}

.header {
    padding: 0 12px;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    background: #fff;
}

.trigger {
    font-size: 18px;
    padding: 0 8px;
    cursor: pointer;
    transition: color 0.3s;
}

.trigger:hover {
    color: #476FFF;
}

.logo {
    height: 32px;
    margin: 16px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.avatar {
    color: #f56a00;
    background-color: #fde3cf;
    cursor: pointer;
    margin-left: 20px;
}

.popover {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 10px 0;
}

.title {
    font-size: 20px;
    color: rgba(31, 31, 31, 0.85);
    font-weight: 620;
    margin-left: 10px;
    overflow: hidden;
}
</style>