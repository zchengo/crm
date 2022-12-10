<template>
    <div>
        <a-space style="margin-bottom: 20px; width: 100%;">
            <a-input v-model:value="keyWord" placeholder="合同编号" style="width: 280px; margin-right: 50px;">
                <template #suffix>
                    <search-outlined style="color: rgba(0, 0, 0, 0.45)" @click="onSearch" />
                </template>
            </a-input>
            <a-button type="primary" @click="onContracts">全部合同</a-button>
            <a-button type="primary" @click="onDelete" :disabled="disabled" danger>删除</a-button>
            <a-button type="primary" @click="onCreate">新建</a-button>
        </a-space>
        <a-table rowKey="id"
            :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectedConteactIds, getCheckboxProps: defaultSelected }"
            :columns="columns" :data-source="data.contractList"
            :pagination="{ current: pagination.current, pageSize: pagination.pageSize, total: pagination.total, onChange: onPagination }"
            :scroll="{ y: '59vh' }" class="ant-table-striped"
            :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)" bordered>
            <template #bodyCell="{ column, text, record }">
                <template v-if="column.dataIndex === 'id'">
                    <a @click="onEdit(record)">{{ text }}</a>
                </template>
                <template v-if="column.dataIndex === 'amount'">
                    <span style="color: #ff991f">{{ text }}</span>
                </template>
                <template v-if="column.dataIndex === 'status'">
                    <a-tag v-if="text == 1" color="green">已签约</a-tag>
                    <a-tag v-if="text == 2" color="red">未签约</a-tag>
                </template>
            </template>
        </a-table>
        <!-- 新建、编辑合同 -->
        <a-modal v-model:visible="visible" :title="title" @ok="onSave" @cancel="onCancel" cancelText="取消" okText="保存"
            width="800px" style="top: 80px">
            <div style="height: 55vh;overflow-y: scroll;padding: 0 15px;">
                <a-form ref="contractFormRef" :model="contract" layout="vertical" name="contract" :rules="rules">
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="合同编号" name="id">
                                <a-input v-model:value="contract.id" :disabled="true" placeholder="根据编号规则自动生成" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="合同名称" name="name">
                                <a-input v-model:value="contract.name" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="客户名称" name="cid">
                                <a-select v-model:value="contract.cid" style="width: 100%" placeholder="请选择"
                                    :fieldNames="{ label: 'name', value: 'id' }" :options="data.customerOption"
                                    @focus="getCustomerOption" @change="changeCustomerOption"></a-select>
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="合同金额" name="amount">
                                <a-input-number v-model:value="contract.amount" style="width: 100%" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="合同开始时间" name="beginTime">
                                <a-date-picker v-model:value="contract.beginTime" placeholder="选择日期"
                                    style="width: 100%" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="合同结束时间" name="overTime">
                                <a-date-picker v-model:value="contract.overTime" placeholder="选择日期"
                                    style="width: 100%" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="合同状态" name="status">
                                <a-select v-model:value="contract.status" placeholder="请选择">
                                    <a-select-option :value="1">已生效</a-select-option>
                                    <a-select-option :value="2">未生效</a-select-option>
                                </a-select>
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="备注" name="remarks">
                                <a-input v-model:value="contract.remarks" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="24">
                            <a-form-item label="产品" name="product">
                                <a-button type="primary" @click="onAddProduct"
                                    style="float: right;margin-bottom: 10px;">
                                    添加产品</a-button>
                                <a-table rowKey="id" :columns="productColumns" :data-source="data.addedProductList"
                                    :scroll="{ y: '59vh' }" class="ant-table-striped"
                                    :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)"
                                    :pagination="false" bordered>
                                    <template #bodyCell="{ column, text, record }">
                                        <template v-if="column.dataIndex === 'type'">
                                            <span v-if="text == 1">默认</span>
                                        </template>
                                        <template v-if="column.dataIndex === 'price'">
                                            <span style="color: #ff991f">{{ text }}</span>
                                        </template>
                                        <template v-if="column.dataIndex === 'count'">
                                            <a-input-number v-model:value="record.count" @change="calculatedAmount" />
                                        </template>
                                        <template v-if="column.dataIndex === 'total'">
                                            <span>{{ record.total = record.price * record.count }}</span>
                                        </template>
                                        <template v-if="column.dataIndex === 'operation'">
                                            <a @click="delProduct(record)">删除</a>
                                        </template>
                                    </template>
                                </a-table>
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row>
                        <a-col :span="24">
                            <div style="float: right;margin: 0 20px;">
                                <span>已选择产品：{{ data.addedProductList.length }}种&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;总金额：</span>
                                <a-input-number v-model:value="contract.amount" style="width: 200px;" />
                            </div>
                        </a-col>
                    </a-row>
                </a-form>
            </div>
        </a-modal>
        <!-- 添加产品 -->
        <a-modal v-model:visible="productListVisible" title="添加产品" @cancel="onCancelProductList" @ok="onConfirm"
            cancelText="取消" okText="确定" width="800px" style="top: 80px" :getCheckboxProps="defaultSelected">
            <div style="height: 55vh;padding: 0 15px;">
                <a-table rowKey="id"
                    :row-selection="{ selectedRowKeys: data.defaultSelectedIds, onChange: onSelectedProductIds }"
                    :columns="productListColumns" :data-source="data.productList"
                    :pagination="{ current: pagination.current, pageSize: pagination.pageSize, total: pagination.total, onChange: onPaginationProduct }"
                    :scroll="{ y: '40vh' }" class="ant-table-striped"
                    :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)" bordered>
                    <template #bodyCell="{ column, text, record }">
                        <template v-if="column.dataIndex === 'name'">
                            <a @click="onEdit(record)">{{ text }}</a>
                        </template>
                        <template v-if="column.dataIndex === 'status'">
                            <span v-if="text == 1">上架</span>
                            <span v-if="text == 2">下架</span>
                        </template>
                        <template v-if="column.dataIndex === 'type'">
                            <span v-if="text == 1">默认</span>
                        </template>
                        <template v-if="column.dataIndex === 'price'">
                            <span style="color: #ff991f">{{ text }}</span>
                        </template>
                    </template>
                </a-table>
            </div>
        </a-modal>
    </div>
