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
      <span slot="created_at" slot-scope="text">
        {{ text | timestampToString }}
      </span>
      <span slot="is_ssl" slot-scope="text">
        <a-badge :status="text | isSSLTypeFilter" :text="text | isSSLFilter" />
      </span>
      <span slot="description" slot-scope="text">
        <ellipsis :length="4" tooltip>{{ text }}</ellipsis>
      </span>

      <span slot="action" slot-scope="text, record">
        <template>
          <a-divider type="vertical" />
          <a @click="handleEdit(record)">配置</a>
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
      :quillEditor="quillEditor"
      :selectSendInfo="selectSendInfo"
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
import dayjs from 'dayjs'
import {
 delMessage, getMessageList, RequestConFactory, setMessage, setSuperSendOnline
} from '@/api/super_send'
import CreateForm from './modules/CreateForm'
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
    title: '标题',
    dataIndex: 'title'
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
      deep: true // 深度监听
      // immediate: true 立即触发一次监听器
    }
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
         if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
          const requestParameters = Object.assign({}, parameter, this.queryParam)
          console.log('loadData request parameters:', requestParameters)
          console.log('loadData', this.selectSendInfo)
          const con = RequestConFactory(this.selectSendInfo)
          return getMessageList(con, requestParameters)
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
      selectedRows: [],
      selectSendInfo: {}
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
    console.log('created')
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
      console.log('mdl', this.mdl)
    },
    handleDelete (id) {
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
        const con = RequestConFactory(this.selectSendInfo)
        delMessage(con, id).then(res => {
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
      const formData = this.$refs.createModal.getFormData()
      console.log('表单数据', formData)
      const form = this.$refs.createModal.form
      this.confirmLoading = true
      form.validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
          if (values.id > 0) {
            // 修改 e.g.
            setMessage(values).then(res => {
              if (res.status === 200) {
                this.visible = false
                this.confirmLoading = false
                // 重置表单数据
                this.$refs.createModal.resetFields()
                // 刷新表格
                this.$refs.table.refresh()

                this.$message.success(res.message)
              } else {
                this.confirmLoading = false
                this.$message.error(res.message)
              }
            })
          } else {
            if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
              const con = RequestConFactory(this.selectSendInfo)
              // 新增
              setMessage(con, formData).then(res => {
                if (res.status === 200) {
                  this.visible = false
                  this.confirmLoading = false
                  // 重置表单数据
                  this.$refs.createModal.resetFields()
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
