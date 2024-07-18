<template>
  <UserFormShow @sub="login" :isLoading="isLoading" />
</template>

<script>
import UserFormShow from "@/components/public/UserFormShow.vue";

import notify from "@/assets/js/notify.js";

import { path, api } from "@/data.js";

export default {
  name: "LoginFormContainer",
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
      fetch(api.login, {
        method: "POST",
        body: JSON.stringify({ email, password }),
      })
        .then((res) => {
          this.isLoading = false;
          if (res.ok) {
            notify("info", "登录成功", () => {
              window.location.href = path.account;
            });
          } else {
            notify("error", "登录失败");
          }
        })
        .catch(() => {
          this.isLoading = false;
          notify("error", "登录失败");
        });
    },
  },
};
</script>
