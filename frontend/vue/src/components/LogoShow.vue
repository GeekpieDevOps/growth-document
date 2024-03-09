<script setup>
import { ref } from "vue";
const datas = ref([
  {
    imgURL: "/src/assets/上道书院.png",
    webURL: "https://shangdao.shanghaitech.edu.cn/",
    imgColor: "#0527af",
    name: "上道书院",
  },
  {
    imgURL: "/src/assets/科道书院.png",
    webURL: "https://kedao.shanghaitech.edu.cn/",
    imgColor: "#b188bb",
    name: "科道书院",
  },
  {
    imgURL: "/src/assets/大道书院.png",
    webURL: "https://dadao.shanghaitech.edu.cn/",
    imgColor: "#688a51",
    name: "大道书院",
  },
]);
</script>

<template>
  <sylogo
    v-for="set in datas"
    :imgURL="set.imgURL"
    :webURL="set.webURL"
    :imgColor="set.imgColor"
    :name="set.name"
  />
</template>
