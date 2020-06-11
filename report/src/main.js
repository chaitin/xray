import Vue from "vue";
import App from "./App.vue";
import VueClipboard from 'vue-clipboard2'
import {Button, Card, Descriptions, Input, Layout, message, Modal, Table, Switch} from "ant-design-vue";

Vue.use(VueClipboard)
Vue.component(Layout.name, Layout)
Vue.component(Layout.Content.name, Layout.Content)
Vue.component(Layout.Header.name, Layout.Header)
Vue.component(Layout.Footer.name, Layout.Footer)
Vue.component(Button.name, Button)
Vue.component(Card.name, Card)
Vue.component(Table.name, Table)
Vue.component(Table.Column.name, Table.Column)
Vue.component(Modal.name, Modal)
Vue.component(Input.name, Input)
Vue.component(Input.Search.name, Input.Search)
Vue.component(Input.TextArea.name, Input.TextArea)
Vue.component(Descriptions.name, Descriptions)
Vue.component(Descriptions.Item.name, Descriptions.Item)
Vue.component(Switch.name, Switch)

Vue.config.productionTip = false;
Vue.prototype.$message = message
Vue.prototype.$copy = function (data) {
  this.$copyText(data).then(function (e) {
    Vue.prototype.$message.success('Copied', 3)
  }, function (e) {
    message.error('Can not copy', 3)
  })
}

Vue.filter('humanize', function (timestamp) {
  let date = new Date(timestamp);

  let month = date.getMonth() + 1;
  let day = date.getDate();
  let hour = date.getHours();
  let min = date.getMinutes();
  let sec = date.getSeconds();

  month = (month < 10 ? "0" : "") + month;
  day = (day < 10 ? "0" : "") + day;
  hour = (hour < 10 ? "0" : "") + hour;
  min = (min < 10 ? "0" : "") + min;
  sec = (sec < 10 ? "0" : "") + sec;

  return date.getFullYear() + "-" + month + "-" + day + " " + hour + ":" + min + ":" + sec;
})

Vue.filter('capitalize', function (value) {
  if (!value) return ''
  value = value.toString()
  return value.charAt(0).toUpperCase() + value.slice(1)
})

Vue.prototype.$debounce = function debounce(fn, delay) {
  let timer
  return function () {
    // 保存函数调用时的上下文和参数，传递给 fn
    let context = this
    let args = arguments
    // 每次这个返回的函数被调用，就清除定时器，以保证不执行 fn
    clearTimeout(timer)
    // 当返回的函数被最后一次调用后（也就是用户停止了某个连续的操作），
    // 再过 delay 毫秒就执行 fn
    timer = setTimeout(function () {
      fn.apply(context, args)
    }, delay)
  }
}

new Vue({
  render: h => h(App)
}).$mount("#app");
