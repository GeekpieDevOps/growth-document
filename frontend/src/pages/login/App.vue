<template>
  <v-app>
    <v-main class="ma-0 pa-0">
      <UserBar />
      <v-row
        class="ma-0 d-flex justify-center align-center"
        :style="{
          width: '100vw',
          height: '100vh',
        }"
        ><v-col cols="12" md="6" class="bg-indigo-lighten-5 d-none d-md-block">
          <v-container
            class="d-flex justify-center align-center"
            style="height: 100vh"
          >
            <v-img
              src="../../../public/imgs/物质塔.jpg"
              :max-width="400"
            ></v-img> </v-container></v-col
        ><v-col cols="12" md="6" class="pa-0"> <UserFormShow /> </v-col></v-row
    ></v-main>
  </v-app>
</template>

<script>
import UserBar from "@/components/public/UserBar.vue";
import UserFormShow from "@/components/public/UserFormShow.vue";

export default {
  name: "App",
  components: {
    UserBar,
    UserFormShow,
  },
};
</script>
