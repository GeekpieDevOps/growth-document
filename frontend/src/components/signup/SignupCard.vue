<template>
  <v-card
    :loading="isLoading"
    class="px-1 px-md-3 px-lg-1 py-5 py-sm-7 mx-lg-6"
  >
    <v-card-title class="text-center">注册</v-card-title>
    <v-card-text>
      <v-form @submit.prevent="signupSubmit">
        <v-row>
          <v-col cols="0" sm="1"></v-col>
          <v-col cols="12" sm="10"
            ><v-text-field
              prepend-icon="mdi-email"
              v-model="Email"
              :rules="rules.Email"
              label="Email"
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
            <v-col cols="0" sm="1"></v-col></v-col
        ></v-row>

        <v-row>
          <v-col cols="3"></v-col>
          <v-col>
            <v-btn class="mt-2 bg-blue" type="submit" block>注册</v-btn></v-col
          ><v-col cols="3"></v-col
        ></v-row>
      </v-form>
    </v-card-text>
    <v-overlay :model-value="isMistake" class="align-center justify-center">
      <v-alert type="warning" title="注册失败" closable></v-alert> </v-overlay
  ></v-card>
</template>

<script>
import { path, api } from "@/data.js";

export default {
  name: "SignupCard",
  data() {
    return {
      Email: "",
      pwd: "",
      showPwd: false,
      isLoading: false,
      isMistake: false,
      rules: {
        Email: [(v) => !!v || "邮箱不能为空！"],
        pwd: [(v) => !!v || "密码不能为空！"],
      },
    };
  },
  methods: {
    checkForm() {
      const items = ["Email", "pwd"];
      for (let item of items) {
        for (let rule of this.rules[item]) {
          console.log(item, this[item]);
          if (!rule(this[item])) {
            return false;
          }
        }
      }
    },
    signupSubmit() {
      const url = api.signup;
      if (this.checkForm()) return;

      console.log(this.Email, this.pwd);
      const fetchPromis = fetch(url, {
        method: "POST",
        body: JSON.stringify({
          email: this.Email,
          password: this.pwd,
        }),
      });
      this.isLoading = true;
      fetchPromis
        .then(() => {
          this.isLoading = false;
        })
        .catch(() => {
          this.isLoading = false;
          this.isMistake = true;
          setTimeout(() => {
            this.isMistake = false;
          }, 3000);
        });
    },
  },
};
</script>
