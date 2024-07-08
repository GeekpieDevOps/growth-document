<template>
  <v-navigation-drawer
    :temporary="!isPermanent"
    :permanent="isPermanent"
    class="pa-3"
    width="270"
    v-model="isdraw"
    ><v-list-item rounded height="100">
      <v-row class="d-flex align-center">
        <v-col>
          <v-img
            width="45"
            src="/public/imgs/上道书院.png"
            style="
              background-color: #0527af;
              clip-path: circle(49.1% at 50% 50.5%);
            "
          ></v-img
        ></v-col>
        <v-col>
          <v-list-item-title class="text-h6 font-weight-bold"
            >上道书院</v-list-item-title
          ></v-col
        ></v-row
      >
    </v-list-item>
    <v-list>
      <v-list-subheader class="font-weight-bold text-black-50"
        >我的</v-list-subheader
      >
      <v-list direction="vertical">
        <v-list-item
          v-for="i in 10"
          class="my-1 rounded-lg"
          color="primary"
          prepend-icon="mdi-account"
          :key="i"
          :value="i"
          link
          >个人信息
          <template v-slot:append>
            <v-avatar append color="light-blue" size="25">12</v-avatar>
          </template></v-list-item
        >
      </v-list>
    </v-list> </v-navigation-drawer
  ><v-app-bar elevation="0"
    ><v-app-bar-nav-icon @click="isdraw = !isdraw"> </v-app-bar-nav-icon>
    <v-spacer></v-spacer>
    <v-menu
      v-model="showProfileMenu"
      :close-on-content-click="false"
      location="bottom"
    >
      <template v-slot:activator="{ props }">
        <v-btn class="rounded-circle" height="62" width="30" v-bind="props">
          <v-avatar
            image="https://randomuser.me/api/portraits/men/78.jpg"
          ></v-avatar></v-btn
      ></template>
      <v-card title="我的" min-width="340" class="px-7 py-5 rounded-lg">
        <v-row class="my-2 d-flex align-center"
          ><v-col cols="4"
            ><v-avatar
              size="80"
              image="https://randomuser.me/api/portraits/men/78.jpg"
            ></v-avatar> </v-col
          ><v-spacer></v-spacer
          ><v-col cols="7"
            ><v-row class="text-h6">我是谁</v-row>
            <v-row class="text-grey"
              ><v-icon>mdi-email-outline</v-icon>email</v-row
            ></v-col
          ></v-row
        >
        <v-divider></v-divider>
        <v-hover v-for="item in profileMenu">
          <template v-slot:default="{ isHovering, props }">
            <v-card variant="flat" v-bind="props" :href="item.url" class="px-3">
              <v-row class="my-2 d-flex align-center">
                <v-col cols="3">
                  <v-icon class="rounded-lg" color="blue-lighten-2" size="50">{{
                    item.icon
                  }}</v-icon></v-col
                ><v-spacer></v-spacer
                ><v-col cols="8">
                  <v-row
                    :class="`text-subtitle-1 font-weight-bold ${
                      isHovering ? 'text-blue' : ''
                    }`"
                    >{{ item.title }}</v-row
                  >
                  <v-row class="text-grey">{{ item.subtitle }}</v-row></v-col
                >
              </v-row></v-card
            ></template
          ></v-hover
        >
        <v-row class="d-flex justify-end ma-3">
          <v-btn color="blue-lighten-1" width="80" @click="signout"
            >登出</v-btn
          ></v-row
        ></v-card
      ></v-menu
    >
  </v-app-bar>
</template>
<script>
import { ref, computed } from "vue";
import { useDisplay } from "vuetify/lib/framework.mjs";
export default {
  setup() {
    const { name } = useDisplay();
    const isPermanent = computed(() => {
      return name.value === "sm" || name.value === "xs" ? false : true;
    });
    const showProfileMenu = ref(false);
    const isdraw = ref(
      name.value === "sm" || name.value === "xs" ? false : true
    );
    const profileMenu = [
      {
        title: "我的信息",
        subtitle: "Setting my profile",
        icon: "mdi-account-box",
        url: "#",
      },
      {
        title: "我的任务",
        subtitle: "Checking my task",
        icon: "mdi-calendar-check",
        url: "#",
      },
      {
        title: "我的提交",
        subtitle: "Checking my submit",
        icon: "mdi-send",
        url: "#",
      },
    ];
    function signout() {
      fetch();
    }
    return {
      isPermanent,
      showProfileMenu,
      isdraw,
      profileMenu,
      signout,
    };
  },
};
</script>
