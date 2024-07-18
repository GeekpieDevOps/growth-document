<template>
  <v-sheet
    class="ma-0 pa-0 h-screen d-flex justify-center align-center userFormShow-v-sheet"
  >
    <v-card variant="text" width="500px" :loading="isLoading">
      <v-card-title class="font-weight-bold"
        >欢迎来到书院综合素质平台</v-card-title
      >
      <v-card-text class="mt-7">
        <v-row class="mb-6">
          <v-col>
            <v-btn
              variant="outlined"
              class="d-flex align-center"
              :style="{
                borderColor: '#f0f0f0',
              }"
              size="large"
            >
              <div class="userFormShow-E-icon">E</div>
              使用Egate登录
            </v-btn>
          </v-col>
        </v-row>
        <v-row class="mb-16"> <v-divider></v-divider></v-row>
        <v-form @submit.prevent="sub" ref="form" class="mb-10">
          <v-row class="mb-4">
            <v-text-field
              label="邮箱"
              type="email"
              variant="outlined"
              :rules="rules.email"
              v-model="email"
            ></v-text-field> </v-row
          ><v-row class="mb-4">
            <v-text-field
              label="密码"
              type="password"
              variant="outlined"
              :rules="rules.password"
              v-model="password"
            ></v-text-field
          ></v-row>
          <v-row>
            <v-btn
              width="500px"
              variant="flat"
              color="info"
              type="submit"
              height="50px"
              size="large"
              >{{ btnText }}</v-btn
            >
          </v-row>
        </v-form>
        <v-row>
          <slot name="information"></slot>
        </v-row>
      </v-card-text> </v-card
  ></v-sheet>
</template>

<script>
export default {
  name: "UserFormShow",
  props: {
    isLoading: {
      type: Boolean,
      default: false,
    },
    btnText: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      rules: {
        email: [
          (v) => !!v || "邮箱不能为空",
          (v) => /.+@shanghaitech.edu.cn/.test(v) || "请使用上海科技大学邮箱",
        ],
        password: [(v) => !!v || "密码不能为空"],
      },
      email: "",
      password: "",
    };
  },
  methods: {
    sub() {
      // @ts-ignore
      this.$refs.form.validate().then(({ valid }) => {
        if (!valid) return;

        this.$emit("sub", {
          email: this.email,
          password: this.password,
        });
      });
    },
  },
};
</script>

<style>
.userFormShow-v-sheet {
  background-color: white !important;
}

.userFormShow-E-icon {
  width: 1.5em !important;
  height: 1.5em !important;
  border-radius: 50% !important;
  background-color: lightblue !important;
  line-height: 1.5em !important;
  text-align: center !important;
  color: white !important;
  margin-right: 0.5em !important;
}

.v-field--error:not(.v-field--disabled) .v-field__outline,
.v-field--error:not(.v-field--disabled) .v-label.v-field-label,
.v-input--error:not(.v-input--disabled) .v-input__details > .v-icon,
.v-input--error:not(.v-input--disabled) .v-input__details .v-messages,
.v-input--error:not(.v-input--disabled) .v-input__prepend > .v-icon,
.v-input--error:not(.v-input--disabled) .v-input__prepend .v-messages,
.v-input--error:not(.v-input--disabled) .v-input__append > .v-icon,
.v-input--error:not(.v-input--disabled) .v-input__append .v-messages {
  color: rgb(250, 137, 107) !important;
}
</style>
