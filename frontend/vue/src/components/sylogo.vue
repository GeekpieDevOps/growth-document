<script setup>
const props = defineProps({
  imgURL: String,
  webURL: String,
  imgColor: String,
  name: String,
});
</script>
<template>
  <div class="img">
    <v-tooltip activator="parent" location="top">{{ name }}</v-tooltip>
    <a :href="webURL">
      <img :src="imgURL" :style="`background-color: ${imgColor}; width: 100px`"
    /></a>
  </div>
</template>

<style scoped>
.img {
  width: fit-content;
  height: fit-content;
  filter: drop-shadow(0 0 10px rgba(0, 0, 0, 0.5));
}
img {
  clip-path: circle(49.1% at 50% 50.5%);
  position: relative;
  animation: floating infinite 2s linear;
}

img:hover {
  transform: scale(1.2);
  animation: none;
  transition: 0.1s;
}

@keyframes floating {
  0% {
    top: 0px;
  }
  25% {
    top: 3px;
  }
  50% {
    top: 0px;
  }
  75% {
    top: -3px;
  }
  100% {
    top: 0px;
  }
}
</style>
