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
    }
  },
  methods: {
    getlisting(id) {
      this.listing = 
  }
},
  watch: {
    '$route'() {
      this.getPost(this.id);
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