</template>

<script>
import { ref, reactive, onMounted, createVNode } from 'vue';
import { SearchOutlined, ExclamationCircleOutlined, UpCircleOutlined, DownCircleOutlined } from '@ant-design/icons-vue';
import moment from 'moment'
import { createContract, updateContract, queryContractList, queryContractInfo, deleteContract, queryContractPlist } from '../api/contract';
import { queryProductList } from "../api/product";
import { queryCustomerOption } from "../api/customer";
import { message, Modal } from 'ant-design-vue';

export default {
    components: {
        SearchOutlined,
        UpCircleOutlined,
        DownCircleOutlined
    },
    setup() {

        // 合同表格字段
        const columns = [{
            title: '合同编号',
            dataIndex: 'id',
            width: 100,
            fixed: 'left',
            ellipsis: true,
        }, {
            title: '合同名称',
            dataIndex: 'name',
            width: 100,
            fixed: 'left',
            ellipsis: true,
        }, {
            title: '客户名称',
            dataIndex: 'cname',
            width: 240,
        }, {
            title: '合同金额',
            dataIndex: 'amount',
            width: 100,
        }, {
            title: '合同开始时间',
            dataIndex: 'beginTime',
            width: 150,
        }, {
            title: '合同结束时间',
            dataIndex: 'overTime',
            width: 150,
        }, {
            title: '备注',
            dataIndex: 'remarks',
            width: 240,
            ellipsis: true,
        }, {
            title: '签约状态',
            dataIndex: 'status',
            width: 100,
            ellipsis: true,
        }, {
            title: '创建时间',
            dataIndex: 'created',
            width: 185,
            customRender: text => {
                let m = moment(text.value * 1000).format('YYYY-MM-DD HH:mm:ss')
                return m == 'Invalid date' ? '' : m
            }
        }, {
            title: '更新时间',
            dataIndex: 'updated',
            width: 185,
            customRender: text => {
                let m = moment(text.value * 1000).format('YYYY-MM-DD HH:mm:ss')
                return m == 'Invalid date' ? '' : m
            }
        }];

        // 新建或编辑合同，已添加产品表格字段
        const productColumns = [{
            title: '产品名称',
            dataIndex: 'name',
            width: 100,
        }, {
            title: '产品类别',
            dataIndex: 'type',
            width: 90,
        }, {
            title: '单位',
            dataIndex: 'unit',
            width: 80,
        }, {
            title: '价格',
            dataIndex: 'price',
            width: 100,
        }, {
            title: '数量',
            dataIndex: 'count',
            width: 120,
        }, {
            title: '合计',
            dataIndex: 'total',
            width: 100,
        }, {
            title: '操作',
            dataIndex: 'operation',
            width: 100,
        }]

        // 产品表格字段
        const productListColumns = [{
            title: '产品名称',
            dataIndex: 'name',
            width: 100,
            fixed: 'left',
            ellipsis: true,
        }, {
            title: '是否上下架',
            dataIndex: 'status',
            width: 120,
        }, {
            title: '产品类型',
            dataIndex: 'type',
            width: 100,
        }, {
            title: '产品单位',
            dataIndex: 'unit',
            width: 100,
        }, {
            title: '产品编码',
            dataIndex: 'code',
            width: 150,
        }, {
            title: '价格',
            dataIndex: 'price',
            width: 150,
        }, {
            title: '产品描述',
            dataIndex: 'description',
            width: 240,
            ellipsis: true,
        }, {
            title: '创建时间',
            dataIndex: 'created',
            width: 185,
            customRender: text => {
                let m = moment(text.value * 1000).format('YYYY-MM-DD HH:mm:ss')
                return m == 'Invalid date' ? '' : m
            }
        }, {
            title: '更新时间',
            dataIndex: 'updated',
            width: 185,
            customRender: text => {
                let m = moment(text.value * 1000).format('YYYY-MM-DD HH:mm:ss')
                return m == 'Invalid date' ? '' : m
            }
        }, {
            title: '创建人',
            dataIndex: 'creator',
            width: 150,
        }];

        // 表单校验
        const rules = {
            name: [{ required: true, message: '请输入合同名称', trigger: 'blur' }],
            cid: [{ required: true, message: '请选择客户', trigger: 'blur' }],
            amount: [{ required: true, message: '请输入合同金额', trigger: 'blur' }],
            status: [{ required: true, message: '请选择合同状态' }]
        };

        // 合同属性
        let contract = reactive({
            id: undefined,
            name: undefined,
            amount: undefined,
            beginTime: undefined,
            cid: undefined,
            overTime: undefined,
            remarks: undefined,
            status: undefined,
            productlist: [],
        });

        const data = reactive({
            contractList: [],
            contractIds: [],
            productList: [],
            productIds: [],
            addedProductList: [],
            customerOption: [],
            defaultSelectedIds: []
        })

        // 表格分页
        let pagination = reactive({
            current: 1,
            pageSize: 10,
            total: undefined,
        })

        // 点击搜索
        const onSearch = () => { getContractList() };

        const title = ref('');
        const visible = ref(false);
        const disabled = ref(true)
        const operation = ref(0);
        const contractFormRef = ref();
        const keyWord = ref('')
        const productListVisible = ref(false);

        // 点击新建合同
        const onCreate = () => {
            title.value = '新建合同'
            operation.value = 1
            visible.value = true
        }

        // 点击编辑合同
        const onEdit = (row) => {
            title.value = '编辑合同'
            operation.value = 2
            getCustomerOption()
            let param = { id: row.id }
            queryContractInfo(param).then((res) => {
                if (res.data.code == 0) {
                    let p = res.data.data
                    contract.id = p.id
                    contract.name = p.name
                    contract.cid = p.cid
                    contract.amount = p.amount
                    contract.beginTime = moment(new Date(p.beginTime))
                    contract.overTime = moment(new Date(p.overTime))
                    contract.remarks = p.remarks
                    contract.status = p.status
                    contract.productlist = p.productlist
                    data.addedProductList = p.productlist
                }
            })
            visible.value = true
        }

        // 点击保存合同
        const onSave = () => {
            contractFormRef.value.validateFields().then(() => {
                if (operation.value == 1) {
                    let param = {
                        name: contract.name,
                        cid: contract.cid,
                        amount: contract.amount,
                        beginTime: moment(contract.beginTime).format('YYYY-MM-DD'),
                        overTime: moment(contract.overTime).format('YYYY-MM-DD'),
                        remarks: contract.remarks,
                        status: contract.status,
                        productlist: data.addedProductList,
                    }
                    createContract(param).then((res) => {
                        if (res.data.code == 0) {
                            message.success('保存成功')
                            data.defaultSelectedIds = []
                            getContractList()
                        }
                    })
                }
                if (operation.value == 2) {
                    let param = {
                        id: contract.id,
                        name: contract.name,
                        cid: contract.cid,
                        amount: contract.amount,
                        beginTime: moment(contract.beginTime).format('YYYY-MM-DD'),
                        overTime: moment(contract.overTime).format('YYYY-MM-DD'),
                        remarks: contract.remarks,
                        status: contract.status,
                        productlist: data.addedProductList,
                    }
                    updateContract(param).then((res) => {
                        if (res.data.code == 0) {
                            message.success('保存成功')
                            data.defaultSelectedIds = []
                            getContractList()
                        }
                    })
                }
                contractFormRef.value.resetFields()
                visible.value = false;
            });
        };

        // 点击删除合同
        const onDelete = () => {
            let param = {
                ids: data.contractIds
            }
            Modal.confirm({
                title: '确定删除选中的' + data.contractIds.length + '项吗?',
                icon: createVNode(ExclamationCircleOutlined),
                centered: true,
                cancelText: '取消',
                okText: '确定',
                onOk() {
                    deleteContract(param).then((res) => {
                        if (res.data.code == 0) {
                            getContractList()
                            disabled.value = true
                            message.success('删除成功')
                        }
                    })
                },
                onCancel() {
                    console.log('Cancel');
                },
            });
        }

        // 初始化数据
        onMounted(() => { getContractList() })

        // 点击全部合同
        const onContracts = () => {
            keyWord.value = ''
            getContractList()
        }

        // 分页查询合同列表
        const onPagination = (page) => {
            pagination.current = page
            getContractList()
        }
        const getContractList = () => {
            let param = {
                id: parseInt(keyWord.value == '' ? '0' : keyWord.value),
                pageNum: pagination.current,
                pageSize: pagination.pageSize
            }
            queryContractList(param).then((res) => {
                if (res.data.code == 0) {
                    pagination.total = res.data.data.total
                    data.contractList = res.data.data.list
                }
            })
        }

        // 点击添加产品
        const onAddProduct = () => {
            let param = {
                pageNum: pagination.current,
                pageSize: pagination.pageSize
            }
            queryProductList(param).then((res) => {
                if (res.data.code == 0) {
                    pagination.total = res.data.data.total
                    data.productList = res.data.data.list
                }
            })
            data.defaultSelectedIds = []
            if (data.addedProductList.length > 0) {
                for (let i = 0; i < data.addedProductList.length; i++) {
                    data.defaultSelectedIds[i] = data.addedProductList[i].id
                }
            }
            productListVisible.value = true
        }

        // 分页查询产品列表
        const onPaginationProduct = (page) => {
            pagination.current = page
            let param = {
                pageNum: pagination.current,
                pageSize: pagination.pageSize
            }
            queryProductList(param).then((res) => {
                if (res.data.code == 0) {
                    pagination.total = res.data.data.total
                    data.productList = res.data.data.list
                }
            })
        }

        // 已选中的合同ID
        const onSelectedConteactIds = selectedRowKeys => {
            data.contractIds = selectedRowKeys
            if (data.contractIds.length !== 0) {
                disabled.value = false
            } else {
                disabled.value = true
            }
        };

        // 已选中的产品ID
        const onSelectedProductIds = selectedRowKeys => {
            data.productIds = selectedRowKeys
            data.defaultSelectedIds = selectedRowKeys
        };

        // 删除选中的产品
        const delProduct = (row) => {
            for (let i = 0; i < data.addedProductList.length; i++) {
                if (data.addedProductList[i].id == row.id) {
                    data.addedProductList.splice(i, 1);
                }
            }
            calculatedAmount()
        }

        // 点击确定选中的产品ID
        const onConfirm = () => {
            console.log("xzx", data.productIds)
            let param = {
                ids: data.productIds
            }
            queryContractPlist(param).then((res) => {
                if (res.data.code == 0) {
                    data.addedProductList = res.data.data
                }
            })
            productListVisible.value = false
        }

        // 查询客户选项
        const getCustomerOption = () => {
            queryCustomerOption().then((res) => {
                if (res.data.code == 0) {
                    data.customerOption = res.data.data
                    console.log("zxxzc", data.customerOption)
                }
            })
        }

        const changeCustomerOption = (value) => {
            contract.cid.value = value
        }

        // 计算金额
        const calculatedAmount = () => {
            contract.amount = 0
            let totalAmount = 0
            for (let i = 0; i < data.addedProductList.length; i++) {
                totalAmount = totalAmount + (data.addedProductList[i].price * data.addedProductList[i].count)
            }
            contract.amount = totalAmount
        }

        // 点击合同取消按钮
        const onCancel = () => {
            contractFormRef.value.resetFields()
            data.addedProductList = []
            visible.value = false
        };

        // 点击取消产品列表
        const onCancelProductList = () => {
            productListVisible.value = false
            pagination.current = 1,
                pagination.total = undefined
        }

        return {
            columns,
            productColumns,
            productListColumns,
            rules,
            data,
            onSelectedConteactIds,
            onSelectedProductIds,
            onSearch,
            contract,
            title,
            visible,
            disabled,
            productListVisible,
            operation,
            onCreate,
            onEdit,
            contractFormRef,
            onSave,
            onCancel,
            onDelete,
            onCancelProductList,
            getContractList,
            keyWord,
            onConfirm,
            onAddProduct,
            getCustomerOption,
            changeCustomerOption,
            calculatedAmount,
            delProduct,
            pagination,
            onPagination,
            onContracts,
            onPaginationProduct,
        };
    },
}
</script>

<style scoped>
.ant-table-striped :deep(.table-striped) td {
    background-color: #fafafa;
}
</style>