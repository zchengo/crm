<template>
    <div :style="{ padding: '20px 20px 12px 20px' }">
        <div style="display: flex;justify-content: space-between;margin-bottom: 20px;">
            <a-space>
                <a-input v-model:value="keyWord" placeholder="客户名称" style="width: 280px; margin-right: 50px;">
                    <template #suffix>
                        <search-outlined style="color: rgba(0, 0, 0, 0.45)" @click="onSearch" />
                    </template>
                </a-input>
                <a-button type="primary" @click="onCustomers">全部客户</a-button>
                <a-button type="primary" @click="onDelete" :disabled="disabled" danger>删除</a-button>
                <a-button type="primary" @click="onCreate">新建</a-button>
            </a-space>
            <div>
                <a-button type="primary" @click="onExport">
                    <template #icon>
                        <ExportOutlined />
                    </template>导出</a-button>
            </div>
        </div>
        <a-table rowKey="id" :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
            :columns="columns" :data-source="data.customerList"
            :pagination="{ current: pagination.current, pageSize: pagination.pageSize, total: pagination.total, onChange: onPagination }"
            :scroll="{ y: '59vh' }" class="ant-table-striped"
            :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)" bordered>
            <template #bodyCell="{ column, text, record }">
                <template v-if="column.dataIndex === 'name'">
                    <a @click="onEdit(record)">{{ text }}</a>
                </template>
                <template v-if="column.dataIndex === 'operation'">
                    <a-button type="link"><template #icon>
                            <MailTwoTone two-tone-color="#476FFF" @click="onMail(record.name, record.email)" />
                        </template></a-button>
                </template>
                <template v-if="column.dataIndex === 'status'">
                    <Spot v-if="text == 1" type="success" title="已成交" />
                    <Spot v-if="text == 2" type="danger" title="未成交" />
                </template>
                <template v-if="column.dataIndex === 'address'">
                    <span>{{ record.region + " " + record.address }}</span>
                </template>
            </template>
        </a-table>

        <!-- 新建、编辑客户 -->
        <a-modal v-model:visible="visible" :title="title" @ok="onSave" @cancel="onCancel" cancelText="取消" okText="保存"
            width="800px" :centered="true">
            <div style="height: 55vh;overflow-y: scroll;padding: 0 15px;">
                <a-form ref="customerFormRef" :model="customer" layout="vertical" name="customer" :rules="rules">
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="客户名称" name="name">
                                <a-input v-model:value="customer.name" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="客户来源" name="source">
                                <a-select v-model:value="customer.source" placeholder="请选择">
                                    <a-select-option value="促销">促销</a-select-option>
                                    <a-select-option value="搜索引擎">搜索引擎</a-select-option>
                                    <a-select-option value="广告">广告</a-select-option>
                                    <a-select-option value="转介绍">转介绍</a-select-option>
                                    <a-select-option value="线上注册">线上注册</a-select-option>
                                    <a-select-option value="电话咨询">电话咨询</a-select-option>
                                    <a-select-option value="邮件咨询">邮件咨询</a-select-option>
                                </a-select>
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="手机号" name="phone">
                                <a-input v-model:value="customer.phone" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="邮箱" name="email">
                                <a-input v-model:value="customer.email" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="客户行业" name="industry">
                                <a-select v-model:value="customer.industry" placeholder="请选择">
                                    <a-select-option value="互联网">互联网</a-select-option>
                                    <a-select-option value="金融业">金融业</a-select-option>
                                    <a-select-option value="政府">政府</a-select-option>
                                    <a-select-option value="房地产">房地产</a-select-option>
                                    <a-select-option value="文化传媒">文化传媒</a-select-option>
                                    <a-select-option value="生产">生产</a-select-option>
                                    <a-select-option value="物流运输">物流运输</a-select-option>
                                </a-select>
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="客户级别" name="level">
                                <a-select v-model:value="customer.level" placeholder="请选择">
                                    <a-select-option value="重点客户">重点客户</a-select-option>
                                    <a-select-option value="普通客户">普通客户</a-select-option>
                                    <a-select-option value="非优先客户">非优先客户</a-select-option>
                                </a-select>
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="所在地区" name="region">
                                <a-cascader v-model:value="customer.region" @change="selectedOptions" :options="options"
                                    placeholder="请选择" style="width: 100%" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="详细地址" name="address">
                                <a-input v-model:value="customer.address" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="备注" name="remarks">
                                <a-textarea v-model:value="customer.remarks" :auto-size="{ minRows: 3, maxRows: 3 }" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                </a-form>
            </div>
        </a-modal>
        <!-- 发送邮件对话框 -->
        <a-modal v-model:visible="visibleMail" title="发送邮件" @ok="onSend" @cancel="onCancel" cancelText="取消" okText="发送"
            width="600px" :centered="true">
            <div style="height: 55vh;overflow-y: scroll;padding: 0 15px;">
                <a-form ref="mailSendFormRef" :model="mail" layout="vertical" :rules="mailRules">
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="客户名称" name="customerName">
                                <a-input v-model:value="mail.customerName" disabled />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="收件人" name="receiver">
                                <a-input v-model:value="mail.receiver" disabled />
                            </a-form-item>
                        </a-col>
                        <a-col :span="24">
                            <a-form-item label="邮件主题" name="subject">
                                <a-input v-model:value="mail.subject" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="24">
                            <a-form-item label="邮件内容" name="content">
                                <a-textarea v-model:value="mail.content" placeholder="邮件正文支持HTML语法"
                                    :auto-size="{ minRows: 6, maxRows: 100 }" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="24">
                            <a-form-item label="上传附件" name="attachment">
                                <a-upload-dragger v-model:fileList="fileList" name="file" :multiple="true"
                                    :headers="{ 'X-Requested-With': null }" :action="action" @change="upload"
                                    @drop="upload" @remove="remove">
                                    <p class="ant-upload-drag-icon">
                                        <inbox-outlined></inbox-outlined>
                                    </p>
                                    <p style="font-size: 14px;color: #00000073;">单击或拖动文件到此区域</p>
                                </a-upload-dragger>
                            </a-form-item>
                        </a-col>
                    </a-row>
                </a-form>
            </div>
        </a-modal>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted, createVNode } from 'vue';
