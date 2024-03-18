<template>
  <v-app>
    <v-main>
      <v-container>
        <MemberFram /><v-container class="pa-0">
          <MemberTitle />
          <AccountWindow /> </v-container
      ></v-container>
    </v-main>
  </v-app>
</template>
<script setup>
import AccountWindow from "@/components/AccountWindow.vue";
import { ref, computed } from "vue";
import { useDisplay } from "vuetify/lib/framework.mjs";
const { name } = useDisplay();
const isPermanent = computed(() => {
  return name.value === "sm" || name.value === "xs" ? false : true;
});
</script>
<!-- const searchItems = ref([
  {
    title: "个人信息",
    subtitle: "https://vuetifyjs.com/en/components/icons/#buttons",
  },
  { title: "我的任务" },
  { title: "我的提交" },
]); -->
<!-- <v-menu
            v-model="showSearchMenu"
            :close-on-content-click="false"
            location="bottom"
            ><template v-slot:activator="{ props }">
              <v-btn icon="mdi-magnify" v-bind="props"></v-btn
            ></template>
            <v-card
              min-width="390"
              variant="elevated"
              style="height: fit-content"
            >
              <v-autocomplete
                class="mx-3"
                placeholder="搜索"
                :items="searchItems"
                variant="plain"
                clearable
              ></v-autocomplete> </v-card></v-menu> -->
