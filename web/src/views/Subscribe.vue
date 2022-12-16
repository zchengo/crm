<template>
    <div>
        <a-alert message="点击立即购买后，会跳转到支付宝支付页面（沙箱环境，账户名 emrpqt1589@sandbox.com 登录密码/支付密码 111111）" type="info"
            show-icon /><br />
        <a-row :gutter="30">
            <a-col :span="8">
                <a-card :class="version == 1 ? 'selected-free-card' : 'card'" :bordered="false">
                    <span class="member-tag" style="color: #33B9FE;">基础版</span>
                    <h2 class="title">免费
                        <check-circle-filled v-if="ver == 1" class="icon" />
                    </h2>
                    <div class="content">满足基础功能需求，永久免费</div><br />
                    <a-button size="large" class="btn-free" shape="round">免费使用</a-button>
                    <br />
                    <div class="subscribe-list" v-for="item in ['客户管理', '合同管理', '产品管理']">
                        <check-circle-filled style="color: #33B9FE;font-size: 20px;;" />
                        <span style="margin-left: 10px;">{{ item }}</span>
                    </div>
                </a-card>
            </a-col>
            <a-col :span="8">
                <a-card :class="version == 2 ? 'selected-card' : 'card'" :bordered="false">
                    <span class="member-tag" style="color: #3788FF;">专业版</span>
                    <h2 class="title">按月付费，每月18元</h2>
                    <div class="content">能力不设限，新功能优先体验</div><br />
                    <a-button v-if="version == 1 || version == 3" type="primary" size="large" class="btn-buy"
                        @click="onPay(2)" shape="round" :disabled="disabled">立即购买</a-button>
                    <a-button v-if="version == 2" type="primary" size="large" class="btn-buy" shape="round">{{ expired
                    }} 到期</a-button>
                    <br />
                    <div class="subscribe-list" v-for="item in ['客户管理', '合同管理', '产品管理', '仪表盘可体验30天']">
                        <check-circle-filled style="color: #476FFF;font-size: 20px;;" />
                        <span style="margin-left: 10px;">{{ item }}</span>
                    </div>
                </a-card>
            </a-col>
            <a-col :span="8">
                <a-card :class="version == 3 ? 'selected-card' : 'card'" :bordered="false">
                    <span class="member-tag" style="color: #3788FF;">高级版</span>
                    <h2 class="title">按年付费，每年198元</h2>
                    <div class="content">能力不设限，新功能优先体验</div><br />
                    <a-button v-if="version == 1 || version == 2" type="primary" size="large" class="btn-buy"
                        @click="onPay(3)" shape="round" :disabled="disabled">立即购买</a-button>
                    <a-button v-if="version == 3" type="primary" size="large" class="btn-buy" shape="round">{{ expired
                    }} 到期</a-button>
                    <br />
                    <div class="subscribe-list" v-for="item in ['客户管理', '合同管理', '产品管理', '仪表盘可体验365天']">
                        <check-circle-filled style="color: #476FFF;font-size: 20px;;" />
                        <span style="margin-left: 10px;">{{ item }}</span>
                    </div>
                </a-card>
            </a-col>
        </a-row>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { CheckCircleFilled } from '@ant-design/icons-vue';
import { subscribePay, getSubscribeInfo } from '../api/subscribe';
import { useRouter } from 'vue-router'
import moment from 'moment'

export default {
    components: {
        CheckCircleFilled
    },
    setup() {

        const router = useRouter()

        const version = ref(0)
        const expired = ref(undefined)
        const payUrl = ref()
        const visible = ref(false)
        const disabled = ref(false)
        const active = ref(30)

        const payResult = ref(false)
        const title = ref('')
        const buttonText = ref(undefined)

        const isClick = (index) => {
            active.value = index
        }

        // 初始化数据
        onMounted(() => { subscribeInfo() })

        // 点击支付
        const onPay = (ver) => {
            let param = {
                version: ver
            }
            subscribePay(param).then((res) => {
                if (res.data.code == 0) {
                    visible.value = false
                    payResult.value = true
                    window.open(res.data.data.payUrl, '_self')
                }
            })
        }

        // 获取用户订阅信息
        const subscribeInfo = () => {
            getSubscribeInfo().then((res) => {
                if (res.data.code == 0) {
                    version.value = res.data.data.version
                    expired.value = moment(res.data.data.expired * 1000).format('YYYY-MM-DD')
                    if (res.data.data.version !== 1) {
                        disabled.value = true
                    }
                }
            })
        }

        return {
            version,
            expired,
            onPay,
            payUrl,
            visible,
            disabled,
            payResult,
            title,
            active,
            buttonText,
            isClick,
            subscribeInfo,
        }
    }
}
</script>

<style scoped>
.title {
    color: #121212;
    line-height: 47px;
    margin-top: 4px;
    margin-bottom: 16px;
    font-size: 24px;
}

.icon {
    color: #3bc8b6;
    font-size: 30px;
    margin-left: 8px;
}

.content {
    font-size: 14px;
    color: rgba(27, 27, 27, .5);
    line-height: 24px;
}

.price {
    height: 54px;
    font-size: 35px;
    line-height: 54px;
    font-weight: 700;
    color: #212930;
}

.card {
    min-height: 450px;
    box-shadow: 0 1px 16px 0 rgb(33 41 48 / 5%);
}

.selected-free-card {
    min-height: 450px;
    box-shadow: 0 1px 16px 0 rgb(33 41 48 / 5%);
    border: 2px solid #ceebfa;
    background: #f0faff;
}

.selected-card {
    min-height: 450px;
    box-shadow: 0 1px 16px 0 rgb(33 41 48 / 5%);
    border: 2px solid #d6ddf9;
    background: #f3f6fd;
}

.subscribe-list {
    display: flex;
    flex-direction: row;
    align-items: center;
    font-size: 16px;
    padding: 5px;
}

.member-tag {
    font-weight: 500;
    line-height: 24px;
    font-size: 16px;
}

.btn-free {
    width: 100%;
    height: 50px;
    background-color: #33B9FE;
    border-color: #33B9FE;
    display: flex;
    font-size: 20px;
    font-weight: 500;
    color: #fff;
    align-items: center;
    justify-content: center;
}

.btn-buy {
    width: 100%;
    height: 50px;
    display: flex;
    font-size: 20px;
    font-weight: 500;
    align-items: center;
    justify-content: center;
}
</style>