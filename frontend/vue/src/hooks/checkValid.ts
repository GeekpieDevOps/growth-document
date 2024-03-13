import { ref } from "vue";
export const primeRulesForId = ref([
  (value: string) => {
    if (value) return true;
    return "用户ID不能为空！";
  },
]);
export const primeRulesForPwd = ref([
  (value: string) => {
    if (value) return true;
    return "密码不能为空！";
  },
]);
