<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true">
      <el-form-item label="文章名称" prop="article_title">
        <el-input v-model="queryParams.article_title" placeholder="文章名称" />
      </el-form-item>
      <el-form-item label="文章标签" prop="tag_id">
        <el-select v-model="queryParams.tag_id" placeholder="文章标签">
          <el-option :key="-1" :value="-1" label="全部" />
          <el-option
            v-for="(item, index) in tags"
            :key="index"
            :label="item.tag_title"
            :value="item.tag_id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="文章状态" prop="article_status">
        <el-select v-model="queryParams.article_status">
          <el-option :key="-1" :value="-1" label="全部" />
          <el-option :key="1" :value="1" label="已发布" />
          <el-option :key="0" :value="0" label="草稿" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button v-permission="['blog:article:query']" type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permission="['blog:article:del']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除
        </el-button>
        <el-button
          v-permission="['blog:article:add']"
          size="mini"
          type="primary"
          icon="el-icon-plus"
          @click="handleAdd"
        >新增
        </el-button>
      </el-col>
    </el-row>
    <el-table v-loading="loading" :data="articleList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="编号" prop="article_id" width="100" />
      <el-table-column label="文章名称" prop="article_title" />
      <el-table-column label="文章标签" align="center">
        <template slot-scope="scope">
          <el-tag v-for="tag in scope.row.tags" :key="tag.tag_id">{{ tag.tag_title }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="文章状态" prop="article_status" align="center" width="120">
        <template slot-scope="scope">
          <el-tag
            :type="scope.row.article_status ? 'success' : 'info'"
            disable-transitions
          >{{ scope.row.article_status ? '已发布' : '草稿' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="阅读量" prop="article_read" />
      <el-table-column label="修改时间" prop="update_time" width="150">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.update_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" prop="create_time" width="150">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.create_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="150">
        <template slot-scope="scope">
          <el-button
            v-permission="['blog:article:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改
          </el-button>
          <el-button
            v-permission="['blog:article:del']"
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
          >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <pagination
      :total="total"
      :page="queryParams.page"
      :limit="queryParams.limit"
      @pagination="handlePageChange"
    />
  </div>
</template>

<script>
import { listArticle, delArticle } from '@/api/blog/article'
import { listTag } from '@/api/blog/tag'
import Pagination from '@/components/Pagination'

export default {
  name: 'BlogArticle',
  components: { Pagination },
  data() {
    return {
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 查询参数
      queryParams: {
        page: 1,
        limit: 20,
        article_title: '',
        tag_id: -1,
        article_status: -1
      },
      // 总数
      total: 0,
      // 文章列表
      articleList: [],
      // 标签
      tags: [],
      // 表单
      form: {
        article_title: '',
        article_status: 1,
        article_content: ''
      },
      // 表单校验
      rules: {
        article_title: [
          { required: true, message: '文章名称不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getList()
    this.getTags()
  },
  methods: {
    getList() {
      listArticle(this.queryParams).then(response => {
        this.articleList = response.data.articles
        this.total = response.data.total
        this.loading = false
      })
    },
    getTags() {
      listTag({ limit: 50 }).then(response => {
        this.tags = response.data.tags
      })
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.article_id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.resetForm('queryForm')
      this.handleQuery()
    },
    /** 分页改变 */
    handlePageChange(arg) {
      this.queryParams.page = arg.page
      this.queryParams.limit = arg.limit
      this.getList()
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.$router.push({ name: 'BlogArticleAdd' })
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.$router.push({ path: '/blog/article/edit/' + row.article_id })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const article_ids = row.article_id ? [row.article_id] : this.ids
      this.$confirm('是否确认删除文章编号为"' + article_ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delArticle(article_ids)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {
      })
    }
  }
}
</script>

<style scoped>

</style>
