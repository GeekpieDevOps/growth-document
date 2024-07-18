<template>
  <UserFormShow @sub="login" :isLoading="isLoading" btnText="注册">
    <template #information>
      已有账号？<a
        href="/login"
        class="text-h7 text-decoration-none text-indigo signupformcontainer-hover"
        >登录</a
      ></template
    >
  </UserFormShow>
</template>

<script>
import UserFormShow from "@/components/public/UserFormShow.vue";

import notify from "@/assets/js/notify.js";

import { path, api } from "@/data.js";

export default {
  name: "SignupFormContainer",
  components: {
    UserFormShow,
  },
  data() {
    return {
      isLoading: false,
    };
  },
  methods: {
    login({ email, password }) {
      this.isLoading = true;
      fetch(api.signup, {
        method: "POST",
        body: JSON.stringify({ email, password }),
      })
        .then((res) => {
          this.isLoading = false;
          if (res.ok) {
            notify("info", "注册成功", () => {
              window.location.href = path.account;
            });
          } else {
            notify("error", "注册失败");
          }
        })
        .catch(() => {
          this.isLoading = false;
          notify("error", "注册失败");
        });
    },
  },
};
</script>

<style scoped>
.signupformcontainer-hover:hover {
  text-decoration: underline !important;
}
</style>
