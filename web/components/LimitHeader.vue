<template>
<div>
	<div class="nav">
	<el-row type="flex" class="NavHeader" justify="center">
		<el-col :span="20">
			<nuxt-link :to="{name:'index'}" class="logo">Nuxt Sample Blog</nuxt-link>
			<el-menu :default-active="active" class="el-menu-demo navbar" active-text-color="#41b883" text-color="#388"  mode="horizontal" @select="handleSelect">
				<el-menu-item index="index">ホーム</el-menu-item>
                <el-menu-item index="limit">会員限定</el-menu-item>
			</el-menu>
		</el-col>
	</el-row>
	</div>
</div>
</template>

<script>
export default {
	data() {
		return {
			activeIndex:'index'
		}
	},
	methods:{
		handleSelect(key){
			this.$router.push({name:key});
		}
	},
	props:['active']
}
</script>
<style lang="less">
	@import './../assets/less/navheader.less';
</style>
