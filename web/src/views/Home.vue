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
          <router-link :to="item.to">
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
          <!-- 个人信息 -->
          <a-dropdown :trigger="['click']">
            <a-avatar @click="toMine" class="avatar" :size="28">U</a-avatar>
            <template #overlay>
              <a-menu>
                <a-menu-item key="0">
                  <a @click="onMail">
                    <MailOutlined /> 修改邮箱
                  </a>
                </a-menu-item>
                <a-menu-item key="1">
                  <a @click="onDelete">
                    <ClearOutlined /> 注销账号
                  </a>
                </a-menu-item>
                <a-menu-divider />
                <a-menu-item key="2">
                  <a @click="onLogout">
                    <LogoutOutlined /> 退出账号
                  </a>
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </div>
        <!-- 修改邮箱弹出框 -->
        <a-modal v-model:visible="visible" title="修改邮箱" @ok="onSave" @cancel="onCancel" cancelText="取消" okText="保存"
          width="450px" style="top: 120px">
          <a-form :model="user" layout="vertical" @finish="onSubmit" :rules="rules">
            <a-form-item name="email">
              <a-input v-model:value="user.email" size="large" placeholder="邮箱" disabled />
            </a-form-item>
            <a-form-item name="code">
              <a-input v-model:value="user.code" size="large" style="width: 55%;" placeholder="验证码" />
              <a-button @click="onGetCode" size="large" style="width: 40%;float: right;" :loading="loading"
                :disabled="disabled">
                {{ buttonText }}</a-button>
            </a-form-item>
            <a-form-item name="newEmail">
              <a-input v-model:value="user.newEmail" size="large" placeholder="新邮箱" />
            </a-form-item>
          </a-form>
        </a-modal>
        <!-- 注销账号弹出框 -->
        <a-modal v-model:visible="delUserVisible" title="注销账号" @ok="onConfirm" @cancel="onCancel" cancelText="取消"
          okText="注销" width="450px" style="top: 120px">
          <a-form :model="user" layout="vertical" @finish="onSubmit" :rules="rules">
            <a-form-item name="email">
              <a-input v-model:value="user.email" size="large" placeholder="邮箱" disabled />
            </a-form-item>
            <a-form-item name="code">
              <a-input v-model:value="user.code" size="large" style="width: 55%;" placeholder="验证码" />
              <a-button @click="onGetCode" size="large" style="width: 40%;float: right;" :loading="loading"
                :disabled="disabled">
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
import { reactive, ref, onMounted } from 'vue';
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue';
import { getUserInfo, updateMail, getVerifyCode, userDelete, userLogout } from '../api/user';
import { DashboardOutlined, SmileOutlined, MehOutlined, ShoppingOutlined } from '@ant-design/icons-vue';
import { CrownOutlined, MenuUnfoldOutlined, MenuFoldOutlined, QuestionCircleFilled } from '@ant-design/icons-vue';
import { SmileFilled, BellFilled, MailOutlined, ClearOutlined } from '@ant-design/icons-vue';
import { LogoutOutlined, ExclamationCircleOutlined } from '@ant-design/icons-vue';

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
    MailOutlined,
    ClearOutlined,
    LogoutOutlined,
    ExclamationCircleOutlined,
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
      newEmail: [{
        required: true,
        message: '请输入邮箱!',
        trigger: 'blur',
      }, {
        pattern: /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/,
        message: '邮箱格式不正确',
        trigger: 'blur',
      }],
    };

    const selectedKeys = ref(['dashboard'])
    const collapsed = ref(false)

    const router = useRouter()

    const user = reactive({
      name: undefined,
      email: undefined,
      verison: undefined,
      code: undefined,
      newEmail: undefined
    })

    const visible = ref(false)
    const visibleLogo = ref(false)
    const delUserVisible = ref(false)
    const loading = ref(false)
    const disabled = ref(false)
    const buttonText = ref('获取验证码')

    // 初始化数据
    onMounted(() => { userInfo() })

    // 获取用户信息
    const userInfo = () => {
      getUserInfo().then((res) => {
        if (res.data.code == 0) {
          user.name = res.data.data.name
          user.email = res.data.data.email
          user.version = res.data.data.version
        }
      })
    }

    // 点击升级到专业版
    const onUpgrade = () => {
      router.push('/subscribe')
    }

    // 点击修改邮箱
    const onMail = () => {
      visible.value = true
    }

    // 确认修改邮箱
    const onSave = () => {
      let param = {
        email: user.email,
        code: user.code,
        newEmail: user.newEmail
      }
      updateMail(param).then((res) => {
        if (res.data.code == 0) {
          router.push('/')
          message.success('修改成功，请重新登录！')
        }
        if (res.data.code == 10005) {
          message.error('验证码错误');
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
      userLogout().then((res) => {
        if (res.data.code == 0) { router.push('/') }
      })
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
      delUserVisible,
      loading,
      disabled,
      buttonText,
      userInfo,
      onDelete,
      onLogout,
      onMail,
      onGetCode,
      onSave,
      onConfirm,
      onCancel,
      onUpgrade,
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