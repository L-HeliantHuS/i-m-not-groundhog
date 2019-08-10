import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import 'element-ui/lib/theme-chalk/display.css';
import {
    Menu,
    Submenu,
    MenuItem,
    MenuItemGroup,
    Input,
    InputNumber,
    Form,
    FormItem,
    Icon,
    Row,
    Col,
    Upload,
    Card,
    Transfer,
    Container,
    Header,
    Main,
    Link,
    Backtop,
    PageHeader,
    Loading,
    // Notification,
    Button,
    ButtonGroup,
} from 'element-ui';

Vue.use(Menu);
Vue.use(Submenu);
Vue.use(MenuItem);
Vue.use(MenuItemGroup);
Vue.use(Input);
Vue.use(InputNumber);
Vue.use(Form);
Vue.use(FormItem);
Vue.use(Button);
Vue.use(ButtonGroup);
Vue.use(Icon);
Vue.use(Row);
Vue.use(Col);
Vue.use(Upload);
Vue.use(Card);
Vue.use(Main);
Vue.use(Transfer);
Vue.use(Container);
Vue.use(Header);
Vue.use(Link);
Vue.use(Backtop);
Vue.use(PageHeader);
Vue.use(Loading);
// Vue.use(Notification);
Vue.use(Button);
Vue.use(ButtonGroup);


import './plugins/element.js'

Vue.config.productionTip = false;

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app');
