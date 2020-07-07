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
        <div class="inputs-main">
          <div class="inputs-left">
            <!-- <div class="email-group">
              <FormulateInput
                type="email"
                name="email"
                label="email:"
                v-model='newListing.email'
                validation="^required|email" />
            </div> -->
            <div class="listing-name-group">
              <FormulateInput
                type="text"
                name="name"
                label="title of your listing:"
                v-model='newListing.name'
                validation="^required|min:5" />
            </div>
            <div class="price-group">
              <FormulateInput
                type="number"
                name="price"
                label="selling price:"
                v-model='newListing.price'
                validation="^required|number|min:1" />
            </div>
            <div class="description-group">
              <FormulateInput
                type="textarea"
                name="description"
                label="provide a short description:"
                v-model='newListing.description'
                validation="^required|min:1" />
            </div>
          </div>
          <div class="inputs-right">
            <div class="photo-group">
              <FormulateInput
                type="image"
                name="listing-photos"
                label="Select images to upload"
                v-model='newListing.photos'
                help="png, jpg or gif"
                multiple
              />
            </div>
          </div>
        </div>
        <br>
        <FormulateInput
          class="submit-button"
          type="submit"
          label="post"
          @click='finalSubmit'
        />
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
  props: ['user', 'newListingAddress'],
  name:'NewListingButton',
  mixins: [ModalMixin, MetamaskMixin],
  data() {
    return {
      modalOpen: false,
      currentAddress: this.newListingAddress,
      listingEndpoint: 'http://localhost:8080/listing',
      cloudinaryEndpoint: 'https://api.cloudinary.com/v1_1/duueqba0z/upload',
      newListing: {
        name: null,
        price: null,
        description: null,
        photos: [],
        seller: this.newListingAddress
      },
    }
  },
  methods: {
    async finalSubmit() {
      //upload images to cloudinary
      const urls = await this.uploadPhotos();
      this.postData(urls);
      this.resetState();
    },
     postData(urlArray) {
      // map urls into photos prop
      this.newListing.photos = urlArray;
      this.newListing.price = Number(this.newListing.price);
      fetch(this.listingEndpoint, {
        method: 'POST',
        body: JSON.stringify(this.newListing)
      })
      this.removeOverlay();
      this.$root.$emit('fetchListings', "hi");
    },
     async uploadPhotos() {
      const urlArray = [];
      const files = this.newListing.photos.files

      for await (const f of files) {
        const formData = new FormData();
        formData.append('upload_preset','iwc9hf9e');
        formData.append('file', f.file);
        let res = await fetch(this.cloudinaryEndpoint, {
          method: 'POST',
          body: formData
        });
        let data = await res.json();
        urlArray.push(data.secure_url);
      }
        return urlArray;
    },
    resetState() {
      this.newListing = {
        name: null,
        price: null,
        description: null,
        photos: [],
        seller: this.newListingAddress
      }
    },
  },
  watch: {
    modalOpen: function () {
      if (this.modalOpen) {
        this.loginAndSetAddress();
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
    width: 80%;
    max-width: 700px;
    height: 600px;
    background-color: #FFF;
    border-radius: 16px;
    padding: 20px;
    h1 {
      text-align: center;
      margin-bottom: 10px;
    }
  }

  .meta-logo img {
    margin-top: 5rem;
    height: 200px;
    width:200px;
  }

  .label {
    text-align: left !important;
  }

  .inputs-main {
    display: flex;
    justify-content: space-between;
    align-items: center;
    max-width: 650px;
    margin: 0 auto;

  }
  .inputs-left {
    display: flex;
    display: flex;
    flex-direction: column;
    width: 100;
    margin-left: 2rem;
    padding-bottom: 5rem;
    div {
      display: flex;
      align-items: baseline;
      margin-top: 1rem;
      p {
        margin-right: 1rem;
      }
    }
  }

  .inputs-right {
    display: flex;
    flex-direction: column;
  }

  .submit-button {
    width: 200px;
  }
</style>
