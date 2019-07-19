<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="ユーザー名前">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="ユーザーメール">
        <el-input v-model="form.email" />
      </el-form-item>
      <el-form-item label="タグ">
        <el-select v-model="form.role" placeholder="please select user role">
          <el-option label="Zone one" value="Role1" />
          <el-option label="Zone two" value="Role2" />
        </el-select>
      </el-form-item>
      <el-form-item label="パスワード">
        <el-input v-model="form.password" />
      </el-form-item>
      <el-form-item label="パスワード確認">
        <el-input v-model="form.subpassword" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">作成</el-button>
        <el-button @click="onCancel">リセット</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>

export default {
  data() {
    return {
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      }
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
.tag-title{
  margin-bottom: 5px;
}
</style>
