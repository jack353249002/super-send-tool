<template>
  <a-card :bordered="false">
    <div class="table-page-search-wrapper">
      <a-form layout="inline">
        <a-row :gutter="48">
          <a-col :md="8" :sm="24">
            <a-form-item label="关键词">
              <a-input v-model="queryParam.keyWords" placeholder=""/>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <a-form-item label="ssl状态">
              <a-select v-model="queryParam.is_ssl" placeholder="请选择" default-value="-1">
                <a-select-option value="-1">全部</a-select-option>
                <a-select-option value="0">禁用</a-select-option>
                <a-select-option value="1">启用</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :md="!advanced && 8 || 24" :sm="24">
            <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
              <a-button type="primary" @click="$refs.table.refresh(true)">查询</a-button>
              <a-button style="margin-left: 8px" @click="() => queryParam = {}">重置</a-button>
            </span>
          </a-col>
        </a-row>
      </a-form>
    </div>
    <div class="table-operator">
      <a-button type="primary" icon="plus" @click="handleAdd">新建</a-button>
      <a-dropdown v-action:edit v-if="selectedRowKeys.length > 0">
        <a-menu slot="overlay">
          <a-menu-item key="1"><a-icon type="delete" />删除</a-menu-item>
        </a-menu>
      </a-dropdown>
    </div>

    <s-table
      ref="table"
      size="default"
      rowKey="key"
      :columns="columns"
      :data="loadData"
      :alert="true"
      :rowSelection="rowSelection"
      showPagination="auto"
    >
      <span slot="serial" slot-scope="text, record, index">
        {{ index + 1 }}
      </span>
      <span slot="is_ssl" slot-scope="text">
        <a-badge :status="text | isSSLTypeFilter" :text="text | isSSLFilter" />
      </span>
      <span slot="description" slot-scope="text">
        <ellipsis :length="4" tooltip>{{ text }}</ellipsis>
      </span>

      <span slot="action" slot-scope="text, record">
        <template>
          <a @click="handleOnline(record.id, record.online === 1 ? 0 : 1)">{{ record.online === 1 ? '关闭连接' : '连接' }}</a>
          <a-divider type="vertical" />
          <a @click="handleEdit(record)">配置</a>
          <a-divider type="vertical" />
          <a @click="handleRegister(record)">注册</a>
          <a-divider type="vertical" />
          <a-popconfirm
            title="确定要删除此记录吗？"
            ok-text="确定"
            cancel-text="取消"
            @confirm="handleDelete(record.id)"
          >
            <a-button type="link" danger>删除</a-button>
          </a-popconfirm>
        </template>
      </span>
    </s-table>
    <create-form
      ref="createModal"
      :visible="visible"
      :loading="confirmLoading"
      :model="mdl"
      @cancel="handleCancel"
      @ok="handleOk"
    />
    <step-by-step-modal ref="modal" @ok="handleOk"/>
  </a-card>
</template>

<script>
import { STable } from '@/components'
import {
  addSuperSend,
  deleteSuperSend,
  getSuperSendList, login, register,
  setSuperSendOnline,
  updateSuperSend
} from '@/api/super_send'
import CreateForm from './modules/CreateForm'
import { UPDATE_MENU_TYPE } from '@/store/mutation-types'
// import router from '@/router'
// import store from '@/store'
const columns = [
  {
    title: 'id',
    dataIndex: 'id'
  },
  {
    title: '地址',
    dataIndex: 'address'
  },
  {
    title: '是否开启ssl',
    dataIndex: 'is_ssl',
    scopedSlots: { customRender: 'is_ssl' }
  },
  {
    title: '用户名',
    dataIndex: 'username'
  },
  {
    title: '密码',
    dataIndex: 'password'
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '150px',
    scopedSlots: { customRender: 'action' }
  }
]

const isSSLMap = {
  0: {
    status: 'default',
    text: '否'
  },
  1: {
    status: 'success',
    text: '是'
  }
}

export default {
  name: 'TableList',
  components: {
    STable,
    CreateForm
  },
  data () {
    this.columns = columns
    return {
      // create model
      visible: false,
      confirmLoading: false,
      mdl: null,
      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: { 'keyWords': '', 'is_ssl': '-1' },
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        const requestParameters = Object.assign({}, parameter, this.queryParam)
        console.log('loadData request parameters:', requestParameters)
        return getSuperSendList(requestParameters)
          .then(res => {
            return res.result
          })
      },
      selectedRowKeys: [],
      selectedRows: []
    }
  },
  filters: {
    isSSLFilter (type) {
      return isSSLMap[type].text
    },
    isSSLTypeFilter (type) {
      return isSSLMap[type].status
    }
  },
  created () {
    // getRoleList({ t: new Date() })
  },
  computed: {
    rowSelection () {
      return {
        selectedRowKeys: this.selectedRowKeys,
        onChange: this.onSelectChange
      }
    }
  },
  methods: {
    updateMenu () {
      this.$store.commit(UPDATE_MENU_TYPE)
    },
    handleRegister (record) {
      register(record.token, record.id, record).then(res => {
          if (res.status === 200) {
            this.$message.success(res.message)
          } else {
            this.$message.error(res.message)
          }
      })
    },
    handleOnline (id, online) {
      setSuperSendOnline(id, online).then(res => {
        if (res.status === 200) {
          if (online === 1) {
            login(id).then(res => {
              if (res.status === 200) {
                this.$message.success(res.message)
              } else {
                this.$message.error(res.message)
              }
            })
          }
          this.$refs.table.refresh()
          this.updateMenu()
          this.$message.success(res.message)
        } else {
          this.$message.error(res.message)
        }
      })
    },
    handleAdd () {
      this.mdl = null
      this.visible = true
    },
    handleEdit (record) {
      this.visible = true
      record.is_ssl = record.is_ssl.toString()
      this.mdl = { ...record }
    },
    handleDelete (id) {
      deleteSuperSend(id).then(res => {
        if (res.status === 200) {
          this.$refs.table.refresh()
          this.$message.success(res.message)
        } else {
          this.$message.error(res.message)
        }
      })
    },
    handleOk () {
      const form = this.$refs.createModal.form
      this.confirmLoading = true
      form.validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
          if (values.id > 0) {
            // 修改 e.g.
            updateSuperSend(values).then(res => {
              if (res.status === 200) {
                this.visible = false
                this.confirmLoading = false
                // 重置表单数据
                form.resetFields()
                // 刷新表格
                this.$refs.table.refresh()

                this.$message.success(res.message)
              } else {
                this.$message.error(res.message)
              }
            })
          } else {
            // 新增
             addSuperSend(values).then(res => {
               if (res.status === 200) {
                 this.visible = false
                 this.confirmLoading = false
                 // 重置表单数据
                 form.resetFields()
                 // 刷新表格
                 this.$refs.table.refresh()

                 this.$message.success(res.message)
               } else {
                 this.$message.error(res.message)
               }
            })
          }
        } else {
          this.confirmLoading = false
        }
      })
    },
    handleCancel () {
      this.visible = false

      const form = this.$refs.createModal.form
      form.resetFields() // 清理表单数据（可不做）
    },
    handleSub (record) {
      if (record.status !== 0) {
        this.$message.info(`${record.no} 订阅成功`)
      } else {
        this.$message.error(`${record.no} 订阅失败，规则已关闭`)
      }
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    resetSearchForm () {
      this.queryParam = {
        keywords: ''
      }
    }
  }
}
</script>
