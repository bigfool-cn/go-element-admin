<template>
  <blog-articel-form action="edit" :form-data="article" />
</template>

<script>
import BlogArticelForm from './components/form'
import { getArticle } from '@/api/blog/article'
export default {
  name: 'BlogArticleEdit',
  components: { BlogArticelForm },
  data() {
    return {
      article: {}
    }
  },
  created() {
    const article_id = this.$route.params.article_id
    getArticle(article_id).then(res => {
      const { article_id, article_title, article_status, article_desc, article_content } = res.data
      this.article = {
        article_id,
        article_title,
        article_status,
        article_desc,
        article_content
      }
      this.article.tag_ids = res.data.tags.map(tag => {
        return tag.tag_id
      })
    })
  }
}
</script>

<style scoped>

</style>
