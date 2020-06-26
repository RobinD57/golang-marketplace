<template lang="html">
  <div class="listing" v-if="listing">
    <h1 class="listing__title">{{ listing.title }}</h1>
    <p class="listing__body">{{ listing.body }}</p>
    <p class="listing__id">{{ listing.id }}</p>
  </div>
</template>

<script>


export default {
  props: ['id'],
  data() {
    return {
      listing: null,
      endpoint: 'http://localhost:8080/listings/',
    };
  },
  methods: {
    async fetchListing(id) {
      let res = await fetch(this.endpoint + id)
      let data = await res.json()
      console.log(data)
      return this.setResults(data);
    },
    setResults (results) {
      this.listings = results;
    }
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
    &__title {
      position: relative;
      text-transform: uppercase;
      z-index: 1;
    }
    &__body {
      position: relative;
      z-index: 1;
    }
    &__id {
      position: absolute;
      font-size: 280px;
      bottom: -50px;
      margin: 0;
      color: #eeeeee;
      right: -20px;
      line-height: 1;
      font-weight: 900;
      z-index: 0;
    }
  }
</style>
