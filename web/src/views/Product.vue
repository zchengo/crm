<template>
    <div>
        <a-space style="margin-bottom: 20px; width: 100%;">
            <a-input v-model:value="keyWord" placeholder="产品名称" style="width: 280px; margin-right: 50px;">
                <template #suffix>
                    <search-outlined style="color: rgba(0, 0, 0, 0.45)" @click="onSearch" />
                </template>
            </a-input>
            <a-button type="primary" @click="onProducts">全部产品</a-button>
            <a-button type="primary" @click="onDelete" :disabled="disabled" danger>删除</a-button>
            <a-button type="primary" @click="onCreate">新建</a-button>
        </a-space>
        <a-table rowKey="id" :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
            :columns="columns" :data-source="data.productList"
            :pagination="{ current: pagination.current, pageSize: pagination.pageSize, total: pagination.total, onChange: onPagination }"
            :scroll="{ y: '59vh' }" class="ant-table-striped"
            :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)" bordered>
            <template #bodyCell="{ column, text, record }">
                <template v-if="column.dataIndex === 'name'">
                    <a @click="onEdit(record)">{{ text }}</a>
                </template>
                <template v-if="column.dataIndex === 'status'">
                    <a-tag v-if="text == 1" color="green">上架</a-tag>
                    <a-tag v-if="text == 2" color="blue">下架</a-tag>
                </template>
                <template v-if="column.dataIndex === 'type'">
                    <span v-if="text == 1">默认</span>
                </template>
                <template v-if="column.dataIndex === 'price'">
                    <span style="color: #ff991f">{{ text }}</span>
                </template>
            </template>
        </a-table>
        <!-- 新建、编辑产品 -->
        <a-modal v-model:visible="visible" :title="title" @ok="onSave" @cancel="onCancel" cancelText="取消" okText="保存"
            width="800px" style="top: 80px">
            <a-form ref="productFormRef" :model="product" layout="vertical" name="product" :rules="rules">
                <a-row :gutter="16">
                    <a-col :span="12">
                        <a-form-item label="产品名称" name="name">
                            <a-input v-model:value="product.name" />
                        </a-form-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-item label="产品类型" name="type">
                            <a-select v-model:value="product.type" placeholder="请选择">
                                <a-select-option :value="1">默认</a-select-option>
                            </a-select>
                        </a-form-item>
                    </a-col>
                </a-row>
                <a-row :gutter="16">
                    <a-col :span="12">
                        <a-form-item label="产品单位" name="unit">
                            <a-select v-model:value="product.unit" placeholder="请选择">
                                <a-select-option value="个">个</a-select-option>
                                <a-select-option value="只">只</a-select-option>
                                <a-select-option value="块">块</a-select-option>
                                <a-select-option value="瓶">瓶</a-select-option>
                                <a-select-option value="盒">盒</a-select-option>
                                <a-select-option value="台">台</a-select-option>
                                <a-select-option value="箱">箱</a-select-option>
                                <a-select-option value="吨">吨</a-select-option>
                            </a-select>
                        </a-form-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-item label="产品编码" name="code">
                            <a-input v-model:value="product.code" />
                        </a-form-item>
                    </a-col>
                </a-row>
                <a-row :gutter="16">
                    <a-col :span="12">
                        <a-form-item label="价格" name="price">
                            <a-input-number v-model:value="product.price" style="width: 100%" />
                        </a-form-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-item label="是否上下架" name="status">
                            <a-select v-model:value="product.status" placeholder="请选择">
                                <a-select-option :value="1">上架</a-select-option>
                                <a-select-option :value="2">下架</a-select-option>
                            </a-select>
                        </a-form-item>
                    </a-col>
                </a-row>
                <a-row :gutter="16">
                    <a-col :span="24">
                        <a-form-item label="产品描述" name="description">
                            <a-textarea v-model:value="product.description" :rows="4" />
                        </a-form-item>
                    </a-col>
                </a-row>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
import { ref, reactive, onMounted, createVNode } from 'vue';
import { SearchOutlined, ExclamationCircleOutlined } from '@ant-design/icons-vue';
import moment from 'moment'
import { createProduct, updateProduct, queryProductList, deleteProduct, queryProductInfo } from '../api/product';
import { message, Modal } from 'ant-design-vue';

