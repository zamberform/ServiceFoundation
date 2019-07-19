<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="タイトル">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="タグ">
        <el-select v-model="form.tag" placeholder="please select your tag">
          <el-option label="Zone one" value="Tag1" />
          <el-option label="Zone two" value="Tag2" />
        </el-select>
      </el-form-item>
      <el-form-item label="配布時間設定">
        <el-col :span="11">
          <el-date-picker v-model="form.date1" type="date" placeholder="Pick a date" style="width: 100%;" />
        </el-col>
        <el-col :span="2" class="line">-</el-col>
        <el-col :span="11">
          <el-time-picker v-model="form.date2" type="fixed-time" placeholder="Pick a time" style="width: 100%;" />
        </el-col>
      </el-form-item>
      <el-form-item label="コメントフラグ">
        <el-switch v-model="form.delivery" />
      </el-form-item>
      <el-form-item label="内容">
        <div class="editor-container">
          <markdown-editor v-model="form.content" height="300px" />
        </div>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">作成</el-button>
        <el-button @click="onCancel">リセット</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor'

export default {
  components: { MarkdownEditor },
  data() {
    return {
      form: {
        name: '',
        tag: '',
        date1: '',
        date2: '',
        delivery: false,
        content: ''
      }
    }
  },
  created() {
    // const route_id = this.$route.params.id
    this.form.name = 'test'
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
.editor-container{
  margin-bottom: 30px;
}
.tag-title{
  margin-bottom: 5px;
}
</style>
