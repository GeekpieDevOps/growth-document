<script setup>
import { ref, computed, reactive } from "vue";
import { useDisplay } from "vuetify/lib/framework.mjs";
const { name } = useDisplay();
const isPermanent = computed(() => {
  return name.value === "sm" || name.value === "xs" ? false : true;
});
const isdraw = ref(false);
const isShowNav = computed(() => {
  return isPermanent.value ? true : isdraw.value;
});
const items = ref([
  { text: "Real-Time", icon: "mdi-clock" },
  { text: "个人信息", icon: "mdi-account" },
  { text: "综合素质", icon: "mdi-flag" },
]);
const user = ref({
  name: "wzj",
  email: "123@email",
});
const tabIndex = ref(1);
</script>

<template>
  <v-app>
    <v-main>
      <v-container>
        <v-app-bar color="blue" class="px-3">
          <v-app-bar-nav-icon
            class="d-flex d-md-none"
            @click="isdraw = !isdraw"
          ></v-app-bar-nav-icon>
          <v-toolbar-title>欢迎：{{ user.name }}</v-toolbar-title>
        </v-app-bar>
        <v-row>
          <v-col cols="0">
            <v-navigation-drawer
              color="blue"
              dark
              location="left"
              transition="none"
              :permanent="isPermanent"
              v-model="isShowNav"
              class="py-3"
            >
              <v-tabs v-model="tabIndex" direction="vertical">
                <v-tab v-for="(item, i) in items" :key="i" :value="i">
                  <v-icon start>{{ item.icon }}</v-icon>

                  <v-list-item-title v-text="item.text"></v-list-item-title>
                </v-tab>
              </v-tabs> </v-navigation-drawer
          ></v-col>
          <v-col cols="12">
            <v-window v-model="tabIndex"
              ><v-window-item :value="0">
                <v-card>
                  <v-card-title>Real-Time</v-card-title>
                  <v-card-text>Real-Time</v-card-text>
                </v-card>
              </v-window-item>
              <v-window-item :value="1">
                <MemberWindow />
              </v-window-item>
              <v-window-item :value="2">
                <v-card>
                  <v-card-title>综合素质</v-card-title>
                  <v-card-text>综合素质</v-card-text>
                </v-card>
              </v-window-item>
            </v-window>
          </v-col></v-row
        >
      </v-container>
    </v-main>
  </v-app>
</template>
