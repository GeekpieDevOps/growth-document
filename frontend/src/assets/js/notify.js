import { createApp } from "vue";
import { registerPlugins } from "@/plugins";

import SubmitAlert from "@/assets/js/SubmitAlert.vue";

const styleMap = {
  info: {
    titleMsg: "提示",
    titleColor: "primary",
    isCancelShow: false,
    isPersistent: false,
  },
  alert: {
    titleMsg: "注意",
    titleColor: "red",
    isCancelShow: true,
    isPersistent: true,
  },
};

/**
 * @type {import("../../types").notify}
 * @example notify("info", "Are you sure to submit?", () => {console.log('submit')})
 */
export default function notify(type, msg, callback) {
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
