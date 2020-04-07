const HTMLInlineSourceWebpackPlugin = require('html-webpack-inline-source-plugin');
const webpack = require('webpack');
const path = require('path')

const vueConfig = {
  configureWebpack: {
    optimization: {
      splitChunks: false // makes there only be 1 js file - leftover from earlier attempts but doesn't hurt
    },
    plugins: [
      new webpack.ContextReplacementPlugin(/moment[/\\]locale$/, /cn|en/),
      new HTMLInlineSourceWebpackPlugin()
    ],
    resolve: {
      alias: {
        '@ant-design/icons/lib/dist$': path.join(__dirname, './src/icons.js')
      }
    },
  },
  chainWebpack: config => {
    config
      .plugin('html')
      .tap(args => {
        args[0].inlineSource = '.(js|css)$'
        return args
      })
  },
  css: {
    extract: false,
    loaderOptions: {
      less: {
        modifyVars: {
          // less vars，customize ant design theme

          // 'primary-color': '#F5222D',
          // 'link-color': '#F5222D',
          // 'border-radius-base': '4px'
        },
        javascriptEnabled: true
      }
    }
  },

  devServer: {
    // development server port 8000
    port: 8000
    // proxy: {
    //   '/api': {
    //     target: 'https://mock.ihx.me/mock/5baf3052f7da7e07e04a5116/antd-pro',
    //     ws: false,
    //     changeOrigin: true
    //   }
    // }
  },

  // disable source map in production
  productionSourceMap: false,
};

module.exports = vueConfig;
