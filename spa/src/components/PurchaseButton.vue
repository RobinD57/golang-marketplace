<template lang="html">
  <button @click='InitPurchase' class="msg-button" type="button" name="button">request to purchase</button>
</template>

<script>
export default {
  props: ['seller'],
  name:'PurchaseButton',
  data() {
    return {

    }
  },
  methods: {
    InitPurchase() {

    }
  }
}
</script>

<style lang="css" scoped>
</style>
