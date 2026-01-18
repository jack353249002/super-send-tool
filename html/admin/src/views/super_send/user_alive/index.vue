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
          <a @click="handleEdit(record)">编辑</a>
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
  addCheckAlive, delCheckAlive,
  getCheckAliveList,
  login, register, setCheckAlive,
  setSuperSendOnline
} from '@/api/super_send'
import CreateForm from './modules/CreateForm'
import { UPDATE_MENU_TYPE } from '@/store/mutation-types'
import dayjs from 'dayjs'
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
    title: '签到时间',
    dataIndex: 'day_login_first_time_text',
    customRender: (text) => {
      return text
    }
  },
  {
    title: '消息id',
    dataIndex: 'message_id'
  },
  {
    title: '距离签到时间超时时长(秒)',
    dataIndex: 'send_email_action_timeout'
  },
  {
    title: '连接器地址',
    dataIndex: 'super_send_conn_info_address'
  },
  {
    title: '最后通信时间',
    dataIndex: 'last_ping_time',
    customRender: (text) => {
      if (text === 0) {
        return '无'
      }
      return dayjs(text * 1000).format('YYYY-MM-DD HH:mm:ss')
    }
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
      queryParam: { 'keyWords': '' },
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        const requestParameters = Object.assign({}, parameter, this.queryParam)
        console.log('loadData request parameters:', requestParameters)
        return getCheckAliveList(requestParameters)
          .then(res => {
            if (res != null) {
              for (let i = 0; i < res.result.data.length; i++) {
                if (res.result.data[i].day_login_first_time === 0) {
                  res.result.data[i].day_login_first_time_text = ''
                } else {
                  res.result.data[i].day_login_first_time_text = dayjs(res.result.data[i].day_login_first_time * 1000).format('YYYY-MM-DD HH:mm:ss')
                }
              }
              return res.result
            } else {
              return { data: [], pageSize: 0, total: 0, pageNo: 0, totalPage: 0, totalCount: 0 }
            }
          }).catch(error => {
            console.log(error)
            this.$message.error('网络错误')
            return { data: [], pageSize: 0, total: 0, pageNo: 0, totalPage: 0, totalCount: 0 }
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
      }).catch(error => {
        console.log(error)
        this.$message.error('网络错误')
      })
    },
    handleAdd () {
      this.mdl = null
      this.visible = true
    },
    handleEdit (record) {
      record.day_login_first_time = dayjs.unix(record.day_login_first_time).format('YYYY-MM-DD HH:mm:ss')
      record.password = ''
      this.visible = true
      this.mdl = { ...record }
    },
    handleDelete (id) {
      delCheckAlive(id).then(res => {
        if (res.status === 200) {
          this.$refs.table.refresh()
          this.$message.success(res.message)
        } else {
          this.$message.error(res.message)
        }
      }).catch(error => {
        console.log(error)
        this.$message.error('网络错误')
      })
    },
    handleOk () {
      const form = this.$refs.createModal.form
      this.confirmLoading = true
      form.validateFields((errors, values) => {
        if (!errors) {
          values.send_id = parseInt(values.send_id)
          values.super_send_conn_info_id = parseInt(values.super_send_conn_info_id)
          values.send_email_action_timeout = parseInt(values.send_email_action_timeout)
          values.id = parseInt(values.id)
          values.message_id = parseInt(values.message_id)
          console.log('values', values)
          // 创建 Date 对象
          const date = new Date(values.day_login_first_time_text)
          // 获取毫秒级时间戳，然后除以1000转为秒级，并转为整数
          values.day_login_first_time = Math.floor(date.getTime() / 1000)
          if (values.id > 0) {
            // 修改 e.g.
            setCheckAlive(values).then(res => {
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
            }).catch(error => {
              console.log(error)
              this.$message.error('网络错误')
              this.confirmLoading = false
            })
          } else {
            // 新增
             addCheckAlive(values).then(res => {
               if (res.status === 200) {
                 this.visible = false
                 this.confirmLoading = false
                 // 重置表单数据
                 this.$refs.createModal.resetFields()
                 // 刷新表格
                 this.$refs.table.refresh()

                 this.$message.success(res.message)
               } else {
                 this.$message.error(res.message)
                 this.confirmLoading = false
               }
            }).catch(error => {
              console.log(error)
               this.$message.error('网络错误')
               this.confirmLoading = false
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
