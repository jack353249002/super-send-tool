<template>
  <a-modal
    title="设置smtp"
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
        <!-- 现有的表单项 -->
      </a-form>
      <!-- 进度日志区域 -->
      <a-list item-layout="horizontal" :data-source="logItems">
        <a-list-item slot="renderItem" slot-scope="item">
          <a-list-item-meta>
            <span slot="title">{{ item.message }}</span>
          </a-list-item-meta>
        </a-list-item>
      </a-list>
    </a-spin>
  </a-modal>
</template>

<script>

import { reloadSmtpServerCon, RequestConFactory } from '@/api/super_send'

export default {
  components: {
  },
  watch: {
  },
  methods: {
    ReloadStart (token) {
      this.logItems = []
      if (this.selectSendInfo !== undefined && this.selectSendInfo.token !== '' && this.selectSendInfo.id !== 0) {
        const con = RequestConFactory(this.selectSendInfo)
        reloadSmtpServerCon(con, token)
          .then(res => {
            if (res.status === 200) {
              if (res.result !== null) {
                if (token === '') {
                  this.ReloadStart(res.result)
                } else {
                  for (let i = 0; i < res.result.length; i++) {
                    this.logItems.push({ message: res.result[i].message })
                  }
                  setTimeout(() => {
                    this.ReloadStart(token)
                  }, 1000)
                }
                // this.messageListInfo = { list: res.result.data, total: res.result.totalPage, page: res.result.pageNo, size: res.result.pageSize }
              } else {
                this.logItems.push({ message: '已结束' })
              }
            } else if (res.status === 401) {
              this.logItems.push({ message: res.message })
            } else {
              this.logItems.push({ message: '已结束' })
            }
          })
      } else {
        this.logItems.push({ message: '错误' })
      }
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
  data () {
    return {
      logItems: [
      ]
    }
  }
}
</script>
