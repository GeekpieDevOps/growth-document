<template>
  <v-app>
    <v-main>
      <v-container>
        <MemberFram /><v-container class="pa-0">
          <MemberTitle
            :Items="[
              { title: '我的', disabled: false, href: '#' },
              { title: '我的信息', disabled: true, href: '#' },
            ]" />
          <TeacherUnscoredContainer /></v-container
      ></v-container>
    </v-main>
  </v-app>
</template>

<script>
import MemberFram from "@/components/MemberFram.vue";
import MemberTitle from "@/components/MemberTitle.vue";
import TeacherUnscoredContainer from "@/components/teacher/unscored/TeacherUnscoredContainer.vue";

export default {
  components: {
    MemberFram,
    MemberTitle,
    TeacherUnscoredContainer,
  },
};
</script>
