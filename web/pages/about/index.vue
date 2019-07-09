<template>
<div>
	<nav-header :active="active"></nav-header>
	<el-row type="flex" justify="center" class="about_content">
		<el-col :span="15" class="about">
			<el-card class="box-card">
				<h1>About Blog</h1>
				<hr class="line"/>
    			<p>技術内容：</p>
    			<p class="about_title">Frontend：</p>
    			<p class="about_text"><b>SSR</b>,Vue.js ベースフレームワーク:<b>Nuxt.js + Element-UI + @nuxtjs/axios + Less</b></p>
    			<p class="about_title">Backend:</p>
    			<p class="about_text">Admin： <b>Vue + Vue-Router + Vue-cli + Vue + Axios + Less + Iview</b></p>
    			<p class="about_text">サーバー： <b>Golang + Mysql + MongoDB + Redis</b></p>
    			<h1>About Me</h1>
				<hr class="line"/>
				<p class="about_title">紹介</p>
				<ul>
					<li>名前：JunJun</li>
					<li>座右銘：できないものはなし</li>
				</ul>
				<p class="about_title">他技術</p>
				<ul>
					<li>HTML、CSS、Javacript(ES next)、JQuery、Vue</li>
					<li>Bootstrap、Layer-UI、WEUI、Muse-UI、Vux、MUI+HUI、Element-UI、Mint-UI、Iview</li>
					<li>NodeJs、MongoDB、Golang</li>
				</ul>
				<p class="about_title">アピール</p>
				<p class="about_text">サービスの作りは止まらない</p>
				<p class="about_title">メール</p>
				<p class="about_text">brightzamber@gmail.com</p>
			</el-card>
		</el-col>
	</el-row>
</div>
</template>

<script>
import NavHeader from '@/components/NavHeader.vue'
export default {
	data() {
		return {
			active:'about'
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
	}
}
</script>

<style lang="less" scoped>
	@import './../../assets/less/about.less';
</style>
