<template>
  <div class="narbar">
    <el-menu class="el-menu-demo" mode="horizontal" router :default-active="$route.path">  <!-- 根据$route.path来动态激活标签 -->
      <el-menu-item index="/"><i class="el-icon-s-home"></i><span class="hidden-xs-only">首页</span></el-menu-item>
      <el-menu-item index="/pushvideo"><i class="el-icon-upload2"></i><span class="hidden-xs-only">投稿</span></el-menu-item>
      <el-menu-item index="/about"><i class="el-icon-warning-outline"></i><span class="hidden-xs-only">关于</span></el-menu-item>
      <el-submenu index="2" class="users" >
        <template slot="title" v-if="!isLogin"><i class="el-icon-user-solid"></i><span class="hidden-xs-only">登陆/注册</span></template>
        <template slot="title" v-if="isLogin">NO.{{ id }} | {{ nickname }}</template>
        <el-menu-item index="/user/login" v-show="!isLogin">登陆</el-menu-item>
        <el-menu-item index="/user/register" v-show="!isLogin">注册</el-menu-item>
        <el-menu-item @click="pushToLiveRomm" v-show="isLogin">我的直播间</el-menu-item>
        <el-menu-item @click="pushToRoom" v-show="isLogin">个人中心</el-menu-item>
        <el-menu-item @click="logout" v-show="isLogin">注销</el-menu-item>
      </el-submenu>
    </el-menu>
  </div>
</template>

<script>
  import * as API from '../api'

  export default {
    name: 'NavBar',
    data() {
      return {
        isLogin: false,
        id: "",
        nickname: "",
      }
    },
    beforeCreate() {
      API.userStatus().then(response => {
        if (response.data !== null) {
          this.isLogin = true;
          this.id = response.data.id;
          this.nickname = response.data.nickname;
        }
      })
    },
    methods: {
      logout() {
        API.userLogout().then(() => {
          this.$router.push('/');
          location.reload()
        })
      },
      pushToLiveRomm() {
        this.$router.push(`/live/${this.id}`);
      },
      pushToRoom() {
        this.$router.push('/user/home');
      }
    }
  }
</script>

<style scoped>
  .users {
    float: right;
  }
</style>
