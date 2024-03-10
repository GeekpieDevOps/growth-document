<template>
  <v-app class="bg-blue-lighten-2">
    <v-main>
      <LoginNav :tab="1" />
      <v-container
        fluid
        fill-height
        style="width: 100vw; height: 100vh"
        class="d-flex flex-column pa-10 justify-center"
      >
        <v-spacer></v-spacer>
        <v-row>
          <v-col cols="0" sm="2" md="3"></v-col>
          <v-col cols="12" sm="8" md="6"> <SignupCard /></v-col>
          <v-col cols="0" sm="2" md="3"></v-col
        ></v-row>
        <v-row class="d-flex justify-space-around align-center pt-5">
          <LogoShow />
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import LoginNav from "@/components/LoginNav.vue";

//
</script>
