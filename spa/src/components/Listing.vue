<template lang="html">
  <div class="listing" v-if="listing">
    <button class="user-info" @click="showModal = true">
      {{listing.user}}
    </button>
    <transition name="fade" appear>
      <div class="modal-overlay" v-if="showModal" @click="showModal = false"></div>
    </transition>
    <transition name="slide" appear>
      <div class="modal" v-if="showModal">
        <h1>{{listing.user}}</h1>
        <p>Member since:</p>
        <button class="user-info" @click="showModal = false" type="button" name="button">
          close
        </button>
      </div>
    </transition>
    <h1 class="listing_name">{{ listing.name }}</h1>
    <p class="listing_description">{{ listing.description }}</p>
    <p class="listing_price">${{ listing.price }}</p>
  </div>
</template>

<script>


export default {
  props: ['id'],
  data() {
    return {
      listing: null,
      endpoint: 'http://localhost:8080/listing/',
      showModal: false
    };
  },
  methods: {
    async fetchListing(id) {
      let res = await fetch(this.endpoint + id)
      let data = await res.json()
      return this.setResults(data);
    },
    setResults (results) {
      this.listing = results;
    },
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
      line-height: 1;
      font-weight: 900;
      z-index: 0;
  }

    .user-info {
      appearance:none;
      outline: none;
      border: none;
      background: none;
      cursor: pointer;
      display: inline-block;
      margin-top: 2rem;
      padding: 5px 10px;
      border-radius: 5px;
      background-color: rgba(255, 255, 255, 0.25);
      font-size: 18px;
      font-weight: 700;
      box-shadow:  2px 2px rgba(0, 0, 0, 0.4);
      transition: 0.4s ease-out;
      display: flex
  }
  .button:hover {
      box-shadow: 4px 4px rgba(0, 0, 0, 0.6);
  }

  .modal-overlay {
    position: absolute;
    top:0;
    left:0;
    right: 0;
    bottom: 0;
    z-index: 98;
    background-color: rgba(0, 0, 0, 0.1);
  }

  .fade-enter-active,
  .fade-leave-active {
    transition: opacity 0.5s;
  }

  .fade-enter,
  .fade-leave-to {
    opacity: 0;
  }

  .modal {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 99;

    width: 100%;
    height: 50vh;
    max-width: 400px;
    background-color: #FFF;
    border-radius: 16px;

    padding: 25px;
  }
  .modal h1 {
    color: #222;
    font-size: 32px;
    font-weight: 900;
    margin-bottom: 15px;
   }

  .modal p {
    color: #666;
    font-size: 18px;
    font-weight: 400;
    margin-bottom: 15px;
   }

  .black {
    color: black;
  }
</style>
