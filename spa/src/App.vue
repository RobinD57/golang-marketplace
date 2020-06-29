<template>
  <div id="app">
    <header-nav></header-nav>
    <main>
      <aside class="sidebar">
        <router-link
          :key="listing.title"
          v-for="listing in listings"
          class="link"
          :to="{ name: 'listing', params: { id: listing.id } }">
          {{ listing.description }}
          <br>
          {{listing.name}}
        </router-link>
      </aside>
      <div class="content">
        <router-view></router-view>
      </div>
    </main>
  </div>
</template>

<script>
import HeaderNav from '@/components/HeaderNav';

export default {
  name: 'app',
  props: ['id'],
  components: {
    HeaderNav
  },
  data() {
    return {
      listings: [],
      endpoint: 'http://localhost:8080/listings'
    };
  },

  created() {
    this.fetchListings();
  },

  methods: {
    async fetchListings() {
      let res = await fetch(this.endpoint)
      let data = await res.json()
      return this.setResults(data);
    },
    setResults (results) {
      this.listings = results;
    }
  }
}
</script>

<style>
  body {
    margin: 0;
    padding: 0;
  }
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
  }
  h1, h2 {
    font-weight: normal;
  }
  ul {
    list-style-type: none;
    padding: 0;
  }
  li {
    display: inline-block;
    margin: 0 10px;
  }

  main {
    display: flex;
    height: calc(100vh - 90px);
    max-width: 1700px;
    margin-top: 90px;
    margin-left: auto;
    margin-right: auto;
    overflow: hidden;
  }
  aside {
    flex: 1 0 20%;
    height: 100%;
    overflow-y: auto;
    width: 20%;
    padding: 50px 0 ;
    padding-right: 55px;
    box-sizing: border-box;
    border-right: 1px solid #42b983;
  }
  .content {
    flex: 1 1 80%;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .link {
    display: block;
    text-decoration: none;
    margin-bottom: 10px;
    color: #2c3e50;
  }
</style>
