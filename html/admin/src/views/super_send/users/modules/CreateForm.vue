<template>
  <a-modal
    title="super_send账号"
    :width="640"
    :visible="visible"
    :confirmLoading="loading"
    @ok="() => { $emit('ok') }"
    @cancel="() => { $emit('cancel') }"
  >
    <a-spin :spinning="loading">
      <a-form :form="form" v-bind="formLayout">
        <!-- 检查是否有 id 并且大于0，大于0是修改。其他是新增，新增不显示主键ID -->
        <a-form-item v-show="model && model.id > 0" label="主键ID">
          <a-input v-decorator="['id', { initialValue: 0 }]" disabled />
        </a-form-item>
        <a-form-item label="是否开启ssl">
          <a-radio-group v-decorator="['is_ssl', { initialValue: '0', rules: [{ required: true, message: '请选择状态' }] }]">
            <a-radio value="0">关闭</a-radio>
            <a-radio value="1">开启</a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item label="地址">
          <a-input v-decorator="['address', {rules: [{required: true, message: '请输入地址'}]}]" />
        </a-form-item>
        <a-form-item label="用户名">
          <a-input v-decorator="['username', {rules: [{required: true, message: '请输入用户名'}]}]" />
        </a-form-item>
        <a-form-item label="密码">
          <a-input v-decorator="['password', {rules: [{required: true, message: '请输入密码'}]}]" />
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>
import pick from 'lodash.pick'

// 表单字段
const fields = ['address', 'is_ssl', 'username', 'password', 'id']

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
