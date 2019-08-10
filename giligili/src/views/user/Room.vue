<template>
  <div class="room">
    <div class="info">
      <el-row :gutter="24">
        <el-col :span="18">iD: {{ id }}</el-col>
      </el-row>
      <el-row>
        <el-col :span="18">账号: {{ username }}</el-col>
      </el-row>
      <el-row>
        <el-col :span="18">昵称: {{ nickname }}</el-col>
      </el-row>
      <el-row>
        <el-col :span="18">串流地址: {{ live_stream }}</el-col>
      </el-row>
      <el-row>
        <el-col :span="18">最后那串密文是串流密钥 obs填写 服务器: rtmp://......:1935/live    穿流密钥就是后面那串</el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
  import * as API from '../../api'
  export default {
    name: "Room",
    data() {
      return {
        id: null,
        username: "",
        nickname: "",
        live_stream: "",
      }
    },
    created() {
      API.userStatus().then(response => {
        if (response.status !== 0) {
          this.$router.push('/user/login')
        } else {
          this.id = response.data.id;
          this.username = response.data.user_name;
          this.live_stream = response.data.live_stream;
          this.nickname = response.data.nickname
        }
      })
    },
  }
</script>

<style scoped>
  .info {
    font-size: 25px;
  }
</style>
