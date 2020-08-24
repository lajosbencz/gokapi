
export default {
  mode: 'spa',
  target: 'static',
  srcDir: 'src/',
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/admin/favicon.ico' }
    ]
  },
  css: [
    '~/styles/index.scss',
  ],
  loading: '~/components/loading.vue',
  layoutTransition: {
    name: 'page',
    mode: 'out-in'
  },
  plugins: [
    '~/plugins/wamp.client.js',
    '~/plugins/iview',
  ],
  components: false,
  buildModules: [
  ],
  modules: [
    // ['bootstrap-vue/nuxt', {
    //   icons: true,
    // }],
    '@nuxtjs/pwa',
    '@nuxtjs/style-resources',
  ],
  router: {
    base: '/admin/',
  },
  styleResources: {
    scss: [
      'styles/global/*.scss',
    ]
  },
  build: {
    extend (config, { isDev, isClient }) {
      config.node = {
        fs: 'empty'
      }
      if(isDev && isClient) {
        // config.module.rules.push({
        //   enforce: 'pre',
        //   test: /\.(js|vue)$/,
        //   loader: 'eslint-loader',
        //   exclude: /(node_modules)/
        // })
      }
    }
  }
}
