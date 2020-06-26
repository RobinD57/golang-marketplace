<template lang="html">
  <div class="listing" v-if="listing">
    Sold by: {{listing.user}}
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
    max-width: 500px;
    margin: 0 auto;
    padding: 50px 20px 70px;
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
      position: absolute;
      font-size: 50px;
      bottom: -50px;
      margin: 0;
      color: #eeeeee;
      right: -20px;
      line-height: 1;
      font-weight: 900;
      z-index: 0;
  }
</style>
