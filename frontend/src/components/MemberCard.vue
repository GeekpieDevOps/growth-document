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
    <v-card
      :loading="isLoading"
      prepend-icon="mdi-pencil"
      :title="'编辑：' + name"
    >
      <v-card-text>
        <v-text-field
          :label="name"
          :placeholder="content"
          v-model="changeValue"
        ></v-text-field
      ></v-card-text>
      <v-card-actions class="d-flex justify-end"
        ><v-btn @click="isShowDialog = false" variant="plain" hover>取消</v-btn
        ><v-btn color="primary" @click="submitChange" variant="tonal"
          >确定</v-btn
        >
      </v-card-actions>
    </v-card>
    <v-dialog v-model="changeWrong" max-width="300px">
      <v-alert
        type="warning"
        title="修改失败"
        closable
        v-model="changeWrong"
      ></v-alert></v-dialog
    ><v-dialog v-model="changeSuccess" max-width="300px">
      <v-alert
        type="success"
        title="修改成功"
        closable
        v-model="changeSuccess"
      ></v-alert></v-dialog
  ></v-dialog>
</template>

<script setup>
import { defineProps, ref } from "vue";
const datasGet = defineProps({
  name: String,
  content: String,
});
const name = ref(datasGet.name);
const content = ref(datasGet.content);
const changeValue = ref("");
const isLoading = ref(false);
const changeWrong = ref(false);
const changeSuccess = ref(false);
function submitChange() {
  isLoading.value = true;
  const fetchPromis = fetch(
    "http://127.0.0.1:4523/m1/4078973-0-default/api/v1/user/1",
    {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ value: changeValue.value, name: name.value }),
    }
  );
  fetchPromis
    .then((res) => {
      isLoading.value = false;
      changeSuccess.value = true;
      content.value = changeValue.value;
      return res.json();
    })
    .then((data) => {
      console.log(data);
    })
    .catch((err) => {
      isLoading.value = false;
      changeWrong.value = true;
    });
}
const isShowDialog = ref(false);
</script>
