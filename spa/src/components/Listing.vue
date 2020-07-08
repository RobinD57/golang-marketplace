<template lang="html">
  <div class="listing" v-if="listing">
    <div class="image-container" ref="gallery">
      <img
        v-for='photo in listing.photos'
        :src='photo'
        alt=""
        draggable="false"
        v-bind:class="[
          listing.photos[0] == photo ?
          { 'secondary-image': true, 'main-image': true }
          :
          { 'secondary-image': true, 'main-image': false }
          ]"
        @click="resetImagesAndGrow"
        v-bind:key='photo'
      >
    </div>
    <div class="listing-details">
      <div class="listing-product">
        <h1 class="listing_name">{{ listing.name }}</h1>
        <p class="listing_description">{{ listing.description }}</p>
        <p class="listing_price">${{ listing.price }}</p>
        <PurchaseButton :listing='this.listing' />
      </div>
      <div class="user-sidebar" >
        <user-side-bar
          :seller='listing.seller'
          :posting-date='listing.createdAt.slice(0,10)'
        >
        </user-side-bar>
      </div>
    </div>
  </div>
</template>

<script>

import UserSideBar from './ListingUserSideBar';
import PurchaseButton from './ListingPurchaseButton';

export default {
  props: {
    id: String
  },
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

</style>