import { SearchOutlined, ExclamationCircleOutlined, ExportOutlined, MailTwoTone, InboxOutlined } from '@ant-design/icons-vue';
import moment from 'moment'
import { createCustomer, updateCustomer, sendMailToCustomer, queryCustomerList, queryCustomerInfo, deleteCustomer, customerExport } from '../api/customer';
import { message, Modal } from 'ant-design-vue';
import Spot from '../components/Spot.vue';
import { fileRemove } from '../api/common';
import regionData from '../assets/region';

// 表格字段
const columns = [{
    title: '客户名称',
    dataIndex: 'name',
    width: 200,
    fixed: 'left',
    ellipsis: true,
}, {
    title: '客户来源',
    dataIndex: 'source',
    width: 150,
}, {
    title: '手机号',
    dataIndex: 'phone',
    width: 150,
}, {
    title: '邮箱',
    dataIndex: 'email',
    width: 200,
}, {
    title: '客户行业',
    dataIndex: 'industry',
    width: 150,
}, {
    title: '客户级别',
    dataIndex: 'level',
    width: 150,
}, {
    title: '备注',
    dataIndex: 'remarks',
    width: 150,
    ellipsis: true,
}, {
    title: '成交状态',
    dataIndex: 'status',
    width: 150,
}, {
    title: '详细地址',
    dataIndex: 'address',
    width: 240,
    ellipsis: true,
}, {
    title: '创建时间',
    dataIndex: 'created',
    width: 185,
    customRender: text => {
        return text == 0 ? '' : moment(text.value * 1000).format('YYYY-MM-DD HH:mm:ss')
    }
}, {
    title: '更新时间',
    dataIndex: 'updated',
    width: 185,
    customRender: text => {
        return text.value == 0 ? '' : moment(text.value * 1000).format('YYYY-MM-DD HH:mm:ss')
    }
}, {
    title: '操作',
    dataIndex: 'operation',
    width: 65,
    fixed: 'right',
    ellipsis: true,
}];

// 表单校验
const rules = {
    name: [{ required: true, message: '请输入客户名称', trigger: 'blur' }],
    phone: [{
        pattern: /^(0|86|17951)?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$/,
        message: '手机格式不正确',
        trigger: 'blur',
    }],
    email: [{
        pattern: /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/,
        message: '邮箱格式不正确',
        trigger: 'blur',
    }],
};

const data = reactive({
    customerList: [],
    selectedIds: []
})


const onSelectChange = selectedRowKeys => {
    data.selectedIds = selectedRowKeys
    if (data.selectedIds.length !== 0) {
        disabled.value = false
    } else {
        disabled.value = true
    }
};

// 点击搜索
const onSearch = () => {
    getCustomerList()
};

// 点击全部客户
const onCustomers = () => {
    keyWord.value = ''
    getCustomerList()
}

// 客户属性
let customer = reactive({
    id: undefined,
    name: undefined,
    source: undefined,
    phone: undefined,
    email: undefined,
    industry: undefined,
    level: undefined,
    remarks: undefined,
    region: undefined,
    address: undefined,
    status: undefined,
});

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
const customerFormRef = ref();
const keyWord = ref('')
const visibleMail = ref(false)

// 点击新建客户
const onCreate = () => {
    title.value = '新建客户'
    operation.value = 1
    visible.value = true
}

