<template>
  <div class="tinymce-box">
    <editor
      v-model="content"
      :init="init"
      :disabled="disabled"
      @onClick="onClick"
    />
  </div>
</template>

<script>
import tinymce from 'tinymce/tinymce' // tinymce默认hidden，不引入不显示
import Editor from '@tinymce/tinymce-vue'

import 'tinymce/themes/silver'
import 'tinymce/icons/default'

// 编辑器插件plugins
// 更多插件参考：https://www.tiny.cloud/docs/plugins/
import 'tinymce/plugins/image' // 插入上传图片插件
import 'tinymce/plugins/media' // 插入视频插件
import 'tinymce/plugins/table' // 插入表格插件
import 'tinymce/plugins/lists' // 列表插件
import 'tinymce/plugins/wordcount' // 字数统计插件
import 'tinymce/plugins/autosave' // 自动保存
import 'tinymce/plugins/fullscreen' // 全屏
import 'tinymce/plugins/codesample' // 插入代码
import 'tinymce/plugins/preview' // 预览
import 'tinymce/plugins/searchreplace' // 查找替换
import 'tinymce/plugins/hr' // 水平分割线

export default {
  name: 'Tinymce',
  components: {
    Editor
  },
  props: {
    value: {
      type: String,
      default: ''
    },
    height: {
      type: Number,
      default: 300
    },
    disabled: {
      type: Boolean,
      default: false
    },
    plugins: {
      type: [String, Array],
      default: 'lists image media table wordcount autosave fullscreen codesample preview searchreplace hr'
    },
    toolbar: {
      type: [String, Array],
      default: 'searchreplace | formatselect | bold italic forecolor backcolor removeformat codesample blockquote hr | alignleft aligncenter alignright alignjustify outdent indent | bullist numlist outdent indent | lists image media table | preview restoredraft fullscreen'
    }
  },
  data() {
    return {
      init: {
        language_url: '/tinymce/langs/zh_CN.js',
        language: 'zh_CN',
        skin_url: '/tinymce/skins/ui/oxide',
        // skin_url: 'tinymce/skins/ui/oxide-dark',//暗色系
        height: this.height,
        plugins: this.plugins,
        toolbar: this.toolbar,
        branding: false,
        menubar: false,
        // 此处为图片上传处理函数，这个直接用了base64的图片形式上传图片，
        // 如需ajax上传可参考https://www.tiny.cloud/docs/configure/file-image-upload/#images_upload_handler
        images_upload_handler: (blobInfo, success, failure) => {
          const img = 'data:image/jpeg;base64,' + blobInfo.base64()
          success(img)
        }
      },
      content: this.value
    }
  },
  watch: {
    value(newValue) {
      this.content = newValue
    },
    content(newValue) {
      this.$emit('input', newValue)
    }
  },
  mounted() {
    tinymce.init({})
  },
  methods: {
    // 添加相关的事件，可用的事件参照文档=> https://github.com/tinymce/tinymce-vue => All available events
    // 需要什么事件可以自己增加
    onClick(e) {
      this.$emit('onClick', e, tinymce)
    },
    // 可以添加一些自己的自定义事件，如清空内容
    clear() {
      this.content = ''
    }
  }
}

</script>
<style scoped>

</style>
