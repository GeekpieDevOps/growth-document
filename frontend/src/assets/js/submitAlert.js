import { createApp } from "vue";
import { registerPlugins } from "@/plugins";
import SubmitAlert from "@/components/public/SubmitAlert.vue";

// import SubmitAlert from "@/components/public/SubmitAlert.vue";

const SubmitAlert = {
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
  render(ctx) {
    const { $props, $emit } = ctx;
    const { msg } = $props;
    const { cancel, confirm } = ctx;
    return () => (
      <v-container>
        <v-dialog
          vModel={dialog}
          persistent
          max-width="500px"
          max-height="300px"
        >
          {{
            default: () => (
              <v-card title="注意" color="blue-accent-4" variant="flat">
                <v-card-text>{msg}</v-card-text>
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn text="取消" onClick={cancel} variant="tonal"></v-btn>
                  <v-btn text="确认" onClick={confirm} variant="tonal"></v-btn>
                </v-card-actions>
              </v-card>
            ),
          }}
        </v-dialog>
      </v-container>
    );
  },
};

/**
 * @param {*} confirmCallBack
 * @param {*} msg
 * @example submitAlert(() => {console.log('submit')}, '确认提交吗？')
 */
export default function submitAlert(
  confirmCallBack,
  msg = "确认提交吗？提交后不可修改！"
) {
  const app = createApp(SubmitAlert, {
    msg,
    onCancel() {
      document.body.removeChild(root);
      app.unmount();
    },
    onConfirm() {
      confirmCallBack();
    },
  });

  registerPlugins(app);
  const root = document.createElement("div");
  document.body.appendChild(root);
  app.mount(root);
}
