<template>
  <a-modal
    title="user alive 账号"
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
        <a-form-item label="用户名">
          <a-input  v-decorator="['username', {rules: [{required: true, message: '用户名'}]}]" />
        </a-form-item>
        <a-form-item label="密码">
          <a-input-password v-decorator="['password', {rules: []}]"/>
        </a-form-item>
        <a-form-item label="签到时间">
          <a-date-picker
            format="YYYY-MM-DD HH:mm:ss"
            show-time
            v-decorator="['day_login_first_time', {rules: []}]"
          />
        </a-form-item>
        <a-form-item label="发送器id">
          <a-input v-decorator="['send_id', {rules: []}]" />
        </a-form-item>
        <a-form-item label="超级发送连接id">
          <a-input v-decorator="['super_send_conn_info_id', {rules: []}]" />
        </a-form-item>
        <a-form-item label="超时时长(秒)">
          <a-input v-decorator="['send_email_action_timeout', {rules: []}]" />
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>
import pick from 'lodash.pick'

// 表单字段
const fields = ['id', 'username', 'password', 'day_login_first_time', 'send_id', 'super_send_conn_info_id', 'send_email_action_timeout']

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
  methods: {
    resetFields () {
      this.form.resetFields()
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
