<template>
  <div class="app-container">
    <el-form :inline="true">
      <el-form-item label="接口名称">
        <el-input
          v-model="queryParams.name"
          placeholder="请输入接口名称"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="接口路径">
        <el-input
          v-model="queryParams.path"
          placeholder="请输入接口路径"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item>
        <el-button v-permission="['system:path:query']" type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button
          v-permission="['system:path:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增
        </el-button>
      </el-form-item>
    </el-form>

    <el-table
      v-loading="loading"
      :data="pathList"
      row-key="path_id"
      :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
    >
      <el-table-column prop="name" label="接口名称" :show-overflow-tooltip="true" />
      <el-table-column prop="path" label="接口路径" :show-overflow-tooltip="true">
        <template slot-scope="scope">
          <span>{{ scope.row.path ? scope.row.path : '--' }}</span>
        </template>
      </el-table-column>
      <el-table-column label="接口类型" align="center" prop="type" width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.type === 'M' ? '目录' : '接口' }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="method" label="接口方法" align="center" width="100">
        <template slot-scope="scope">
          <span v-if="scope.row.type === 'M'">--</span>
          <el-tag
            v-if="scope.row.type === 'J'"
            :type="methodTypes[scope.row.method]"
            disable-transitions
          >{{ scope.row.method }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="修改时间" align="center" prop="createTime" width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.update_time ? parseTime(scope.row.update_time) : '--' }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center" prop="createTime" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.create_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permission="['system:path:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改
          </el-button>
          <el-button
            v-permission="['system:path:add']"
            size="mini"
            type="text"
            icon="el-icon-plus"
            @click="handleAdd(scope.row)"
          >新增
          </el-button>
          <el-button
            v-permission="['system:path:del']"
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
          >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 添加或修改对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="600px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-row>
          <el-col :span="24">
            <el-form-item label="上级目录">
              <treeselect
                v-model="form.parent_id"
                :options="pathOptions"
                :normalizer="normalizer"
                :show-count="true"
                placeholder="选择上级目录"
              />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="接口类型" prop="type">
              <el-radio-group v-model="form.type">
                <el-radio label="M">目录</el-radio>
                <el-radio label="J">接口</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="接口名称" prop="title">
              <el-input v-model="form.name" placeholder="请输入接口名称" />
            </el-form-item>
          </el-col>
          <el-col v-if="form.type === 'J'" :span="24">
            <el-form-item label="接口路径" prop="title">
              <el-input v-model="form.path" placeholder="请输入接口路径" />
            </el-form-item>
          </el-col>
          <el-col v-if="form.type === 'J'" :span="24">
            <el-form-item label="接口方法" prop="type">
              <el-radio-group v-model="form.method">
                <el-radio label="GET">GET</el-radio>
                <el-radio label="POST">POST</el-radio>
                <el-radio label="PUT">PUT</el-radio>
                <el-radio label="DELETE">DELETE</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listPath, getPath, delPath, addPath, updatePath } from '@/api/system/path'
import Treeselect from '@riophae/vue-treeselect'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'

export default {
  name: 'Interface',
  components: { Treeselect },
  data() {
    return {
      // 遮罩层
      loading: true,
      // 接口表格树数据
      pathList: [],
      // 树选项
      pathOptions: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 查询参数
      queryParams: {
        name: undefined,
        path: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        name: [
          { required: true, message: '接口名称不能为空', trigger: 'blur' }
        ]
      },
      methodTypes: {
        'GET': '',
        'POST': 'success',
        'PUT': 'warning',
        'DELETE': 'danger'
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询接口列表 */
    getList() {
      this.loading = true
      if (this.queryParams.name === '') {
        this.queryParams.name = undefined
      }
      listPath(this.queryParams).then(response => {
        this.pathList = response.data
        this.loading = false
      })
    },
    /** 转换数据结构 */
    normalizer(node) {
      if (!node.children || !node.children.length) {
        delete node.children
        return {
          id: node.path_id,
          label: node.name
        }
      } else {
        return {
          id: node.path_id,
          label: node.name,
          children: node.children
        }
      }
    },
    /** 查询下拉树结构 */
    getTreeselect() {
      listPath().then(response => {
        this.pathOptions = []
        const path = { path_id: 0, name: '主类目', children: [] }
        path.children = response.data
        this.pathOptions.push(path)
      })
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        path_id: undefined,
        parent_id: 0,
        type: 'J',
        name: undefined,
        path: undefined,
        method: undefined
      }
      this.resetForm('form')
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.getList()
    },
    /** 新增按钮操作 */
    handleAdd(row) {
      this.reset()
      this.getTreeselect()
      if (row != null) {
        this.form.parent_id = row.path_id
      }
      this.open = true
      this.title = '添加接口'
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      this.getTreeselect()
      getPath(row.path_id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改接口'
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.type === 'J') {
            if (!this.form.path) {
              this.msgError('请输入接口路径')
              return false
            }
            if (!this.form.method) {
              this.msgError('请选择接口方法')
              return false
            }
          }
          if (this.form.path_id !== undefined) {
            updatePath(this.form.path_id, this.form).then(response => {
              this.msgSuccess(response.message)
              this.open = false
              this.getList()
            })
          } else {
            addPath(this.form).then(response => {
              this.msgSuccess('新增成功')
              this.open = false
              this.getList()
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      this.$confirm('是否确认删除名称为"' + row.name + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delPath([row.path_id])
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
