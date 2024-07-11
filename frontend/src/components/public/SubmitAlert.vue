<template>
  <v-container>
    <v-dialog v-model="dialog" persistent max-width="500px" max-height="300px">
      <template v-slot:default>
        <v-card title="注意" color="blue-accent-4" variant="flat">
          <v-card-text> {{ msg }} </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn text="取消" @click="cancel" variant="tonal"></v-btn>
            <v-btn text="确认" @click="confirm" variant="tonal"></v-btn>
          </v-card-actions>
        </v-card>
      </template> </v-dialog
  ></v-container>
</template>

<script>
export default {
  name: "SubmitAlert",
  props: {
    msg: {
      type: String,
      default: "确认提交吗？提交后不可撤回！",
    },
  },
  data() {
    return {
      dialog: true,
    };
  },
  methods: {
    cancel() {
      this.dialog = false;
      setTimeout(() => {
        this.$emit("cancel");
      }, 500);
    },
    confirm() {
      this.dialog = false;
      this.$emit("confirm");
      setTimeout(() => {
        this.$emit("cancel");
      }, 500);
    },
  },
};
</script>
