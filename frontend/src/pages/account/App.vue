<template>
  <v-app>
    <v-main>
      <v-container>
        <MemberFram /><v-container class="pa-0">
          <MemberTitle
            :Items="[
              { title: '我的', disabled: false, href: '#' },
              { title: '我的信息', disabled: true, href: '#' },
            ]"
          />
          <AccountWindow /> </v-container
      ></v-container>
    </v-main>
  </v-app>
</template>