// 点击客户名称
const onEdit = (row) => {
    title.value = '编辑客户'
    operation.value = 2
    let param = { id: row.id }
    queryCustomerInfo(param).then((res) => {
        if (res.data.code == 0) {
            let p = res.data.data
            customer.id = p.id
            customer.name = p.name
            customer.source = p.source
            customer.phone = p.phone
            customer.email = p.email
            customer.industry = p.industry
            customer.level = p.level
            customer.remarks = p.remarks
            customer.region = p.region
            customer.region = p.region.split(',')
            customer.address = p.address
            customer.status = p.status
        }
    })
    visible.value = true
}

// 点击保存客户
const onSave = () => {
    customerFormRef.value.validateFields().then(() => {
        if (customer.region !== undefined) {
            customer.region = customer.region.toString()
        }
        if (operation.value == 1) {
            createCustomer(customer).then((res) => {
                if (res.data.code == 0) {
                    message.success('保存成功')
                    customerFormRef.value.resetFields()
                    visible.value = false;
                    getCustomerList()
                }
                if (res.data.code == 20001) {
                    message.error('客户名称已经存在')
                }
            })
        }
        if (operation.value == 2) {
            updateCustomer(customer).then((res) => {
                if (res.data.code == 0) {
                    message.success('保存成功')
                    customerFormRef.value.resetFields()
                    visible.value = false;
                    getCustomerList()
                }
            })
        }
    });
};

// 点击删除客户
const onDelete = () => {
    let param = {
        ids: data.selectedIds
    }
    Modal.confirm({
        title: '确定删除选中的' + data.selectedIds.length + '项吗?',
        icon: createVNode(ExclamationCircleOutlined),
        centered: true,
        cancelText: '取消',
        okText: '确定',
        onOk() {
            deleteCustomer(param).then((res) => {
                if (res.data.code == 0) {
                    getCustomerList()
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

// 分页查询客户列表
const onPagination = (page) => {
    pagination.current = page
    getCustomerList()
}

// 初始化数据
onMounted(() => { getCustomerList() })

const getCustomerList = () => {
    let param = {
        name: keyWord.value,
        pageNum: pagination.current,
        pageSize: pagination.pageSize,
    }
    queryCustomerList(param).then((res) => {
        if (res.data.code == 0) {
            pagination.total = res.data.data.total
            data.customerList = res.data.data.list
        }
    })
}

// 导出表格
const onExport = () => {
    customerExport().then((res) => {
        if (res.data.type == 'application/json') {
            message.error('导出错误！')
        } else {
            let blob = new Blob([res.data], {
                type: "application/vnd.ms-excel"
            })
            let a = document.createElement('a')
            a.setAttribute("download", "客户信息.xlsx");
            a.href = window.URL.createObjectURL(blob)
            a.click()
            window.URL.revokeObjectURL(a.href)
        }
    })
}

const mail = reactive({
    customerName: '',
    receiver: '',
    subject: '',
    content: '',
    attachment: undefined
})

const mailSendFormRef = ref()

// 邮件发送表单校验
const mailRules = {
    receiver: [{
        required: true,
        pattern: /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/,
        message: '邮箱格式不正确',
        trigger: 'blur',
    }],
    subject: [{ required: true, message: '请输入邮件主题', trigger: 'blur' }],
    content: [{ required: true, message: '请输入邮件内容', trigger: 'blur' }],
};

// 点击邮件
const onMail = (cname, email) => {
    mail.customerName = cname
    mail.receiver = email
    visibleMail.value = true
}

// 上传附件
const action = ref(import.meta.env.VITE_FILE_UPLOAD_URL)
const upload = (file) => {
    if (file.file.status == 'done') {
        if (file.file.response.code == 0) {
            mail.attachment = file.file.response.data.url
            fileName.value = file.file.response.data.name
        } else {
            message.error('附件上传失败')
        }
    }
}

// 移除附件
const fileName = ref(undefined)
const remove = (file) => {
    if (file.status == 'done') {
        fileRemove({ name: fileName.value })
    }
}

// 点击发送邮件
const onSend = () => {
    mailSendFormRef.value.validateFields().then(() => {
        let param = {
            receiver: mail.receiver,
            subject: mail.subject,
            content: mail.content,
            attachment: mail.attachment
        }
        sendMailToCustomer(param).then((res) => {
            if (res.data.code == 0) {
                message.success("邮件已发送")
                visibleMail.value = false
            }
            if (res.data.code == 50002) {
                message.error("邮件发送失败")
            }
            if (res.data.code == 50003) {
                message.warn("邮件服务未开启")
            }
        })
    })
}

// 点击取消按钮
const onCancel = () => {
    customerFormRef.value.resetFields()
    visible.value = false
};

const options = regionData

const selectedOptions = (value) => {
    customer.region = value
}
</script>

<style scoped>
.ant-table-striped :deep(.table-striped) td {
    background-color: #fafafa;
}
</style>