<template>
  <v-card
    :loading="isLoading"
    class="px-1 px-md-3 px-lg-1 py-5 py-sm-7 mx-lg-6"
  >
    <v-card-title class="text-center">登录</v-card-title>
    <v-card-text>
      <v-form @submit.prevent="loginSubmit">
        <v-row>
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
            <v-col cols="0" sm="1"></v-col></v-col
          ><v-row class="text-center">
            <v-col
              ><a
                class="text-blue text-decoration-none"
                href="#"
                target="_blank"
                >注册</a
              ></v-col
            >
            <v-col>
              <a class="text-blue text-decoration-none" href="#" target="_blank"
                >忘记密码</a
              ></v-col
            >
          </v-row></v-row
        >

        <v-row>
          <v-col cols="3"></v-col>
          <v-col>
            <v-btn class="mt-2 bg-blue" type="submit" block>登录</v-btn></v-col
          ><v-col cols="3"></v-col
        ></v-row>
      </v-form>
    </v-card-text>
    <v-overlay :model-value="isMistake" class="align-center justify-center">
      <v-alert type="warning" title="登陆失败" closable></v-alert> </v-overlay
  ></v-card>
</template>

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
  console.log(name.value);
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
  const url = "http://127.0.0.1:4523/m1/4078973-0-default/api/v1/login";
  if (!Object.values(formIsChecked.value).every((value) => value === true)) {
    return false;
  }
  const fetchPromis = fetch(url, {
    method: "POST",
    body: JSON.stringify({
      username: ID.value,
      password: pwd.value,
    }),
  });
  isLoading.value = true;
  fetchPromis
    .then((response) => {
      isLoading.value = false;
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
