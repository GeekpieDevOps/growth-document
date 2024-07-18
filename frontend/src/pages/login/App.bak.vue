<!-- <template>
  <v-app class="bg-blue-lighten-2">
    <v-main>
      <LoginNav :tab="0" />
      <v-container
        fluid
        fill-height
        style="width: 100vw; height: 100vh"
        class="d-flex flex-column pa-1 pa-sm-10 justify-center"
      >
        <v-spacer></v-spacer>
        <v-row>
          <v-col cols="0" sm="2" md="3"></v-col>
          <v-col cols="12" sm="8" md="6"> <LoginCard /></v-col>
          <v-col cols="0" sm="2" md="3"></v-col
        ></v-row>
        <v-row class="d-flex justify-space-around align-center py-5 px-2">
          <LogoShow />
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template> -->

<template>
  <v-app style="background-color: #cfd8dc">
    <v-main>
      <!-- Navigation -->
      <LoginNav :tab="0" />
      <!-- Contents -->

      <v-row class="rounded rounded-md pa-0" style="height: 90vh">
        <v-col
          cols="6"
          cols-sm="12"
          class="d-none d-lg-flex align-center justify-center ma-2"
          style="
            min-height: 300px;
            background-image: url('../../../public/imgs/IMG_0638-已增强-NR33.jpg');
            background-size: cover;
            background-position: center;
            box-shadow: inset 0px 0px 20px 10px rgba(0, 0, 0, 0.6);
            border-radius: 10px;
          "
        >
          <div
            style="
              position: absolute;
              bottom: 20px;
              width: 60%;
              display: flex;
              justify-content: space-around;
            "
          >
            <LogoShow />
          </div>
        </v-col>

        <v-col
          class="d-flex align-center justify-center pa-0"
          style="min-height: 300px"
        >
          <!-- <v-card v-bind="props" class="elevation-12 mx-auto pa-6" width="500px" style="border:1px solid black;">
                <v-card-title tag="h2">
                  <h2>欢迎来到综合素质培养平台</h2> 
                  <p>您的书院管理系统</p>
                </v-card-title>

              <v-card-text>
                <v-form ref="form" v-model="valid" lazy-validation>
                  <v-text-field
                    v-model="email"
                    :rules="emailRules"
                    label="邮箱"
                    required
                  ></v-text-field>


                  <v-text-field
                    v-model="password"
                    :rules="passwordRules"
                    label="密码"
                    type="password"
                    required
                  ></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-btn color="primary" @click="login"  width="200px" class="mx-auto elevation-4 font-weight-bold" :hover="true">登录</v-btn>
              </v-card-actions>
              <v-card-actions>
                <v-spacer></v-spacer>
                还没有账户？<v-btn @click="register" text color="primary">注册</v-btn>
              </v-card-actions>
            </v-card> -->

          <LoginCard />
        </v-col>
      </v-row>
    </v-main>
  </v-app>
</template>

<!-- <script>
export default {
  data: () => ({
    valid: true,
    email: '',
    emailRules: [
      v => !!v || '请填写邮箱地址',
      v => /.+@.+\..+/.test(v) || '邮箱地址格式错误',
    ],
    password: '',
    passwordRules: [
      v => !!v || '请输入密码',
    ],
  }),
  methods: {
    login() {
      if (this.$refs.form.validate()) {
        // TODO: Perform login action
        console.log(this.email, this.password);
      }
    },
    register() {
      window.location.href = '../../signup.html'
    },
  },
};
</script> -->

<script setup>
import LoginCard from "@/components/login/LoginCard.vue";
</script>
