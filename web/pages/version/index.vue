<template>
  <div>
  	<nav-header :active="active"></nav-header>
	<el-row type="flex" justify="center">
		<el-col :span="10">
			<h1>タイムライン</h1>
		</el-col>
	</el-row>
	<el-row type="flex" class="version" justify="center" v-for="item in versions" :key="item._id">
		<el-col :span="10">
			<el-card class="box-card">
			  <div slot="header" class="clearfix">
			    <span style="font-weight:bold;">{{item.version}}</span>
			    <span style="float:right">{{item.created_at}}</span>
			  </div>
			  <div v-html="item.name">
                  {{item.name}}
              </div>
			</el-card>
		</el-col>
	</el-row>
  </div>
</template>

<script>
import NavHeader from '~/components/NavHeader.vue'

export default {
    data() {
        return {
            active:'version'
        }
    },
    async asyncData({app}) {
        let result = await app.$axios.$get(`/api/version/list`);
        let {error,versions} = result;
        return {versions}
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

<style>
    .version {
        margin-bottom:2rem;
    }
</style>
