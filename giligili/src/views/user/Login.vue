<template>
  <div class="login">
    <el-form ref="userdata" :rules="rules" :model="userdata" label-width="80px" status-icon>
      <el-form-item label="账号" prop="user_name" required>
        <el-input type="text" v-model="userdata.user_name" placeholder="请输入账号"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password" required>
        <el-input type="password" v-model="userdata.password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('userdata')" round :loading="button_loading">点我登陆</el-button>
        <span class="gotoregister">没有账号? <router-link to="/user/register">点我注册</router-link></span>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
  import * as API from '../../api'
  export default {
    name: "Login",
    data() {
      return {
        button_loading: false,
        userdata: {
          user_name: "",
          password: ""
        },
        rules: {
          user_name: [
            {required: true, message: '请输入账号', trigger: 'blur'},
            {min: 5, max: 30, message: '长度在 5 到 30 个字符', trigger: 'blur'}
          ],
          password: [
            {required: true, message: '请输入密码', trigger: 'blur'},
            {min: 8, max: 40, message: '长度在 8 到 40 个字符', trigger: 'blur'}
          ],
        }
      }
    },

    methods: {
      submitForm(refs) {

        this.$refs[refs].validate(valid => {
          if (valid) {
            this.button_loading = true;
            API.userLogin(this.userdata).then(response => {
              if (response.status === 40001) {
                this.$notify.error({
                  message: response.msg,
                });
                this.button_loading = false;
              } else if (response.status === 0) {
                this.$notify({
                  type: 'success',
                  message: `欢迎回来！${response.data.nickname}`
                });
                this.$router.push('/');
                location.reload()
              }
            })
          } else {
            return false
          }
        })
      }
    }

  }
</script>

<style scoped>
  .gotoregister {
    margin-left: 20px;
  }
</style>
