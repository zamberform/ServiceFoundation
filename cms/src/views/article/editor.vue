<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="タイトル">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="タグ">
        <el-select v-model="form.tag_id" placeholder="please select your tag">
          <el-option v-for="item in tags" :key="item.id" :label="item.name" :value="item.id" />
        </el-select>
      </el-form-item>
      <!-- <el-form-item label="配布時間設定">
        <el-col :span="11">
          <el-date-picker v-model="form.date1" type="date" placeholder="Pick a date" style="width: 100%;" />
        </el-col>
        <el-col :span="2" class="line">-</el-col>
        <el-col :span="11">
          <el-time-picker v-model="form.date2" type="fixed-time" placeholder="Pick a time" style="width: 100%;" />
        </el-col>
      </el-form-item> -->
      <el-form-item label="コメントフラグ">
        <el-switch v-model="form.comment_flg" />
      </el-form-item>
      <el-form-item label="内容">
        <div class="editor-container">
          <markdown-editor v-model="form.content_desc" height="300px" />
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
import { getList } from '@/api/tag'
import { updateArticle, addArticle, findArticle } from '@/api/article'

export default {
  components: { MarkdownEditor },
  data() {
    return {
      tags: null,
      isEditMode: false,
      form: {
        title: '',
        tag_id: null,
        comment_flg: false,
        content_desc: ''
      }
    }
  },
  created() {
    getList().then(response => {
      this.tags = response.tags
    })
    const route_id = this.$route.params.id
    if (route_id) {
      this.isEditMode = true
      findArticle(route_id).then(response => {
        this.form = response.article
      })
    } else {
      this.isEditMode = false
    }
    // this.form.name = 'test'
  },
  methods: {
    onSubmit() {
      if (this.isEditMode) {
        updateArticle(this.form).then(response => {
          this.$message('submit!')
        })
      } else {
        addArticle(this.form).then(response => {
          this.$message('submit!')
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
.editor-container{
  margin-bottom: 30px;
}
.tag-title{
  margin-bottom: 5px;
}
</style>
