<template>
  <div class="register">
    <el-form ref="usertable" :rules="rules" :model="usertable" label-width="80px" status-icon>
      <el-form-item label="昵称" prop="nickname" required>
        <el-input type="text" v-model="usertable.nickname" placeholder="请输入昵称"></el-input>
      </el-form-item>
      <el-form-item label="账号" prop="user_name" required>
        <el-input type="text" v-model="usertable.user_name" placeholder="请输入账号"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password" required>
        <el-input type="password" v-model="usertable.password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item label="确认密码" prop="password_confirm" required>
        <el-input type="password" v-model="usertable.password_confirm" placeholder="再次输入密码"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" round @click="submitForm('usertable')">点我注册</el-button>
        <div id="TencentCaptcha"></div>
      </el-form-item>
    </el-form>
  </div>
</template>
<!---->

<script>
  import * as API from '../../api'

  export default {
    name: "Register",
    data() {
      var checkpassword = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== this.usertable.password) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };
      return {
        captchaOption: {
          // 各平台的参数，具体请参阅个平台文档
          // 以下为腾讯验证码的参数
          appid: '2014850464',
          // 以下为极验验证码的参数
        },
        button_loading: false,
        usertable: {
          nickname: "",
          user_name: "",
          password: "",
          password_confirm: "",
          ticket: "",
          randstr: ""
        },
        rules: {
          nickname: [
            {required: true, message: '请输入昵称', trigger: 'blur'},
            {min: 2, max: 30, message: '长度在 2 到 30 个字符', trigger: 'blur'}
          ],
          user_name: [
            {required: true, message: '请输入账号', trigger: 'blur'},
            {min: 5, max: 30, message: '长度在 5 到 30 个字符', trigger: 'blur'}
          ],
          password: [
            {required: true, message: '请输入密码', trigger: 'blur'},
            {min: 8, max: 40, message: '长度在 8 到 40 个字符', trigger: 'blur'}
          ],
          password_confirm: [
            {required: true, message: '输入的密码不匹配', validator: checkpassword, trigger: 'blur'},
            {min: 8, max: 40, message: '长度在 8 到 40 个字符', trigger: 'blur'}
          ],
        }
      }
    },
    methods: {
      // 提交用户名和密码
      submitForm(refs) {
        this.$refs[refs].validate(valid => {
          this.button_loading = true;
          if (valid) {
            var captcha1 = new TencentCaptcha('2014850464', (res) => {
              // console.log(res)
              this.usertable.ticket = res.ticket;
              this.usertable.randstr = res.randstr;
              API.userRegister(this.usertable).then(response => {
                if (response.data == null) {
                  this.$notify.error({
                    message: response.msg,
                  });
                  this.button_loading = false;
                } else {
                  this.$notify({
                    type: 'success',
                    message: `${response.data.nickname}你好~ 你是本站的第${response.data.id}名用户~ \n
                    已经自动为你开启直播间. 请登录后去个人中心查看。`,
                    duration: 0
                  });
                  this.$router.push('/user/login')
                }
              })
            });
            captcha1.show(); // 显示验证码

          } else {
            return false
          }
        })
      }
    }

  }
</script>

<style scoped>

</style>