export default {
    components: {
        SearchOutlined,
    },
    setup() {
        // 表格字段
        const columns = [{
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
        }];

        // 表单校验
        const rules = {
            name: [{ required: true, message: '请输入产品名称', trigger: 'blur' }],
            type: [{ required: true, message: '请选择产品类型' }],
            code: [{ pattern: /^\d+$/, message: '产品编码格式不正确', trigger: 'blur' }],
            price: [{ required: true, message: '请输入产品价格' }],
            status: [{ required: true, message: '请选择是否上下架' }]
        };

        // 产品属性
        let product = reactive({
            id: undefined,
            name: undefined,
            type: undefined,
            unit: undefined,
            code: undefined,
            price: undefined,
            status: undefined,
            description: undefined,
        });

        // 产品列表
        const data = reactive({
            productList: [],
            selectedIds: []
        })

        // 表格分页
        let pagination = reactive({
            current: 1,
            pageSize: 10,
            total: undefined,
        })

        const title = ref('');
        const visible = ref(false);
        const disabled = ref(true)
        const operation = ref(0);
        const productFormRef = ref();
        const keyWord = ref('')

        // 初始化数据
        onMounted(() => { getProductList() })

        // 表格记录多选
        const onSelectChange = selectedRowKeys => {
            data.selectedIds = selectedRowKeys
            if (data.selectedIds.length !== 0) {
                disabled.value = false
            } else {
                disabled.value = true
            }
        };

        // 点击搜索
        const onSearch = () => { getProductList() };

        // 点击全部产品
        const onProducts = () => {
            keyWord.value = ''
            getProductList()
        }

        // 点击新建产品
        const onCreate = () => {
            title.value = '新建产品'
            operation.value = 1
            visible.value = true
        }

        // 点击产品名称
        const onEdit = (row) => {
            title.value = '编辑产品'
            operation.value = 2
            let param = { id: row.id }
            queryProductInfo(param).then((res) => {
                if (res.data.code == 0) {
                    let p = res.data.data
                    product.id = p.id
                    product.name = p.name
                    product.type = p.type
                    product.unit = p.unit
                    product.code = p.code
                    product.price = p.price
                    product.status = p.status
                    product.description = p.description
                }
            })
            visible.value = true
        }

        // 点击保存产品
        const onSave = () => {
            productFormRef.value.validateFields().then(() => {
                if (operation.value == 1) {
                    createProduct(product).then((res) => {
                        if (res.data.code == 0) {
                            message.success('保存成功')
                            getProductList()
                        }
                    })
                }
                if (operation.value == 2) {
                    updateProduct(product).then((res) => {
                        if (res.data.code == 0) {
                            message.success('保存成功')
                            getProductList()
                        }
                    })
                }
                productFormRef.value.resetFields()
                visible.value = false;
            });
        };

        // 点击删除产品
        const onDelete = () => {
            Modal.confirm({
                title: '确定删除选中的' + data.selectedIds.length + '项吗?',
                icon: createVNode(ExclamationCircleOutlined),
                centered: true,
                cancelText: '取消',
                okText: '确定',
                onOk() {
                    deleteProduct({ ids: data.selectedIds }).then((res) => {
                        if (res.data.code == 0) {
                            getProductList()
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

        // 分页查询产品列表
        const onPagination = (page) => {
            pagination.current = page
            getProductList()
        }

        // 查询产列表
        const getProductList = () => {
            let param = {
                name: keyWord.value,
                pageNum: pagination.current,
                pageSize: pagination.pageSize,
            }
            queryProductList(param).then((res) => {
                if (res.data.code == 0) {
                    pagination.total = res.data.data.total
                    data.productList = res.data.data.list
                }
            })
        }

        // 点击取消按钮
        const onCancel = () => {
            productFormRef.value.resetFields()
            visible.value = false
        };

        return {
            columns,
            rules,
            data,
            onSelectChange,
            onSearch,
            product,
            title,
            visible,
            disabled,
            operation,
            onProducts,
            onCreate,
            onEdit,
            productFormRef,
            onSave,
            onCancel,
            onDelete,
            getProductList,
            keyWord,
            pagination,
            onPagination,
        };
    },
}
</script>

<style scoped>
.ant-table-striped :deep(.table-striped) td {
    background-color: #fafafa;
}
</style>