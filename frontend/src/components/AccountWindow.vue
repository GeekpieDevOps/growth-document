<template>
  <v-card>
    <v-tabs v-model="tab" color="#5d87ff" height="65">
      <v-tab
        :value="1"
        prepend-icon="mdi-account-circle-outline"
        width="140"
        class="text-subtitle-1"
      >
        账户</v-tab
      ></v-tabs
    >
    <v-window>
      <v-window-item :value="1">
        <v-container class="py-7 px-4 px-sm-7"
          ><v-row
            ><v-col cols="12" sm="6"
              ><v-card
                elevation="2"
                class="d-flex flex-column justify-center align-center pa-4"
                ><v-card-title class="font-weight-bold align-self-start">
                  换头像
                </v-card-title>
                <p
                  class="align-self-start text-grey text-subtitle-1 mx-4"
                  style="line-height: 1.2em"
                >
                  Change your profile picture from here
                </p>
                <v-avatar
                  size="120"
                  image="https://randomuser.me/api/portraits/men/78.jpg"
                  class="mt-7 mb-10"
                ></v-avatar>
                <v-row>
                  <v-dialog persistent max-width="500">
                    <template v-slot:activator="{ props: activatorProps }">
                      <v-btn
                        class="mx-3"
                        color="#5d87ff"
                        width="80"
                        v-bind="activatorProps"
                        >上传</v-btn
                      ></template
                    >
                    <template v-slot:default="{ isActive }">
                      <v-card title="选择图像" class="px-12 py-5">
                        <v-card-action>
                          <v-form
                            @submit.prevent
                            class="my-5 d-flex flex-column"
                          >
                            <v-row>
                              <v-file-input
                                label="选择文件"
                                prepend-icon="mdi-file-image"
                                accept="image/png, image/jpeg"
                                clearable
                              ></v-file-input
                            ></v-row>
                            <v-row class="align-self-end">
                              <v-btn
                                text="关闭"
                                @click="isActive.value = false"
                                variant="text"
                                color="rgb(250, 137, 107)"
                              ></v-btn
                              ><v-btn
                                type="submit"
                                text="提交"
                                @click="isActive.value = false"
                                variant="text"
                                color="rgb(19, 222, 185)"
                              ></v-btn
                            ></v-row> </v-form></v-card-action></v-card></template
                  ></v-dialog>
                  <v-dialog max-width="350" persistent
                    ><template v-slot:activator="{ props: activatorProps }"
                      ><v-btn
                        class="mx-3"
                        color="rgb(250, 137, 107)"
                        variant="outlined"
                        width="80"
                        v-bind="activatorProps"
                        >重置</v-btn
                      ></template
                    >
                    <template v-slot:default="{ isActive }">
                      <Alert
                        title="确认重置"
                        text="重置后将无法撤回！"
                        activeText="重置"
                        :isActive="isActive"
                        :active="resetImg"
                      />
                    </template>
                  </v-dialog>
                </v-row>
                <v-row class="my-10 px-8"
                  ><p class="text-center text-grey">
                    允许使用.jpg, .png类型的文件。图片不大于800K
                  </p></v-row
                >
              </v-card></v-col
            ><v-col cols="12" sm="6">
              <v-card class="pa-4" elevation="2">
                <v-card-title class="font-weight-bold"> 修改密码 </v-card-title>
                <p
                  class="align-self-start text-grey text-subtitle-1 mx-4 mb-5"
                  style="line-height: 1.2em"
                >
                  To change your password please confirm here
                </p>
                <v-card-subtitle class="text-subtitle-1"
                  >修改密码</v-card-subtitle
                >
                <v-text-field
                  variant="outlined"
                  v-model="passwordModify"
                  color="rgb(93,135,255)"
                  type="password"
                  class="mt-3 mx-4"
                ></v-text-field>
                <v-card-subtitle class="text-subtitle-1"
                  >确认密码</v-card-subtitle
                ><v-text-field
                  variant="outlined"
                  v-model="passwordConfirm"
                  color="rgb(93,135,255)"
                  type="password"
                  class="mt-3 mx-4"
                ></v-text-field>
                <v-col cols="12" class="d-flex justify-center">
                  <v-dialog max-width="350" persistent
                    ><template v-slot:activator="{ props: activatorProps }"
                      ><v-btn color="#5d87ff" width="80" v-bind="activatorProps"
                        >确认</v-btn
                      >
                    </template>
                    <template v-slot:default="{ isActive }">
                      <Alert
                        title="确认修改"
                        text="修改后将无法撤回！"
                        activeText="修改"
                        :isActive="isActive"
                        :active="changePWD" /></template></v-dialog
                ></v-col>
              </v-card> </v-col
          ></v-row>
          <v-row>
            <v-col>
              <v-card class="pa-4" elevation="2">
                <v-card-title class="font-weight-bold">个人信息</v-card-title>
                <p
                  class="align-self-start text-grey text-subtitle-1 mx-4 mb-5"
                  style="line-height: 1.2em"
                >
                  To change your password please confirm here
                </p>
                <v-form @reset.prevent="userInfoFormRest">
                  <v-row>
                    <v-col cols="12" sm="6" v-for="(value, key) in accountInfo">
                      <v-card-subtitle class="text-subtitle-1">{{
                        key
                      }}</v-card-subtitle>
                      <v-text-field
                        v-if="!value.options"
                        v-model="accountInfo[key].content"
                        variant="outlined"
                        color="rgb(93,135,255)"
                        class="mt-3 mx-4"
                        :disabled="accountInfo[key].disabled"
                      ></v-text-field>
                      <v-select
                        v-model="accountInfo[key].content"
                        variant="outlined"
                        color="rgb(93,135,255)"
                        class="mt-3 mx-4"
                        v-if="value.options"
                        :items="value.options"
                      ></v-select> </v-col></v-row
                  ><v-row class="d-flex justify-end align-center mb-5 mx-5">
                    <v-spacer></v-spacer>
                    <v-dialog max-width="350" persistent
                      ><template v-slot:activator="{ props: activatorProps }"
                        ><v-btn
                          class="mx-4"
                          color="#5d87ff"
                          width="80"
                          v-bind="activatorProps"
                          >修改</v-btn
                        ></template
                      >
                      <template v-slot:default="{ isActive }">
                        <Alert
                          title="确认修改"
                          text="修改后将无法撤回！"
                          activeText="修改"
                          :isActive="isActive"
                          :active="userInfoFormSubmit"
                      /></template> </v-dialog
                    ><v-btn
                      type="reset"
                      class="mx-4"
                      color="rgb(250, 137, 107)"
                      variant="flat"
                      width="80"
                      >还原</v-btn
                    ></v-row
                  ></v-form
                ></v-card
              ></v-col
            ></v-row
          ></v-container
        >
      </v-window-item>
    </v-window></v-card
  >
