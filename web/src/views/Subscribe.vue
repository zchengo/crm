<template>
    <div style="margin: 20vh auto;">
        <a-row :gutter="30">
            <a-col :span="6" :offset="6">
                <a-card class="card" :bordered="false">
                    <h2 class="title">免费版
                        <check-circle-filled v-if="ver == 1" class="icon" />
                    </h2>
                    <div class="content">满足基础功能需求，永久免费</div><br />
                    <div class="price">￥0</div><br />
                    <a-button size="large" class="btn">免费使用</a-button>
                </a-card>
            </a-col>
            <a-col :span="6">
                <a-card class="card" :bordered="false">
                    <h2 class="title">个人版
                        <check-circle-filled v-if="ver == 2" class="icon" />
                    </h2>
                    <div class="content">能力不设限，新功能优先体验</div><br />
                    <div class="price">￥28<span style="font-size: 20px;font-weight: 500;"> / 月</span></div><br />
                    <a-button v-if="ver == 1" type="primary" size="large" @click="onBuy" style="width: 100%;">立即订阅
                    </a-button>
                    <a-button v-if="ver == 2" type="primary" size="large" style="width: 100%;">{{ expired }} 到期
                    </a-button>
                </a-card>
            </a-col>
        </a-row>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { message } from 'ant-design-vue';
import { CheckCircleFilled } from '@ant-design/icons-vue';
import { getUserInfo, userBuy } from '../api/user';
import moment from 'moment'

export default {
    components: {

        CheckCircleFilled
    },
    setup() {

        const ver = ref(0)
        const expired = ref(undefined)

        // 初始化数据
        onMounted(() => { sysVersion() })

        // 点击订阅
        const onBuy = () => {
            userBuy().then((res) => {
                if (res.data.code == 0) {
                    ver.value = res.data.data.version
                    expired.value = moment(res.data.data.expired * 1000).format('YYYY-MM-DD')
                    message.success('恭喜你！订阅成功')
                }
            })
        }

        // 获取用户系统版本
        const sysVersion = () => {
            getUserInfo().then((res) => {
                if (res.data.code == 0) {
                    ver.value = res.data.data.version
                    expired.value = moment(res.data.data.expired * 1000).format('YYYY-MM-DD')
                }
            })
        }

        return {
            ver,
            expired,
            onBuy,
            sysVersion,
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
    font-size: 28px;
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
    font-size: 40px;
    line-height: 54px;
    font-weight: 700;
    color: #212930;
}

.card {
    box-shadow: 0 1px 16px 0 rgb(33 41 48 / 5%);
}

.btn {
    width: 100%;
    float: right;
    color: #476FFF;
    border-color: #476FFF;
}
</style>