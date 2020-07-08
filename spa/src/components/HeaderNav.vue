<template lang="html">
  <div id="app">
    <header>
      <div class="header-contents">
        <img src="http://acmelogos.com/images/logo-1.svg" alt="">
        <div class="main-right">
          <div class="right-links">
            <new-listing-button :new-listing-address='currentAddress'> </new-listing-button>
          </div>
          <div class="metamask">
            <metamask
              class="alignment"
              :meta-address='currentAddress'
              >
            </metamask>
          </div>
        </div>
      </div>
    </header>
  </div>
</template>

<script>
import Metamask from './HeaderNavMetamask';
import MetamaskMixin from '../mixins/MetamaskMixin';
import NewListingButton from './HeaderNavNewListingButton';
import Web3 from 'web3';

export default {
  name: 'header-nav',
  mixins: [MetamaskMixin],
  components: {
    Metamask,
    NewListingButton
  },
  data() {
    return {
      currentAddress: null
    }
  },
  created() {
    if (window.ethereum) {
      this.currentAddress = new Web3(window.ethereum).givenProvider.selectedAddress;
    }
  }
}
</script>

<style lang="css" scoped>

header {
  padding: 0px 1.5rem;
  display: flex;
  position: fixed;
  top: 0;
  width: 100%;
  min-height: 90px;
  border-bottom: 2px solid rgba(246, 246, 246, 1);
  text-align: center;
  background: #ffffff;
  z-index: 3;
  align-items: center;
}

.header-contents {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.main-right {
  display: flex;
  align-items: center;
}

.right-links {
  position: relative;
}

header p {
  padding-right: 4rem;
}

</style>
