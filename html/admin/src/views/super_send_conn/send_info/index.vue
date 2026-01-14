<template>
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
                <a-select-option value="0">未开始</a-select-option>
                <a-select-option value="1">进行中</a-select-option>
                <a-select-option value="2">已结束</a-select-option>
                <a-select-option value="3">暂停</a-select-option>
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
          <a @click="handleEdit(record)">配置</a>
          <a-divider type="vertical" />
          <a @click="sendListShow(record)">接收账号</a>
          <a-divider type="vertical" />
          <a-select
            :value="record.selectedActionStatus"
            @change="sendInfoAction(record, $event)"
          >
            <a-select-option value="-1">停止</a-select-option>
            <a-select-option value="1">启动</a-select-option>
            <a-select-option value="-2">暂停</a-select-option>
            <!-- 添加更多选项 -->
          </a-select>
        </template>
      </span>
    </s-table>
    <create-form
      ref="createModal"
      :selectSendInfo="selectSendInfo"
      :visible="visible"
      :loading="confirmLoading"
      :model="mdl"
      @cancel="handleCancel"
      @ok="handleOk"
    />
    <send-list
      ref="sendListModal"
      :selectSendInfo="selectSendInfo"
      :visible="sendListvisible"
      :loading="sendListConfirmLoading"
      :sendID="sendID"
      @cancel="handleSendListCancel"
    />
    <step-by-step-modal ref="modal" @ok="handleOk"/>
  </a-card>
</template>

<script>
import { STable } from '@/components'
import dayjs from 'dayjs'
import {
  deleteSuperSend,
  getSuperSendListCon,
  RequestConFactory, sendInfoActionCon, setSuperSendInfoCon, setSuperSendListCon,
  setSuperSendOnline
} from '@/api/super_send'
import CreateForm from './modules/CreateForm'
import { UPDATE_MENU_TYPE } from '@/store/mutation-types'
import VueDraggableResizable from 'vue-draggable-resizable'
import 'vue-draggable-resizable/dist/VueDraggableResizable.css'
import { mapState } from 'vuex'
import SendList from '@/views/super_send_conn/send_info/modules/SendList.vue'
// import router from '@/router'
// import store from '@/store'
const columns = [
  {
    title: 'id',
    dataIndex: 'id'
  },
  {
    title: '配置标题',
    dataIndex: 'title'
  },
  {
    title: '消息标题',
    dataIndex: 'message_title'
  },
  {
    title: '发送服务器标题',
    dataIndex: 'send_server_text'
  },
  {
    title: '发送数量',
    dataIndex: 'send_list_count'
  },
  {
    title: '发送成功数量',
    dataIndex: 'success_count'
  },
  {
    title: '参数',
    dataIndex: 'params'
  },
  {
    title: '状态',
    dataIndex: 'status',
    scopedSlots: { customRender: 'statusBox' }
  },
  {
    title: '创建时间',
    dataIndex: 'createtime',
    customRender: (text) => {
      return dayjs(text * 1000).format('YYYY-MM-DD HH:mm:ss')
    }
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '250px',
    scopedSlots: { customRender: 'action' }
  }
]

const statusMap = {
  0: {
    status: 'default', // 或者 'not_started'
    text: '未开始'
  },
  1: {
    status: 'default',
    text: '进行中'
  },
  2: {
    status: 'default',
    text: '已结束'
  },
  3: {
    status: 'default',
    text: '暂停'
  }
}

export default {
  name: 'TableList',
  components: {
    SendList,
    STable,
    CreateForm,
    VueDraggableResizable
  },
  watch: {
    nowSendInfo: {
      handler (newVal, oldVal) {
        this.selectSendInfo = newVal
        console.log('nowSendInfo changed:', newVal)
        this.$refs.table.refresh(true)
      },
      deep: true, // 深度监听
      immediate: false // 立即触发一次监听器
    }
  },
  data () {
    this.columns = columns
    return {
      sendID: 0,
      // create model
      visible: false,
      sendListvisible: false,
      sendListConfirmLoading: false,
      confirmLoading: false,
      mdl: null,
      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: { 'keyWords': '', 'status': null },
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
         if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
          const requestParameters = Object.assign({}, parameter, this.queryParam)
          console.log('loadData request parameters:', requestParameters)
          console.log(this.selectSendInfo)
          const con = RequestConFactory(this.selectSendInfo)
          return getSuperSendListCon(con, requestParameters)
            .then(res => {
              if (res == null) {
                return { data: [], pageSize: 0, total: 0, pageNo: 0, totalPage: 0, totalCount: 0 }
              }
              if (res.status === 200) {
                if (res.result.data !== null) {
                  res.result.data.forEach(item => {
                    if (item.status == null) {
                      item.status = 0
                    }
                    if (item.status === 0 || item.status === 2) {
                      item.selectedActionStatus = '-1'
                    }
                    if (item.status === 1) {
                      item.selectedActionStatus = '1'
                    }
                    if (item.status === 3) {
                      item.selectedActionStatus = '-2'
                    }
                  })
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
      selectedRows: [],
      selectSendInfo: {}
    }
  },
  filters: {
    statusFilter (type) {
      if (type === undefined || type === null || type === '') {
        return statusMap[0].text
      } else {
        return statusMap[type].text
      }
    },
    statusTypeFilter (type) {
      if (type === undefined || type === null || type === '') {
        return statusMap[0].status
      } else {
        return statusMap[type].status
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
    sendInfoAction (record, value) {
      // 更新记录中的选项值
     const sendId = record.id
     const status = Number(value)
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
        const con = RequestConFactory(this.selectSendInfo)
        sendInfoActionCon(con, { 'send_id': sendId, 'status': status }).then(res => {
          if (res.status === 200) {
            this.$refs.table.refresh()
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
      if (record.params == null) {
          record.params = ''
      }
      console.log('record', record)
      this.visible = true
      this.mdl = { ...record }
      console.log('record2', this.mdl)
    },
    sendListShow (record) {
      this.sendID = record.id
      this.sendListvisible = true
    },
    handleSendListCancel () {
      this.sendID = 0
      this.sendListvisible = false
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
            if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0) {
              const con = RequestConFactory(this.selectSendInfo)
              setSuperSendInfoCon(con, formData).then(res => {
                if (res.status === 200) {
                  // 重置表单数据
                  this.$refs.createModal.resetFields()
                  this.visible = false
                  this.confirmLoading = false
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
            if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0) {
              const con = RequestConFactory(this.selectSendInfo)
              setSuperSendListCon(con, formData).then(res => {
                if (res.status === 200) {
                  // 重置表单数据
                  this.$refs.createModal.resetFields()
                  this.visible = false
                  this.confirmLoading = false
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
