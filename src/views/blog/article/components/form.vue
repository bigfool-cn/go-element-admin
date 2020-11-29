<template>
  <div class="app-container">
    <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
      <el-form-item label="文章标题" prop="article_title">
        <el-input v-model="formData.article_title" placeholder="请输入文章标题" :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="文章状态" prop="article_status">
        <el-select v-model="formData.article_status" placeholder="请选择文章状态" :style="{width: '100%'}">
          <el-option label="发布" :value="1" />
          <el-option label="草稿" :value="0" />
        </el-select>
      </el-form-item>
      <el-form-item label="文章标签" prop="tag_ids">
        <el-select v-model="formData.tag_ids" placeholder="请选择文章标签" multiple :style="{width: '100%'}" @change="$forceUpdate()">
          <el-option
            v-for="(item, index) in tag_idsOptions"
            :key="index"
            :label="item.tag_title"
            :value="item.tag_id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="文章描述" prop="article_desc">
        <el-input
          v-model="formData.article_desc"
          type="textarea"
          placeholder="请输入文章描述"
          resize="none"
          rows="4"
          maxlength="300"
          :style="{width: '100%'}"
        />
      </el-form-item>
      <el-form-item label="文章内容" prop="article_content">
        <tinymce v-model="formData.article_content" placeholder="请输入文章内容" :height="500" />
      </el-form-item>
      <el-form-item size="large">
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="resetArticleForm">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import Tinymce from '@/components/Tinymce/index'
import { listTag } from '@/api/blog/tag'
import { addArticle, updateArticle } from '@/api/blog/article'
export default {
  name: 'BlogArticelForm',
  components: { Tinymce },
  props: {
    action: {
      type: String,
      required: true
    },
    formData: {
      type: Object,
      default: () => {
        return {
          article_title: undefined,
          tag_ids: [],
          article_status: 1,
          article_desc: undefined,
          article_content: ''
        }
      }
    }
  },
  data() {
    return {
      rules: {
        article_title: [{
          required: true,
          message: '请输入文章标题',
          trigger: 'blur'
        }],
        tag_ids: [{
          required: true,
          type: 'array',
          message: '请至少选择一个标签',
          trigger: 'change'
        }],
        article_content: [{
          required: true,
          message: '请输入文章内容',
          trigger: 'blur'
        }]
      },
      tag_idsOptions: []
    }
  },
  created() {
    listTag({ limit: 50 }).then(res => {
      this.tag_idsOptions = res.data.tags
    })
  },
  methods: {
    submitForm() {
      this.$refs['elForm'].validate(valid => {
        if (valid) {
          if (this.action === 'add') {
            addArticle(this.formData).then(res => {
              this.msgSuccess(res.message)
              this.resetArticleForm()
            })
          } else if (this.action === 'edit') {
            updateArticle(this.formData.article_id, this.formData).then(res => {
              this.msgSuccess(res.message)
            })
          } else {
            this.msgError('无效操作')
          }
        }
      })
    },
    resetArticleForm() {
      this.resetForm('elForm')
    }
  }
}
</script>

<style scoped>

</style>
