<template lang="html">
  <div class="listing" v-if="listing">
      <div class="image-container" ref="gallery">
        <img v-bind:class="{ 'secondary-image': true, 'main-image': true,  }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5/9b2e37e8-78f3-42bb-b31d-b0cf629c9627/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
        <img v-bind:class="{ 'main-image': false, 'secondary-image': true }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/ef03abcb-b457-4d0c-81ef-74ddd36bf82b/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
        <img v-bind:class="{ 'main-image': false, 'secondary-image': true }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/a74cd00d-7272-4bd5-a898-dd03f233623c/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
        <img v-bind:class="{ 'main-image': false, 'secondary-image': true }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/a18ffc7e-f172-47f2-87eb-1b4de4163b10/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
        <img v-bind:class="{ 'main-image': false, 'secondary-image': true }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/7d2b8bee-a73a-4021-b839-5191ad7709e4/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
        <img v-bind:class="{ 'main-image': false, 'secondary-image': true }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/3e8834aa-2f0a-47d2-a7f6-b283447639c8/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
      </div>
    <div class="listing-details">
      <div class="">
        <h1 class="listing_name">{{ listing.name }}</h1>
        <p class="listing_description">{{ listing.description }}</p>
        <p class="listing_price">${{ listing.price }}</p>
      </div>
      <UserSideBar :user='listing.user' :postingDate='"2020 05 10"' :reviews='reviews' />
    </div>
  </div>
</template>

<script>

import UserSideBar from './UserSideBar';

export default {
  props: ['id'],
  components: {
    UserSideBar
  },
  data() {
    return {
      listing: null,
      reviews: [],
      endpoint: 'http://localhost:8080/listing/',
      showModal: false,
      isClicked: false
    }
  },
  methods: {
    async fetchData(attr, id, dest = "" ) {
      let res = await fetch(`${this.endpoint}${id}${dest}`);
      let data = await res.json()
      console.log(data);
      return this.setResults(attr, data);
    },

    setResults(attr, results) {
      attr == "listing" ? this.listing = results : this.reviews = results
    },
    resetImagesAndGrow(e) {
      this.$refs.gallery.children.forEach((img) => {
        img.classList.remove('main-image');
      })
      e.currentTarget.classList.add('main-image');
    },
    onLoadOrChange() {
      this.fetchData("listing", this.id);
      this.fetchData("reviews", this.id, "/reviews");
    }
  },
  created() {
    this.onLoadOrChange()
  },
  watch: {
    '$route'() {
      this.onLoadOrChange()
    }
  }
}

</script>

<style scoped>
  .listing {
    position: relative;
    width: 90%;
    margin: 0 auto;
    padding: 50px 20px 70px;
    height: 80vh;
    margin-top: 3.5rem;
  }

  .listing-contents {
    display: flex;
    flex-direction: row;
  }

  .listing-images {
    display: flex;
  }

  .listing-details {
    display: flex;
    justify-content: space-between;
    margin-top: 5rem;
  }

  .listing .listing_name {
    position: relative;
    text-transform: uppercase;
    z-index: 1;
    margin-top: 0;
  }
  .listing .listing_description {
      position: relative;
      z-index: 1;
    }
  .listing .listing_price {
      font-size: 16px;
      opacity: .9;
      margin: 0;
      line-height: 2;
      font-weight: 900;
      z-index: 0;
  }

  .image-container{
    min-height: 500px;
    display: flex;
    width: 60%;
  }

  .image-container img {
    border-radius: 5px;
    object-fit: cover;
    z-index: 0;
  }

  .main-image {
    width: 30rem !important;
    margin-left: 0 !important;
  }

  .secondary-image {
    width: 8.7rem;
    margin-left: 10px;
    cursor: pointer;
  }
</style>
