<template>
  <TeacherUnscoredShow :activities="activities" @submit="submit" />
</template>

<script>
import TeacherUnscoredShow from "./TeacherUnscoredShow.vue";
export default {
  name: "TeacherUnscoredContainer",
  components: {
    TeacherUnscoredShow,
  },
  data() {
    return {
      activities: [
        {
          student: { name: "Anna", uuid: "6523" },
          activity: { name: "xx摄影比赛", uuid: "1132" },
          skills: ["能力1", "能力2"],
        },
        {
          student: { name: "John", uuid: "652123" },
          activity: { name: "xx摄影比赛", uuid: "1132" },
          skills: ["能力1", "能力2"],
        },
      ],
    };
  },
  methods: {
    submit() {
      console.log("submit");
    },
  },
};
</script>
