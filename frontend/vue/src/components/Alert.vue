<template>
  <v-card
    class="pa-7 rounded-lg d-flex flex-column align-center"
    color="#5d87ff"
  >
    <v-card-title class="text-h6 align-self-start"
      ><v-icon class="mr-3" prepend>mdi-information</v-icon
      >{{ title }}</v-card-title
    >
    <v-card-text class="text-h6 mb-10">{{ text }}</v-card-text>
    <v-row class="align-self-stretch">
      <v-spacer></v-spacer>
      <v-col cols="4">
        <v-btn width="80" @click="isActive.value = false" variant="flat"
          >取消</v-btn
        ></v-col
      >
      <v-spacer></v-spacer
      ><v-col
        ><v-btn width="80" cols="4" @click="active(isActive)" variant="flat">{{
          activeText
        }}</v-btn></v-col
      >
      <v-spacer></v-spacer
    ></v-row>
  </v-card>
</template>
<script>
export default {
  name: "Alert",
  props: {
    title: {
      type: String,
      default: "注意",
    },
    text: {
      type: String,
      default: "这是一个警告",
    },
    activeText: {
      type: String,
      default: "确定",
    },
    isActive: {},
    active: {
      type: Function,
      default: () => {},
    },
  },
};
</script>
