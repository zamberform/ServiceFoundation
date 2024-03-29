<template>
  <div class="comment">
    <nav-header :active="active"></nav-header>
    <el-row type="flex" justify="center">
      <el-col :span="14" class="detail_title">
          <div>{{title}}</div>
          <div class="time">更新時間：{{updated_at}}&nbsp;&nbsp;&nbsp;&nbsp;</div>
          <div class="time">タグ：{{tag}}&nbsp;&nbsp;&nbsp;&nbsp;</div>
      </el-col>

    </el-row>
    <el-row type="flex" justify="center">
      <el-col :span="14" class="detail_content">
        <div v-show="!content_desc">データがない</div>
        <div v-html="content_desc" class="md markdown-body"></div>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center">
      <el-col :span="14">
        <h2 style="color:#3D5064;border-top:1px dashed #3D5064;padding-top:15px;margin-top:30px;">コメント：</h2>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center" class="detail_content">
      <el-col :span="14">
        <el-card class="box-card" v-show="commentList.length !== 0" v-for="(item, index) in commentList" :key="index">
          <div slot="header" class="clearfix">
            <span style="font-weight: bold;">{{item.user.name}} <el-tag type="success" v-show="author.includes(item.user.name)">が</el-tag> コメント：</span>
            <span style="float: right; padding: 3px 0;font-weight: bold;"><Time :time="item.created_at" :interval="1" /></span>
          </div>
          <div>
            {{item.comment}}
          </div>
        </el-card>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center">
      <el-col :span="15" class="detail_content" style="margin-left:-63px;">
        <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
          <el-form-item v-if="$store.getters['userName'] === null" label="名前" prop="username">
            <el-input type="text" v-model="ruleForm.username" @change="usernameChange" autocomplete="off" placeholder="ユーザー名"></el-input>
          </el-form-item>
          <el-form-item v-if="$store.getters['userName'] === null" label="メール" prop="email">
            <el-input type="text" v-model="ruleForm.email" autocomplete="off" placeholder="メールを入力してください"></el-input>
          </el-form-item>
          <el-form-item label="コメント" prop="content">
            <el-input type="textarea" :rows="8" v-model="ruleForm.content" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')">OK</el-button>
            <el-button @click="resetForm('ruleForm')">Reset</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>

var qs = require('qs');

import NavHeader from '~/components/NavHeader.vue'
import Time from '~/components/Time.vue'
import { log } from 'util';

export default {
	data() {
    var checkUsername = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('ユーザー名を入力してください'));
      } else {
        callback()
      }
    };
    var validateEmail = (rule, value, callback) => {
      const reg = /^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$/
      if (value === '') {
        callback(new Error('メールを入力してください'));
      } else {
        if (reg.test(value)) {
          callback()
        } else {
          callback(new Error('メール形式ではない'));
        }
        // callback();
      }
    };
    var validateContent = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('内容を入れてください'));
      } else {
        callback();
      }
    };
    return {
      active:'index',
      ruleForm: {
        username: this.$store.getters['userName'],
        email: '',
        content: ''
      },
      rules: {
        username: [
          { validator: checkUsername, trigger: 'change' }
        ],
        email: [
          { validator: validateEmail, trigger: 'blur' }
        ],
        content: [
          { validator: validateContent, trigger: 'change' }
        ]
      },
      commentList: [],
      count: 0,
      author: ['admin', 'editor'],
      authorStatus: false
    }
	},
	async asyncData({app,params}) {
		let json = {id:params.id}
    let result = await app.$axios.$post(`/api/article/` + params.id);
    let info = result.article;
    let {content_desc,tag,updated_at,title} = info;
		return {title,content_desc,tag,updated_at}
	},
    head() {
		return {
			title:this.title,
            meta:[
                {hid:'description',name:'description',content:`${this.des}`},
                {hid:'author',content:'zamberform'}
            ]
            }
	},
    components:{
        NavHeader,
        Time
    },
    created () {
        this.commentLists(this.$route.params.id)
    },
    methods: {
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.ruleForm.articleId = this.$route.params.id
                    if (this.$store.getters['userName'] !== null) {
                      this.ruleForm.username = this.$store.getters['userName']
                    }
                    let json = Object.assign({}, this.ruleForm)
                    this.commentsSubmit(json, formName)
                } else {
                    this.$notify({
                        title: '失敗',
                        message: '内容を入れてください',
                        type: 'error'
                    });
                    return false;
                }
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields()
        },
        async commentsSubmit (json, formName) {
            try {
                let {status} = await this.$axios.$post(`/api/comment`, json)
                if (status == 200) {
                    this.$refs[formName].resetFields()
                    this.$notify({
                        title: '成功',
                        message: 'コメント確認してください',
                        type: 'success'
                    });
                    this.commentLists(this.$route.params.id)
                } else {
                    this.$notify({
                        title: '失敗',
                        message: 'コメント失敗',
                        type: 'error'
                    });
                    return false
                }
            } catch (error) {
                // handle error
                return false
            }
        },
        async commentLists (id) {
            try {
                let {count, comments} = await this.$axios.$post(`/api/comment/list/` + id)
                this.count = count
                if (count > 0) {
                  this.commentList = comments.reverse()
                }
            } catch (error) {
                // handle error
                console.log(error)
            }
        },
        usernameChange (val) {
            this.authorStatus = this.author.includes(val)
            console.log(`status:${this.authorStatus}`)
        },
  }
}
</script>
<style lang="less">
  @import './../../assets/less/detail.less';
  .comment {
    .clearfix:before,
    .clearfix:after {
      display: table;
      content: "";
    }
    .clearfix:after {
      clear: both
    }
    .box-card {
      border:1px solid #dcdfe6 !important;
      border-radius: 5px;
      margin-bottom:1rem;
    }
    .el-card__body {
      background:rgb(248, 248, 248) !important;
    }
    .el-tag {
      padding:0 6px !important;
      height:25px !important;
      line-height: 25px !important;
    }
    /*分页*/
    .pagination {
      float:right;
      margin-top:1rem;
    }
    .el-pagination.is-background .el-pager li:not(.disabled).active {
      background-color:#41B883;
    }
  }
</style>
