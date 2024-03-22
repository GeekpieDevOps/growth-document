<template>
  <v-card class="pa-5 rounded-lg mb-5" variant="flat" color="#ecf2ff">
    <v-card-title class="text-h6 font-weight-bold">我的信息</v-card-title
    ><v-breadcrumbs :items="Items" class="text-subtitle-2">
      <template v-slot:divider>
        <v-icon icon="mdi-chevron-right"></v-icon>
      </template>
    </v-breadcrumbs>
  </v-card>
</template>
<script setup>
import { ref } from "vue";
defineProps(["Items"]);
</script>
