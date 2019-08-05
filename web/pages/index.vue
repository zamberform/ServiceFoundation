<template>
<div>
	<nav-header :active="active"></nav-header>
	<el-row type="flex" justify="center" class="content-blog">
		<el-col :span="10">
			<nuxt-link  v-for="item in articles" :key="item.title" :to="{name:'article-id',params:{id:item.id}}" class="box-href">
				<el-card class="box-card" shadow="hover">
					<h2 class="box-title">{{item.title}}</h2>
					<div class="box-icon">
						<span><i class="el-icon-date"></i>&nbsp;{{item.updated_at}}</span>
					</div>
					<div class="box-content">{{item.content_desc}}</div>
					<div class="box-content">{{item.tag}}</div>
				</el-card>
			</nuxt-link>
			<el-pagination class="pagination" @current-change="pagination" background layout="prev, pager, next" :page-size="5" :total="count" v-show="count >= 5"></el-pagination>
		</el-col>
    <el-col :span="5" :offset="1">
			<el-card class="about">
				<div class="about-title">About Me</div>
				<div class="about-name">
					<img src="~/static/images/name.png" alt="brianlee">
				</div>
				<div class="about-content">
					<p>名前：JunJun</p>
					<p>Github：Zamberform</p>
					<p>メール：brightzamber@gmail.com</p>
				</div>
			</el-card>
			<el-card class="article">
				<div class="article-title">最近</div>
				<hr>
				<nuxt-link v-for="item in lately" :key="item.id" :to="{name:'article-id',params:{id:item.id}}" class="article-link">
					<i class="el-icon-edit"></i>&nbsp;&nbsp;{{item.title}}
				</nuxt-link>
			</el-card>
			<el-card class="link">
				<div class="link-title">その他</div>
				<hr>
				<div class="link-content">
					<a href="/" target="_blank" class="link-url">Product1</a>
					<a href="/" target="_blank" class="link-url">Product2</a>
          <a href="/" target="_blank" class="link-url">Product3</a>
					<a href="/" target="_blank" class="link-url">Product4</a>
				</div>
			</el-card>
		</el-col>
	</el-row>
	<nav-footer />
</div>
</template>

<script>
import NavHeader from '~/components/NavHeader.vue'
import NavFooter from '~/components/Footer.vue'

export default {
	data() {
		return {
			active:'index'
		}
  },
  
	async asyncData({app}) {
    let json = {page:1,pagesize:5}
    let data = await app.$axios.$get(`/api/article/list`,{params:json});
    let {articles,count} = data;
    let lately = articles.slice(0,4);

		return {articles,count,lately}
  },
	methods: {
		pagination(page) {
			let json = {page,pagesize:5}
			this.$axios.$get(`/api/article/list`,{params:json}).then(res=>{
			  let {error,count,articles} = res.data;
        	  this.articles =articles;
        
			});
		}
  },
  
	components: {
		NavHeader,
		NavFooter
	},
	head() {
		return {
			title:'Nuxt Blog Sample',
			meta:[
				{hid:'description',name:'description',content:'A blog system frontend made by nuxt.js'},
				{hid:'keywords',name:'keywords',content:'blog,nuxt,nuxtjs,vue,vuejs'},
				{hid:'author',content:'zamberform'}
			]
		}
	}
}
</script>

<style lang="less">
    @import './../assets/less/content.less';
</style>
