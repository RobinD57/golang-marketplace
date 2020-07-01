<template lang="html">
  <div class="listing" v-if="listing">
      <div class="image-container" ref="gallery">
        <img draggable="false" v-bind:class="{ 'secondary-image': true, 'main-image': true,  } "
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5/9b2e37e8-78f3-42bb-b31d-b0cf629c9627/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
        <img draggable="false" v-bind:class="{ 'main-image': false, 'secondary-image': true }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/ef03abcb-b457-4d0c-81ef-74ddd36bf82b/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
        <img draggable="false" v-bind:class="{ 'main-image': false, 'secondary-image': true }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/7d2b8bee-a73a-4021-b839-5191ad7709e4/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
        <img draggable="false" v-bind:class="{ 'main-image': false, 'secondary-image': true }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/a18ffc7e-f172-47f2-87eb-1b4de4163b10/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
        <img draggable="false"  v-bind:class="{ 'main-image': false, 'secondary-image': true }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/a74cd00d-7272-4bd5-a898-dd03f233623c/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
        <img draggable="false" v-bind:class="{ 'main-image': false, 'secondary-image': true }"
        @click="resetImagesAndGrow"
        src="https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5,q_80/3e8834aa-2f0a-47d2-a7f6-b283447639c8/react-infinity-run-flyknit-mens-running-shoe-zX42Nc.jpg"
        alt="">
      </div>
    <div class="listing-details">
      <div class="listing-product">
        <h1 class="listing_name">{{ listing.name }}</h1>
        <p class="listing_description">{{ listing.description }}</p>
        <p class="listing_price">${{ listing.price }}</p>
        <PurchaseButton :seller='listing.seller' />
      </div>
      <div class="user-sidebar">
        <UserSideBar :seller='listing.seller' :postingDate='listing.createdAt.slice(0,10)'/>
      </div>
    </div>
  </div>
</template>

<script>

import UserSideBar from './UserSideBar';
import PurchaseButton from './PurchaseButton';

export default {
  props: ['id'],
  components: {
    UserSideBar,
    PurchaseButton
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
    async fetchData(attr, id) {
      let res = await fetch(`${this.endpoint}${id}`);
      let data = await res.json()
      return this.setResults(attr, data);
    },

    setResults(attr, results) {
      this.listing = results
    },
    resetImagesAndGrow(e) {
      this.$refs.gallery.children.forEach((img) => {
        img.classList.remove('main-image');
      })
      e.currentTarget.classList.add('main-image');
    }
  },
  created() {
    this.fetchData("listing", this.id);
  },
  watch: {
    '$route'() {
      this.fetchData("listing", this.id);
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

  .listing-images {
    display: flex;
  }

  .listing-details {
    display: flex;
    margin-top: 5rem;
    margin-left: 10px;
    width: 100%;
  }

  .listing-product {
    flex: 0 1 70%;
  }

  .user-sidebar {
    flex: 0 1 30%;
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
      margin-bottom: 5rem;
      line-height: 2;
      font-weight: 900;
      z-index: 0;
  }

  .image-container{
    min-height: 500px;
    display: flex;
    width: 100%;
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
    width: 8.7rem;
    margin-left: 10px;
    cursor: pointer;
  }

  .msg-button {
    display: flex;
    margin: 0 auto;
    margin-top: 3rem;
    margin-bottom: 3rem;
    width: 250px;
    font-family: 'Crimson Text', serif;
    font-size: 20px;
    opacity: .7;
    border: none;
    padding: 1rem;
    border-radius: 5px;
    justify-content: center;
    cursor: pointer;
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
    box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
    outline: none;
  }

  .msg-button:hover {
    opacity: .5;
  }

</style>
