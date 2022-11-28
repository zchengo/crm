<template>
    <div>
        <a-row :gutter="16">
            <a-col :span="6">
                <a-card class="card">
                    <a-statistic :value="data.customers" style="margin-right: 50px">
                        <template #title>
                            <span>全部客户</span>
                            <a-tooltip placement="right">
                                <template #title>
                                    <span>客户数量，单位（人）</span>
                                </template>
                                <question-circle-two-tone style="margin-left: 5px" />
                            </a-tooltip>
                        </template>
                    </a-statistic>
                </a-card>
            </a-col>
            <a-col :span="6">
                <a-card class="card">
                    <a-statistic :value="data.contracts" style="margin-right: 50px">
                        <template #title>
                            <span>全部合同</span>
                            <a-tooltip placement="right">
                                <template #title>
                                    <span>合同数量，单位（份）</span>
                                </template>
                                <question-circle-two-tone style="margin-left: 5px" />
                            </a-tooltip>
                        </template>
                    </a-statistic>
                </a-card>
            </a-col>
            <a-col :span="6">
                <a-card class="card">
                    <a-statistic :value="data.contractAmount" style="margin-right: 50px">
                        <template #title>
                            <span>合同金额</span>
                            <a-tooltip placement="right">
                                <template #title>
                                    <span>实际完成合同金额，单位（元）</span>
                                </template>
                                <question-circle-two-tone style="margin-left: 5px" />
                            </a-tooltip>
                        </template>
                    </a-statistic>
                </a-card>
            </a-col>
            <a-col :span="6">
                <a-card class="card">
                    <a-statistic :value="data.products" style="margin-right: 50px">
                        <template #title>
                            <span>全部产品</span>
                            <a-tooltip placement="right">
                                <template #title>
                                    <span>产品数量，单位（个）</span>
                                </template>
                                <question-circle-two-tone style="margin-left: 5px" />
                            </a-tooltip>
                        </template>
                    </a-statistic>
                </a-card>
            </a-col>
        </a-row>
        <a-row :gutter="16">
            <a-col :span="24">
                <a-card class="card">
                    <div id="main" style="width: 100%; height: 400px"></div>
                </a-card>
            </a-col>
        </a-row>
    </div>
</template>

<script>
import { QuestionCircleTwoTone } from '@ant-design/icons-vue'
import * as echarts from "echarts";
import { reactive, onMounted } from 'vue';
import { getSummary } from "../api/dashboard";
import { getUserInfo } from "../api/user";
import { useRouter } from 'vue-router'

export default {
    components: {
        QuestionCircleTwoTone
    },
    setup() {
        onMounted(() => {
            checkVersion();
            initChart();
        });

        const router = useRouter()

        const data = reactive({
            customers: 0,
            contracts: 0,
            contractAmount: 0.00,
            products: 0,
        })

        const initChart = () => {
            getSummary().then((res) => {
                if (res.data.code == 0) {
                    data.customers = res.data.data.customers
                    data.contracts = res.data.data.contracts
                    data.contractAmount = res.data.data.contractAmount
                    data.products = res.data.data.products
                    echarts.init(document.getElementById("main")).setOption({
                        xAxis: {
                            type: 'category',
                            data: res.data.data.date
                        },
                        tooltip: {
                            trigger: 'axis',
                        },
                        legend: {
                            data: ['实际完成金额']
                        },
                        yAxis: {
                            type: 'value'
                        },
                        series: [
                            {
                                name: '实际完成金额',
                                data: res.data.data.amount,
                                type: 'line',
                                smooth: true
                            }
                        ]
                    })
                }
            })
        }

        // 查看用户系统版本
        const checkVersion = () => {
            getUserInfo().then((res) => {
                if (res.data.code == 0 && res.data.data.version == 1) {
                    router.push('/result')
                }
            })
        }
        return {
            data,
            initChart,
            checkVersion
        }
    }
}
</script>

<style scoped>
.card {
    margin-top: 20px;
    border: none;
    box-shadow: 0 1px 16px 0 rgb(33 41 48 / 5%);
}
</style>