<template>
  <a-layout style="min-height: 100vh">
    <a-layout-sider width="180" class="sider" v-model:collapsed="collapsed" :trigger="null" collapsible>
      <div class="logo">
        <div><img src="../assets/logo.svg" style="width: 32px;height: 32px;filter: drop-shadow(2px 2px 6px #79bbff);" />
        </div>
        <div v-if="collapsed == false" class="title"><b>Z</b>O<b style="color: #1283FF;">C</b>RM</div>
      </div>
      <a-menu style="border-right: none;" v-model:selectedKeys="selectedKeys" mode="inline">
        <a-menu-item :key="item.key" v-for="item in menuItem">
          <router-link :to="item.to" @click="store.selectedKeys = item.key" replace>
            <component :is="item.icon" />
            <span>{{ item.name }}</span>
          </router-link>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="header">
        <div>
          <menu-unfold-outlined v-if="collapsed" class="trigger" @click="() => (collapsed = !collapsed)" />
          <menu-fold-outlined v-else class="trigger" @click="() => (collapsed = !collapsed)" />
        </div>
        <div style="display: flex;align-items: center;justify-items: center;">
          <a-tooltip title="帮助">
            <a-button type="text" shape="circle">
              <template #icon>
                <QuestionCircleFilled style="color: #909399; font-size: 18px;" />
              </template>
            </a-button>
          </a-tooltip>
          <a-tooltip title="消息">
            <a-button type="text" shape="circle">
              <template #icon>
                <BellFilled style="color: #909399; font-size: 20px;" />
              </template>
            </a-button>
          </a-tooltip>
          <a-popover trigger="click" v-model:visible="popoverVisible" placement="bottomRight">
            <template #title>
              <div
                style="display: flex;align-items: center;justify-content:space-between;flex-direction: row;padding: 10px 0;">
                <a-avatar @click="toMine" class="avatar" :size="50">U</a-avatar>
                <div
                  style="display: flex;align-items:flex-start;flex-direction: column;margin: 0 10px;font-weight: normal;">
                  <span>{{ user.email }}</span>
                  <sapn style="font-size: 12px;color: #476FFF;"><crown-two-tone style="font-size:15px;"
                      two-tone-color="#476FFF" />&nbsp;{{ user.versionText }}</sapn>
                </div>
              </div>
            </template>
            <template #content>
              <div style="display: flex;align-items: center;justify-content:space-between;flex-direction: row;">
                <a-button type="text" style="color: #476FFF;" @click="onDelete">注销账号</a-button>
                <a-divider type="vertical" />
                <a-button type="text" @click="onLogout">退出登录</a-button>
              </div>
            </template>
            <a-avatar @click="popoverVisible = true" class="avatar" :size="28">U</a-avatar>
          </a-popover>
        </div>
        <!-- 注销账号弹出框 -->
        <a-modal v-model:visible="delUserVisible" title="注销账号" @ok="onConfirm" @cancel="onCancel" cancelText="取消"
          okText="注销" width="360px" style="top: 120px">
          <a-alert message="账号注销后，会清空账号相关的所有数据" banner /><br />
          <a-form :model="user" layout="vertical" @finish="onSubmit" :rules="rules">
            <a-form-item name="email">
              <a-input v-model:value="user.email" placeholder="邮箱" disabled />
            </a-form-item>
            <a-form-item name="code">
              <a-input v-model:value="user.code" style="width: 55%;" placeholder="验证码" />
              <a-button @click="onGetCode" style="width: 40%;float: right;" :loading="loading" :disabled="disabled">
                {{ buttonText }}</a-button>
            </a-form-item>
          </a-form>
        </a-modal>
      </a-layout-header>
      <a-layout-content :style="{ margin: '10px', padding: '18px 18px 12px 18px', background: '#fff' }">
        <transition name="fade">
          <router-view v-slot="{ Component }">
            <component :is="Component" />
          </router-view>
        </transition>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script>
import { reactive, ref, onBeforeMount } from 'vue';
import { useRouter } from 'vue-router'
import { useStore } from '../store/index';
import { message } from 'ant-design-vue';
import { getUserInfo, getVerifyCode, userDelete } from '../api/user';
import { DashboardOutlined, SmileOutlined, MehOutlined, ShoppingOutlined } from '@ant-design/icons-vue';
import { CrownOutlined, MenuUnfoldOutlined, MenuFoldOutlined, QuestionCircleFilled } from '@ant-design/icons-vue';
import { SmileFilled, BellFilled } from '@ant-design/icons-vue';
import { ExclamationCircleOutlined, CrownTwoTone } from '@ant-design/icons-vue';

export default {
  components: {
    DashboardOutlined,
    SmileOutlined,
    MehOutlined,
    ShoppingOutlined,
    CrownOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    QuestionCircleFilled,
    SmileFilled,
    BellFilled,
    ExclamationCircleOutlined,
    CrownTwoTone,
  },
  setup() {
    // 菜单选项
    const menuItem = reactive([{
      key: "dashboard",
      to: "/dashboard",
      icon: "dashboard-outlined",
      name: "仪表盘"
    }, {
      key: "customer",
      to: "/customer",
      icon: "smile-outlined",
      name: "客户"
    }, {
      key: "contract",
      to: "/contract",
      icon: "meh-outlined",
      name: "合同"
    }, {
      key: "product",
      to: "/product",
      icon: "shopping-outlined",
      name: "产品"
    }, {
      key: "subscribe",
      to: "/subscribe",
      icon: "crown-outlined",
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
    const popoverVisible = ref(false)
    const visibleLogo = ref(false)
    const delUserVisible = ref(false)
    const loading = ref(false)
    const disabled = ref(false)
    const buttonText = ref('获取验证码')

    // 初始化数据
    onBeforeMount(() => {
      userInfo()
      store.selectedKeys = 'dashboard'
      router.push('dashboard')
    })

    // 获取用户信息
    const userInfo = () => {
      getUserInfo().then((res) => {
        if (res.data.code == 0) {
          user.name = res.data.data.name
          user.email = res.data.data.email
          switch (res.data.data.version) {
            case 1:
              user.versionText = '基础版'
              break;
            case 2:
              user.versionText = '专业版'
              break;
            case 3:
              user.versionText = '高级版'
              break;
            default:
              user.versionText = ''
              break;
          }
        }
      })
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

    // 点击注销账号
    const onDelete = () => {
      popoverVisible.value = false
      delUserVisible.value = true
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
      delUserVisible.value = false
    };

    return {
      menuItem,
      rules,
      selectedKeys,
      collapsed,
      user,
      visible,
      visibleLogo,
      popoverVisible,
      delUserVisible,
      loading,
      disabled,
      buttonText,
      userInfo,
      onDelete,
      onLogout,
      onGetCode,
      onConfirm,
      onCancel,
      store,
    };
  },
}
</script>

<style scoped>
.sider {
  background: #fff;
  border-right: 0.5px solid #F0F2F5;
}

.header {
  padding: 0 12px;
  display: flex;
  align-items: center;
  justify-content: space-between;
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
  margin-left: 10px;
}

.popover {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 10px 0;
}

.title {
  font-size: 25px;
  color: rgba(31, 31, 31, 0.85);
  font-weight: 620;
  margin-left: 10px;
  overflow: hidden;
}
</style>