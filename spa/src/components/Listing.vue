<template lang="html">
  <div class="listing" v-if="listing">
    <div class="listing-contents">
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
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/dbd113ff-1516-417b-ae5a-1a393bc716a8/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
      </div>
      <UserSideBar :user='listing.user' :postingDate='"2020 05 10"' />
    </div>
    <h1 class="listing_name">{{ listing.name }}</h1>
    <p class="listing_description">{{ listing.description }}</p>
    <p class="listing_price">${{ listing.price }}</p>
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
      endpoint: 'http://localhost:8080/listing/',
      showModal: false,
      isClicked: false
    }
  },
  methods: {
    async fetchListing(id) {
      let res = await fetch(this.endpoint + id)
      let data = await res.json()
      return this.setResults(data);
    },
    setResults(results) {
      this.listing = results;
    },
    resetImagesAndGrow(e) {

      this.$refs.gallery.children.forEach((img) => {
        img.classList.remove('main-image');
      })
      e.currentTarget.classList.add('main-image');
    }
  },
  created() {
    this.fetchListing(this.id);
  },
  watch: {
    '$route'() {
      this.fetchListing(this.id);
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
    justify-content: space-between;
  }

  .listing .listing_name {
    position: relative;
    text-transform: uppercase;
    z-index: 1;
  }
  .listing .listing_description {
      position: relative;
      z-index: 1;
    }
  .listing .listing_price {
      position: relative;
      font-size: 30px;
      opacity: .7;
      bottom: -50px;
      margin: 0;
      right: -20px;
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
  }

  .secondary-image {
    width: 9rem;
    margin-left: 10px;
    cursor: pointer;
  }
</style>
