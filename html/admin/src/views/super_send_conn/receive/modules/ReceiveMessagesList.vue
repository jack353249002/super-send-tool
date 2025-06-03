<template>
  <a-modal
    title="接收邮件列表"
    width="100%"
    wrap-class-name="full-modal"
    :visible="visible"
    :confirmLoading="loading"
    class="draggable-modal"
    ref="headerUploadRef"
    :footer="null"
    @cancel="() => { $emit('cancel') }"
  >
    <a-card :bordered="false">
      <div class="table-page-search-wrapper">
        <a-form layout="inline">
          <a-row :gutter="48">
            <a-col :md="8" :sm="24">
              <!--            <a-form-item label="关键词">
              <a-input v-model="queryParam.keyWords" placeholder=""/>
            </a-form-item>-->
            </a-col>
            <a-col :md="!advanced && 8 || 24" :sm="24">
              <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
                <a-button type="primary" @click="$refs.table.refresh(true)">刷新</a-button>
              </span>
            </a-col>
          </a-row>
        </a-form>
      </div>
      <div class="table-operator">
      </div>
      <div class="table-container">
        <s-table
          ref="table"
          size="default"
          rowKey="key"
          :columns="columns"
          :data="loadData"
          :alert="true"
          :rowSelection="rowSelection"
          :showPagination="false"
        >
          <span slot="serial" slot-scope="text, record, index">
            {{ index + 1 }}
          </span>
          <span slot="created_at" slot-scope="text">
            {{ text | timestampToString }}
          </span>
          <span slot="statusBox" slot-scope="text">
            <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
          </span>
          <span slot="description" slot-scope="text">
            <ellipsis :length="4" tooltip>{{ text }}</ellipsis>
          </span>
          <span slot="action" slot-scope="text, record">
            <template>
              <a-divider type="vertical" />
              <a @click="handleEdit(record)">消息详情</a>
            </template>
          </span>
        </s-table>
          <receive-message-info
            ref="createModal"
            :selectSendInfo="selectSendInfo"
            :receiveMessageInfoVisible="receiveMessageInfoVisible"
            :loading="confirmLoading"
            :model="mdl"
            @cancel="handleCancel"
          />
      </div>
    </a-card>
  </a-modal>
</template>

<script>
import { STable } from '@/components'
import {
  deleteSuperSend, getReceiveMessagesInfo, getReceiveMessagesList,
  RequestConFactory, sendInfoActionCon, setSuperSendInfoCon, setSuperSendListCon,
  setSuperSendOnline
} from '@/api/super_send'
import { UPDATE_MENU_TYPE } from '@/store/mutation-types'
import VueDraggableResizable from 'vue-draggable-resizable'
import 'vue-draggable-resizable/dist/VueDraggableResizable.css'
import { mapState } from 'vuex'
import CreateForm from '@/views/super_send_conn/receive/modules/CreateForm.vue'
import ReceiveMessageInfo from '@/views/super_send_conn/receive/modules/ReceiveMessageInfo.vue'
// import router from '@/router'
// import store from '@/store'
const columns = [
  {
    title: 'subject',
    dataIndex: 'subject'
  },
  {
    title: 'imap服务器',
    dataIndex: 'imap_server_text'
  },
  {
    title: '发送者邮箱',
    dataIndex: 'sender_address'
  },
  {
    title: '日期',
    dataIndex: 'date'
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '250px',
    scopedSlots: { customRender: 'action' }
  }
]

const statusMap = {
  '0': {
    status: 'default', // 或者 'not_started'
    text: '未发送'
  },
  '1': {
    status: 'default',
    text: '已发送'
  }
}

