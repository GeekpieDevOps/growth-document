<template>
  <TeacherUnscoredShow :activities="activities" />
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
          student_name: "Anna",
          activity_name: "xx摄影比赛",
          skills: ["能力1", "能力2"],
        },
        {
          student_name: "Jane",
          activity_name: "xx摄影比赛",
          skills: ["能力1", "能力2"],
        },
      ],
    };
  },
};
</script>
