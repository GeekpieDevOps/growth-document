<template>
  <v-app>
    <v-main class="ma-0 pa-0">
      <UserBar strong="注册" />
      <v-row
        class="ma-0 d-flex justify-center align-center pa-0"
        :style="{
          width: '100vw',
          height: '100vh',
        }"
        ><v-col
          cols="12"
          md="7"
          class="bg-indigo-lighten-5 d-none d-md-block h-screen pa-0 ma-0"
        >
          <v-container
            class="d-flex justify-center align-center"
            style="height: 100vh"
          >
            <v-img
              src="../../../public/imgs/物质塔.jpg"
              :max-width="400"
            ></v-img> </v-container></v-col
        ><v-col cols="12" md="5" class="pa-0">
          <SignupFormContainer /> </v-col></v-row
    ></v-main>
  </v-app>
</template>

<script>
import UserBar from "@/components/public/UserBar.vue";
import SignupFormContainer from "@/components/signup/SignupFormContainer.vue";

export default {
  name: "App",
  components: {
    UserBar,
    SignupFormContainer,
  },
};
</script>
