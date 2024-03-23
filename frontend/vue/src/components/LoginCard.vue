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
              :prepend-icon="getViewPointName === 'xs' ? null : 'mdi-account'"
              v-model="ID"
              :rules="rulesForId"
              label="ID"
              type="input"
              clearable
              required
            ></v-text-field>
            <v-text-field
              v-model="pwd"
              :prepend-icon="getViewPointName === 'xs' ? null : 'mdi-lock'"
              :append-inner-icon="showPwd ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="rulesForPwd"
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
              type="submit"
              block
              style="border-radius: 10px; font-size: 18px"
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
export default {
  methods: {
    register() {
      window.location.href = "../../signup.html";
    },
  },
};
</script>

<script setup>
import { computed, ref } from "vue";
import { useDisplay } from "vuetify/lib/framework.mjs";
const ID = ref("");
const pwd = ref("");
const formIsChecked = ref({ Id: false, Pwd: false });
const showPwd = ref(false);
const isLoading = ref(false);
const { name } = useDisplay();
const getViewPointName = computed(() => {
  return name.value;
});
const isMistake = ref(false);
function procRulefunction(func, name) {
  function resultFunction(value) {
    let result = func(value);
    if (result === true) {
      formIsChecked.value[name] = true;
      return true;
    }
    formIsChecked.value[name] = false;
    return result;
  }
  return resultFunction;
}
const rulesForId = ref([
  procRulefunction((value) => {
    if (value) return true;
    return "用户ID不能为空！";
  }, "Id"),
]);
const rulesForPwd = ref([
  procRulefunction((value) => {
    if (value) return true;
    return "密码不能为空！";
  }, "Pwd"),
]);
function loginSubmit() {
  const url = "http://10.20.235.113:8080/api/v1/user/sign_in";
  if (!Object.values(formIsChecked.value).every((value) => value === true)) {
    return false;
  }
  const fetchPromis = fetch(url, {
    method: "POST",
    body: JSON.stringify({
      email: ID.value,
      password: pwd.value,
    }),
  });
  isLoading.value = true;
  fetchPromis
    .then((response) => {
      isLoading.value = false;
      console.log("hello");
      return response.json();
    })
    .then((json) => {
      json;
    })
    .catch((error) => {
      isLoading.value = false;
      isMistake.value = true;
      setTimeout(() => {
        isMistake.value = false;
      }, 3000);
    });
}
</script>
