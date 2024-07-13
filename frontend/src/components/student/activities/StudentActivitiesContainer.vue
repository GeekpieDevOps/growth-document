<template>
  <StudentActivitiesShow :activities="activities" />
</template>

<script>
import { api, path } from "@/data.js";
import StudentActivitiesShow from "@/components/student/activities/StudentActivitiesShow.vue";

export default {
  name: "StudentActivitiesContainer",
  components: {
    StudentActivitiesShow,
  },
  data() {
    return {
      activities: [
        // {
        //   name: "活动asdhasdhaskdoajdoiawdyhya1",
        //   updateTime: "2021-01-01",
        //   skills: [
        //     { name: "能力1", val: 10 },
        //     { name: "能力2", val: 20 },
        //   ],
        // },
        // {
        //   name: "活动2",
        //   updateTime: "2021-01-02",
        //   skills: [{ name: "能力2", val: 0 }],
        // },
        // {
        //   name: "活动3",
        //   updateTime: "2021-01-03",
        //   skills: [
        //     { name: "能力1", val: 10 },
        //     { name: "能力3", val: 30 },
        //   ],
        // },
      ],
    };
  },
  mounted() {
    fetch(api.studentActivities, {
      method: "GET",
    })
      .then((response) => response.json())
      .then((data) => {
        this.activities = data.data;
      });
  },
};
</script>
