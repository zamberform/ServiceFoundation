<template>
  <div>
    <limit-header :active="active"></limit-header>
    <div v-if="userName == '' || !isLimit">
        <el-row type="flex" justify="center" class="limit_content">
		<el-col :span="15" class="limit">
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
			<p>今すぐ<nuxt-link to="/login">ログイン</nuxt-link></p>
		</el-col>
	    </el-row>
    </div>
    <div v-else>
	    <img src="~/static/images/limit.jpg" alt="banner" class="limit_img">
    </div>
    <nav-footer />
  </div>
  
</template>

<script>
import LimitHeader from '~/components/LimitHeader.vue'
import NavFooter from '~/components/Footer.vue'

export default {
    layout: 'limit',
	data() {
		return {
            userName: this.$store.state.userName, 
            isLimit: this.$store.state.isLimit, 
			active:'limit',
		}
    },
    components: {
		LimitHeader,
		NavFooter
    },
	methods:{
		handleSelect(key){
			this.$router.push({name:key});
		}
	}
}
</script>

