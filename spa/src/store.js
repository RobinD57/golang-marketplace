import Vuex from 'vuex';
import Vue from 'vue';

Vue.use(Vuex);

export const store = new Vuex.Store({
  state: {
    listingCount: 14
  },
  mutations: {
    incrementListingCount (state) {
      state.listingCount++
    }
  }
})
