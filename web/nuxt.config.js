
export default {
  mode: 'universal',
  /*
  ** Headers of the page
  */
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff' },
  /*
  ** Global CSS
  */
  css: [
    'element-ui/lib/theme-chalk/index.css',
    '~assets/css/common.css',
    '~assets/css/markdown.css'
  ],
  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
    '@/plugins/element-ui',
    '@/plugins/axios',
    {
      src: "~plugins/persistedstate.js",
      ssr: false
    }
  ],
  router: {
    middleware: ['auth']
  },
  /*
  ** Nuxt.js modules
  */
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/proxy'
  ],
  /*
  ** Build configuration
  */
  build: {
    transpile: [/^element-ui/],
    /*
    ** You can extend webpack config here
    */
    extend(config, ctx) {
    },
    loaders: {
      less: {
        javascriptEnabled: true,
        modifyVars: {
          'primary-color': 'rgba(222, 12, 101, 1.0)',
          'component-background': '#ffffff'
        }
      }
    }
  },
  server: {
    port: 3000, // デフォルト: 3000
    host: '0.0.0.0', // デフォルト: localhost
  },
  axios: {
    proxy: true
  },
  proxy: {
    '/api/': { target: 'http://gin:3333/api', pathRewrite: {'^/api/': ''}, changeOrigin: true }
  },
  generate: {
    routes: function () {
      return axios.get('http://gin:3333/api/article')
      .then((res) => {
        return res.data.map((article) => {
          return {
            route: '/article/' + article.id,
            payload: article
          }
        })
      })
    }
  }
}