export default {
  name: 'ReceiveMessagesList',
  components: {
    ReceiveMessageInfo,
    CreateForm,
    STable,
    VueDraggableResizable
  },
  props: {
    receiveID: {
      type: Number,
      required: true
    },
    selectSendInfo: {
      type: Object,
      required: true
    },
    visible: {
      type: Boolean,
      required: true
    },
    loading: {
      type: Boolean,
      default: () => false
    },
    model: {
      type: Object,
      default: () => null
    }
  },
  watch: {
    nowSendInfo: {
      handler (newVal, oldVal) {
        this.selectSendInfo = newVal
        console.log('nowSendInfo changed:', newVal)
        this.$refs.table.refresh(true)
      },
      deep: true // 深度监听
    },
    receiveID: {
      handler (newVal, oldVal) {
        if (newVal !== 0) {
          this.reloadList()
        }
      },
      deep: true
    }
  },
  data () {
    this.columns = columns
    return {
      receiveMessageInfoVisible: false,
      // create model
      confirmLoading: false,
      mdl: null,
      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: { },
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
          const requestParameters = Object.assign({}, parameter, this.queryParam)
          requestParameters.receive_id = this.receiveID
          console.log('loadData request parameters:', requestParameters)
          console.log(this.selectSendInfo)
          const con = RequestConFactory(this.selectSendInfo)
          return getReceiveMessagesList(con, requestParameters)
            .then(res => {
              if (res.status === 200) {
                if (res.result.data !== null) {
                 /* res.result.data.forEach(item => {
                    if (item.status == null) {
                      item.status = 0
                    }
                  }) */
                  return res.result
                } else {
                  return Promise.resolve({ data: [], total: 0 })
                }
              } else if (res.status === 401) {
                return Promise.resolve({ data: [], total: 0 })
              } else {
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
    statusFilter (type) {
      if (type === undefined || type === null || type === '') {
        return statusMap['0'].text
      } else {
        const typeString = String(type)
        return statusMap[typeString].text
      }
    },
    statusTypeFilter (type) {
      if (type === undefined || type === null || type === '') {
        return statusMap['0'].status
      } else {
        const typeString = String(type)
        return statusMap[typeString].status
      }
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
    reloadList () {
      this.queryParam.status = '-1'
      this.$refs.table.refresh(true)
    },
    sendInfoAction (record, value) {
      // 更新记录中的选项值
      const sendId = record.id
      const status = Number(value)
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0) {
        const con = RequestConFactory(this.selectSendInfo)
        sendInfoActionCon(con, { 'send_id': sendId, 'status': status }).then(res => {
          if (res.status === 200) {
            this.$message.success(res.message)
          } else {
            this.$message.error(res.message)
          }
        })
      }
    },
    updateMenu () {
      this.$store.commit(UPDATE_MENU_TYPE)
    },
    handleOnline (id, online) {
      setSuperSendOnline(id, online).then(res => {
        if (res.status === 200) {
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
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0) {
        const con = RequestConFactory(this.selectSendInfo)
        getReceiveMessagesInfo(con, { uid: record.uid, imap_server_id: record.imap_server_id }).then(res => {
          if (res.status === 200) {
            this.mdl = res.result
            this.receiveMessageInfoVisible = true
          }
        })
      }
      // this.mdl = { ...record }
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
      const saveData = this.$refs.createModal.saveData
      form.validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
          if (saveData.id > 0) {
            const formData = this.$refs.createModal.getFormData('update')
            console.log('表单数据', formData)
            // 修改 e.g.
            if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
              const con = RequestConFactory(this.selectSendInfo)
              setSuperSendInfoCon(con, formData).then(res => {
                if (res.status === 200) {
                  this.visible = false
                  this.confirmLoading = false
                  // 重置表单数据
                  form.resetFields()
                  // 刷新表格
                  this.$refs.table.refresh()

                  this.$message.success(res.message)
                } else {
                  this.confirmLoading = false
                  this.$message.error(res.message)
                }
              })
            }
          } else {
            const formData = this.$refs.createModal.getFormData('add')
            console.log('表单数据', formData)
            if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
              const con = RequestConFactory(this.selectSendInfo)
              setSuperSendListCon(con, formData).then(res => {
                if (res.status === 200) {
                  this.visible = false
                  this.confirmLoading = false
                  // 重置表单数据
                  form.resetFields()
                  // 刷新表格
                  this.$refs.table.refresh()

                  this.$message.success(res.message)
                } else {
                  this.confirmLoading = false
                  this.$message.error(res.message)
                }
              })
            }
          }
        } else {
          this.confirmLoading = false
        }
      })
    },
    handleCancel () {
      this.receiveMessageInfoVisible = false

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
<style scoped>
.draggable-box {
  border: 2px solid red;  /* 添加边框 */
  border-radius: 5px;       /* 可选：设置圆角 */
}
 .full-modal {
   height: 100%;
 }

.draggable-modal {
  height: 100%;
}

.ant-modal-content {
  height: 100%;
}

.ant-modal-body {
  height: 100%;
  padding: 0;
}

.ant-card {
  height: 100%;
}

.table-page-search-wrapper {
  padding: 16px;
}

.table-operator {
  margin-bottom: 16px;
}

.table-container {
  height: 400px; /* 调整高度以适应搜索和操作区域 */
  overflow-y: auto; /* 启用垂直滚动条 */
}
</style>
