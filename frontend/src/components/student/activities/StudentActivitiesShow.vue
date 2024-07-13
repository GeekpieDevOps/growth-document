<template>
  <v-table>
    <thead>
      <tr>
        <th
          class="font-weight-bold"
          v-for="h of heads"
          :class="{
            'd-none d-md-table-cell': h.text === '能力维度',
          }"
        >
          {{ h.text }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="activity of activities" :key="activity.name">
        <td>{{ activity.name }}</td>
        <td>{{ activity.updateTime }}</td>
        <td class="d-none d-md-table-cell">
          <v-chip
            v-for="skill of activity.skills"
            :key="skill"
            variant="tonal"
            color="primary"
            >{{ skill.name }}</v-chip
          >
        </td>
        <td>
          <v-dialog max-width="500">
            <template v-slot:activator="{ props: activatorProps }">
              <v-btn v-bind="activatorProps" color="info">查看</v-btn></template
            >

            <template v-slot:default="{ isActive }">
              <v-card>
                <v-card-title>能力维度</v-card-title>
                <v-card-text>
                  <v-table>
                    <tbody>
                      <tr v-for="skill of activity.skills">
                        <td>{{ skill.name }}</td>
                        <td>{{ skill.val }}</td>
                      </tr>
                    </tbody>
                  </v-table></v-card-text
                >
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn
                    @click="isActive.value = false"
                    color="info"
                    variant="tonal"
                    >关闭</v-btn
                  >
                </v-card-actions>
              </v-card>
            </template>
          </v-dialog>
        </td>
      </tr>
    </tbody></v-table
  >
</template>

<script>
export default {
  name: "StudentActivitiesShow",
  props: {
    activities: {
      type: Array,
      required: true,
      default: [],
    },
  },
  data() {
    return {
      heads: [
        { text: "活动名称" },
        { text: "更新时间" },
        { text: "能力维度" },
        { text: "操作" },
      ],
    };
  },
};
</script>
