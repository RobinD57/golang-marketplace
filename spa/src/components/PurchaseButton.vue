<template lang="html">
  <div class="">
    <button @click='showModal' class="button" type="button" name="button">request to purchase</button>
    <transition name="slide" appear>
      <div class="modal" v-if="modalOpen">
        <button class="button" @click="modalOpen = false" type="button" name="button">
          Close
        </button>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  props: ['seller'],
  name:'PurchaseButton',
  data() {
    return {
      modalOpen: false
    }
  },
  methods: {
    showModal() {
      document.querySelector('header').insertAdjacentHTML("afterend",
        `<transition name="fade" appear>
          <div class='modal-overlay'>
          </div>
        </transition>`);
      this.modalOpen = true;
    }
  }
}
</script>

<style lang="css" scoped>

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

  .fade-enter-active,
  .fade-leave-active {
    transition: 0.5s;
  }

  .fade-enter,
  .fade-leave-to {
    opacity: 0;
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
