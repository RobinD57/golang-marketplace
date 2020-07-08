<template lang="html">
  <div class="">
    <button @click='showModal' class="purchase-button" type="button" name="button">
      request to purchase
    </button>
    <transition name="slide" appear>
      <div class="modal" v-bind:style='{display: "none"}' ref='modal'>
        <button
          class="close-button"
          @click="removeOverlay('proceed')"
          type="button"
          name="button">
          x
        </button>
        <div v-if='!next' class="purchase-main">
          <div class="buying">
            <h2 class="shadowed">Requesting to buy:</h2>
            <div class="product-card">
              <div class="card-wrapper" ref="cWrap"></div>
            </div>
          </div>
          <div class="selling">
            <h2 class='shadowed'>Seller:
              <br>
               <span class="shadowed">
                 {{listing.seller.slice(0,3)}}...{{listing.seller.slice(39)}}
               </span>
             </h2>
             <div class="">
               <h2 class='shadowed'>Price:
                 <br>
                 <span class='shadowed'>${{listing.price}}</span>
               </h2>
             </div>
             <div class="">
               <h2 class='shadowed'>Listing id:
                 <br>
                 <span class='shadowed'>${{listing._id}}</span>
               </h2>
             </div>
          </div>
        </div>
        <div v-else-if='next' class='redirect-metamask'>
          <h3>please complete transaction with metamask -></h3>
          <button class="purchase-button" @click="removeOverlay" type="button" name="button">
            ok
          </button>
        </div>
        <button v-if="!next" class="purchase-button" @click="next = true" type="button" name="button">
          proceed
        </button>
      </div>
    </transition>
  </div>
</template>

<script>
import ModalMixin from '../mixins/ModalMixin';

export default {
  props: {
    listing: Object
  },
  name:'PurchaseButton',
  mixins: [ModalMixin],
  data() {
    return {
      modalOpen: false,
      next: false,
    }
  },
}
</script>

<style scoped>


  .selling {
    display: flex;
    flex-direction: column;
    justify-content: space-around;
  }

  .selling span {
    font-size: 18px;
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
  }

  .buying {
    width: 50%;
    padding-left: 5rem;
  }

  .buying span {
    font-size: 16px;
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
  }

  .purchase-main {
    display: flex;
    margin-top: 3rem;
  }

  .product-card {
    width: 250px;
    height: 250px;
    margin-top: 0rem;
  }

  .purchase-button {
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
    margin-top: 6rem;
  }

  .redirect-metamask {
    display: flex;
    flex-direction: column;
    height: 80%;
    align-items: center;
    justify-content: center;
  }

  .purchase-button:hover {
    opacity: .5;
  }

  .modal {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 2;
    box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
    width: 80%;
    max-width: 700px;
    height: 600px;
    background-color: #FFF;
    border-radius: 16px;

    padding: 25px;
  }

</style>
