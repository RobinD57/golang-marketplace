import Vuex from 'vuex';
import Vue from 'vue';

Vue.use(Vuex);

export const store = new Vuex.Store({
  state: {
    listingCount: null,
    modalOpen: false
  },
  mutations: {
    incrementListingCount (state) {
      state.listingCount++
    }
  }
})
