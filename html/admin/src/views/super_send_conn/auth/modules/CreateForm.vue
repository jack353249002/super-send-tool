<template>
  <a-modal
    title="权限设置"
    :width="640"
    :visible="visible"
    :confirmLoading="loading"
    @ok="() => { $emit('ok') }"
    @cancel="() => { $emit('cancel') }"
  >
    <a-spin :spinning="loading">
      <a-form :form="form" v-bind="formLayout">
        <!-- 检查是否有 id 并且大于0，大于0是修改。其他是新增，新增不显示主键ID -->
        <a-form-item label="role">
          <a-input v-decorator="['role', {rules: [{required: true, message: '请输入role'}]}]" />
        </a-form-item>
        <a-form-item label="path">
          <a-input v-decorator="['path', {rules: [{required: true, message: '请输入path'}]}]" />
        </a-form-item>
        <a-form-item label="action">
          <a-input v-decorator="['action', {rules: [{required: true, message: '请输入action'}]}]" />
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>
import pick from 'lodash.pick'

// 表单字段
const fields = ['role', 'path', 'action']

export default {
  props: {
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
      form: this.$form.createForm(this)
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
