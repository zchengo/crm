<template>
    <div :style="{ padding: '20px 20px 12px 20px' }">
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
            <a-col :span="15">
                <a-card class="card" style="height: 60vh;margin-top: 20px;">
                    <div style="display:flex;align-items: center;justify-content: space-between;">
                        <div style="color: #606266;font-size: 16px;font-weight: 600;">
                            <span>合同金额完成情况</span>
                            <a-tooltip placement="right">
                                <template #title>
                                    <span>实际完成金额，单位（元）</span>
                                </template>
                                <question-circle-two-tone style="margin-left: 5px" />
                            </a-tooltip>
                        </div>
                        <a-select v-model:value="daysRange" style="width: 120px" @focus="initChart" @change="initChart">
                            <a-select-option :value="7">最近7天</a-select-option>
                            <a-select-option :value="14">最近14天</a-select-option>
                            <a-select-option :value="30">最近30天</a-select-option>
                        </a-select>
                    </div>
                    <div id="contract" style="width: 100%; height: 360px;"></div>
                </a-card>
            </a-col>
            <a-col :span="9">
                <a-card class="card" style="height: 60vh;margin-top: 20px;">
                    <div style="color: #606266;font-size: 16px;font-weight: 600;">
                        <span>客户行业分布</span>
                        <a-tooltip placement="right">
                            <template #title>
                                <span>客户所在行业，单位（个）</span>
                            </template>
                            <question-circle-two-tone style="margin-left: 5px" />
                        </a-tooltip>
                    </div>
                    <div id="customer" style="width: 100%; height: 360px;"></div>
                </a-card>
            </a-col>
        </a-row>
    </div>
</template>

<script setup>
import { QuestionCircleTwoTone } from '@ant-design/icons-vue'
import * as echarts from "echarts";
import { reactive, ref, onBeforeMount } from 'vue';
import { getSummary } from "../api/dashboard";
import { getSubscribeInfo } from '../api/subscribe';
import { useRouter } from 'vue-router'

const daysRange = ref(7);

const router = useRouter()

const data = reactive({
    customers: 0,
    contracts: 0,
    contractAmount: 0.00,
    products: 0,
})

onBeforeMount(() => {
    subscribeInfo();
    initChart();
});

const initChart = () => {
    let param = {
        daysRange: daysRange.value
    }
    getSummary(param).then((res) => {
        if (res.data.code == 0) {
            data.customers = res.data.data.customers
            data.contracts = res.data.data.contracts
            data.contractAmount = res.data.data.contractAmount
            data.products = res.data.data.products
            echarts.init(document.getElementById("contract")).setOption({
                xAxis: {
                    type: 'category',
                    data: res.data.data.date,
                },
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {
                        type: 'shadow'
                    }
                },
                legend: {
                    data: ['实际完成金额'],
                    orient: 'vertical',
                    bottom: 10,
                },
                yAxis: {
                    type: 'value',
                },
                series: [
                    {
                        name: '实际完成金额',
                        data: res.data.data.amount,
                        type: 'line',
                        smooth: true,
                        lineStyle: {
                            width: 3
                        }
                    }
                ]
            });
            echarts.init(document.getElementById("customer")).setOption({
                tooltip: {
                    trigger: 'item'
                },
                legend: {
                    top: 'bottom',
                    left: 'center',
                },
                series: [
                    {
                        type: 'pie',
                        bottom: '15%',
                        radius: ['40%', '70%'],
                        avoidLabelOverlap: false,
                        itemStyle: {
                            borderRadius: 10,
                            borderColor: '#fff',
                            borderWidth: 2,
                        },
                        label: {
                            show: false,
                            position: 'center'
                        },
                        emphasis: {
                            label: {
                                show: true,
                                fontSize: 25,
                                fontWeight: 'bold'
                            }
                        },
                        labelLine: {
                            show: false
                        },
                        data: res.data.data.customerIndustry,
                    }
                ]
            })
        }
    })
}

// 获取用户订阅信息
const subscribeInfo = () => {
    getSubscribeInfo().then((res) => {
        if (res.data.code == 0 && res.data.data.version == 1) {
            router.push('/result')
        }
    })
}
</script>

<style scoped>
.card {
    border: none;
    box-shadow: 0 1px 16px 0 rgb(33 41 48 / 5%);
}
</style>