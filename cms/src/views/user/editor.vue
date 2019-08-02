<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="ユーザー名前">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="ユーザーメール">
        <el-input v-model="form.email" />
      </el-form-item>
      <el-form-item label="紹介">
        <el-input v-model="form.introduce" />
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
import { addUser } from '@/api/user'

export default {
  data() {
    return {
      form: {
        name: '',
        email: '',
        introduce: '',
        password: ''
      }
    }
  },
  methods: {
    onSubmit() {
      if (this.form.subpassword === this.form.password) {
        addUser(this.form).then(response => {
          this.$message('submit!')
        }).catch(error => {
          this.$message(error)
        })
      } else {
        this.$message({
          message: '入力エラー',
          type: 'warning'
        })
      }
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
