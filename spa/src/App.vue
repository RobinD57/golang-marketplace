<template>
  <div id="app">
    <header-nav></header-nav>
    <main>
      <aside class="sidebar">
        <div @click='toggleCollapse' class="collapse">
          <font-awesome-icon :icon="['fas', 'chevron-left']"/>
        </div>
        <div class="" v-if="!asideShrunk">
          <router-link
            :key="listing.title"
            v-for="listing in listings"
            class="link"
            :to="{ name: 'listing', params: { id: listing.id } }">
            <div class="main-card">
            </div>
            <div class="blur">
              <div class="card-details">
                <span>
                  {{listing.name.length > 15 ? listing.name.slice(0,15) + "..." : listing.name}}
                </span>
                <span>${{listing.price}}</span>
              </div>
            </div>
          </router-link>
        </div>
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
      asideShrunk: false,
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
      if (!this.asideShrunk) {
         document.querySelector('aside').style.flex = '0 1 3%';
         document.querySelector('.collapse').style.transform = 'rotate(180deg)';
         document.querySelector('.listing').style.marginLeft = '10rem'
         document.querySelector('.listing-product').style.flex = '0 1 59%'
         this.asideShrunk = true;
    } else {
        document.querySelector('aside').style.flex = '';
        document.querySelector('.collapse').style.transform = '';
        document.querySelector('.listing').style.marginLeft = '';
        document.querySelector('.listing-product').style.flex = '0 1 70%'
        this.asideShrunk = false;
      }
    }
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Crimson+Text&family=Mukta:wght@200;400&family=Noto+Sans&display=swap');

  .main-card {
    background-image: url(https://process.fs.grailed.com/AJdAgnqCST4iPtnUxiGtTz/auto_image/cache=expiry:max/rotate=deg:exif/resize=height:1400,fit:scale/output=quality:90/compress/https://cdn.fs.grailed.com/api/file/fpe0IssNQEWfiLXRbMX1);
    background-size: cover;
    height: 200px;
    width: 70%;
    margin: 0 auto;
    margin-top: 30px;
    border-radius: 5px;
    box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
  }

  .blur {
    background: rgba(245, 245, 245, 0.9);
    margin: 0 auto;
    height: 50px;
    width: 70%;
    margin-top: -4px;
    padding: 2px;
    border-radius: 5px;
  }

  .card-details {
    position: relative;
    display: flex;
    justify-content: space-around;
    margin-top: 2.5px;
  }

  .card-details span {
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
    margin-top: 0.8rem;
  }

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
    overflow-x: hidden;
    width: 20%;
    padding: 50px 0 ;
    padding-right: 55px;
    box-sizing: border-box;
    border-right: 2px solid rgba(246, 246, 246, 1);
    border-bottom: 2px solid rgba(246, 246, 246, 1);
    border-radius: 5px;

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
    left: 109%;
    opacity: .5;
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
