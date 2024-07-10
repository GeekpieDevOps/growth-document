<template>
  <v-card
    :loading="isLoading"
    class="px-1 px-md-3 px-lg-1 py-5 py-sm-7 mx-lg-6"
    width="500px"
    style="background-color: #eceff1"
  >
    <v-card-title tag="h2">
      <h2 style="text-align: center; padding: 5px">欢迎来到综合素质培养平台</h2>
      <p style="text-align: center">您的书院管理系统</p>
    </v-card-title>
    <v-card-text>
      <v-row>
        <v-col class="d-flex justify-center align-center ma-2">
          <v-btn prepend-icon="$vuetify"> egate统一登陆 </v-btn>
        </v-col>
        <v-col class="d-flex justify-center align-center ma-4">
          <v-btn prepend-icon="$vuetify"> ShanghaiTech登录 </v-btn>
        </v-col>
      </v-row>
      <v-divider class="my-4"></v-divider>

      <v-form @submit.prevent="loginSubmit">
        <v-row style="margin-top: 20px">
          <v-col cols="0" sm="1"></v-col>
          <v-col cols="12" sm="10"
            ><v-text-field
              prepend-icon="mdi-account"
              v-model="ID"
              :rules="rules.ID"
              label="ID"
              type="input"
              clearable
              required
            ></v-text-field>
            <v-text-field
              v-model="pwd"
              prepend-icon="mdi-lock"
              :append-inner-icon="showPwd ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="rules.pwd"
              label="密码"
              :type="showPwd ? 'text' : 'password'"
              @click:append-inner="showPwd = !showPwd"
              required
            ></v-text-field>
            <v-col cols="0" sm="1"></v-col
          ></v-col>
        </v-row>
        <v-row>
          <v-checkbox
            label="记住该用户"
            style="margin-left: 20px"
            class="d-flex justify-center align-center"
            :model-value="true"
          ></v-checkbox>
          <a
            href="#"
            style="position: absolute; right: 50px; transform: translateY(18px)"
            class="text-blue text-decoration-none"
            target="_blank"
            >忘记密码</a
          >
        </v-row>

        <v-row>
          <v-col cols="3"></v-col>
          <v-col>
            <v-btn
              class="mt-2 bg-blue font-weight-bold"
              block
              style="border-radius: 10px; font-size: 18px"
              type="submit"
              >登录</v-btn
            ></v-col
          ><v-col cols="3"></v-col
        ></v-row>

        <v-card-actions>
          <v-spacer></v-spacer>
          还没有账户？<v-btn @click="register" text color="primary"
            >去注册</v-btn
          >
        </v-card-actions>
      </v-form>
    </v-card-text>
    <v-dialog :model-value="isMistake" max-width="450">
      <v-alert type="warning" title="登陆失败" closable></v-alert> </v-dialog
  ></v-card>
</template>

<script>
import { path, api } from "@/data.js";

export default {
  name: "LoginCard",
  data() {
    return {
      ID: "",
      pwd: "",
      showPwd: false,
      isLoading: false,
      isMistake: false,
      rules: {
        ID: [
          (v) => !!v || "ID不能为空",
          (v) => (v && v.length <= 10) || "ID长度不能超过10",
        ],
        pwd: [
          (v) => !!v || "密码不能为空",
          (v) => (v && v.length <= 10) || "密码长度不能超过10",
        ],
      },
    };
  },
  methods: {
    checkForm() {
      const items = ["ID", "pwd"];
      for (let item of items) {
        for (let rule of this.rules[item]) {
          if (!rule(this[item])) return false;
        }
      }
    },
    register() {
      window.location.href = path.signup;
    },
    loginSubmit() {
      const url = api.login;
      console.log(this.ID, this.pwd);
      if (!this.checkForm()) return;
      fetch(url, {
        method: "POST",
        body: JSON.stringify({
          email: this.ID,
          password: this.pwd,
        }),
      })
        .then(() => {
          this.isLoading = false;
        })
        .catch((error) => {
          this.isLoading = false;
          this.isMistake = true;
          setTimeout(() => {
            this.isMistake = false;
          }, 3000);
        });
      this.isLoading = true;
    },
  },
};
</script>
