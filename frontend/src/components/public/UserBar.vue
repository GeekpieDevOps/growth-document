<template>
  <v-app-bar class="userbar-v-app-bar">
    <template v-slot:prepend><v-icon class="ml-3">mdi-menu</v-icon> </template>
    <v-app-bar-title>
      <v-btn
        v-for="item of contents"
        variant="text"
        :href="item.link"
        class="text-h7"
        :color="item.title === strong ? 'primary' : ''"
        >{{ item.title }}</v-btn
      >
    </v-app-bar-title>
  </v-app-bar>
</template>

<script>
export default {
  name: "UserBar",
  props: {
    strong: { type: String },
  },
  data() {
    return {
      contents: [
        { title: "主页", link: "/" },
        { title: "书院", link: "https://life.shanghaitech.edu.cn/" },
        { title: "登录", link: "/login" },
        { title: "注册", link: "/signup" },
      ],
    };
  },
};
</script>

<style scope>
.userbar-v-app-bar {
  max-width: 50vw !important;
  border-bottom-right-radius: 25px !important;
}

@media (max-width: 700px) {
  .userbar-v-app-bar {
    max-width: 100vw !important;
    border-bottom-right-radius: 0 !important;
  }
}
</style>
