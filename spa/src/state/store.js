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
    },
    setListingCount (state, count) {
      state.listingCount = count;
    },
    setModalOpen (state, bool) {
      state.modalOpen = bool
    },
  },
  getters: {
    listingCount: state => state.listingCount,
    modalOpen: state => state.modalOpen
  }
})
