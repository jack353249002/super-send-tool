<template>
  <a-modal
    title="接收列表"
    width="100%"
    wrap-class-name="full-modal"
    :visible="visible"
    :confirmLoading="loading"
    class="draggable-modal"
    ref="headerUploadRef"
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
              <a-form-item label="状态">
                <a-select v-model="queryParam.status" placeholder="请选择" default-value="-1">
                  <a-select-option value="-1">全部</a-select-option>
                  <a-select-option value="0">未发送</a-select-option>
                  <a-select-option value="1">已发送</a-select-option>
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
        <span slot="created_at" slot-scope="text">
          {{ text | timestampToString }}
        </span>
        <span slot="statusBox" slot-scope="text">
          <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
        </span>
        <span slot="description" slot-scope="text">
          <ellipsis :length="4" tooltip>{{ text }}</ellipsis>
        </span>
      </s-table>
    </a-card>
  </a-modal>
</template>

<script>
import { STable } from '@/components'
import {
  deleteSuperSend, getSendListCon,
  RequestConFactory, sendInfoActionCon, setSuperSendInfoCon, setSuperSendListCon,
  setSuperSendOnline
} from '@/api/super_send'
import { UPDATE_MENU_TYPE } from '@/store/mutation-types'
import VueDraggableResizable from 'vue-draggable-resizable'
import 'vue-draggable-resizable/dist/VueDraggableResizable.css'
import { mapState } from 'vuex'
// import router from '@/router'
// import store from '@/store'
const columns = [
  {
    title: 'id',
    dataIndex: 'id'
  },
  {
    title: '接收者账号',
    dataIndex: 'receive'
  },
  {
    title: '邮件服务器',
    dataIndex: 'send_server_text'
  },
  {
    title: '状态',
    dataIndex: 'status',
    scopedSlots: { customRender: 'statusBox' }
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
  name: 'SendList',
  components: {
    STable,
    VueDraggableResizable
  },
  props: {
    sendID: {
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
    sendID: {
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
      // create model
      confirmLoading: false,
      mdl: null,
      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: { 'keyWords': '', 'status': '-1' },
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
          const requestParameters = Object.assign({}, parameter, this.queryParam)
          requestParameters.send_id = this.sendID
          console.log('loadData request parameters:', requestParameters)
          console.log(this.selectSendInfo)
          const con = RequestConFactory(this.selectSendInfo)
          return getSendListCon(con, requestParameters)
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
      console.log('record', record)
      this.visible = true
      this.mdl = { ...record }
      console.log('record2', this.mdl)
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
<style scoped>
.draggable-box {
  border: 2px solid red;  /* 添加边框 */
  border-radius: 5px;       /* 可选：设置圆角 */
}
</style>
