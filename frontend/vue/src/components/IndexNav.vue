<template>
<v-container>
  <v-app-bar app elevation="0" flat color="#2b3447">
    <template v-slot:prepend >
      <v-icon class="ml-12">
        <v-img
          src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
        ></v-img
      ></v-icon>
      <v-app-bar-title class="ml-6" color="#304FFE">Growth Document</v-app-bar-title>         
    </template>



    <template v-slot >
    <v-menu >
      <template v-slot:activator="{ props }">
        <v-btn color="blue-grey-lighten-4" v-bind="props"> <strong>书院风采</strong> <v-icon icon="mdi-chevron-down"></v-icon></v-btn>
      </template>

      <v-list>
        <v-list-item v-for="(item, index) in items" :key="index">
          <v-list-item-title><a href="https://modernize-vue3-dark.netlify.app/" style="text-decoration: none ;color: inherit">{{ item.title }}</a></v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>


    <v-menu >
      <template v-slot:activator="{ props }">
        <v-btn color="blue-grey-lighten-4" v-bind="props"> <strong>近期活动</strong> <v-icon icon="mdi-chevron-down"></v-icon></v-btn>
      </template>

      <v-list>
        <v-list-item v-for="(item, index) in items" :key="index">
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <v-menu>
      <template v-slot:activator="{ props }">
        <v-btn color="blue-grey-lighten-4" v-bind="props"> <strong>123</strong> <v-icon icon="mdi-chevron-down"></v-icon></v-btn>
      </template>

      <v-list>
        <v-list-item v-for="(item, index) in items" :key="index">
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <v-btn color="blue" rounded="XL" @click="registerButton"> <strong>登录/注册</strong></v-btn>

    <v-menu  rounded>
        <template v-slot:activator="{ props }">
          <v-btn icon v-bind="props">
            <v-avatar color="blue" size="large">
              <v-img src="https://cdn.vuetifyjs.com/images/john.jpg" alt="John"></v-img>
            </v-avatar>
          </v-btn>
        </template>
        <v-card>
          <v-card-text>
            <div class="mx-auto text-center">
              <v-avatar color="brown">
                <span class="text-h5">{{ user.initials }}</span>
              </v-avatar>
              <h3>{{ user.fullName }}</h3>
              <p class="text-caption mt-1">{{ user.email }}</p>
              <v-divider class="my-3"></v-divider>
              <v-btn variant="text" rounded> Edit Account </v-btn>
              <v-divider class="my-3"></v-divider>
              <v-btn variant="text" rounded> Disconnect </v-btn>
            </div>
          </v-card-text>
        </v-card>
      </v-menu>
    </template>
  </v-app-bar>
</v-container>
</template>


<script>
  export default {
    data: () => ({
      items: [
        { title: 'Click Me' },
        { title: 'Click Me' },
        { title: 'Click Me' },
        { title: 'Click Me 2' },
      ],
      user: {
        initials: '王',
        fullName: 'John Doe',
        email: 'john.doe@doe.com',
      },
    }),

    methods: {
      goToProfile() {
        window.location.href = 'https://www.baidu.com';
      },
      registerButton() {
        window.location.href = '../login.html';
      },
    },
  }
</script>