</template>
<script>
import { ref } from "vue";
export default {
  name: "AccountWindow",

  setup() {
    const tab = ref(1);
    const passwordModify = ref("");
    const passwordConfirm = ref("");
    const getAccountInfo = {
      姓名: "John Doe",
      Email: "1234.shanghai.edu.cn",
      书院: "上道书院",
      学院: "信息学院",
    };
    const accountInfo = ref({
      姓名: { content: "", disabled: false },
      Email: { content: "", disabled: true },
      书院: {
        content: "",
        disabled: true,
        options: ["上道书院", "科道书院", "大道书院"],
      },
      学院: {
        content: "",
        disabled: true,
        options: ["信息学院", "数学所", "物质学院"],
      },
    });
    const userInfoChangeState = ref("");
    userInfoFormRest();

    function resetImg() {
      // fetch
    }
    function changePWD() {
      // fetch
    }
    function checkUserInfoChange() {
      let changeKeys = [];
      for (let key in accountInfo.value) {
        if (accountInfo.value[key].content !== getAccountInfo[key]) {
          changeKeys.push(key);
        }
      }
      return changeKeys;
    }

    function userInfoFormSubmit(isActive) {
      let changeKeys = checkUserInfoChange();
      if (changeKeys.length === 0) {
        isActive.value = false;
        return;
      }
      for (let key of changeKeys) {
        fetch(
          "http://10.20.235.113:8080/api/v1/user/6608eea3-ac62-41d3-af88-e55df834cb1c",
          {
            method: "PUT",
            body: JSON.stringify({
              token:
                "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjM2FjMTBkNy02NTI2LTQ2ODQtODZhOC03MDc0YTNlMGJkMzgiLCJhdWQiOlsiNjYwOGVlYTMtYWM2Mi00MWQzLWFmODgtZTU1ZGY4MzRjYjFjIl0sImp0aSI6IjlkZjRmNjY2LTAwMjUtNGU0ZC1hMmE5LWU3NTJmZDQwMTI4OCJ9.Fmp0PsBsS4UzUb6EBOAoiYWz_QCV-BW0vg8TWW0RvDo",
              name: key,
              value: accountInfo.value[key].content,
            }),
          }
        ).then((res) => {
          if (res.status === 200) {
            getAccountInfo[key] = accountInfo.value[key].content;
            isActive.value = false;
          }
        });
      }
    }

    function userInfoFormRest() {
      for (let key in accountInfo.value) {
        accountInfo.value[key].content = getAccountInfo[key];
      }
    }
    return {
      tab,
      passwordConfirm,
      passwordModify,
      accountInfo,
      userInfoChangeState,
      checkUserInfoChange,
      userInfoFormRest,
      userInfoFormSubmit,
      resetImg,
      changePWD,
    };
  },
};
</script>
