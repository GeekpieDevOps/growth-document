<template>
  <v-container class="d-flex justify-center align-center">
    <v-app-bar
      :collapse="isCollapse"
      @mouseenter="ChangeCollapse(false)"
      @mouseleave="ChangeCollapse(true)"
      :class="{ 'userbar-v-app-bar': !isCollapse }"
    >
      <template v-slot:prepend
        ><v-icon class="ml-3">mdi-menu</v-icon>
      </template>
      <v-app-bar-title>
        <v-btn
          v-for="item of contents"
          variant="text"
          :href="item.link"
          class="text-h6"
          >{{ item.title }}</v-btn
        >
      </v-app-bar-title>
    </v-app-bar>
    <v-img src="../../../public/imgs/物质塔.jpg" :max-width="400"></v-img>
  </v-container>
</template>

<script>
export default {
  name: "UserBar",
  data() {
    return {
      isCollapse: true,
      contents: [
        { title: "主页", link: "/" },
        { title: "书院", link: "https://life.shanghaitech.edu.cn/" },
        { title: "登录", link: "/login" },
        { title: "注册", link: "/signup" },
      ],
    };
  },
  methods: {
    ChangeCollapse(value) {
      this.isCollapse = value;
    },
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
