module.exports = {
  devServer: {
    proxy: { // 配置跨域
      '^/api': {
        target: 'http://127.0.0.1/api',
        ws: true,
        changeOrigin: true,
      },
    },
  },
  configureWebpack: {
    // 把原本需要写在webpack.config.js中的配置代码 写在这里 会自动合并
    externals: {
      'axios': 'axios',
      'vuex': 'Vuex',
      'vue-router': 'VueRouter',
    }
  }
};
