<template>
  <v-carousel cycle interval="3000" width="200"  show-arrows="hover" hide-delimiters>
    <v-carousel-item
      :src="imageUrl"
      cover
      rounded="xl"
    ></v-carousel-item>

    <v-carousel-item
      :src="imageUrl"
      cover
      rounded="xl"
    ></v-carousel-item>

    <v-carousel-item
      :src="imageUrl"
      cover
      rounded="xl"
    ></v-carousel-item>
  </v-carousel>



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