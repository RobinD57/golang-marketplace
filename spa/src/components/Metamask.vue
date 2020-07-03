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
    <div class="address"
      v-if="currentAddress"
      @mouseover="hover = true"
      @mouseleave="hover = false">
      <div class="short-address">
      <span>{{currentAddress.slice(0,3)}}<span v-if="!hover">...</span>
      <span v-if="!hover">{{currentAddress.slice(39)}} </span> </span>
    </div>
    <span class="full-address" v-if="hover">{{currentAddress.slice(4)}}</span>
    </div>
  </div>
</template>

<script>
import MetamaskMixin from '../mixins/MetamaskMixin';
  export default {
    props: ['MetaAddress'],
    name: "Metamask",
    mixins: [MetamaskMixin],
    data() {
      return {
        currentAddress: this.MetaAddress,
        hover: false
      }
    },
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
      margin-right: 4rem;
      text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
    }

    #connect span {
      margin-top: 1.2rem;
      margin-right: 0.5rem;
      text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
    }

    .Login-button img {
      width: 28px;
      height: 28px;
    }

    .Login-button:hover {
        opacity: 0.9;
    }

    .address {
      display: flex;
      align-items: center;
      margin-right: 8rem;
      padding: 10px;
      transition: .5s;
    }

    .address span {
      font-size: 14px;
      text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
    }

    #Login-mm {
      background-color: rgba(0, 0, 0, 0.2);
    }

    #connect {
      display: flex;
      margin-right: 4rem;
    }
</style>
