<template>
  <v-app  class="bg-pink-lighten-5">
    <v-main>
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
    <div class="bg-blue-lighten-5">
    <v-container class="mt-12 ">
      <v-row justify="center" align="center" class="flex-container mb-6" >
        <v-chip variant="elevated" class="mt-12">
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
  </div>
    <v-container class="mt-12"></v-container>



    <v-container class="mt-2">
      <v-row justify="center">
    
    
        <v-card
    class="mx-auto"
    max-width="500"
  >
    <v-container fluid>
      <v-row dense>
        <v-col
          v-for="card in cards"
          :key="card.title"
          :cols="card.flex"
        >
          <v-card>
            <v-img
              :src="card.src"
              class="align-end"
              gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
              height="200px"
              cover
            >
              <v-card-title class="text-white" v-text="card.title"></v-card-title>
            </v-img>

            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn color="surface-variant" icon="mdi-heart" size="small" variant="text"></v-btn>

              <v-btn color="surface-variant" icon="mdi-bookmark" size="small" variant="text"></v-btn>

              <v-btn color="surface-variant" icon="mdi-share-variant" size="small" variant="text"></v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-card>


  <v-card
    class="mx-auto"
    max-width="500"
  >
    <v-container fluid>
      <v-row dense>
        <v-col
          v-for="card in cards"
          :key="card.title"
          :cols="card.flex"
        >
          <v-card>
            <v-img
              :src="card.src"
              class="align-end"
              gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
              height="200px"
              cover
            >
              <v-card-title class="text-white" v-text="card.title"></v-card-title>
            </v-img>

            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn color="surface-variant" icon="mdi-heart" size="small" variant="text"></v-btn>

              <v-btn color="surface-variant" icon="mdi-bookmark" size="small" variant="text"></v-btn>

              <v-btn color="surface-variant" icon="mdi-share-variant" size="small" variant="text"></v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</v-row>
</v-container>

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




    <!-- <v-container class="mt-12 bg-c blue-accent-2 height=10vh"></v-container>  -->
    <!-- <div class="bg-blue-lighten-3" >
      <v-container class="mt-12"  >
        <v-row >
          <v-col class="align-start">
            
          <v-container class="mt-12"></v-container>  
          <v-sheet class="bg-blue-lighten-3 text-center" >
            <p style="font-size: 40px; color: white" class="mb-6">欢迎光临我们的书院 </p>
            <p style="font-size: 30px; color: white">欢迎光临我们的书院 </p>
            <v-container class="mt-12"></v-container>
          </v-sheet>
          </v-col>
          <v-col >
            <img src="../../../public/imgs/tower.svg" alt="Welcome to our college" height="100vh">
        </v-col>

        </v-row>
      </v-container>
    </div> -->

    <!-- <v-content>
  <v-container>
    <v-row 
           align="right" 
           align-content="right" 
           class="text-right"
           >
      <v-col class="text-right">
        <v-img 
               max-height="200px" 
               max-width="200px"
               src="https://picsum.photos/200/300"
               align="right"
               ></v-img>
      </v-col>
    </v-row>  
  </v-container>
</v-content>


    <v-container class="mt-12"></v-container>  <v-container class="mt-12"></v-container>  <v-container class="mt-12"></v-container> -->


  <v-footer
    class="bg-indigo-lighten-1 text-center d-flex flex-column"
  >
    <div>
      <v-btn
        v-for="icon in footer_icons"
        :key="icon"
        icon
        class="mx-4 mb-4 mt-4"
        variant="text"
       
      ><img :src="icon" width="50px"></v-btn>
    </div>

    <div class="pt-0">
      Phasellus feugiat arcu sapien, et iaculis ipsum elementum sit amet. Mauris cursus commodo interdum. Praesent ut risus eget metus luctus accumsan id ultrices nunc. Sed at orci sed massa consectetur dignissim a sit amet dui. Duis commodo vitae velit et faucibus. Morbi vehicula lacinia malesuada. Nulla placerat augue vel ipsum ultrices, cursus iaculis dui sollicitudin. Vestibulum eu ipsum vel diam elementum tempor vel ut orci. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.
    </div>
    <div>
      And thanks for our team!
    </div>
    <v-btn>
      <v-icon>mdi-arrow-up-bold</v-icon>
    </v-btn>

    <v-divider></v-divider>


  </v-footer>

</v-main>
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

      cards: [
        {
          title: 'Pre-fab homes',
          src: 'https://cdn.vuetifyjs.com/images/cards/house.jpg',
          flex: 12,
        },
        {
          title: 'Favorite road trips',
          src: 'https://cdn.vuetifyjs.com/images/cards/road.jpg',
          flex: 6,
        },
        {
          title: 'Best airlines',
          src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg',
          flex: 6,
        },
      ],
        footer_icons:["../../../public/imgs/大道书院.png","../../../public/imgs/科道书院.png","../../../public/imgs/上道书院.png"]
    })

  }

</script>

