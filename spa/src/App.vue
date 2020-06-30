<template>
  <div id="app">
    <header-nav></header-nav>
    <main>
      <aside class="sidebar">
        <div class="collapse">
          <font-awesome-icon @click='toggleCollapse' :icon="['fas', 'chevron-left']"/>
        </div>
        <router-link
          :key="listing.title"
          v-for="listing in listings"
          class="link"
          :to="{ name: 'listing', params: { id: listing.id } }">
          <div class="">
            {{ listing.description }}
            <br>
            {{listing.name}}
          </div>
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
    },
    toggleCollapse() {
      document.querySelector('aside').style.flex = '0 1 3%';
    }
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Crimson+Text&family=Mukta:wght@200;400&family=Noto+Sans&display=swap');
  body {
    margin: 0;
    padding: 0;
  }
  #app {
    font-family: 'Crimson Text', serif;
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
    margin: 90px 0;
    margin-left: auto;
    margin-right: auto;

  }
  aside {
    flex: 0 1 20%;
    height: 100%;
    overflow-y: auto;
    width: 20%;
    padding: 50px 0 ;
    padding-right: 55px;
    box-sizing: border-box;
    border-right: 2px solid rgba(246, 246, 246, 1);
    border-bottom: 2px solid rgba(246, 246, 246, 1);
    transition: 1s;
  }
  .collapse {
    position: relative;
    width: 30px;
    height: 30px;
    border-radius: 50%;
    align-items: center;
    display: flex;
    justify-content: center;
    top: 50%;
    left: 110%;
    opacity: .7;
    cursor: pointer;
    padding: 2px;
  }

  .content {
    position: relative;
    flex: 1 1 80%;
    display: flex;
    align-items: center;
    justify-content: center;
    right: 0.5%;

  }
  .link {
    display: block;
    text-decoration: none;
    margin-bottom: 10px;
    color: #2c3e50;
  }
</style>
