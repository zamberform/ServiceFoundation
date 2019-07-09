<template>
<div>
	<nav-header :active="active"></nav-header>
    <div v-if="userName == ''">
    {{ userName }}
    </div>
    <div v-else>
    <el-row type="flex" justify="center" class="about_content">
		<el-col :span="15" class="about">
			<el-card class="box-card">
				<el-form style="margin-top:20px" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
                    <el-form-item label="ユーザー" prop="username">
                        <el-input v-model="ruleForm.username"></el-input>
                    </el-form-item>
                <el-form-item label="パスワード" prop="password">
                        <el-input v-model="ruleForm.password"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="submit">ログイン</el-button>
                        <el-button @click="cancel">キャンセル</el-button>
                    </el-form-item>
                </el-form>
                <p><nuxt-link to="/signup">サイン</nuxt-link></p>
			</el-card>
		</el-col>
	</el-row>
    </div>
	
</div>
</template>

<script>
import NavHeader from '@/components/NavHeader.vue'
export default {
	data(){
        
        return {
            userName: this.$store.state.userName, 
            active:'login',
            ruleForm:{username:'',password:''},
            rules:{
                username:[{required:true,trigger:'blur',message:'ユーザー名を入れてください'}],
                password:[{required:true,trigger:'blur',message:'パスワード入れてください'}]
            }
        }
    },
	components:{
		NavHeader
	},
	head() {
		return {
			title:'Sample Nuxt Blog',
			meta:[
				{hid:'description',name:'description',content:'Use Nuxt to build Web Page'},
				{hid:'keywords',name:'keywords',content:'nuxt,nuxtjs,vue,vuejs'},
				{hid:'author',content:'zamberform'}
			]
		}
    },
    methods:{
        submit(){
            this.$refs['ruleForm'].validate(async (valide)=>{
                if(valide){
                    let {data} =  await this.$axios.post('/api/user/login',this.ruleForm)
                    
                    if(data.status == 0){
                        this.$store.commit('setLoginState', true)
                        this.$store.commit('token', data.user.token)
                        this.$store.commit('user', data.user.name)
                        this.$store.commit('limit', data.user.isLimit)
                    }
                }else{
                    return false
                }
            })
        },
        cancel(){
            window.location.href = '/'
        }
    }
}
</script>

<style lang="less" scoped>
	@import './../../assets/less/about.less';
</style>