<template lang="html">
  <div class="">
    <button @click='showModal' class="button" type="button" name="button">request to purchase</button>
    <transition name="slide" appear>
      <div class="modal" v-bind:style='{display: "none"}' ref='modal'>
        <div class="purchase-main">
          <div class="">
            <span>buyer</span>
            <span>seller</span>
          </div>
          <div class="card-wrapper" ref="cWrap"></div>
        </div>
        <button class="button" @click="removeOverlay" type="button" name="button">
          Close
        </button>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  props: ['seller', 'id'],
  name:'PurchaseButton',
  data() {
    return {
      modalOpen: false
    }
  },
  methods: {
    showModal() {
      this.modalOpen = true;
      if (this.modalOpen) {
        document.querySelector('header').insertAdjacentHTML("afterend",
        `<transition id='overlay' appear>
        <div class='modal-overlay'>
        </div>
        </transition>`);
      }
      this.$refs.modal.style.display = ""
      const cardHTML = document.querySelector(`[data-id="${this.id}"]`).innerHTML;
      this.$refs.cWrap.innerHTML = cardHTML;
    },
    removeOverlay() {
      this.modalOpen = false;
      document.querySelector('#overlay').style.display = 'none';
      this.$refs.modal.style.display = "none"
    },
    created() {
    }
  }
}
</script>

<style lang="css" scoped>

  .main-card {
    background-image: url(https://process.fs.grailed.com/AJdAgnqCST4iPtnUxiGtTz/auto_image/cache=expiry:max/rotate=deg:exif/resize=height:1400,fit:scale/output=quality:90/compress/https://cdn.fs.grailed.com/api/file/fpe0IssNQEWfiLXRbMX1);
    background-size: cover;
    height: 200px;
    width: 70%;
    margin: 0 auto;
    margin-top: 30px;
    border-radius: 5px;
    box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
  }

  .blur {
    background: rgba(245, 245, 245, 0.9);
    margin: 0 auto;
    height: 50px;
    width: 70%;
    margin-top: -4px;
    padding: 2px;
    border-radius: 5px;
  }

  .card-details {
    position: relative;
    display: flex;
    justify-content: space-around;
    margin-top: 2.5px;
  }

  .card-details span {
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
    margin-top: 0.8rem;
  }

  .button {
    display: flex;
    margin: 0 auto;
    margin-top: 3rem;
    margin-bottom: 3rem;
    width: 250px;
    font-family: 'Crimson Text', serif;
    font-size: 20px;
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

  .msg-button:hover {
    opacity: .5;
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
