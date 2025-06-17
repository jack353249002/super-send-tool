<template>
  <a-modal
    title="消息设置"
    width="100%"
    wrap-class-name="full-modal"
    :visible="visible"
    :confirmLoading="loading"
    class="draggable-modal"
    ref="headerUploadRef"
    @ok="() => { $emit('ok') }"
    @cancel="() => { $emit('cancel') }"
    :bodyStyle="{ maxHeight: '600px', overflow: 'auto' }"
  >
    <a-form :form="form" v-bind="formLayout">
      <!-- 检查是否有 id 并且大于0，大于0是修改。其他是新增，新增不显示主键ID -->
      <a-form-item v-show="model && model.id > 0" label="主键ID">
        <a-input v-model="saveData.id" disabled />
      </a-form-item>
      <a-form-item label="标题">
        <a-input v-model="saveData.title" />
      </a-form-item>
      <a-form-item label="内容类型">
        <a-input v-model="saveData.content_type" />
      </a-form-item>
      <a-form-item label="内容" style="height: 400px;">
        <quill-editor ref="myQuillEditor" v-model="saveData.body" style="height: 300px;" />
      </a-form-item>
      <a-form-item label="附件内容">
        <a-input v-model="saveData.attach" @change="handleAttachChange" />
      </a-form-item>
      <a-form-item label="文件上传">
        <a-upload
          name="file"
          :file-list="fileList"
          list-type="picture"
          multiple
          :custom-request="uploadFileRequest"
          :remove="(file)=>{handleRemove(file, 'header')}"
        >
          <a-button>
            <a-icon type="upload" />
          </a-button>
        </a-upload>
      </a-form-item>
    </a-form>

    <a-spin :spinning="loading"></a-spin>
  </a-modal>
</template>

<script>
import VueDraggableResizable from 'vue-draggable-resizable'
import 'vue-draggable-resizable/dist/VueDraggableResizable.css'
import { quillEditor } from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import axios from 'axios'
// 表单字段
const fields = ['id', 'body', 'title', 'attach', 'content_type']
export default {
  components: {
    VueDraggableResizable,
    quillEditor
  },
  watch: {
    fileList: {
      handler (newVal) {
        if (!this.isEditingAttach) {
          this.saveData.attach = newVal.map(file => file.path).join(',')
        }
      },
      deep: true // 深度监听
    },
    'saveData.attach': {
      handler (newVal) {
        if (this.isEditingAttach) {
          const attrArray = newVal ? newVal.split(',') : []
          const newPathsSet = new Set(attrArray)
          const oldPathsSet = new Set(this.fileList.map(file => file.path))

          const addedPaths = attrArray.filter(path => !oldPathsSet.has(path))
          const removedPaths = this.fileList.filter(file => !newPathsSet.has(file.path)).map(file => file.path)

          addedPaths.forEach(path => {
            this.attrCount++
            this.fileList.push({
              uid: this.attrCount.toString(),
              name: path.split('/').pop(),
              status: 'done',
              response: 'ok',
              url: path,
              path: path
            })
          })

          removedPaths.forEach(path => {
            const index = this.fileList.findIndex(file => file.path === path)
            if (index !== -1) {
              this.fileList.splice(index, 1)
            }
          })
        }
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
    // 当 model 发生改变时，为表单设置值
    this.$watch('model', () => {
      // 设置表单的初始值
      // this.model && this.form.setFieldsValue(pick(this.model, fields))
      console.log('model', this.model)
      if (this.model && this.model.id > 0) {
        this.saveData = this.model
      } else {
        this.saveData = {}
      }
    })
  },
  methods: {
    resetFields () {
      this.saveData = {}
    },
    getFormData () {
      // 获取编辑器内容
      const editorContent = this.$refs.myQuillEditor.quill.root.innerHTML
      // 返回表单数据
      return {
        id: Number(this.saveData.id),
        title: this.saveData.title,
        attach: this.saveData.attach,
        body: editorContent,
        content_type: this.saveData.content_type
      }
    },
    handleRemove (file) {
      // eslint-disable-next-line prefer-const
      this.fileList.forEach((item, index) => {
        if (item.uid === file.uid) {
          this.fileList.splice(index, 1)
        }
      })
    },
    handleChange (info) {
      console.log(info)
      // this.fileList = info.fileList
    },
    handleAttachChange () {
      this.isEditingAttach = true
      this.$nextTick(() => {
        this.isEditingAttach = false
      })
    },
    uploadFileRequest (options) {
      const { onSuccess, onError, file } = options
      const formData = new FormData()
      formData.append('file', file)

      // eslint-disable-next-line vue/no-async-in-computed-properties
      axios.post(process.env.VUE_APP_API_BASE_URL + '/super_send/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
          'Token': this.selectSendInfo.token, // 添加你的自定义头信息
          'Super_send_id': this.selectSendInfo.id
        }
      })
        .then(response => {
          console.log(response)
          if (response.data.status === 200) {
            this.attrCount++
            this.fileList.push({
              uid: this.attrCount.toString(),
              name: file.name,
              status: 'done',
              response: 'ok', // custom error message to show
              url: response.data.result.url,
              path: response.data.result.path
            })
            console.log(this.fileList)
            onSuccess(response.data.result, file)
          } else {
            onError('上传失败')
          }
        })
        .catch(error => {
          onError(error)
        })
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
      saveData: { 'content_type': 'text/html' },
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
      isEditingAttach: false
    }
  },
  created () {
    console.log('custom modal created')
    // 防止表单未注册
    fields.forEach(v => this.form.getFieldDecorator(v))
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
