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
      <span slot="is_root" slot-scope="text">
        <a-badge :status="text | isRootTypeFilter" :text="text | isRootFilter" />
      </span>
      <span slot="description" slot-scope="text">
        <ellipsis :length="4" tooltip>{{ text }}</ellipsis>
      </span>

      <span slot="action" slot-scope="text, record">
        <template>
          <a @click="handelSetUserPassword(record)">设置密码</a>
          <a-divider type="vertical" />
          <a @click="handelSetRole(record)">设置用户角色</a>
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
    <set-user-password
      ref="setUserPasswordModal"
      :visible="visibleUserPassword"
      :loading="confirmLoadingUserPassword"
      :model="mdlSetUserPassword"
      @cancel="handleSetUserPasswordCancel"
      @ok="handleSetUserPasswordOk"
    />
    <set-role
      ref="setRoleModal"
      :visible="visibleSetRole"
      :loading="confirmLoadingRole"
      :model="mdlSetRole"
      :selectSendInfo="selectSendInfo"
      @cancel="handleRoleCancel"
      @ok="false"
    />
  </a-card>
</template>

<script>
import { STable } from '@/components'
import {
  createServerUserCon,
  delServerUserCon, getServerUsersCon,
  login, register, RequestConFactory, setServerUserPasswordCon,
  setSuperSendOnline
} from '@/api/super_send'
import CreateForm from './modules/CreateForm'
import { UPDATE_MENU_TYPE } from '@/store/mutation-types'
import { mapState } from 'vuex'
import SetUserPassword from '@/views/super_send_conn/server_user/modules/SetUserPassword.vue'
import SetRole from '@/views/super_send_conn/server_user/modules/SetRole.vue'
// import router from '@/router'
// import store from '@/store'
const columns = [
  {
    title: 'id',
    dataIndex: 'id'
  },
  {
    title: '用户名',
    dataIndex: 'username'
  },
  {
    title: '是否是root',
    dataIndex: 'is_root',
    scopedSlots: { customRender: 'is_root' }
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '150px',
    scopedSlots: { customRender: 'action' }
  }
]

const isRootMap = {
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
    SetRole,
    SetUserPassword,
    STable,
    CreateForm
  },
  watch: {
    nowSendInfo: {
      handler (newVal, oldVal) {
        this.selectSendInfo = newVal
        console.log('nowSendInfo changed:', newVal)
        this.$refs.table.refresh(true)
      },
      deep: true // 深度监听
    }
  },
  data () {
    this.columns = columns
    return {
      selectSendInfo: {},
      // create model
      visible: false,
      confirmLoading: false,
      mdl: null,
      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: { 'keyWords': '' },
      visibleUserPassword: false,
      confirmLoadingUserPassword: false,
      mdlSetUserPassword: null,
      visibleSetRole: false,
      confirmLoadingRole: false,
      mdlSetRole: null,
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        const requestParameters = Object.assign({}, parameter, this.queryParam)
        console.log('loadData request parameters:', requestParameters)
        if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
          const requestParameters = Object.assign({}, parameter, this.queryParam)
          console.log('loadData request parameters:', requestParameters)
          console.log(this.selectSendInfo)
          const con = RequestConFactory(this.selectSendInfo)
          return getServerUsersCon(con, requestParameters)
            .then(res => {
              if (res.status === 200) {
                if (res.result.data !== null) {
                  return res.result
                } else {
                  return Promise.resolve({ data: [], total: 0 })
                }
              } else if (res.status === 401) {
                this.$message.error(res.message)
                return Promise.resolve({ data: [], total: 0 })
              } else {
                this.$message.error(res.message)
                return Promise.resolve({ data: [], total: 0 })
              }
            })
        } else {
          return Promise.resolve({ data: [], total: 0 })
        }
      },
      selectedRowKeys: [],
      selectedRows: []
    }
  },
  filters: {
    isRootFilter (type) {
      if (type === undefined || type === null || type === '') {
        return '否'
      }
      return isRootMap[type].text
    },
    isRootTypeFilter (type) {
      if (type === undefined || type === null || type === '') {
        return 'default'
      }
      return isRootMap[type].status
    }
  },
  created () {
    // getRoleList({ t: new Date() })
  },
  computed: {
    ...mapState({
      nowSendInfo: state => state.app.nowSendInfo
    }),
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
    handelSetUserPassword (record) {
      this.visibleUserPassword = true
      this.mdlSetUserPassword = { ...record }
    },
    handelSetRole (record) {
      this.$refs.setRoleModal.setRoleStr(record.rolestr)
      this.visibleSetRole = true
      this.mdlSetRole = { ...record }
    },
    handleAdd () {
      this.mdl = null
      this.visible = true
    },
    handleEdit (record) {
      this.visible = true
      this.mdl = { ...record }
    },
    handleDelete (id) {
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
        const con = RequestConFactory(this.selectSendInfo)
        delServerUserCon(con, { 'id': id }).then(res => {
          if (res.status === 200) {
            this.$refs.table.refresh()
            this.$message.success(res.message)
          } else {
            this.$message.error(res.message)
          }
        })
      }
    },
    handleOk () {
      const form = this.$refs.createModal.form
      this.confirmLoading = true
      form.validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
          if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0) {
            const con = RequestConFactory(this.selectSendInfo)
            // 新增
            createServerUserCon(con, values).then(res => {
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
    handleSetUserPasswordOk () {
      const form = this.$refs.setUserPasswordModal.form
      this.confirmLoading = true
      form.validateFields((errors, values) => {
        if (!errors) {
          if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0) {
            const con = RequestConFactory(this.selectSendInfo)
            // 新增
            setServerUserPasswordCon(con, values).then(res => {
              if (res.status === 200) {
                this.visibleUserPassword = false
                this.confirmLoadingUserPassword = false
                // 重置表单数据
                form.resetFields()

                this.$message.success(res.message)
              } else {
                this.$message.error(res.message)
              }
            })
          }
        } else {
          this.confirmLoadingUserPassword = false
        }
      })
    },
    handleSetUserPasswordCancel () {
      this.visibleUserPassword = false
      const form = this.$refs.setUserPasswordModal.form
      form.resetFields() // 清理表单数据（可不做）
    },
    handleRoleCancel () {
      this.visibleSetRole = false
      // const form = this.$refs.setRoleModal.form
      // form.resetFields() // 清理表单数据（可不做）
      this.$refs.setRoleModal.clearSelectPage()
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
