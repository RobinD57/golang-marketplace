<template lang="html">
  <div id="app">
    <header>
      <div class="header-contents">
        <img src="http://acmelogos.com/images/logo-1.svg" alt="">
        <div class="main-right">
          <div class="right-links">
            <NewListingButton :NewListingAddress='currentAddress' />
          </div>
          <div class="metamask">
            <Metamask
              class="alignment"
              @click="loginAndSetAddress"
              :MetaAddress='currentAddress'
              >
            </Metamask>
          </div>
        </div>
      </div>
    </header>
  </div>
</template>

<script>
import Metamask from './Metamask';
import NewListingButton from './NewListingButton';
import Web3 from 'web3';

export default {
  name: 'header-nav',
  components: {
    Metamask,
    NewListingButton
  },
  data() {
    return {
      currentAddress: null
    }
  },
  methods: {
    async loadWeb3() {
      // Modern dapp browsers...
        if (window.ethereum) {
          window.web3 = new Web3(window.ethereum);
          await window.ethereum.enable();
          this.loadBlockchainData();
        }
          // Legacy dapp browsers...
        else if (window.web3) {
          window.web3 = new Web3(window.web3.currentProvider);
        }
        // Non-dapp browsers...
        else {
          window.alert('Non-Ethereum browser detected. You should consider trying MetaMask!');
        }
      },
    async loadBlockchainData() {
      const web3 = window.web3
      const accounts = await web3.eth.getAccounts()
      this.currentAddress = accounts[0];
    },
    loginAndSetAddress() {
      if (!this.currentAddress) {
        this.loadWeb3();
      }
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
