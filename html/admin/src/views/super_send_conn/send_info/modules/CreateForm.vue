<template>
  <a-modal
    title="设置发送配置"
    width="100%"
    wrap-class-name="full-modal"
    :visible="visible"
    :confirmLoading="loading"
    class="draggable-modal"
    ref="headerUploadRef"
    @ok="() => { $emit('ok') }"
    @cancel="() => { $emit('cancel') }"
  >
    <a-spin :spinning="loading">
      <a-form :form="form" v-bind="formLayout">
        <!-- 检查是否有 id 并且大于0，大于0是修改。其他是新增，新增不显示主键ID -->
        <a-form-item v-show="model && model.id > 0" label="主键ID">
          <a-input v-model="saveData.id" disabled />
        </a-form-item>
        <a-form-item label="标题" v-show="!model">
          <a-input v-model="saveData.title" />
        </a-form-item>
        <a-form-item label="选择消息模板" v-show="!model">
          <v-selectpage
            placeholder="请选择一个选项"
            field="name"
            value-field="id"
            v-model="saveData.message_id"
            :data="messageListPath"
            :params="selectSendInfo"
            :result-format="formatResult"
            ref="messageSelectPage"
          >
          </v-selectpage>
        </a-form-item>
        <a-form-item label="选择服务器">
          <v-selectpage
            placeholder="请选择一个选项"
            field="name"
            value-field="id"
            v-model="saveData.send_server_id"
            :data="smtpList"
            :params="selectSendInfo"
            :multiple="true"
            :pagination="false"
            ref="serverSelectPage"
          >
          </v-selectpage>
        </a-form-item>
        <a-form-item label="接收者账号(逗号分割)" v-show="!model">
          <a-input v-model="saveData.receive" />
        </a-form-item>
        <a-form-item label="参数(json{变量名:变量值})">
          <a-input v-model="saveData.params" />
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
import { getMessageList, getSmtpServerList, RequestConFactory, superSendApi } from '@/api/super_send'
// const fields = ['message_id', 'title', 'send_model', 'send_server_id', 'dispatch_model', 'receive']// 表单字段
export default {
  components: {
  },
  watch: {
    selectSendInfo: {
      handler (newVal, oldVal) {
        this.getSmtpList()
      },
      deep: true // 深度监听
    }
  },
  props: {
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
  mounted () {
  },
  methods: {
    resetFields () {
      this.saveData.send_server_id = null
      this.saveData.message_id = null
      this.saveData.params = null
      this.saveData = {}
    },
    getFormData (type) {
      // 返回表单数据
      if (type === 'add') {
        return {
          title: this.saveData.title,
          message_id: this.saveData.message_id,
          send_server_id: this.saveData.send_server_id,
          params: this.saveData.params,
          receive: this.saveData.receive
        }
      } else {
        return {
          id: this.saveData.id,
          send_server_id: this.saveData.send_server_id,
          params: this.saveData.params
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
    SearchMessageList (keywords, pageNo) {
      const requestParameters = { 'keyWords': keywords, 'pageSize': 10, 'pageNo': pageNo }
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0) {
        const con = RequestConFactory(this.selectSendInfo)
        getMessageList(con, requestParameters)
          .then(res => {
            console.log(res)
            if (res.status === 200) {
              if (res.result.data !== null) {
                const listTemp = []
                for (let i = 0; i < res.result.data.length; i++) {
                  listTemp.push({ id: res.result.data[i].id, name: res.result.data[i].title })
                }
                // this.total = res.result.totalCount
                this.list = listTemp
                this.page = pageNo
                this.total = 50
                // this.messageListInfo = { list: res.result.data, total: res.result.totalPage, page: res.result.pageNo, size: res.result.pageSize }
              } else {
                this.list = []
                this.total = 0
                this.page = 1
              }
            } else if (res.status === 401) {
              this.list = []
              this.total = 0
              this.page = 1
            } else {
              this.list = []
              this.total = 0
              this.page = 1
            }
          })
      } else {
        this.total = 0
        this.page = 1
        this.list = []
      }
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
      saveData: { },
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
     // this.SearchMessageList('', 1)
    // 防止表单未注册
    // eslint-disable-next-line no-unreachable
    // fields.forEach(v => this.form.getFieldDecorator(v))
    // 当 model 发生改变时，为表单设置值
    this.$watch('model', () => {
      // 设置表单的初始值
      // this.model && this.form.setFieldsValue(pick(this.model, fields))
      this.saveData.title = ''
      this.saveData.id = 0
      this.saveData.message_id = ''
      this.saveData.send_server_id = ''
      this.saveData.receive = ''
      this.saveData.params = ''
      if (!this.model) {
        this.$refs.messageSelectPage.remove()
        this.$refs.serverSelectPage.remove()
      }
      console.log('model', this.model)
      if (this.model && this.model.id > 0) {
        this.saveData.title = this.model.title
        this.saveData.id = this.model.id
        // this.saveData.message_id = String(this.model.message_id)
        this.saveData.send_server_id = String(this.model.send_server_id)
        this.saveData.params = String(this.model.params)
        console.log('saveData', this.saveData)
      }
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
</style>
