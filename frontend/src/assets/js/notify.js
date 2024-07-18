import { createApp } from "vue";
import { registerPlugins } from "@/plugins";

import SubmitAlert from "@/assets/js/SubmitAlert.vue";

const styleMap = {
  info: {
    titleMsg: "提示",
    titleColor: "primary",
    isCancelShow: false,
    isPersistent: false,
    icon: "mdi-lightbulb-outline",
  },
  error: {
    titleMsg: "错误",
    titleColor: "error",
    isCancelShow: false,
    isPersistent: false,
    icon: "mdi-alert-circle-outline",
  },
  alert: {
    titleMsg: "注意",
    titleColor: "red",
    isCancelShow: true,
    isPersistent: true,
    icon: "mdi-alert-octagram",
  },
};

/**
 * @type {import("../../types/types.d.ts").notify}
 * @example notify("info", "Are you sure to submit?", () => {console.log('submit')})
 */
function notify(type, msg, callback = () => {}) {
  const app = createApp(SubmitAlert, {
    styleOptions: styleMap[type],
    msg: msg,
    onCancel() {
      document.body.removeChild(root);
      app.unmount();
    },
    onConfirm() {
      callback();
    },
  });

  registerPlugins(app);
  const root = document.createElement("div");
  document.body.appendChild(root);
  app.mount(root);
}

export default notify;
