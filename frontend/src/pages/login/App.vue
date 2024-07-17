<template>
  <v-app>
    <v-main>
      <v-container>
        <v-row
          ><v-col cols="12" md="6"> <UserBar /> </v-col
          ><v-col cols="12" md="6">asdasda</v-col></v-row
        ></v-container
      >
    </v-main>
  </v-app>
</template>

<script>
import UserBar from "@/components/public/UserBar.vue";

export default {
  name: "App",
  components: {
    UserBar,
  },
};
</script>
