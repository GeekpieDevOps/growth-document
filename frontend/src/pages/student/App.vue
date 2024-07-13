<template>
  <v-app>
    <v-main>
      <v-container>
        <MemberFram /><v-container class="pa-0">
          <MemberTitle
            :Items="[
              { title: '我的', disabled: false, href: '#' },
              { title: '我的信息', disabled: true, href: '#' },
            ]"
          />
          <StudentActivitiesContainer /> </v-container
      ></v-container>
    </v-main>
  </v-app>
</template>

<script>
import MemberFram from "@/components/MemberFram.vue";
import MemberTitle from "@/components/MemberTitle.vue";
import StudentActivitiesContainer from "@/components/student/activities/StudentActivitiesContainer.vue";

export default {
  components: {
    MemberFram,
    MemberTitle,
    StudentActivitiesContainer,
  },
};
</script>
