<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true">
      <el-form-item label="标签名称" prop="tag_title">
        <el-input v-model="queryParams.tag_title" placeholder="标签名称" />
      </el-form-item>
      <el-form-item label="标签状态" prop="tag_status">
        <el-select v-model="queryParams.tag_status">
          <el-option :key="-1" :value="-1" label="全部" />
          <el-option :key="1" :value="1" label="启用" />
          <el-option :key="0" :value="0" label="停用" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button v-permission="['blog:tag:query']" type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permission="['blog:tag:del']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除
        </el-button>
        <el-button
          v-permission="['blog:tag:add']"
          size="mini"
          type="primary"
          icon="el-icon-plus"
          @click="handleAdd"
        >新增
        </el-button>
      </el-col>
    </el-row>
    <el-table v-loading="loading" :data="tagList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="编号" prop="tag_id" width="100" />
      <el-table-column label="标签名称" prop="tag_title" />
      <el-table-column label="标签状态" align="center" width="120">
        <template slot-scope="scope">
          <el-tag
            :type="scope.row.tag_status ? 'success' : 'info'"
            disable-transitions
          >{{ scope.row.tag_status ? '启用' : '停用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="修改时间" prop="update_time" width="150" />
      <el-table-column label="创建时间" prop="create_time" width="150" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="150">
        <template slot-scope="scope">
          <el-button
            v-permission="['blog:tag:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改
          </el-button>
          <el-button
            v-permission="['blog:tag:del']"
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
    <!-- 添加或修改标签对话框 -->
    <el-dialog v-el-drag-dialog :title="title" :visible.sync="open" width="400px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="标签名称" prop="tag_title">
          <el-input v-model="form.tag_title" placeholder="请输入标签名称" />
        </el-form-item>
        <el-form-item label="标签状态" prop="tag_status">
          <el-radio-group v-model="form.tag_status">
            <el-radio :key="1" :label="1">启用</el-radio>
            <el-radio :key="0" :label="0">停用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listTag, addTag, updateTag, delTag } from '@/api/blog/tag'
import Pagination from '@/components/Pagination'
import ElDragDialog from '@/directive/el-drag-dialog'

export default {
  name: 'BlogTag',
  directives: { ElDragDialog },
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
        tag_title: '',
        tag_status: -1
      },
      // 总数
      total: 0,
      // 标签列表
      tagList: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 表单
      form: {
        tag_title: '',
        tag_status: 1
      },
      // 表单校验
      rules: {
        tag_title: [
          { required: true, message: '标签名称不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      listTag(this.queryParams).then(response => {
        this.tagList = response.data.tags
        this.total = response.data.total
        this.loading = false
      })
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.tag_id)
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
    /** 表单重置 */
    reset() {
      this.form = {
        tag_title: '',
        tag_status: 1
      }
      this.resetForm('form')
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加标签'
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      this.form = {
        tag_id: row.tag_id,
        tag_title: row.tag_title,
        tag_status: row.tag_status
      }
      this.open = true
      this.title = '修改标签'
    },
    /** 取消按钮 */
    cancel() {
      this.open = false
      this.reset()
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.tag_id !== undefined) {
            updateTag(this.form.tag_id, this.form).then(response => {
              this.msgSuccess(response.message)
              this.open = false
              this.getList()
            })
          } else {
            addTag(this.form).then(response => {
              this.msgSuccess(response.message)
              this.open = false
              this.getList()
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const tag_ids = row.tag_id ? [row.tag_id] : this.ids
      this.$confirm('是否确认删除标签编号为"' + tag_ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delTag(tag_ids)
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
