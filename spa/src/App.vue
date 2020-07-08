<template>
  <div id="app">
    <header-nav></header-nav>
    <transition id='overlay' appear>
      <div class='modal-overlay'></div>
    </transition>
    <main>
      <aside class="sidebar">
        <div class="">
          <h3 class='count' v-if='this.listings'>Listings: {{this.listingCount}}</h3>
        </div>
        <div @click='toggleCollapse' class="collapse">
          <font-awesome-icon :icon="['fas', 'chevron-left']"/>
        </div>
        <div class="" v-if="!asideShrunk">
          <router-link
            :key="listing.title"
            v-for="listing in listings"
            class="link"
            :to="{ name: 'listing', params: { id: listing.id } }">
            <main-card
              :id='listing.id'
              :name='listing.name'
              :price='listing.price'
              :photos='listing.photos'
              >
            </main-card>
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
import MainCard from '@/components/MainCard';
import ModalMixin from '../src/mixins/ModalMixin';

export default {
  name: 'app',
  props: ['id'],
  mixins: [ModalMixin],
  components: {
    HeaderNav,
    MainCard,
  },
  data() {
    return {
      listings: [],
      asideShrunk: false,
      endpoint: 'http://localhost:8080/listings'
    }
  },
  computed: {
    listingCount () {
      return this.$store.state.listingCount
    }
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
      this.$store.state.listingCount = results.length;
    },
    toggleCollapse() {
      if (!this.asideShrunk) {
         document.querySelector('aside').style.flex = '0 1 3%';
         document.querySelector('.collapse').style.transform = 'rotate(180deg)';
         document.querySelector('.collapse').style.left = '45%';
         document.querySelector('.listing').style.marginLeft = '10rem'
         document.querySelector('.listing-product').style.flex = '0 1 59%'
         document.querySelector('.count').style.display = 'none';
         this.asideShrunk = true;
    } else {
        document.querySelector('aside').style.flex = '';
        document.querySelector('.collapse').style.transform = '';
        document.querySelector('.collapse').style.left = '85%';
        document.querySelector('.listing').style.marginLeft = '';
        document.querySelector('.listing-product').style.flex = '0 1 70%'
        document.querySelector('.count').style.display = '';
        this.asideShrunk = false;
      }
    }
  },
  watch: {
    listingCount () {
      this.fetchListings();
    }
  }
}
</script>

<style lang='scss'>
@import url('https://fonts.googleapis.com/css2?family=Lato&display=swap');
  .modal-overlay {
    display: none;
    position: absolute;
    top:0;
    left:0;
    right: 0;
    bottom: 0;
    z-index: 2;
    background-color: rgba(0, 0, 0, 0.3);
    min-height: 2000px;
}

  body {
    margin: 0;
    padding: 0;
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
  }
  #app {
    font-family: 'Lato', sans-serif;
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
    box-sizing: border-box;
    border-right: 2px solid rgba(246, 246, 246, 1);
    border-bottom: 2px solid rgba(246, 246, 246, 1);
    border-radius: 5px;
    flex-direction: column;
    justify-content: center;
    h3 {
      margin: 0;
    }
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
    left: 85%;
    opacity: .5;
    cursor: pointer;
    padding: 2px;
    z-index: auto;
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


  .close-button {
    display: flex;
    align-items: center;
    width: 25px;
    height: 25px;
    font-family: 'Lato', sans-serif;
    font-size: 12px;
    font-weight: bold;
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
  .shadowed {
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
  }

  .modal {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 99;
    box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 1000px;
    height: 600px;
    background-color: #FFF;
    border-radius: 16px;

    padding: 25px;
}
</style>
