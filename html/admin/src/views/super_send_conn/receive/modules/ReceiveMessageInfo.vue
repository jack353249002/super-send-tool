<template>
  <a-modal
    title="接收消息详情"
    width="100%"
    wrap-class-name="full-modal"
    :visible="receiveMessageInfoVisible"
    :confirmLoading="loading"
    class="draggable-modal"
    ref="headerUploadRef"
    :footer="null"
    @cancel="() => { $emit('cancel') }"
  >
    <a-spin :spinning="loading">
      <a-form class="box" :form="form" v-bind="formLayout">
        <!-- 检查是否有 id 并且大于0，大于0是修改。其他是新增，新增不显示主键ID -->
        <a-form-item label="标题">
          <a-input v-decorator="['subject', {rules: [{required: true, message: '标题'}]}]" />
        </a-form-item>
        <a-form-item label="imap服务器">
          <a-input v-decorator="['imap_server_text', {rules: [{required: true, message: 'imap服务器'}]}]" />
        </a-form-item>
        <a-form-item label="sender_address">
          <a-input v-decorator="['sender_address', {rules: [{required: true, message: 'sender_address'}]}]" />
        </a-form-item>
        <a-form-item label="发送时间">
          <a-input v-decorator="['date', {rules: [{required: true, message: '发送时间'}]}]" />
        </a-form-item>
        <a-form-item label="message_id">
          <a-input v-decorator="['message_id', {rules: [{required: true, message: 'message_id'}]}]" />
        </a-form-item>
        <a-form-item label="邮件内容">
          <div class="email-content"  v-html="this.model.body"></div>
        </a-form-item>
        <a-form-item label="附件列表">
          <ul>
            <li v-for="(attachment, index) in this.model.file_paths" :key="index">
              <a :href="attachment" :download="getFileName(attachment)">{{attachment}}</a>
            </li>
          </ul>
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>
import 'vue-draggable-resizable/dist/VueDraggableResizable.css'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import { getSmtpServerList, RequestConFactory, superSendApi } from '@/api/super_send'
import pick from 'lodash.pick'
const fields = ['subject', 'sender_address', 'date', 'message_id', 'body', 'file_paths', 'imap_server_text']// 表单字段
export default {
  components: {
  },
  watch: {
  },
  props: {
    selectSendInfo: {
      type: Object,
      required: true
    },
    receiveMessageInfoVisible: {
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
  mounted () {
    // this.getSmtpList()
  },
  methods: {
    getFormData (type) {
      // 返回表单数据
      if (type === 'add') {
        return {
          title: this.saveData.title,
          message_id: this.saveData.message_id,
          send_server_id: this.saveData.send_server_id,
          receive: this.saveData.receive
        }
      } else {
        return {
          id: this.saveData.id,
          send_server_id: this.saveData.send_server_id
        }
      }
    },
    formatResult (response) {
      console.log('format', response)
      const listTemp = []
      for (let i = 0; i < response.result.data.length; i++) {
        listTemp.push({ id: response.result.data[i].id, name: response.result.data[i].title })
      }
      return {
        list: listTemp, // 数据列表
        totalRow: response.result.totalCount // 总条目数（分页用）
      }
    },
    smtpFormatResult (response) {
      console.log('format', response)
      const listTemp = []
      for (let i = 0; i < response.result.data.length; i++) {
        listTemp.push({ id: response.result.data[i].id, name: response.result.data[i].name })
      }
      return {
        list: listTemp, // 数据列表
        totalRow: response.result.totalCount // 总条目数（分页用）
      }
    },
    dataLoad (a1, a2, a3) {
      console.log(a1, a2, a3)
    },
    handlePageChange (pageNo) {
      // this.SearchMessageList(this.keywords, pageNo)
    },
    handleSearch (keywords) {
      this.keywords = keywords
      // this.SearchMessageList(keywords, 1)
    },
    getSmtpList () {
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
        const con = RequestConFactory(this.selectSendInfo)
        getSmtpServerList(con, {})
          .then(res => {
            console.log(res)
            if (res.status === 200) {
              if (res.result.data !== null) {
                const listTemp = []
                for (let i = 0; i < res.result.data.length; i++) {
                  listTemp.push({ id: res.result.data[i].id, name: res.result.data[i].name })
                }
                // this.total = res.result.totalCount
                this.smtpList = listTemp
                // this.messageListInfo = { list: res.result.data, total: res.result.totalPage, page: res.result.pageNo, size: res.result.pageSize }
              } else {
                this.smtpList = []
              }
            } else if (res.status === 401) {
              this.smtpList = []
            } else {
              this.smtpList = []
            }
          })
      } else {
        this.smtpList = []
      }
    },
    getFileName (url) {
      // 提取文件名
      const parts = url.split('/')
      return parts[parts.length - 1]
    }
  },
  data () {
    this.formLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      }
    }
    return {
      saveData: {},
      messageListPath: superSendApi.GetMessageList,
      smtpServerListPath: superSendApi.GetSmtpServerList,
      selectedValue: null, // 用于绑定选择的值
      pageSize: 10, // 每页条数
      keywords: '',
      list: [],
      filteredOptions: {},
      form: this.$form.createForm(this),
      width: 500,
      height: 500,
      x: 0,
      y: 0,
      fileList: [],
      editorOption: {
        // 编辑器选项
      },
      attrCount: 0,
      isEditingAttach: false,
      serverIds: null,
      smtpList: [],
      messageSelectPage: {},
      serverSelectPage: {}
    }
  },
  created () {
    console.log('custom modal created')

    // 防止表单未注册
    fields.forEach(v => this.form.getFieldDecorator(v))

    // 当 model 发生改变时，为表单设置值
    this.$watch('model', () => {
      // 设置表单的初始值
      this.model && this.form.setFieldsValue(pick(this.model, fields))
    })
  }
}
</script>

<style lang="less">
.full-modal {
  .ant-modal {
    max-width: 100%;
    top: 0;
    padding-bottom: 0;
    margin: 0;
  }
  .ant-modal-content {
    display: flex;
    flex-direction: column;
    height: calc(100vh);
  }
  .ant-modal-body {
    flex: 1;
  }
}
.box {
  height: 600px; /* 调整高度以适应搜索和操作区域 */
  overflow-y: auto; /* 启用垂直滚动条 */
}
.email-content {
  width: 100%;
  height: 500px;
  overflow-y: auto;
  border: 1px solid #ccc;
  padding: 10px;
  background-color: #f9f9f9;
}
</style>
