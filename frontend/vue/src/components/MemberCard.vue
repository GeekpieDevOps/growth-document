<template>
  <v-card class="my-1" variant="tonal" color="primary">
    <v-row>
      <v-col>
        <v-card-title class="text">{{ name }}</v-card-title>
        <v-card-text>{{ content }}</v-card-text></v-col
      >
      <v-col class="d-flex pr-12 align-center justify-end">
        <v-btn color="primary" @click="isShowDialog = !isShowDialog"
          >编辑</v-btn
        ></v-col
      ></v-row
    >
  </v-card>
  <v-dialog v-model="isShowDialog" max-width="450px">
    <v-card :loading="isLoading">
      <v-card-title class="text-center">编辑 </v-card-title>
      <v-card-subtitle>{{ name }}</v-card-subtitle>
      <v-card-text>{{ content }} </v-card-text>
      <v-card-text
        ><v-text-field
          :label="name"
          :placeholder="content"
          v-model="changeValue"
        ></v-text-field
      ></v-card-text>
      <v-card-actions class="d-flex justify-end"
        ><v-btn color="primary" @click="submitChange">确定修改</v-btn>
        <v-btn color="primary" @click="isShowDialog = false"
          >取消</v-btn
        ></v-card-actions
      >
    </v-card></v-dialog
  >
</template>

<script setup>
import { defineProps, ref } from "vue";
const datas = defineProps({
  name: String,
  content: String,
});
const changeValue = ref("");
const isLoading = ref(false);
function submitChange() {
  isLoading.value = true;
  fetch("", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ value: changeValue.value, name: datas.name }),
  })
    .then((res) => {
      datas.name = changeValue.value;
      isLoading.value = false;
      return res.json();
    })
    .then((data) => {
      console.log(data);
    })
    .catch((err) => {
      isLoading.value = false;
      console.log(err);
    });
}
const isShowDialog = ref(false);
</script>
