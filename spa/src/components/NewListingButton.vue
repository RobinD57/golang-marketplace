<template lang="html">
  <div class="">
    <button
    @click='showModal'
    class='newlisting-button'
    type="button"
    name="button">
    post ad</button>
    <transition name="slide" appear>
      <div class="modal" v-bind:style='{display: "none"}' ref='modal'>
        <button
          class="close-button"
          @click="removeOverlay"
          type="button"
          name="button">
          x
        </button>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  props: ['user'],
  name:'NewListingButton',
  data() {
    return {
      modalOpen: false
    }
  },
  methods: {
    showModal() {
      this.modalOpen = true;
      if (this.modalOpen) {
        document.querySelector('.modal-overlay').style.display = "block";
        this.$refs.modal.style.display = "";
      }
    },
    removeOverlay() {
      this.modalOpen = false;
      document.querySelector('.modal-overlay').style.display = 'none';
      this.$refs.modal.style.display = "none"
    },
  }
}
</script>

<style lang="css" scoped>

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
  .shadowed {
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
  }

  .modal {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 10000;
    box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 1000px;
    height: 600px;
    background-color: #FFF;
    border-radius: 16px;

    padding: 25px;
  }

</style>
