<template>
  <div class="app-container">
    <el-form ref="form">
      <el-form-item>
        <el-button type="primary" @click="handleCreate">追加</el-button>
      </el-form-item>
      <el-form-item>
        <el-table
          v-loading="listLoading"
          :data="list"
          element-loading-text="Loading"
          border
          fit
          highlight-current-row
        >
          <el-table-column align="center" label="ID" width="95">
            <template slot-scope="scope">
              {{ scope.$index }}
            </template>
          </el-table-column>
          <el-table-column label="名前">
            <template slot-scope="scope">
              {{ scope.row.name }}
            </template>
          </el-table-column>
          <el-table-column label="色" width="110" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.color }}</span>
            </template>
          </el-table-column>
          <el-table-column align="center" prop="created_at" label="Display_time" width="200">
            <template slot-scope="scope">
              <i class="el-icon-time" />
              <span>{{ scope.row.create_time }}</span>
            </template>
          </el-table-column>
          <el-table-column label="Actions" align="center" width="300" class-name="small-padding fixed-width">
            <template slot-scope="{row}">
              <el-button size="mini" type="success" @click="handleUpdate(row)">
                編集
              </el-button>
              <el-button size="mini" type="danger" @click="deleteTag(row)">
                削除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-form-item>
    </el-form>

    <el-dialog title="タグ編集" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px" style="width: 120px; margin-left:50px;">
        <el-form-item label="タグ名前">
          <el-input v-model="temp.name" />
        </el-form-item>
        <el-form-item label="カーラー">
          <el-color-picker v-model="temp.color" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="onCancel">
          リセット
        </el-button>
        <el-button type="primary" @click="onSubmit">
          作成
        </el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="dialogConfirmVisible" title="削除確認">
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogConfirmVisible = false">Confirm</el-button>
        <el-button type="block" @click="dialogConfirmVisible = false">Cancel</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { getList } from '@/api/tag'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true,
      dialogFormVisible: false,
      dialogConfirmVisible: false,
      temp : {
        id: undefined,
        name: '',
        color: '#000000',
        timestamp: new Date()
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getList().then(response => {
        this.list = response.items
        this.listLoading = false
      })
    },
    onSubmit() {
      this.dialogFormVisible = true
      this.$message('submit!')
    },
    onCancel() {
      this.dialogFormVisible = false
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        name: '',
        color: '#000000',
        timestamp: new Date()
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row)
      this.temp.timestamp = new Date(this.temp.timestamp)
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    deleteTag(row) {
      this.dialogConfirmVisible = true
    }
  }
}
</script>
