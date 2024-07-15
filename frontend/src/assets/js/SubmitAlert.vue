<template>
  <v-container>
    <v-dialog
      v-model="dialog"
      :persistent="styleOptions.isPersistent"
      max-width="500px"
      max-height="300px"
    >
      <template v-slot:default>
        <v-card variant="flat">
          <v-card-title :class="`text-${styleOptions.titleColor}`">
            {{ styleOptions.titleMsg }}
          </v-card-title>
          <v-card-text> {{ msg }} </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              v-if="styleOptions.isCancelShow"
              text="取消"
              @click="cancel"
              variant="text"
              color="primary"
            ></v-btn>
            <v-btn
              text="确认"
              @click="confirm"
              variant="text"
              color="red"
            ></v-btn>
          </v-card-actions>
        </v-card>
      </template> </v-dialog
  ></v-container>
</template>

<script>
export default {
  name: "SubmitAlert",
  props: {
    styleOptions: {
      type: Object,
    },
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
