<template>
<div>
	<nav-header :active="active"></nav-header>
	<div id="signup">
    <el-row type="flex" justify="center" class="about_content">
		<el-col :span="15" class="about">
			<el-card class="box-card">
				<el-form
      ref="registerform"
      :model="registerform"
      :rules="rules"
    >
      <el-form-item prop="username">
        <el-input
          v-model="registerform.username"
          placeholder="ユーザー名"
        />
      </el-form-item>
      <el-form-item prop="useremail">
        <el-input
          v-model="registerform.useremail"
          placeholder="メールアドレス"
        />
      </el-form-item>
      <el-form-item prop="userpwd">
        <el-input
          v-model="registerform.userpwd"
          type="password"
          placeholder="パスワード"
        />
      </el-form-item>
      <el-form-item prop="userpwd2">
        <el-input
          v-model="registerform.userpwd2"
          type="password"
          placeholder="もう一度パスワード"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="success"
          plain
          class="allw"
          @click="submitForm('registerform')"
        >
          サインイン
        </el-button>
      </el-form-item>
      </el-form>
			</el-card>
      <p>今すぐ<nuxt-link to="/login">ログイン</nuxt-link></p>
		</el-col>
	</el-row>
    
  </div>
</div>
</template>

<script>
import NavHeader from '@/components/NavHeader.vue'
export default {
	data() {
    var validateuserpwd2 = (rule, value, callback) => {
      if (value !== this.registerform.userpwd) {
        callback(new Error('パスワード一致していません！'))
      } else {
        callback()
      }
    }
    return {
      registerform: {
        username: '',
        useremail: '',
        userpwd: '',
        userpwd2: ''
      },
      rules: {
        username: [
          {
            required: true,
            message: 'ユーザー名を入力してください',
            trigger: 'blur'
          }
        ],
        useremail: [
          {
            required: true,
            message: 'メールアドレス入力してください',
            trigger: 'blur'
          },
          {
            type: 'email',
            message: 'メールアドレスエラー',
            trigger: 'blur'
          }
        ],
        userpwd: [
          {
            required: true,
            message: 'パスワード設定してください',
            trigger: 'blur'
          },
          {
            min: 6,
            max: 18,
            message: 'パスワードを６以上１８文字以内',
            trigger: 'blur'
          }
        ],
        userpwd2: [
          {
            required: true,
            message: 'もう一度パスワード入力してください',
            trigger: 'blur'
          },
          {
            validator: validateuserpwd2,
            trigger: 'blur'
          }
        ]
      }
    }
  },
  components:{
    NavHeader
  },
  methods: {
    submitForm(registerform) {
      this.$refs[registerform].validate((valid, obj) => {
        if (valid) {
          let { username, useremail, userpwd } = this.registerform
          setTimeout(() => {
            this.$axios
              .$post('api/auth/register', {
                username: username,
                email: useremail,
                password: userpwd
              })
              .then(res => {
                if (res.code === 0) {
                  this.$message({
                    showClose: true,
                    message: 'サイン成功、メールを確認してください',
                    type: 'success'
                  })
                  this.$router.push('/user/signin')
                } else {
                  this.$message({
                    showClose: true,
                    message: 'サイン失敗',
                    type: 'error'
                  })
                }
              })
              .catch(err => {
                console.log(err)
                this.$message({
                  showClose: true,
                  message: err.response.data.message,
                  type: 'error'
                })
              })
          }, 100)
        } else {
          this.$message({
            showClose: true,
            message: 'ユーザーの情報を正しく入力してください',
            type: 'error'
          })
          return false
        }
      })
      this.$refs[registerform].resetFields()
    }
  }  
}
</script>

<style lang="less" scoped>
	@import './../../assets/less/about.less';
</style>