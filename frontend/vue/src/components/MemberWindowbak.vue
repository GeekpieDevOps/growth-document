<template>
  <v-container>
    <v-row class="d-block d-md-none">
      <MemberCard
        v-for="(value, key) in userDatas"
        :key="key"
        :name="key"
        :content="value"
    /></v-row>
    <v-row class="d-none d-md-flex">
      <v-col>
        <MemberCard
          v-for="(value, key) in userDatasTwoCol[0]"
          :key="key"
          :name="key"
          :content="value"
        />
      </v-col>
      <v-col>
        <MemberCard
          v-for="(value, key) in userDatasTwoCol[1]"
          :key="key"
          :name="key"
          :content="value"
        />
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { computed, ref } from "vue";
import MemberCard from "./MemberCard.vue";
// fetch("http://localhost:3000/user", {
//   method: "GET",
//   headers: {
//     "Content-Type": "application/json",
//   },
//   body: JSON.stringify({ email: "123@email" }),
// })
//   .then((res) => res.json())
//   .then((data) => {
//     var userDatas = data;
//   });
const userDatas = ref({
  姓名: "张三",
  年龄: "18",
  性别: "男",
  学院: "信息学院",
});
const userDatasTwoCol = computed(() => {
  const keys = Object.keys(userDatas.value);
  const values = Object.values(userDatas.value);
  const half = Math.ceil(keys.length / 2);
  const first = keys.slice(0, half);
  const second = keys.slice(half, keys.length);
  const firstValue = values.slice(0, half);
  const secondValue = values.slice(half, keys.length);
  const obj1 = {};
  const obj2 = {};
  first.forEach((key, i) => {
    obj1[key] = firstValue[i];
  });
  second.forEach((key, i) => {
    obj2[key] = secondValue[i];
  });
  return [obj1, obj2];
});
</script>
