<template lang="html">
  <div class="">
    <button
    @click='showModal'
    class='newlisting-button'
    type="button"
    name="button">
    post ad</button>
    <transition name="slide" appear>
      <div class="modal" v-if='currentAddress' v-bind:style='{display: "none"}' ref='modal'>
        <button
          class="close-button"
          @click="removeOverlay"
          type="button"
          name="button">
          x
        </button>
        <h1>post listing</h1>
        <div class="inputs">
          <div class="email-group">
            <p>email:</p>
              <FormulateInput
              type="email"
              name="email"
              validation="required|email" />
          </div>

          <div class="listing-name-group">
            <p>listing name:</p>
            <FormulateInput
              type="text"
              name="name"
              validation="required|min:5" />
          </div>
          <div class="price-group">
            <p class="">price:</p>
            <FormulateInput
              type="text"
              name="price"
              validation="required|number|min:1" />
          </div>
          <div class="description-group">
            <p class="">description:</p>
            <FormulateInput
              type="textarea"
              name="description"
              validation="required|min:1" />
          </div>
          <FormulateInput
            type="image"
            name="listing-images"
            label="Select an image to upload"
            help="png, jpg or gif"
          />
        </div>
      </div>
      <div class="modal" v-if='!currentAddress' v-bind:style='{display: "none"}' ref='modal'>
        <button
          class="close-button"
          @click="removeOverlay"
          type="button"
          name="button">
          x
        </button>
        <h1>please login with metamask to continue</h1>
        <div class="meta-logo">
          <img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fcdn.freebiesupply.com%2Flogos%2Flarge%2F2x%2Fmetamask-logo-png-transparent.png&f=1&nofb=1" alt="">
        </div>
      </div>
    </transition>
  </div>

</template>

<script>
import ModalMixin from '../mixins/ModalMixin';
import MetamaskMixin from '../mixins/MetamaskMixin';
import '../assets/scss/main.scss';


export default {
  props: ['user', 'NewListingAddress'],
  name:'NewListingButton',
  mixins: [ModalMixin, MetamaskMixin],
  data() {
    return {
      modalOpen: false,
      currentAddress: this.NewListingAddress,

    }
  },
  watch: {
    modalOpen: function () {
      if (this.modalOpen) {
        this.loginAndSetAddress()
      }
    }
  }
}
</script>

<style lang="scss" scoped>

  .newlisting-button {
    display: flex;
    align-items: center;
    width: 75px;
    height: 25px;
    font-family: 'Crimson Text', serif;
    font-size: 12px;
    font-weight: bold;
    opacity: .7;
    border: none;
    border-radius: 5px;
    justify-content: center;
    cursor: pointer;
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
    box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
    outline: none;
    margin-right: 2rem;
  }

  .newlisting-button:hover {
    opacity: .5;
  }

  .close-button {
    display: flex;
    align-items: center;
    width: 25px;
    height: 25px;
    font-family: 'Crimson Text', serif;
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
  .v {
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
  }

  .modal {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 10;
    box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 1000px;
    height: 600px;
    background-color: #FFF;
    border-radius: 16px;
    padding: 25px;
    h1 {
      text-align: center;
    }
  }

  .meta-logo img {
    margin-top: 5rem;
    height: 200px;
    width:200px;
  }

  .inputs {
    width: 100;
    div {
      display: flex;
      align-items: baseline;
      margin-top: 1rem;
      p {
        margin-right: 1rem;
      }
    }
  }
</style>
