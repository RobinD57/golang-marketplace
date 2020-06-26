<template lang="html">
  <div class="">
    <div class="" v-if="!currentAddress">
      <div id="connect">
        <span>login</span>
        <button type="button" class="Login-button" id="Login-mm" @click="loginAndSetAddress">
          <img src="https://cdn.freebiesupply.com/logos/thumbs/1x/metamask-logo.png" alt="">
        </button>
      </div>
    </div>
    <div class="" v-if="currentAddress">
      <p @event="revealAddress">{{currentAddress}}</p>
    </div>
  </div>
</template>

<script>
    import Web3 from 'web3';
    export default {
        name: "Metamask",
        data() {
            return {
                currentAddress: "",
            }
        },
        methods: {
          async loadWeb3() {
            // Modern dapp browsers...
            if (window.ethereum) {
              window.web3 = new Web3(window.ethereum);
              await window.ethereum.enable()
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
            this.currentAddress = accounts[0]
          },
          loginAndSetAddress() {
            this.loadWeb3();
            this.loadBlockchainData();
          },
          revealAddress() {

          }
        },
        // created() {
        //     this.loadWeb3();
        //     this.loadBlockchainData();
        // },
    }
</script>

<style scoped>

    .Login-button {
        color: white;
        display: block;
        font-size: 16px;
        height: 40px;
        margin: 10px auto;
        width: 40px;
        border-radius: 50%;
        margin-right: 4rem;
        border: none;
        box-shadow: 0px 1px 2px rgba(0, 0, 0, 0.18);
        cursor: pointer;
    }
    p {
      padding-right: 4rem;
      text-shadow: 1px 1px 3px rgba(0,0,0,0.2);
    }

    span{
      margin-top: 1rem;
      margin-right: 1rem;
      text-shadow: 1px 1px 3px rgba(0,0,0,0.2);
    }

    .Login-button img {
      width: 28px;
      height: 28px;
    }

    .Login-button:hover {
        opacity: 0.9;
    }

    #Login-mm {
        background-color: rgba(0, 0, 0, 0.2);
    }

    #connect {
      display: flex;
    }
</style>
