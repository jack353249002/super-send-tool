<template>
  <a-modal
    title="设置角色"
    width="100%"
    wrap-class-name="full-modal"
    :visible="visible"
    :confirmLoading="loading"
    :footer="false"
    class="draggable-modal"
    ref="headerUploadRef"
    @ok="() => { $emit('ok') }"
    @cancel="() => { $emit('cancel') }"
  >
    <a-spin :spinning="loading">
      <a-form :form="form" v-bind="formLayout">
        <!-- 检查是否有 id 并且大于0，大于0是修改。其他是新增，新增不显示主键ID -->
        <a-form-item v-show="model && model.id > 0" label="主键ID">
          <a-input v-model="model.id" disabled />
        </a-form-item>
        <a-form-item v-show="model && model.id > 0" label="用户名">
          <a-input v-model="model.username" disabled />
        </a-form-item>
        <a-form-item label="选择角色">
          <v-selectpage
            placeholder="请选择一个选项"
            field="name"
            value-field="id"
            :data="roleList"
            :params="selectSendInfo"
            v-model="rolestr"
            :multiple="true"
            :pagination="false"
            ref="roleSelectPage"
          >
          </v-selectpage>
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
import {
  addServerRolesFroUserCon, delServerUserRoleCon,
  getMessageList, getServerRolesCon,
  getServerUserInfoCon,
  RequestConFactory,
  superSendApi
} from '@/api/super_send'
// const fields = ['message_id', 'title', 'send_model', 'send_server_id', 'dispatch_model', 'receive']// 表单字段
export default {
  components: {
  },
  watch: {
    selectSendInfo: {
      handler (newVal, oldVal) {
        console.log('newVal', newVal)
        if (newVal != null) {
          this.initLoad(newVal.id)
        }
      },
      deep: true // 深度监听
    },
    rolestr: {
      handler (newVal, oldVal) {
        this.rolestrChange(oldVal, newVal)
      },
      deep: true
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
    setRoleStr (newVal) {
      var oldVal = this.rolestr
      if (oldVal === newVal) {
        this.rolestrChange(oldVal, newVal)
      } else {
        this.rolestr = newVal
      }
    },
    rolestrChange (oldVal, newVal) {
      // 将 rolestr 拆分为数组
      console.log('rolestr-in', newVal)
      console.log('rolestr-in2', typeof oldVal)
      console.log('rolestr-in3', this.needRequest)
      const newRoleArray = newVal ? newVal.split(',') : []
      const oldRoleArray = oldVal ? oldVal.split(',') : []
      console.log('newRoleArray', newRoleArray)
      console.log('oldRoleArray', oldRoleArray)
      console.log('firstIn', this.firstIn)
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined && this.model != null && this.model.id !== 0 && this.needRequest) {
        const con = RequestConFactory(this.selectSendInfo)
        // 找出新增的元素
        const addedRoles = newRoleArray.filter(role => !oldRoleArray.includes(role))
        for (const role of addedRoles) {
          addServerRolesFroUserCon(con, { 'role': role, 'username': this.model.username })
        }
        // 找出减少的元素
        const removedRoles = oldRoleArray.filter(role => !newRoleArray.includes(role))
        for (const role of removedRoles) {
          delServerUserRoleCon(con, { 'role': role, 'username': this.model.username })
        }
      }
      this.needRequest = true
      console.log('rolestr-in4', this.needRequest)
    },
    getFormData (type) {

    },
    isNeedRequest () {
      this.needRequest = true
    },
    clearSelectPage () {
      this.needRequest = false
      this.model.id = 0
      this.setRoleStr('')
    },
    initLoad (id) {
      this.getRoles(id, true)
    },
    getUserInfo (id) {
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
        const con = RequestConFactory(this.selectSendInfo)
        getServerUserInfoCon(con, { id: id }).then(res => {
          this.firstIn = true
          if (res.status === 200) {
            if (res.result.roles != null) {
              this.needRequest = false
              this.setRoleStr(res.result.roles.join(','))
            }
            console.log('rolestr', this.rolestr)
          } else {
            this.needRequest = false
            this.setRoleStr('')
          }
        })
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
    getRoles (id, getUser) {
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0 && this.selectSendInfo.id !== undefined) {
        console.log('getRoles')
        const con = RequestConFactory(this.selectSendInfo)
        getServerRolesCon(con, {})
          .then(res => {
            console.log(res)
            if (res.status === 200) {
              if (res.result.data !== null) {
                const listTemp = []
                for (let i = 0; i < res.result.data.length; i++) {
                  listTemp.push({ id: res.result.data[i].role, name: res.result.data[i].role })
                }
                // this.total = res.result.totalCount
                this.roleList = listTemp
                if (getUser) {
                  this.getUserInfo(id)
                }
                // this.messageListInfo = { list: res.result.data, total: res.result.totalPage, page: res.result.pageNo, size: res.result.pageSize }
              } else {
                this.roleList = []
              }
            } else if (res.status === 401) {
              this.roleList = []
            } else {
              this.roleList = []
            }
          })
      } else {
        this.roleList = []
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
      needRequest: true,
      rolestr: '',
      saveData: { 'rolestr': '' },
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
      roleList: [],
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
      if (this.model && this.model.id > 0) {
        console.log('model', this.model)
          this.initLoad(this.model.id)
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
