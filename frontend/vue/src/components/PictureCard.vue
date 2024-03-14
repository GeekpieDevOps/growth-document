<template>
  <v-card class="mx-auto" max-width="400" rounded="xl">
    <v-img
      class="align-end text-white"
      height="200"
      :src="imageUrl"
      cover>
      <v-card-title>科道书院</v-card-title>
    </v-img>
    <v-card-subtitle class="pt-4"> 2024.3.14 </v-card-subtitle>
    <v-card-text>
      <div>科道庭院</div>

      <div>沐浴着午后的阳光，同学们在科道书院写下贺卡，留下祝福</div>
    </v-card-text>

    <v-card-actions class="align-center">
      <!-- <v-btn color="orange"> Share </v-btn> -->

      <v-btn color="orange" variant="tonal" density="comfortable" rounded="xl"> <strong>详情</strong> </v-btn>
    </v-card-actions>
  </v-card>
</template>



<script>
export default {
  props: {
    imageIndex: {
      type: Number,
      required: true
    }
  },
  computed: {
    imageUrl() {
      return this.imageList[this.imageIndex];
    }
  },
  data() {
    return {
      imageList: [
        '../../public/imgs/IMG_0663.jpg',
        '../../public/imgs/IMG_0673.jpg',
        '../../public/imgs/IMG_0663.jpg',
      ]
    };
  }
};
</script>
