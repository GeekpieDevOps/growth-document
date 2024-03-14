<template>
  <v-app  class="indigo-darken-3">
    <IndexNav></IndexNav>

    <v-container fluid color="#2b3447">
      <v-parallax
          height="100vh"
          src="../../../public/imgs/IMG_0638-已增强-NR33.jpg"
          rounded="ml"
        >
          
            <!-- <h1  style="font-size: 120px ;color:antiquewhite"> 上海科技大学 </h1> -->

            <div class="d-flex align-center justify-center" style="height: 100%;font-size: 80px ;">
      <h1 class="display-2" style="color: antiquewhite;"> 上海科技大学 </h1>
    </div>

          
        </v-parallax>
    </v-container>

    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>
    <v-container class="mb-12" >
      <v-carousel cycle interval="3000" width="200"  show-arrows="hover" hide-delimiters>
        <v-carousel-item
          src="../../../public/imgs/_MG_1506-已增强-NR.jpg"
          cover
          rounded="xl"
        ></v-carousel-item>

        <v-carousel-item
          src="../../../public/imgs/IMG_9783_1.jpg"
          cover
          rounded="xl"
        ></v-carousel-item>

        <v-carousel-item
          src="../../../public/imgs/IMG_5278_1.jpg"
          cover
          rounded="xl"
        ></v-carousel-item>
      </v-carousel>
    </v-container>



    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>
    <v-container class="mt-12">
      <v-row justify="center" align="center" class="flex-container mb-6" >
        <v-chip variant="elevated">
          Colorful Life
        </v-chip>
      </v-row>
      <v-row justify="center" align="center" class="flex-container">
        <v-col cols="3.5" class="mr-8">
          <PictureCard :imageIndex="0"/>
        </v-col>
        <v-col cols="3.5" class="mr-8">
          <PictureCard :imageIndex="1"/>
        </v-col>
        <v-col cols="3.5" class="mr-8">
          <PictureCard :imageIndex="2"/>
        </v-col>
      </v-row>
    </v-container>
    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>
      <v-container class="text-center mt-12">
          <h2 class="display-2 font-weight-bold mb-3">VUETIFY FEATURES</h2>

          <v-responsive
            class="mx-auto mb-12"
            width="56"
          >
            <v-divider class="mb-1"></v-divider>

            <v-divider></v-divider>
          </v-responsive>

          <v-row>
            <v-col
              v-for="({ icon, title, text }, i) in features"
              :key="i"
              cols="12"
              md="4"
            >
              <v-card
                class="py-12 px-4"
                color="grey lighten-5"
                flat
              >
                <v-theme-provider dark>
                  <div>
                    <v-avatar
                      color="primary"
                      size="88"
                    >
                      <v-icon
                        large
                        v-text="icon"
                      ></v-icon>
                    </v-avatar>
                  </div>
                </v-theme-provider>

                <v-card-title
                  class="justify-center font-weight-black text-uppercase"
                  v-text="title"
                ></v-card-title>

                <v-card-text
                  class="subtitle-1"
                  v-text="text"
                >
                </v-card-text>
              </v-card>
            </v-col>
          </v-row>
        </v-container>



        <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>    <v-container class="mt-12"></v-container>

    <section id="stats">
        <v-parallax
          height="100vh"
          src="../../../public/imgs/IMG_0638-已增强-NR33.jpg"
          rounded="ml"
        >
          
            <!-- <h1  style="font-size: 120px ;color:antiquewhite"> 上海科技大学 </h1> -->

            <div class="d-flex align-center justify-center" style="height: 100%;font-size: 80px ;">
      <h1 class="display-2" style="color: antiquewhite;"> 上海科技大学 </h1>
    </div>

          
        </v-parallax>
      </section>

      <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>    <v-container class="mt-12"></v-container>
    <v-container class="mt-12"></v-container>
  </v-app>
</template>


<script>
  import PictureCard from '@/components/PictureCard.vue';
  export default{
    components: {
      PictureCard 
  },
    data: () =>({
      features: [
        {
          icon: 'mdi-account-group-outline',
          title: 'Vibrant Community',
          text: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Iusto cupiditate sint possimus quidem atque harum excepturi nemo velit tempora! Enim inventore fuga, qui ipsum eveniet facilis obcaecati corrupti asperiores nam',
        },
        {
          icon: 'mdi-update',
          title: 'Frequent Updates',
          text: 'Sed ut elementum justo. Suspendisse non justo enim. Vestibulum cursus mauris dui, a luctus ex blandit. Lorem ipsum dolor sit amet consectetur adipisicing elit. qui ipsum eveniet facilis obcaecati corrupti consectetur adipisicing elit.',
        },
        {
          icon: 'mdi-shield-outline',
          title: 'Long-term Support',
          text: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Iusto cupiditate sint possimus quidem atque harum excepturi nemo velit tempora! Enim inventore fuga, qui ipsum eveniet facilis obcaecati corrupti asperiores nam',
        },
      ],
      stats: [
          ['24k', 'Github Stars'],
          ['330+', 'Releases'],
          ['1m', 'Downloads/mo'],
          ['5m', 'Total Downloads'],
        ],
    })

  }

</script>

