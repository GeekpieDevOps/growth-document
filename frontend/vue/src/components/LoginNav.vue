<template>
  <v-app-bar
    class="px-1 px-sm-3 bg-blue-darken-2 d-flex justify-center align-center"
  >
    <template v-slot:prepend> 书院综合素质成长档案平台 </template>
    <template v-slot:append>
      <v-tabs v-model="tab" align-tabs="center" class="d-none d-sm-flex">
        <v-tab href="https://www.shanghaitech.edu.cn/" target="_blank">
          登录
        </v-tab>
        <v-tab href="https://www.shanghaitech.edu.cn/" target="_blank">
          注册
        </v-tab>
        <v-tab href="https://www.shanghaitech.edu.cn/" target="_blank">
          学校官网 </v-tab
        ><v-tab href="https://life.shanghaitech.edu.cn/" target="_blank">
          书院的天空
        </v-tab>
      </v-tabs></template
    >
  </v-app-bar>
</template>

<script setup>
import { ref } from "vue";
const props = defineProps({
  tab: Number,
});
const tab = ref(props.tab);
</script>
